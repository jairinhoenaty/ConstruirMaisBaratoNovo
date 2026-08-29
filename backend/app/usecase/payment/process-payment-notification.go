package payment_usecase

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	pkgplan "construir_mais_barato/app/domain/plan"
	pkgprofessional "construir_mais_barato/app/domain/professional"
	pkgstore "construir_mais_barato/app/domain/store"
	pkgsubscription "construir_mais_barato/app/domain/subscription"
	pkgunlockedbudget "construir_mais_barato/app/domain/unlocked-budget"
	"construir_mais_barato/infra/adapters/gateway-payment/mercadopago"

	"gorm.io/gorm"
)

// amountTolerance absorve diferenças de arredondamento entre o valor cobrado e
// o que o MercadoPago devolve.
const amountTolerance = 0.01

// ProcessPaymentNotificationUC confirma um pagamento junto ao MercadoPago e
// libera o fluxo correspondente (premium do profissional, premium do lojista,
// taxa de solicitação do app ou desbloqueio de orçamento).
//
// O corpo do webhook carrega apenas o id do pagamento, nunca o status. Por
// isso este caso de uso sempre reconsulta a API do MercadoPago: uma notificação
// forjada não consegue liberar nada.
type ProcessPaymentNotificationUC struct {
	MercadoPagoClient     *mercadopago.MPClient
	SubscriptionService   pkgsubscription.SubscriptionService
	UnlockedBudgetService pkgunlockedbudget.UnlockedBudgetService
	ProfessionalService   pkgprofessional.ProfessionalService
	StoreService          pkgstore.StoreService
	PlanService           pkgplan.PlanService
}

type ProcessPaymentNotificationUCParams struct {
	MercadoPagoClient     *mercadopago.MPClient
	SubscriptionService   pkgsubscription.SubscriptionService
	UnlockedBudgetService pkgunlockedbudget.UnlockedBudgetService
	ProfessionalService   pkgprofessional.ProfessionalService
	StoreService          pkgstore.StoreService
	PlanService           pkgplan.PlanService
}

// ProcessResult descreve o que foi feito com a notificação. Serve para log e
// para o endpoint de status saber se houve mudança.
type ProcessResult struct {
	PaymentID   int64
	Status      pkgsubscription.PaymentStatus
	Handled     bool // a notificação corresponde a algum fluxo conhecido
	Activated   bool // o fluxo foi liberado agora
	Description string
}

func NewProcessPaymentNotificationUC(params ProcessPaymentNotificationUCParams) *ProcessPaymentNotificationUC {
	return &ProcessPaymentNotificationUC{
		MercadoPagoClient:     params.MercadoPagoClient,
		SubscriptionService:   params.SubscriptionService,
		UnlockedBudgetService: params.UnlockedBudgetService,
		ProfessionalService:   params.ProfessionalService,
		StoreService:          params.StoreService,
		PlanService:           params.PlanService,
	}
}

func (uc *ProcessPaymentNotificationUC) Execute(paymentID int64) (*ProcessResult, error) {
	payment, err := uc.MercadoPagoClient.GetPayment(paymentID)
	if err != nil {
		return nil, fmt.Errorf("erro ao consultar o pagamento %d no mercadopago: %w", paymentID, err)
	}

	// O desbloqueio de orçamento tem tabela própria, que já é gravada no
	// momento do checkout.
	unlocked, err := uc.UnlockedBudgetService.FindByPaymentID(strconv.FormatInt(paymentID, 10))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("erro ao buscar desbloqueio do pagamento %d: %w", paymentID, err)
	}
	if unlocked != nil {
		return uc.handleUnlockedBudget(unlocked, payment)
	}

	// Premium e taxa de solicitação ficam em subscriptions.
	subscription, err := uc.SubscriptionService.FindByPaymentID(paymentID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("erro ao buscar assinatura do pagamento %d: %w", paymentID, err)
	}

	if subscription == nil {
		// Pagamentos criados antes deste recurso não têm registro local.
		// Reconstruímos a partir do external_reference.
		subscription, err = uc.rebuildFromExternalReference(payment)
		if err != nil {
			return nil, err
		}
		if subscription == nil {
			return &ProcessResult{
				PaymentID:   paymentID,
				Status:      normalizeStatus(payment.Status),
				Handled:     false,
				Description: fmt.Sprintf("pagamento sem fluxo correspondente (external_reference=%q)", payment.ExternalReference),
			}, nil
		}
	}

	return uc.handleSubscription(subscription, payment)
}

