package budget_usecase

import (
	pkgbudget "construir_mais_barato/app/domain/budget"
	"fmt"
)

type FindByEmailUC struct {
	Service   pkgbudget.BudgetService
	Assembler FindBudgetByEmailAssembler
}

type FindByEmailUCParams struct {
	Service pkgbudget.BudgetService
}

func NewFindByEmailUC(params FindByEmailUCParams) FindByEmailUC {
	return FindByEmailUC{
		Service: params.Service,
	}
}

func (uc *FindByEmailUC) Execute() (*BudgetPresenter, error) {
	if uc.Assembler.Email == "" {
		return nil, fmt.Errorf("invalid email")
	}

	budget, err := uc.Service.FindByEmail(uc.Assembler.Email)
	if err != nil {
		return nil, err
	}

	budgetPresenter := GenerateBudgetPresenter(budget)
	return &budgetPresenter, nil
}
