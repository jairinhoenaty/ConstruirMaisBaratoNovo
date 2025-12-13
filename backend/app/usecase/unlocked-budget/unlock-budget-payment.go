package unlocked_budget_usecase

import (
	pkgbudget "construir_mais_barato/app/domain/budget"
	pkgprofessional "construir_mais_barato/app/domain/professional"
	"construir_mais_barato/app/domain/store"
	pkgunlockedbudget "construir_mais_barato/app/domain/unlocked-budget"
	pkgunlockprice "construir_mais_barato/app/domain/unlock-price"
	"construir_mais_barato/infra/adapters/gateway-payment/mercadopago"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// UnlockBudgetPaymentUC cria um pagamento para desbloquear um orçamento
type UnlockBudgetPaymentUC struct {
	UnlockedBudgetService pkgunlockedbudget.UnlockedBudgetService
	BudgetService         pkgbudget.BudgetService
	ProfessionalService   pkgprofessional.ProfessionalService
	StoreService          store.StoreService
	UnlockPriceService    pkgunlockprice.UnlockPriceService
	MercadoPagoClient     *mercadopago.MPClient
}

type UnlockBudgetPaymentParams struct {
	UnlockedBudgetService pkgunlockedbudget.UnlockedBudgetService
	BudgetService         pkgbudget.BudgetService
	ProfessionalService   pkgprofessional.ProfessionalService
	StoreService          store.StoreService
	UnlockPriceService    pkgunlockprice.UnlockPriceService
	MercadoPagoClient     *mercadopago.MPClient
}

type UnlockBudgetPaymentInput struct {
	ProfessionalID *uint             `json:"professionalId,omitempty"`
	StoreID        *uint             `json:"storeId,omitempty"`
	UserType       string            `json:"userType"` // "professional" or "store"
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
		StoreService:          params.StoreService,
		UnlockPriceService:    params.UnlockPriceService,
		MercadoPagoClient:     params.MercadoPagoClient,
	}
}

func (uc *UnlockBudgetPaymentUC) Execute(input UnlockBudgetPaymentInput) (*UnlockBudgetPaymentOutput, error) {
	var unlockedBudget pkgunlockedbudget.UnlockedBudget
	var unlockPrice float64
	var userID uint
	var userType pkgunlockedbudget.UserType
	var externalRef string

	// Validate user type
	if input.UserType != "professional" && input.UserType != "store" {
		return nil, errors.New("tipo de usuário inválido")
	}

	// PROFESSIONAL FLOW
	if input.UserType == "professional" {
		if input.ProfessionalID == nil {
			return nil, errors.New("ID do profissional é obrigatório")
		}

		professional, err := uc.ProfessionalService.FindById(*input.ProfessionalID)
		if err != nil {
			return nil, errors.New("profissional não encontrado")
		}

		if professional.IsPremium != nil && *professional.IsPremium {
			return nil, errors.New("profissionais premium não precisam pagar para desbloquear orçamentos")
		}

		// Check if already unlocked
		existingUnlock, err := uc.UnlockedBudgetService.FindPaidByProfessionalAndBudget(*input.ProfessionalID, input.BudgetID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		if existingUnlock != nil && existingUnlock.IsPaid() {
			return nil, errors.New("você já desbloqueou este orçamento")
		}

		// Get unlock price from database
		priceConfig, err := uc.UnlockPriceService.FindActiveByUserType(pkgunlockprice.UserTypeProfessional)
		if err != nil {
			return nil, fmt.Errorf("erro ao buscar preço de desbloqueio: %w", err)
		}
		unlockPrice = priceConfig.Price
		userID = *input.ProfessionalID
		userType = pkgunlockedbudget.UserTypeProfessional
		externalRef = fmt.Sprintf("unlock:budget:%d:professional:%d:%d", input.BudgetID, userID, time.Now().Unix())
	}

	// STORE FLOW
	if input.UserType == "store" {
		if input.StoreID == nil {
			return nil, errors.New("ID da loja é obrigatório")
		}

		storeObj, err := uc.StoreService.FindById(*input.StoreID)
		if err != nil {
			return nil, errors.New("loja não encontrada")
		}

		if storeObj.IsPremiumStore != nil && *storeObj.IsPremiumStore {
			return nil, errors.New("lojas premium não precisam pagar para desbloquear orçamentos")
		}

		// Check if already unlocked
		existingUnlock, err := uc.UnlockedBudgetService.FindPaidByStoreAndBudget(*input.StoreID, input.BudgetID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		if existingUnlock != nil && existingUnlock.IsPaid() {
			return nil, errors.New("você já desbloqueou este orçamento")
		}

		// Get unlock price from database
		priceConfig, err := uc.UnlockPriceService.FindActiveByUserType(pkgunlockprice.UserTypeStore)
		if err != nil {
			return nil, fmt.Errorf("erro ao buscar preço de desbloqueio: %w", err)
		}
		unlockPrice = priceConfig.Price
		userID = *input.StoreID
		userType = pkgunlockedbudget.UserTypeStore
		externalRef = fmt.Sprintf("unlock:budget:%d:store:%d:%d", input.BudgetID, userID, time.Now().Unix())
	}

	// Verify budget exists
	_, err := uc.BudgetService.FindById(input.BudgetID)
	if err != nil {
		return nil, errors.New("orçamento não encontrado")
	}

	// Check for existing pending payment (reuse QR code if within 30 minutes)
	var existingPending *pkgunlockedbudget.UnlockedBudget
	if input.UserType == "professional" {
		existingPending, _ = uc.UnlockedBudgetService.FindByProfessionalAndBudget(*input.ProfessionalID, input.BudgetID)
	} else {
		existingPending, _ = uc.UnlockedBudgetService.FindByStoreAndBudget(*input.StoreID, input.BudgetID)
	}

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

	// Create PIX payment
	idempotencyKey := fmt.Sprintf("unlock-%s-%d-%d-%d", input.UserType, userID, input.BudgetID, time.Now().Unix())

	pixResult, err := uc.MercadoPagoClient.CreatePixPayment(mercadopago.PixPaymentInput{
		Amount:         unlockPrice,
		Description:    fmt.Sprintf("Desbloqueio de orçamento #%d", input.BudgetID),
		ExternalRef:    externalRef,
		IdempotencyKey: idempotencyKey,
		Payer:          input.Payer,
	})
	if err != nil {
		return nil, fmt.Errorf("erro ao criar pagamento: %v", err)
	}

	// Save or update unlocked budget record
	if existingPending != nil {
		unlockedBudget = *existingPending
		unlockedBudget.Status = "pending"
		unlockedBudget.PaymentID = fmt.Sprintf("%d", pixResult.PaymentID)
		unlockedBudget.Amount = unlockPrice
		unlockedBudget.QRCode = pixResult.QRCode
		unlockedBudget.QRCodeBase64 = pixResult.QRCodeBase64
	} else {
		unlockedBudget = pkgunlockedbudget.UnlockedBudget{
			UserType:       userType,
			ProfessionalID: input.ProfessionalID,
			StoreID:        input.StoreID,
			BudgetID:       input.BudgetID,
			Status:         "pending",
			PaymentID:      fmt.Sprintf("%d", pixResult.PaymentID),
			Amount:         unlockPrice,
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
