package payment_usecase

import (
	"fmt"
	"time"

	pkgprofessional "construir_mais_barato/app/domain/professional"
	pkgstore "construir_mais_barato/app/domain/store"
	pkgsubscription "construir_mais_barato/app/domain/subscription"
)

// ExpirePremiumsUC rebaixa quem tinha premium pago e cuja vigência terminou.
//
// Premium ativado manualmente por um administrador fica com premium_expires_at
// nulo e por isso não é afetado — o rebaixamento só atinge assinaturas com
// prazo definido.
type ExpirePremiumsUC struct {
	ProfessionalService pkgprofessional.ProfessionalService
	StoreService        pkgstore.StoreService
	SubscriptionService pkgsubscription.SubscriptionService
}

type ExpirePremiumsUCParams struct {
	ProfessionalService pkgprofessional.ProfessionalService
	StoreService        pkgstore.StoreService
	SubscriptionService pkgsubscription.SubscriptionService
}

type ExpirePremiumsOutput struct {
	Professionals int64
	Stores        int64
	Subscriptions int
}

func NewExpirePremiumsUC(params ExpirePremiumsUCParams) *ExpirePremiumsUC {
	return &ExpirePremiumsUC{
		ProfessionalService: params.ProfessionalService,
		StoreService:        params.StoreService,
		SubscriptionService: params.SubscriptionService,
	}
}

func (uc *ExpirePremiumsUC) Execute() (*ExpirePremiumsOutput, error) {
	now := time.Now()
	output := &ExpirePremiumsOutput{}

	professionals, err := uc.ProfessionalService.ExpirePremiums(now)
	if err != nil {
		return nil, fmt.Errorf("erro ao expirar premium de profissionais: %w", err)
	}
	output.Professionals = professionals

	stores, err := uc.StoreService.ExpirePremiums(now)
	if err != nil {
		return nil, fmt.Errorf("erro ao expirar premium de lojistas: %w", err)
	}
	output.Stores = stores

	expired, err := uc.SubscriptionService.FindExpired(now)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar assinaturas vencidas: %w", err)
	}

	for _, subscription := range expired {
		subscription.MarkAs(pkgsubscription.PaymentStatusExpired)
		if err := uc.SubscriptionService.Update(subscription); err != nil {
			return nil, fmt.Errorf("erro ao expirar assinatura %d: %w", subscription.ID, err)
		}
		output.Subscriptions++
	}

	return output, nil
}