func (uc *ProcessPaymentNotificationUC) handleSubscription(
	subscription *pkgsubscription.Subscription,
	payment *mercadopago.PaymentQueryResult,
) (*ProcessResult, error) {
	status := normalizeStatus(payment.Status)

	result := &ProcessResult{
		PaymentID: subscription.PaymentID,
		Status:    status,
		Handled:   true,
	}

	// Idempotência: o MercadoPago reenvia a mesma notificação várias vezes.
	if subscription.Status == status {
		result.Description = "status inalterado"
		return result, nil
	}

	// Notificações podem chegar fora de ordem. Um "pending" atrasado não pode
	// desfazer uma aprovação que já aconteceu.
	if subscription.IsApproved() && status == pkgsubscription.PaymentStatusPending {
		result.Status = subscription.Status
		result.Description = "notificação atrasada ignorada: pagamento já aprovado"
		return result, nil
	}

	if status == pkgsubscription.PaymentStatusApproved {
		if math.Abs(payment.TransactionAmount-subscription.Amount) > amountTolerance {
			subscription.MarkAs(pkgsubscription.PaymentStatusFailed)
			if err := uc.SubscriptionService.Update(subscription); err != nil {
				return nil, fmt.Errorf("erro ao marcar assinatura %d como inválida: %w", subscription.ID, err)
			}
			result.Status = pkgsubscription.PaymentStatusFailed
			result.Description = fmt.Sprintf(
				"valor divergente: cobrado %.2f, pago %.2f — premium não liberado",
				subscription.Amount, payment.TransactionAmount,
			)
			return result, nil
		}

		return uc.activateSubscription(subscription, payment, result)
	}

	subscription.MarkAs(status)
	if err := uc.SubscriptionService.Update(subscription); err != nil {
		return nil, fmt.Errorf("erro ao atualizar assinatura %d: %w", subscription.ID, err)
	}

	// Estorno/chargeback derruba o premium liberado antes.
	if status == pkgsubscription.PaymentStatusRefunded {
		if err := uc.setPremium(subscription, false, nil); err != nil {
			return nil, err
		}
		result.Description = "pagamento estornado, premium revogado"
		return result, nil
	}

	result.Description = fmt.Sprintf("assinatura atualizada para %q", status)
	return result, nil
}

func (uc *ProcessPaymentNotificationUC) activateSubscription(
	subscription *pkgsubscription.Subscription,
	payment *mercadopago.PaymentQueryResult,
	result *ProcessResult,
) (*ProcessResult, error) {
	duration := time.Duration(0)
	if plan, err := uc.PlanService.FindByID(subscription.PlanID); err == nil && plan != nil && plan.DurationDays > 0 {
		duration = time.Duration(plan.DurationDays) * 24 * time.Hour
	}

	subscription.MarkAsApproved(parsePaymentTime(payment.DateApproved), duration)
	subscription.PaymentMethod = payment.PaymentMethodID

	if err := uc.setPremium(subscription, true, subscription.ExpiresAt); err != nil {
		return nil, err
	}

	if err := uc.SubscriptionService.Update(subscription); err != nil {
		return nil, fmt.Errorf("erro ao aprovar assinatura %d: %w", subscription.ID, err)
	}

	result.Activated = true
	result.Description = fmt.Sprintf("pagamento aprovado para %q", subscription.UserType)
	return result, nil
}

// setPremium liga ou desliga a flag premium da entidade correspondente. A taxa
// de solicitação é avulsa e não altera nenhuma flag: o app libera o fluxo
// consultando o status do pagamento.
func (uc *ProcessPaymentNotificationUC) setPremium(
	subscription *pkgsubscription.Subscription,
	isPremium bool,
	expiresAt *time.Time,
) error {
	switch pkgplan.UserType(subscription.UserType) {
	case pkgplan.UserTypeProfessional:
		if err := uc.ProfessionalService.SetPremium(subscription.UserID, isPremium, expiresAt); err != nil {
			return fmt.Errorf("erro ao atualizar premium do profissional %d: %w", subscription.UserID, err)
		}
	case pkgplan.UserTypeStore:
		if err := uc.StoreService.SetPremium(subscription.UserID, isPremium, expiresAt); err != nil {
			return fmt.Errorf("erro ao atualizar premium da loja %d: %w", subscription.UserID, err)
		}
	}
	return nil
}

