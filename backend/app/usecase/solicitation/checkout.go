package solicitation

import (
	"fmt"
	"os"
	"time"

	pkgplan "construir_mais_barato/app/domain/plan"
	"construir_mais_barato/infra/adapters/gateway-payment/mercadopago"

	"github.com/google/uuid"
)

// separado da assinatura premium do profissional.
type CheckoutUC struct {
	Assembler   CheckoutPayerAssembler
	PlanService pkgplan.PlanService
}

type CheckoutPayerAssembler struct {
	UserID uint              `json:"userId"`
	Payer  mercadopago.Payer `json:"payer"`
}

type CheckoutOutput struct {
	PaymentID    int64   `json:"paymentId"`
	Amount       float64 `json:"amount"`
	QRCode       string  `json:"qr_code"`
	QRCodeBase64 string  `json:"qr_code_base64"`
	Status       string  `json:"status"`
}

type CheckoutUCParams struct {
	PlanService pkgplan.PlanService
}

func NewCheckoutUC(params CheckoutUCParams) *CheckoutUC {
	return &CheckoutUC{
		PlanService: params.PlanService,
	}
}

func (uc *CheckoutUC) Execute() (*CheckoutOutput, error) {
	plan, err := uc.PlanService.FindByUserType(pkgplan.UserTypeSolicitation)
	if err != nil {
		return nil, fmt.Errorf("failed to find solicitation fee plan: %w", err)
	}

	if !plan.IsActive {
		return nil, fmt.Errorf("solicitation fee plan is not active")
	}

	mpClient := mercadopago.NewMPClient(os.Getenv("MERCADOPAGO_ACCESS_TOKEN"), os.Getenv("MERCADOPAGO_BASE_URL_API"))

	price := plan.Price

	idem := uuid.NewString()
	desc := plan.Name
	extRef := fmt.Sprintf("solicitation:%d:%d", uc.Assembler.UserID, time.Now().Unix())

	paymentInput := mercadopago.PixPaymentInput{
		Amount:         price,
		Description:    desc,
		ExternalRef:    extRef,
		IdempotencyKey: idem,
		Payer:          uc.Assembler.Payer,
	}

	res, err := mpClient.CreatePixPayment(paymentInput)
	if err != nil {
		return nil, err
	}

	return &CheckoutOutput{
		PaymentID:    res.PaymentID,
		Amount:       price,
		QRCode:       res.QRCode,
		QRCodeBase64: res.QRCodeBase64,
		Status:       res.Status,
	}, nil
}
