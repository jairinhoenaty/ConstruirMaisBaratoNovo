package payment_usecase

import (
	"errors"
	"fmt"

	pkgsubscription "construir_mais_barato/app/domain/subscription"

	"gorm.io/gorm"
)

// RegisterBypassUC registra que o usuário seguiu o fluxo do app sem que o
// pagamento tivesse sido confirmado.
//
// Serve só para auditoria: não libera premium nem altera qualquer flag. É o que
// permite medir com que frequência a saída de emergência da tela de pagamento
// é usada e decidir depois se ela deve continuar existindo.
type RegisterBypassUC struct {
	SubscriptionService pkgsubscription.SubscriptionService
	Assembler           BypassAssembler
}

// BypassAssembler identifica o pagamento pelo token opaco do checkout. Usar o
// id do MercadoPago aqui permitiria que qualquer pessoa marcasse assinaturas
// alheias como bypassed, corrompendo justamente a auditoria que este registro
// existe para produzir.
type BypassAssembler struct {
	StatusToken    string `json:"statusToken"`
	UserID         uint   `json:"userId"`
	SolicitationID string `json:"solicitationId"`
	ProfessionalID uint   `json:"professionalId"`
	Reason         string `json:"reason"`
}

type RegisterBypassUCParams struct {
	SubscriptionService pkgsubscription.SubscriptionService
}

type RegisterBypassOutput struct {
	Registered bool   `json:"registered"`
	Status     string `json:"status"`
}

func NewRegisterBypassUC(params RegisterBypassUCParams) *RegisterBypassUC {
	return &RegisterBypassUC{
		SubscriptionService: params.SubscriptionService,
	}
}

func (uc *RegisterBypassUC) Execute() (*RegisterBypassOutput, error) {
	if uc.Assembler.StatusToken == "" {
		return nil, ErrPaymentNotFound
	}

	subscription, err := uc.SubscriptionService.FindByStatusToken(uc.Assembler.StatusToken)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("erro ao buscar assinatura pelo token: %w", err)
	}

	// Um pagamento já aprovado não é bypass: o cliente só demorou a receber a
	// confirmação. Não sobrescrevemos o registro.
	if subscription != nil && subscription.IsApproved() {
		return &RegisterBypassOutput{Registered: false, Status: string(subscription.Status)}, nil
	}

	if subscription != nil {
		subscription.MarkAs(pkgsubscription.PaymentStatusBypassed)
		if uc.Assembler.SolicitationID != "" {
			subscription.ReferenceID = uc.Assembler.SolicitationID
		}
		if err := uc.SubscriptionService.Update(subscription); err != nil {
			return nil, fmt.Errorf("erro ao registrar bypass da assinatura %d: %w", subscription.ID, err)
		}
		return &RegisterBypassOutput{Registered: true, Status: string(pkgsubscription.PaymentStatusBypassed)}, nil
	}

	// Token desconhecido: não criamos registro a partir de dado arbitrário,
	// senão o endpoint viraria uma porta para poluir a auditoria com
	// assinaturas inventadas.
	return nil, ErrPaymentNotFound
}
