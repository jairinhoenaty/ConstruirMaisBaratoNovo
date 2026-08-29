package store_usecase

import (
	pkgplan "construir_mais_barato/app/domain/plan"
	pkgstore "construir_mais_barato/app/domain/store"
	pkgsubscription "construir_mais_barato/app/domain/subscription"
	pkgpaymentuc "construir_mais_barato/app/usecase/payment"
	"construir_mais_barato/infra/adapters/gateway-payment/mercadopago"
)

type CheckoutPremiumUC struct {
	Assembler           PayerAssembler
	PlanService         pkgplan.PlanService
	SubscriptionService pkgsubscription.SubscriptionService
	StoreService        pkgstore.StoreService
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
	StoreService        pkgstore.StoreService
	MercadoPagoClient   *mercadopago.MPClient
}

func NewCheckoutPremiumUC(params CheckoutPremiumUCParams) *CheckoutPremiumUC {
	return &CheckoutPremiumUC{
		PlanService:         params.PlanService,
		SubscriptionService: params.SubscriptionService,
		StoreService:        params.StoreService,
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
		UserType: pkgplan.UserTypeStore,
		TargetID: uc.resolveStoreID(),
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

// resolveStoreID descobre o id da loja a partir do e-mail do pagador, pelo
// mesmo motivo descrito no checkout do profissional: o userId enviado pelo
// cliente ora é o id do usuário, ora o id da loja, e é esse id que o webhook
// usa para liberar o premium.
func (uc *CheckoutPremiumUC) resolveStoreID() uint {
	if uc.StoreService == nil || uc.Assembler.Payer.Email == "" {
		return uc.Assembler.UserID
	}

	store, err := uc.StoreService.FindByEmail(uc.Assembler.Payer.Email)
	if err != nil || store == nil {
		return uc.Assembler.UserID
	}

	return store.ID
}
