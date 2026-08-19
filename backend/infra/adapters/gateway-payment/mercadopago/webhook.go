package mercadopago

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// maxSignatureAge limita a janela de replay de uma notificação assinada.
const maxSignatureAge = 5 * time.Minute

// FlexID aceita o campo data.id tanto como string ("123456") quanto como
// número (123456). O MercadoPago usa string nos webhooks novos e número em
// alguns tópicos legados.
type FlexID string

func (f *FlexID) UnmarshalJSON(b []byte) error {
	b = bytes.TrimSpace(b)
	if len(b) == 0 || string(b) == "null" {
		*f = ""
		return nil
	}

	if b[0] == '"' {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		*f = FlexID(s)
		return nil
	}

	*f = FlexID(string(b))
	return nil
}

func (f FlexID) String() string {
	return string(f)
}

// Int64 converte o id para int64. Retorna erro se não for numérico.
func (f FlexID) Int64() (int64, error) {
	return strconv.ParseInt(strings.TrimSpace(string(f)), 10, 64)
}

// WebhookNotification é o corpo enviado pelo MercadoPago.
// Exemplo:
//
//	{"action":"payment.updated","api_version":"v1","data":{"id":"123456"},
//	 "date_created":"2021-11-01T02:02:02Z","id":"123456","live_mode":false,
//	 "type":"payment","user_id":2697177217}
type WebhookNotification struct {
	Action     string `json:"action"`
	APIVersion string `json:"api_version"`
	Data       struct {
		ID FlexID `json:"id"`
	} `json:"data"`
	DateCreated string `json:"date_created"`
	ID          FlexID `json:"id"`
	LiveMode    bool   `json:"live_mode"`
	Type        string `json:"type"`
	UserID      int64  `json:"user_id"`
}

// IsPayment indica se a notificação é sobre um pagamento. Outros tópicos
// (merchant_order, plan, subscription) são ignorados.
//
// Quando o campo type vem preenchido ele é a autoridade; o action só decide no
// formato IPN antigo, que não manda type.
func (n *WebhookNotification) IsPayment() bool {
	if n.Type != "" {
		return n.Type == "payment"
	}
	return strings.HasPrefix(n.Action, "payment.")
}

// ValidateSignature confere o header x-signature do MercadoPago.
//
// O header vem no formato "ts=<timestamp>,v1=<hash>" e o hash é um
// HMAC-SHA256, com a chave secreta da integração, sobre o manifesto:
//
//	id:<data.id>;request-id:<x-request-id>;ts:<ts>;
//
// Partes ausentes (por exemplo quando não há x-request-id) são omitidas do
// manifesto junto com seu rótulo, conforme a documentação do MercadoPago.
func ValidateSignature(xSignature, xRequestID, dataID, secret string) error {
	if secret == "" {
		return fmt.Errorf("segredo do webhook não configurado")
	}
	if xSignature == "" {
		return fmt.Errorf("header x-signature ausente")
	}

	ts, hash := parseSignatureHeader(xSignature)
	if ts == "" || hash == "" {
		return fmt.Errorf("header x-signature malformado")
	}

	if err := checkSignatureAge(ts); err != nil {
		return err
	}

	var manifest strings.Builder
	if dataID != "" {
		// O MercadoPago normaliza ids alfanuméricos para minúsculas.
		manifest.WriteString("id:" + strings.ToLower(dataID) + ";")
	}
	if xRequestID != "" {
		manifest.WriteString("request-id:" + xRequestID + ";")
	}
	manifest.WriteString("ts:" + ts + ";")

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(manifest.String()))
	expected := hex.EncodeToString(mac.Sum(nil))

	if !hmac.Equal([]byte(expected), []byte(strings.ToLower(hash))) {
		return fmt.Errorf("assinatura inválida")
	}

	return nil
}

func parseSignatureHeader(header string) (ts, hash string) {
	for _, part := range strings.Split(header, ",") {
		key, value, found := strings.Cut(strings.TrimSpace(part), "=")
		if !found {
			continue
		}
		switch strings.TrimSpace(key) {
		case "ts":
			ts = strings.TrimSpace(value)
		case "v1":
			hash = strings.TrimSpace(value)
		}
	}
	return ts, hash
}

// checkSignatureAge rejeita notificações antigas reenviadas por terceiros.
// O ts do MercadoPago vem em milissegundos desde a epoch.
func checkSignatureAge(ts string) error {
	millis, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return fmt.Errorf("timestamp da assinatura inválido: %w", err)
	}

	age := time.Since(time.UnixMilli(millis))
	if age < 0 {
		age = -age
	}
	if age > maxSignatureAge {
		return fmt.Errorf("assinatura expirada (idade de %s)", age.Round(time.Second))
	}

	return nil
}
