package unlocked_budget_usecase

import (
	pkgprofessional "construir_mais_barato/app/domain/professional"
	"construir_mais_barato/app/domain/store"
	pkgunlockedbudget "construir_mais_barato/app/domain/unlocked-budget"
	"errors"

	"gorm.io/gorm"
)

// CheckBudgetAccessUC verifica se um profissional ou loja tem acesso a um orçamento
type CheckBudgetAccessUC struct {
	UnlockedBudgetService pkgunlockedbudget.UnlockedBudgetService
	ProfessionalService   pkgprofessional.ProfessionalService
	StoreService          store.StoreService
}

type CheckBudgetAccessParams struct {
	UnlockedBudgetService pkgunlockedbudget.UnlockedBudgetService
	ProfessionalService   pkgprofessional.ProfessionalService
	StoreService          store.StoreService
}

type CheckBudgetAccessInput struct {
	Professional *pkgprofessional.Professional
	Store        *store.Store
	BudgetID     uint `json:"budgetId"`
}

type CheckBudgetAccessOutput struct {
	HasAccess bool   `json:"hasAccess"`
	IsPremium bool   `json:"isPremium"`
	IsPaid    bool   `json:"isPaid"`
	Reason    string `json:"reason,omitempty"`
	UserType  string `json:"userType,omitempty"`
}

func NewCheckBudgetAccessUC(params CheckBudgetAccessParams) *CheckBudgetAccessUC {
	return &CheckBudgetAccessUC{
		UnlockedBudgetService: params.UnlockedBudgetService,
		ProfessionalService:   params.ProfessionalService,
		StoreService:          params.StoreService,
	}
}

// Execute verifica se o profissional ou loja tem acesso ao orçamento
func (uc *CheckBudgetAccessUC) Execute(input CheckBudgetAccessInput) (*CheckBudgetAccessOutput, error) {
	// PROFESSIONAL LOGIC
	if input.Professional != nil && input.Professional.ID > 0 {
		// Check if professional is premium
		if input.Professional.IsPremium != nil && *input.Professional.IsPremium {
			return &CheckBudgetAccessOutput{
				HasAccess: true,
				IsPremium: true,
				IsPaid:    false,
				Reason:    "premium",
				UserType:  "professional",
			}, nil
		}

		// Check if professional has paid to unlock this specific budget
		unlockedBudget, err := uc.UnlockedBudgetService.FindPaidByProfessionalAndBudget(input.Professional.ID, input.BudgetID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		if unlockedBudget != nil && unlockedBudget.IsPaid() {
			return &CheckBudgetAccessOutput{
				HasAccess: true,
				IsPremium: false,
				IsPaid:    true,
				Reason:    "paid",
				UserType:  "professional",
			}, nil
		}

		// No access
		return &CheckBudgetAccessOutput{
			HasAccess: false,
			IsPremium: false,
			IsPaid:    false,
			Reason:    "blocked",
			UserType:  "professional",
		}, nil
	}

	// STORE LOGIC
	if input.Store != nil && input.Store.ID > 0 {
		// Check if store is premium
		if input.Store.IsPremiumStore != nil && *input.Store.IsPremiumStore {
			return &CheckBudgetAccessOutput{
				HasAccess: true,
				IsPremium: true,
				IsPaid:    false,
				Reason:    "premium",
				UserType:  "store",
			}, nil
		}

		// Check if store has paid to unlock this specific budget
		unlockedBudget, err := uc.UnlockedBudgetService.FindPaidByStoreAndBudget(input.Store.ID, input.BudgetID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		if unlockedBudget != nil && unlockedBudget.IsPaid() {
			return &CheckBudgetAccessOutput{
				HasAccess: true,
				IsPremium: false,
				IsPaid:    true,
				Reason:    "paid",
				UserType:  "store",
			}, nil
		}

		// No access
		return &CheckBudgetAccessOutput{
			HasAccess: false,
			IsPremium: false,
			IsPaid:    false,
			Reason:    "blocked",
			UserType:  "store",
		}, nil
	}

	return &CheckBudgetAccessOutput{
		HasAccess: false,
		IsPremium: false,
		IsPaid:    false,
		Reason:    "no_user",
	}, nil
}
