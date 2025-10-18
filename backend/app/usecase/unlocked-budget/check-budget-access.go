package unlocked_budget_usecase

import (
	pkgprofessional "construir_mais_barato/app/domain/professional"
	pkgunlockedbudget "construir_mais_barato/app/domain/unlocked-budget"
	"errors"

	"gorm.io/gorm"
)

// CheckBudgetAccessUC verifica se um profissional tem acesso a um orçamento
type CheckBudgetAccessUC struct {
	UnlockedBudgetService pkgunlockedbudget.UnlockedBudgetService
	ProfessionalService   pkgprofessional.ProfessionalService
}

type CheckBudgetAccessParams struct {
	UnlockedBudgetService pkgunlockedbudget.UnlockedBudgetService
	ProfessionalService   pkgprofessional.ProfessionalService
}

type CheckBudgetAccessInput struct {
	ProfessionalID uint `json:"professionalId"`
	BudgetID       uint `json:"budgetId"`
}

type CheckBudgetAccessOutput struct {
	HasAccess bool   `json:"hasAccess"`
	IsPremium bool   `json:"isPremium"`
	IsPaid    bool   `json:"isPaid"`
	Reason    string `json:"reason,omitempty"`
}

func NewCheckBudgetAccessUC(params CheckBudgetAccessParams) *CheckBudgetAccessUC {
	return &CheckBudgetAccessUC{
		UnlockedBudgetService: params.UnlockedBudgetService,
		ProfessionalService:   params.ProfessionalService,
	}
}

// Execute verifica se o profissional tem acesso ao orçamento
func (uc *CheckBudgetAccessUC) Execute(input CheckBudgetAccessInput) (*CheckBudgetAccessOutput, error) {
	professional, err := uc.ProfessionalService.FindById(input.ProfessionalID)
	if err != nil {
		return nil, errors.New("profissional não encontrado")
	}

	if professional.IsPremium != nil && *professional.IsPremium {
		return &CheckBudgetAccessOutput{
			HasAccess: true,
			IsPremium: true,
			IsPaid:    false,
			Reason:    "premium",
		}, nil
	}

	unlockedBudget, err := uc.UnlockedBudgetService.FindPaidByProfessionalAndBudget(input.ProfessionalID, input.BudgetID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if unlockedBudget != nil && unlockedBudget.IsPaid() {
		return &CheckBudgetAccessOutput{
			HasAccess: true,
			IsPremium: false,
			IsPaid:    true,
			Reason:    "paid",
		}, nil
	}

	return &CheckBudgetAccessOutput{
		HasAccess: false,
		IsPremium: false,
		IsPaid:    false,
		Reason:    "blocked",
	}, nil
}
