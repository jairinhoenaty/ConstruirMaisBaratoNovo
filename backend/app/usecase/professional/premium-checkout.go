package professional_usecase

import (
	pkgplan "construir_mais_barato/app/domain/plan"
	pkgprofessional "construir_mais_barato/app/domain/professional"
	pkgsubscription "construir_mais_barato/app/domain/subscription"
	pkgpaymentuc "construir_mais_barato/app/usecase/payment"
	"construir_mais_barato/infra/adapters/gateway-payment/mercadopago"
)

type CheckoutPremiumUC struct {
	Assembler           PayerAssembler
	PlanService         pkgplan.PlanService
	SubscriptionService pkgsubscription.SubscriptionService
	ProfessionalService pkgprofessional.ProfessionalService
	MercadoPagoClient   *mercadopago.MPClient
}

type PayerAssembler struct {
	UserID uint              `json:"userId"`
	Payer  mercadopago.Payer `json:"payer"`
}

// CheckoutPremiumOutput não devolve o id do pagamento no MercadoPago: o site
// acompanha o pagamento pelo statusToken, que é opaco e não permite deduzir
// nem enumerar pagamentos de outras pessoas.
type CheckoutPremiumOutput struct {
	StatusToken  string  `json:"statusToken"`
	Amount       float64 `json:"amount"`
	QRCode       string  `json:"qr_code"`
	QRCodeBase64 string  `json:"qr_code_base64"`
	Status       string  `json:"status"`
}

type CheckoutPremiumUCParams struct {
	PlanService         pkgplan.PlanService
	SubscriptionService pkgsubscription.SubscriptionService
	ProfessionalService pkgprofessional.ProfessionalService
	MercadoPagoClient   *mercadopago.MPClient
}

func NewCheckoutPremiumUC(params CheckoutPremiumUCParams) *CheckoutPremiumUC {
	return &CheckoutPremiumUC{
		PlanService:         params.PlanService,
		SubscriptionService: params.SubscriptionService,
		ProfessionalService: params.ProfessionalService,
		MercadoPagoClient:   params.MercadoPagoClient,
	}
}

func (uc *CheckoutPremiumUC) Execute() (*CheckoutPremiumOutput, error) {
	checkout := pkgpaymentuc.NewPlanCheckout(pkgpaymentuc.PlanCheckoutParams{
		MercadoPagoClient:   uc.MercadoPagoClient,
		PlanService:         uc.PlanService,
		SubscriptionService: uc.SubscriptionService,
	})

	result, err := checkout.Execute(pkgpaymentuc.PlanCheckoutInput{
		UserType: pkgplan.UserTypeProfessional,
		TargetID: uc.resolveProfessionalID(),
		Payer:    uc.Assembler.Payer,
	})
	if err != nil {
		return nil, err
	}

	return &CheckoutPremiumOutput{
		StatusToken:  result.StatusToken,
		Amount:       result.Amount,
		QRCode:       result.QRCode,
		QRCodeBase64: result.QRCodeBase64,
		Status:       result.Status,
	}, nil
}

// resolveProfessionalID descobre o id do profissional a partir do e-mail do
// pagador.
//
// O userId que chega do cliente é ambíguo: parte das telas envia o id do
// usuário logado e outra parte o id do profissional recém-cadastrado. Como é
// esse id que o webhook usa para liberar o premium, resolvemos pelo e-mail, que
// identifica o cadastro sem ambiguidade, e só caímos para o valor recebido
// quando não há profissional com aquele e-mail.
func (uc *CheckoutPremiumUC) resolveProfessionalID() uint {
	if uc.ProfessionalService == nil || uc.Assembler.Payer.Email == "" {
		return uc.Assembler.UserID
	}

	professional, err := uc.ProfessionalService.FindByEmail(uc.Assembler.Payer.Email)
	if err != nil || professional == nil {
		return uc.Assembler.UserID
	}

	return professional.ID
}