func (uc *ProcessPaymentNotificationUC) handleUnlockedBudget(
	unlocked *pkgunlockedbudget.UnlockedBudget,
	payment *mercadopago.PaymentQueryResult,
) (*ProcessResult, error) {
	status := normalizeStatus(payment.Status)

	result := &ProcessResult{
		PaymentID: payment.PaymentID,
		Status:    status,
		Handled:   true,
	}

	if status == pkgsubscription.PaymentStatusApproved {
		if unlocked.IsPaid() {
			result.Description = "desbloqueio já estava liberado"
			return result, nil
		}

		if math.Abs(payment.TransactionAmount-unlocked.Amount) > amountTolerance {
			unlocked.MarkAsFailed()
			if err := uc.UnlockedBudgetService.Update(unlocked); err != nil {
				return nil, fmt.Errorf("erro ao marcar desbloqueio %d como inválido: %w", unlocked.ID, err)
			}
			result.Status = pkgsubscription.PaymentStatusFailed
			result.Description = fmt.Sprintf(
				"valor divergente: cobrado %.2f, pago %.2f — desbloqueio não liberado",
				unlocked.Amount, payment.TransactionAmount,
			)
			return result, nil
		}

		unlocked.MarkAsPaid()
		if err := uc.UnlockedBudgetService.Update(unlocked); err != nil {
			return nil, fmt.Errorf("erro ao liberar desbloqueio %d: %w", unlocked.ID, err)
		}

		result.Activated = true
		result.Description = fmt.Sprintf("orçamento %d desbloqueado", unlocked.BudgetID)
		return result, nil
	}

	if status == pkgsubscription.PaymentStatusFailed || status == pkgsubscription.PaymentStatusCanceled {
		unlocked.MarkAsFailed()
		if err := uc.UnlockedBudgetService.Update(unlocked); err != nil {
			return nil, fmt.Errorf("erro ao atualizar desbloqueio %d: %w", unlocked.ID, err)
		}
	}

	result.Description = fmt.Sprintf("desbloqueio atualizado para %q", status)
	return result, nil
}

// rebuildFromExternalReference cria o registro que falta para pagamentos
// gerados antes de os checkouts passarem a gravar em subscriptions. O formato
// é "<tipo>:<id>:<timestamp>" — ver os casos de uso de checkout.
func (uc *ProcessPaymentNotificationUC) rebuildFromExternalReference(
	payment *mercadopago.PaymentQueryResult,
) (*pkgsubscription.Subscription, error) {
	parts := strings.Split(payment.ExternalReference, ":")
	if len(parts) < 2 {
		return nil, nil
	}

	userType := pkgplan.UserType(parts[0])
	switch userType {
	case pkgplan.UserTypeProfessional, pkgplan.UserTypeStore, pkgplan.UserTypeSolicitation:
	default:
		return nil, nil
	}

	targetID, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return nil, nil
	}

	plan, err := uc.PlanService.FindByUserType(userType)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar plano %q para o pagamento %d: %w", userType, payment.PaymentID, err)
	}

	subscription, err := uc.SubscriptionService.Save(pkgsubscription.Subscription{
		UserID:            uint(targetID),
		PlanID:            plan.ID,
		UserType:          string(userType),
		PaymentID:         payment.PaymentID,
		ExternalReference: payment.ExternalReference,
		Amount:            payment.TransactionAmount,
		Status:            pkgsubscription.PaymentStatusPending,
		IdempotencyKey:    fmt.Sprintf("recuperado-%d", payment.PaymentID),
		PaymentMethod:     payment.PaymentMethodID,
	})
	if err != nil {
		return nil, fmt.Errorf("erro ao recuperar assinatura do pagamento %d: %w", payment.PaymentID, err)
	}

	return subscription, nil
}

// normalizeStatus traduz o status do MercadoPago para o nosso enum.
func normalizeStatus(mpStatus string) pkgsubscription.PaymentStatus {
	switch mpStatus {
	case "approved":
		return pkgsubscription.PaymentStatusApproved
	case "pending", "in_process", "authorized", "in_mediation":
		return pkgsubscription.PaymentStatusPending
	case "cancelled", "canceled":
		return pkgsubscription.PaymentStatusCanceled
	case "refunded", "charged_back":
		return pkgsubscription.PaymentStatusRefunded
	default: // rejected e demais estados terminais
		return pkgsubscription.PaymentStatusFailed
	}
}

// parsePaymentTime lê a data de aprovação do MercadoPago, caindo para "agora"
// quando o campo vem vazio ou em formato inesperado.
func parsePaymentTime(value string) time.Time {
	if value == "" {
		return time.Now()
	}
	parsed, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return time.Now()
	}
	return parsed
}
