package professional_usecase

import (
	"fmt"
	"os"
	"time"

	"construir_mais_barato/infra/adapters/gateway-payment/mercadopago"

	"github.com/google/uuid"
)

type CheckoutPremiumUC struct {
	Assembler PayerAssembler
}
type PayerAssembler struct {
	UserID uint              `json:"userId"`
	Payer  mercadopago.Payer `json:"payer"`
}

type CheckoutPremiumOutput struct {
	PaymentID    int64   `json:"paymentId"`
	Amount       float64 `json:"amount"`
	QRCode       string  `json:"qr_code"`
	QRCodeBase64 string  `json:"qr_code_base64"`
	Status       string  `json:"status"`
}

func NewCheckoutPremiumUC() *CheckoutPremiumUC {
	return &CheckoutPremiumUC{}
}

func (uc *CheckoutPremiumUC) Execute() (*CheckoutPremiumOutput, error) {
	mpClient := mercadopago.NewMPClient(os.Getenv("MERCADOPAGO_ACCESS_TOKEN"), os.Getenv("MERCADOPAGO_BASE_URL_API"))

	price := 19.90
	// appURL := os.Getenv("APP_PUBLIC_URL")
	// notificationURL := fmt.Sprintf("%s/webhooks/mercadopago", appURL)

	idem := uuid.NewString()
	desc := "Assinatura Premium mensal"
	extRef := fmt.Sprintf("user:%d:%d", uc.Assembler.UserID, time.Now().Unix())

	paymentInput := mercadopago.PixPaymentInput{
		Amount:      price,
		Description: desc,
		ExternalRef: extRef,
		// NotificationURL: notificationURL,
		IdempotencyKey: idem,
		Payer:          uc.Assembler.Payer,
	}

	res, err := mpClient.CreatePixPayment(paymentInput)
	if err != nil {
		return nil, err
	}
	return &CheckoutPremiumOutput{
		PaymentID:    res.PaymentID,
		Amount:       price,
		QRCode:       res.QRCode,
		QRCodeBase64: res.QRCodeBase64,
		Status:       res.Status,
	}, nil
}
