package solicitation

import (
	pkgplan "construir_mais_barato/app/domain/plan"
	pkgsubscription "construir_mais_barato/app/domain/subscription"
	pkgpaymentuc "construir_mais_barato/app/usecase/payment"
	"construir_mais_barato/infra/adapters/gateway-payment/mercadopago"
)

// CheckoutUC gera o PIX da taxa por solicitação paga pelo cliente no app,
// separado da assinatura premium do profissional.
type CheckoutUC struct {
	Assembler           CheckoutPayerAssembler
	PlanService         pkgplan.PlanService
	SubscriptionService pkgsubscription.SubscriptionService
	MercadoPagoClient   *mercadopago.MPClient
}

type CheckoutPayerAssembler struct {
	UserID uint              `json:"userId"`
	Payer  mercadopago.Payer `json:"payer"`
	// SolicitationID e ProfessionalID são opcionais: versões antigas do app não
	// os enviam. Quando vêm, ficam gravados para ligar o pagamento à
	// solicitação correspondente.
	SolicitationID string `json:"solicitationId"`
	ProfessionalID uint   `json:"professionalId"`
}

type CheckoutOutput struct {
	// PaymentID continua na resposta apenas por compatibilidade: as versões do
	// app já instaladas leem esse campo e o gravam no Firebase. Deve ser
	// removido assim que a atualização do app estiver distribuída — o
	// acompanhamento correto é pelo StatusToken.
	PaymentID    int64   `json:"paymentId"`
	StatusToken  string  `json:"statusToken"`
	Amount       float64 `json:"amount"`
	QRCode       string  `json:"qr_code"`
	QRCodeBase64 string  `json:"qr_code_base64"`
	Status       string  `json:"status"`
}

type CheckoutUCParams struct {
	PlanService         pkgplan.PlanService
	SubscriptionService pkgsubscription.SubscriptionService
	MercadoPagoClient   *mercadopago.MPClient
}

func NewCheckoutUC(params CheckoutUCParams) *CheckoutUC {
	return &CheckoutUC{
		PlanService:         params.PlanService,
		SubscriptionService: params.SubscriptionService,
		MercadoPagoClient:   params.MercadoPagoClient,
	}
}

func (uc *CheckoutUC) Execute() (*CheckoutOutput, error) {
	checkout := pkgpaymentuc.NewPlanCheckout(pkgpaymentuc.PlanCheckoutParams{
		MercadoPagoClient:   uc.MercadoPagoClient,
		PlanService:         uc.PlanService,
		SubscriptionService: uc.SubscriptionService,
	})

	result, err := checkout.Execute(pkgpaymentuc.PlanCheckoutInput{
		UserType:    pkgplan.UserTypeSolicitation,
		TargetID:    uc.Assembler.UserID,
		ReferenceID: uc.Assembler.SolicitationID,
		Payer:       uc.Assembler.Payer,
	})
	if err != nil {
		return nil, err
	}

	return &CheckoutOutput{
		PaymentID:    result.PaymentID,
		StatusToken:  result.StatusToken,
		Amount:       result.Amount,
		QRCode:       result.QRCode,
		QRCodeBase64: result.QRCodeBase64,
		Status:       result.Status,
	}, nil
}
