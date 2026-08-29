package payment_usecase

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	pkgplan "construir_mais_barato/app/domain/plan"
	pkgsubscription "construir_mais_barato/app/domain/subscription"
	"construir_mais_barato/infra/adapters/gateway-payment/mercadopago"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// pendingReuseWindow reaproveita o QR Code de um pagamento pendente recente em
// vez de criar outro. Evita cobranças duplicadas quando o usuário volta para a
// tela ou clica duas vezes. Mesma janela usada no desbloqueio de orçamento.
const pendingReuseWindow = 30 * time.Minute

// PlanCheckout concentra o checkout PIX de qualquer plano (premium do
// profissional, premium do lojista e taxa de solicitação do app), que antes
// eram três casos de uso praticamente idênticos.
//
// Além de criar o pagamento, grava a linha em subscriptions — é esse registro
// que permite ao webhook ligar um PIX pago de volta ao usuário.
type PlanCheckout struct {
	MercadoPagoClient   *mercadopago.MPClient
	PlanService         pkgplan.PlanService
	SubscriptionService pkgsubscription.SubscriptionService
}

type PlanCheckoutParams struct {
	MercadoPagoClient   *mercadopago.MPClient
	PlanService         pkgplan.PlanService
	SubscriptionService pkgsubscription.SubscriptionService
}

type PlanCheckoutInput struct {
	UserType pkgplan.UserType
	// TargetID é o id da entidade que recebe o benefício (professionals.id,
	// stores.id) ou o id do usuário, no caso da taxa de solicitação.
	TargetID uint
	// ReferenceID guarda contexto extra, como o id da solicitação no Firebase.
	ReferenceID string
	Payer       mercadopago.Payer
}

type PlanCheckoutOutput struct {
	// PaymentID não é exposto ao cliente: quem faz esse recorte é cada caso de
	// uso, porque o app já instalado ainda depende de recebê-lo.
	PaymentID    int64
	StatusToken  string  `json:"statusToken"`
	Amount       float64 `json:"amount"`
	QRCode       string  `json:"qr_code"`
	QRCodeBase64 string  `json:"qr_code_base64"`
	Status       string  `json:"status"`
}

func NewPlanCheckout(params PlanCheckoutParams) *PlanCheckout {
	return &PlanCheckout{
		MercadoPagoClient:   params.MercadoPagoClient,
		PlanService:         params.PlanService,
		SubscriptionService: params.SubscriptionService,
	}
}

func (c *PlanCheckout) Execute(input PlanCheckoutInput) (*PlanCheckoutOutput, error) {
	plan, err := c.PlanService.FindByUserType(input.UserType)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar o plano %q: %w", input.UserType, err)
	}

	if !plan.IsActive {
		return nil, fmt.Errorf("o plano %q não está ativo", input.UserType)
	}

	if reused := c.reusePending(plan, input); reused != nil {
		return reused, nil
	}

	externalRef := fmt.Sprintf("%s:%d:%d", input.UserType, input.TargetID, time.Now().Unix())
	idempotencyKey := uuid.NewString()
	statusToken := uuid.NewString()

	result, err := c.MercadoPagoClient.CreatePixPayment(mercadopago.PixPaymentInput{
		Amount:          plan.Price,
		Description:     plan.Name,
		ExternalRef:     externalRef,
		NotificationURL: notificationURL(),
		IdempotencyKey:  idempotencyKey,
		Payer:           input.Payer,
	})
	if err != nil {
		return nil, err
	}

	// A falha em gravar não invalida o PIX já criado: devolvemos o QR Code e
	// o webhook ainda consegue se virar pelo external_reference.
	if _, err := c.SubscriptionService.Save(pkgsubscription.Subscription{
		UserID:            input.TargetID,
		PlanID:            plan.ID,
		UserType:          string(input.UserType),
		PaymentID:         result.PaymentID,
		StatusToken:       statusToken,
		ExternalReference: externalRef,
		ReferenceID:       input.ReferenceID,
		Amount:            plan.Price,
		Status:            pkgsubscription.PaymentStatusPending,
		QRCode:            result.QRCode,
		QRCodeBase64:      result.QRCodeBase64,
		IdempotencyKey:    idempotencyKey,
		PaymentMethod:     "pix",
	}); err != nil {
		fmt.Printf("erro ao gravar a assinatura do pagamento %d: %v\n", result.PaymentID, err)
		// Sem registro não há token válido para consultar depois.
		statusToken = ""
	}

	return &PlanCheckoutOutput{
		PaymentID:    result.PaymentID,
		StatusToken:  statusToken,
		Amount:       plan.Price,
		QRCode:       result.QRCode,
		QRCodeBase64: result.QRCodeBase64,
		Status:       result.Status,
	}, nil
}

// reusePending devolve o QR Code de um pagamento pendente ainda dentro da
// janela de reaproveitamento, ou nil quando não houver.
func (c *PlanCheckout) reusePending(plan *pkgplan.Plan, input PlanCheckoutInput) *PlanCheckoutOutput {
	// Cobrança avulsa (taxa por solicitação) só pode ser reaproveitada dentro
	// da mesma solicitação. Sem essa checagem, quem gerou um PIX e não pagou
	// receberia o mesmo QR Code na solicitação seguinte e um único pagamento
	// liberaria as duas.
	avulso := plan.DurationDays == 0
	if avulso && input.ReferenceID == "" {
		return nil
	}

	pending, err := c.SubscriptionService.FindPendingByUserAndPlan(input.TargetID, plan.ID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("erro ao buscar pagamento pendente do usuário %d: %v\n", input.TargetID, err)
		}
		return nil
	}

	if pending == nil || pending.QRCode == "" {
		return nil
	}

	if avulso && pending.ReferenceID != input.ReferenceID {
		return nil
	}

	reference := pending.CreatedAt
	if pending.UpdatedAt.After(reference) {
		reference = pending.UpdatedAt
	}
	if time.Since(reference) >= pendingReuseWindow {
		return nil
	}

	return &PlanCheckoutOutput{
		PaymentID:    pending.PaymentID,
		StatusToken:  pending.StatusToken,
		Amount:       pending.Amount,
		QRCode:       pending.QRCode,
		QRCodeBase64: pending.QRCodeBase64,
		Status:       string(pending.Status),
	}
}

// notificationURL monta o endereço que o MercadoPago chama ao mudar o status do
// pagamento. Sem APP_PUBLIC_URL configurado o campo vai vazio e a confirmação
// passa a depender só do polling do endpoint de status.
func notificationURL() string {
	appURL := strings.TrimSuffix(os.Getenv("APP_PUBLIC_URL"), "/")
	if appURL == "" {
		return ""
	}
	return appURL + "/publica/webhooks/mercadopago"
}
