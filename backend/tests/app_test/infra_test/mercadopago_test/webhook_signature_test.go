package mercadopago_test

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"construir_mais_barato/infra/adapters/gateway-payment/mercadopago"

	"github.com/stretchr/testify/assert"
)

const (
	testSecret    = "segredo-de-teste"
	testDataID    = "123456"
	testRequestID = "d6f0f2d0-1b1a-4f5e-8a1b-000000000000"
)

// signManifest reproduz o cálculo que o MercadoPago faz ao assinar a
// notificação, para conferirmos a validação sem depender da rede.
func signManifest(t *testing.T, ts, dataID, requestID, secret string) string {
	t.Helper()

	manifest := fmt.Sprintf("id:%s;request-id:%s;ts:%s;", dataID, requestID, ts)
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(manifest))

	return hex.EncodeToString(mac.Sum(nil))
}

func currentTimestamp() string {
	return fmt.Sprintf("%d", time.Now().UnixMilli())
}

func TestValidateSignatureAceitaAssinaturaValida(t *testing.T) {
	ts := currentTimestamp()
	header := fmt.Sprintf("ts=%s,v1=%s", ts, signManifest(t, ts, testDataID, testRequestID, testSecret))

	err := mercadopago.ValidateSignature(header, testRequestID, testDataID, testSecret)

	assert.NoError(t, err)
}

func TestValidateSignatureRecusaSegredoErrado(t *testing.T) {
	ts := currentTimestamp()
	header := fmt.Sprintf("ts=%s,v1=%s", ts, signManifest(t, ts, testDataID, testRequestID, "outro-segredo"))

	err := mercadopago.ValidateSignature(header, testRequestID, testDataID, testSecret)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "assinatura inválida")
}

func TestValidateSignatureRecusaPagamentoTrocado(t *testing.T) {
	// Uma assinatura legítima de um pagamento não pode liberar outro.
	ts := currentTimestamp()
	header := fmt.Sprintf("ts=%s,v1=%s", ts, signManifest(t, ts, testDataID, testRequestID, testSecret))

	err := mercadopago.ValidateSignature(header, testRequestID, "999999", testSecret)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "assinatura inválida")
}

func TestValidateSignatureRecusaAssinaturaAntiga(t *testing.T) {
	ts := fmt.Sprintf("%d", time.Now().Add(-10*time.Minute).UnixMilli())
	header := fmt.Sprintf("ts=%s,v1=%s", ts, signManifest(t, ts, testDataID, testRequestID, testSecret))

	err := mercadopago.ValidateSignature(header, testRequestID, testDataID, testSecret)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "expirada")
}

func TestValidateSignatureRecusaHeaderAusenteOuMalformado(t *testing.T) {
	assert.Error(t, mercadopago.ValidateSignature("", testRequestID, testDataID, testSecret))
	assert.Error(t, mercadopago.ValidateSignature("ts=123", testRequestID, testDataID, testSecret))
	assert.Error(t, mercadopago.ValidateSignature("v1=abc", testRequestID, testDataID, testSecret))

	// Sem segredo configurado a validação não deve passar por engano.
	ts := currentTimestamp()
	header := fmt.Sprintf("ts=%s,v1=%s", ts, signManifest(t, ts, testDataID, testRequestID, testSecret))
	assert.Error(t, mercadopago.ValidateSignature(header, testRequestID, testDataID, ""))
}

// O MercadoPago manda data.id ora como texto, ora como número.
func TestWebhookNotificationAceitaDataIDComoTextoOuNumero(t *testing.T) {
	corpoDoEnunciado := `{
		"action": "payment.updated",
		"api_version": "v1",
		"data": {"id":"123456"},
		"date_created": "2021-11-01T02:02:02Z",
		"id": "123456",
		"live_mode": false,
		"type": "payment",
		"user_id": 2697177217
	}`

	var notification mercadopago.WebhookNotification
	assert.NoError(t, json.Unmarshal([]byte(corpoDoEnunciado), &notification))
	assert.Equal(t, "123456", notification.Data.ID.String())
	assert.True(t, notification.IsPayment())

	paymentID, err := notification.Data.ID.Int64()
	assert.NoError(t, err)
	assert.Equal(t, int64(123456), paymentID)

	var numerico mercadopago.WebhookNotification
	assert.NoError(t, json.Unmarshal([]byte(`{"type":"payment","data":{"id":789}}`), &numerico))
	assert.Equal(t, "789", numerico.Data.ID.String())

	var outroTopico mercadopago.WebhookNotification
	assert.NoError(t, json.Unmarshal([]byte(`{"type":"merchant_order","data":{"id":"1"}}`), &outroTopico))
	assert.False(t, outroTopico.IsPayment())
}
