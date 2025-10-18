package unlocked_budget_usecase

import (
	pkgbudget "construir_mais_barato/app/domain/budget"
	pkgprofessional "construir_mais_barato/app/domain/professional"
	pkgunlockedbudget "construir_mais_barato/app/domain/unlocked-budget"
	"construir_mais_barato/infra/adapters/gateway-payment/mercadopago"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

const UNLOCK_BUDGET_PRICE = 30.00 // R$ 30,00 para desbloquear um orçamento

// UnlockBudgetPaymentUC cria um pagamento para desbloquear um orçamento
type UnlockBudgetPaymentUC struct {
	UnlockedBudgetService pkgunlockedbudget.UnlockedBudgetService
	BudgetService         pkgbudget.BudgetService
	ProfessionalService   pkgprofessional.ProfessionalService
	MercadoPagoClient     *mercadopago.MPClient
}

type UnlockBudgetPaymentParams struct {
	UnlockedBudgetService pkgunlockedbudget.UnlockedBudgetService
	BudgetService         pkgbudget.BudgetService
	ProfessionalService   pkgprofessional.ProfessionalService
	MercadoPagoClient     *mercadopago.MPClient
}

type UnlockBudgetPaymentInput struct {
	ProfessionalID uint              `json:"professionalId"`
	BudgetID       uint              `json:"budgetId"`
	Payer          mercadopago.Payer `json:"payer"`
}

type UnlockBudgetPaymentOutput struct {
	UnlockedBudgetID uint    `json:"unlockedBudgetId"`
	PaymentID        string  `json:"paymentId"`
	Amount           float64 `json:"amount"`
	QRCode           string  `json:"qrCode"`
	QRCodeBase64     string  `json:"qrCodeBase64"`
	Status           string  `json:"status"`
}

func NewUnlockBudgetPaymentUC(params UnlockBudgetPaymentParams) *UnlockBudgetPaymentUC {
	return &UnlockBudgetPaymentUC{
		UnlockedBudgetService: params.UnlockedBudgetService,
		BudgetService:         params.BudgetService,
		ProfessionalService:   params.ProfessionalService,
		MercadoPagoClient:     params.MercadoPagoClient,
	}
}

func (uc *UnlockBudgetPaymentUC) Execute(input UnlockBudgetPaymentInput) (*UnlockBudgetPaymentOutput, error) {
	var unlockedBudget pkgunlockedbudget.UnlockedBudget

	professional, err := uc.ProfessionalService.FindById(input.ProfessionalID)
	if err != nil {
		return nil, errors.New("profissional não encontrado")
	}

	if professional.IsPremium != nil && *professional.IsPremium {
		return nil, errors.New("profissionais premium não precisam pagar para desbloquear orçamentos")
	}

	_, err = uc.BudgetService.FindById(input.BudgetID)
	if err != nil {
		return nil, errors.New("orçamento não encontrado")
	}

	existingUnlock, err := uc.UnlockedBudgetService.FindPaidByProfessionalAndBudget(input.ProfessionalID, input.BudgetID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existingUnlock != nil && existingUnlock.IsPaid() {
		return nil, errors.New("você já desbloqueou este orçamento")
	}

	existingPending, err := uc.UnlockedBudgetService.FindByProfessionalAndBudget(input.ProfessionalID, input.BudgetID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Transfomar em função!
	if existingPending != nil && existingPending.IsPending() {
		timeSinceCreation := time.Since(existingPending.CreatedAt)
		if existingPending.UpdatedAt.After(existingPending.CreatedAt) {
			timeSinceCreation = time.Since(existingPending.UpdatedAt)
		}
		if timeSinceCreation < 30*time.Minute {
			return &UnlockBudgetPaymentOutput{
				UnlockedBudgetID: existingPending.ID,
				PaymentID:        existingPending.PaymentID,
				Amount:           existingPending.Amount,
				QRCode:           existingPending.QRCode,
				QRCodeBase64:     existingPending.QRCodeBase64,
				Status:           existingPending.Status,
			}, nil
		}
	}

	externalRef := fmt.Sprintf("unlock:budget:%d:professional:%d:%d", input.BudgetID, input.ProfessionalID, time.Now().Unix())
	idempotencyKey := fmt.Sprintf("unlock-%d-%d-%d", input.ProfessionalID, input.BudgetID, time.Now().Unix())

	pixResult, err := uc.MercadoPagoClient.CreatePixPayment(mercadopago.PixPaymentInput{
		Amount:         UNLOCK_BUDGET_PRICE,
		Description:    fmt.Sprintf("Desbloqueio de orçamento #%d", input.BudgetID),
		ExternalRef:    externalRef,
		IdempotencyKey: idempotencyKey,
		Payer:          input.Payer,
	})
	if err != nil {
		return nil, fmt.Errorf("erro ao criar pagamento: %v", err)
	}

	//Transformar em função!
	if existingPending != nil {
		unlockedBudget = *existingPending
		unlockedBudget.Status = "pending"
		unlockedBudget.PaymentID = fmt.Sprintf("%d", pixResult.PaymentID)
		unlockedBudget.Amount = UNLOCK_BUDGET_PRICE
		unlockedBudget.QRCode = pixResult.QRCode
		unlockedBudget.QRCodeBase64 = pixResult.QRCodeBase64
	} else {
		unlockedBudget = pkgunlockedbudget.UnlockedBudget{
			ProfessionalID: input.ProfessionalID,
			BudgetID:       input.BudgetID,
			Status:         "pending",
			PaymentID:      fmt.Sprintf("%d", pixResult.PaymentID),
			Amount:         UNLOCK_BUDGET_PRICE,
			QRCode:         pixResult.QRCode,
			QRCodeBase64:   pixResult.QRCodeBase64,
		}
	}

	savedUnlock, err := uc.UnlockedBudgetService.Save(unlockedBudget)
	if err != nil {
		return nil, fmt.Errorf("erro ao salvar desbloqueio: %v", err)
	}

	return &UnlockBudgetPaymentOutput{
		UnlockedBudgetID: savedUnlock.ID,
		PaymentID:        savedUnlock.PaymentID,
		Amount:           savedUnlock.Amount,
		QRCode:           savedUnlock.QRCode,
		QRCodeBase64:     savedUnlock.QRCodeBase64,
		Status:           savedUnlock.Status,
	}, nil
}
