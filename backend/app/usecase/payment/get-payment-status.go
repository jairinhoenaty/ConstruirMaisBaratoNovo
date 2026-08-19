package payment_usecase

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	pkgsubscription "construir_mais_barato/app/domain/subscription"
	pkgunlockedbudget "construir_mais_barato/app/domain/unlocked-budget"

	"gorm.io/gorm"
)

// ErrPaymentNotFound indica que o pagamento consultado não pertence a nenhum
// fluxo conhecido. O controller traduz para 404.
var ErrPaymentNotFound = errors.New("pagamento não encontrado")

// GetPaymentStatusUC responde se um pagamento já foi confirmado.
//
// Além de consultar o banco, ele reconcilia: se o registro ainda está pendente,
// reconsulta o MercadoPago e dispara a mesma ativação do webhook. É isso que
// garante que o fluxo destrave mesmo quando a notificação não chega (firewall,
// deploy fora do ar, retentativas esgotadas).
type GetPaymentStatusUC struct {
	SubscriptionService   pkgsubscription.SubscriptionService
	UnlockedBudgetService pkgunlockedbudget.UnlockedBudgetService
	Processor             *ProcessPaymentNotificationUC
	Throttle              *PollThrottle
}

type GetPaymentStatusUCParams struct {
	SubscriptionService   pkgsubscription.SubscriptionService
	UnlockedBudgetService pkgunlockedbudget.UnlockedBudgetService
	Processor             *ProcessPaymentNotificationUC
	// Throttle é compartilhado entre requisições. Nulo desliga o limite.
	Throttle *PollThrottle
}

// PaymentStatusOutput é público e não autenticado: não expõe dados do pagador
// nem o id do pagamento no MercadoPago.
type PaymentStatusOutput struct {
	Status    string     `json:"status"`
	Approved  bool       `json:"approved"`
	Amount    float64    `json:"amount"`
	PaidAt    *time.Time `json:"paidAt,omitempty"`
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
}

func NewGetPaymentStatusUC(params GetPaymentStatusUCParams) *GetPaymentStatusUC {
	return &GetPaymentStatusUC{
		SubscriptionService:   params.SubscriptionService,
		UnlockedBudgetService: params.UnlockedBudgetService,
		Processor:             params.Processor,
		Throttle:              params.Throttle,
	}
}

// Execute recebe o token opaco entregue no checkout, nunca o id do pagamento
// no MercadoPago. Como o endpoint é público, a chave precisa ser
// impossível de adivinhar ou enumerar.
func (uc *GetPaymentStatusUC) Execute(statusToken string) (*PaymentStatusOutput, error) {
	if statusToken == "" {
		return nil, ErrPaymentNotFound
	}

	subscription, unlocked, err := uc.find(statusToken)
	if err != nil {
		return nil, err
	}
	if subscription == nil && unlocked == nil {
		return nil, ErrPaymentNotFound
	}

	paymentID := int64(0)
	if subscription != nil {
		paymentID = subscription.PaymentID
	} else if id, err := strconv.ParseInt(unlocked.PaymentID, 10, 64); err == nil {
		paymentID = id
	}

	if paymentID != 0 && uc.shouldRefresh(paymentID, subscription, unlocked) {
		if _, err := uc.Processor.Execute(paymentID); err != nil {
			// A consulta ao MercadoPago falhou, mas ainda podemos responder o
			// que temos no banco. O cliente segue fazendo polling.
			fmt.Printf("erro ao reconciliar o pagamento %d: %v\n", paymentID, err)
		} else if subscription, unlocked, err = uc.find(statusToken); err != nil {
			return nil, err
		}
	}

	if subscription != nil {
		return &PaymentStatusOutput{
			Status:    string(subscription.Status),
			Approved:  subscription.IsApproved(),
			Amount:    subscription.Amount,
			PaidAt:    subscription.PaidAt,
			ExpiresAt: subscription.ExpiresAt,
		}, nil
	}

	return &PaymentStatusOutput{
		Status:   unlocked.Status,
		Approved: unlocked.IsPaid(),
		Amount:   unlocked.Amount,
		PaidAt:   unlocked.PaidAt,
	}, nil
}

func (uc *GetPaymentStatusUC) find(statusToken string) (
	*pkgsubscription.Subscription,
	*pkgunlockedbudget.UnlockedBudget,
	error,
) {
	subscription, err := uc.SubscriptionService.FindByStatusToken(statusToken)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, fmt.Errorf("erro ao buscar assinatura pelo token: %w", err)
	}
	if subscription != nil {
		return subscription, nil, nil
	}

	unlocked, err := uc.UnlockedBudgetService.FindByStatusToken(statusToken)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, fmt.Errorf("erro ao buscar desbloqueio pelo token: %w", err)
	}

	return nil, unlocked, nil
}

// shouldRefresh só reconsulta o MercadoPago para pagamentos ainda pendentes,
// respeitando o intervalo mínimo entre consultas.
//
// O controle de intervalo é externo (PollThrottle) e não derivado do UpdatedAt
// do registro: um pagamento pendente nunca tem o UpdatedAt alterado, então
// usá-lo bloquearia a primeira consulta e liberaria todas as seguintes —
// exatamente o oposto do desejado.
func (uc *GetPaymentStatusUC) shouldRefresh(
	paymentID int64,
	subscription *pkgsubscription.Subscription,
	unlocked *pkgunlockedbudget.UnlockedBudget,
) bool {
	if uc.Processor == nil {
		return false
	}

	pendente := (subscription != nil && subscription.IsPending()) ||
		(subscription == nil && unlocked != nil && unlocked.IsPending())
	if !pendente {
		return false
	}

	return uc.Throttle.ShouldCheck(paymentID)
}
