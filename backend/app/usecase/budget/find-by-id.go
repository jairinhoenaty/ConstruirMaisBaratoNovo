package budget_usecase

import (
	pkgbudget "construir_mais_barato/app/domain/budget"
	"fmt"
)

type FindByIdUC struct {
	Service pkgbudget.BudgetService
	ID      *uint
}

type FindByIdUCParams struct {
	Service pkgbudget.BudgetService
}

func NewFindByIdUC(params FindByIdUCParams) FindByIdUC {
	return FindByIdUC{
		Service: params.Service,
	}
}

func (uc *FindByIdUC) Execute() (*BudgetPresenter, error) {
	if uc.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	budget, err := uc.Service.FindById(*uc.ID)
	if err != nil {
		return nil, err
	}

	budgetPresenter := GenerateBudgetPresenter(budget)
	return &budgetPresenter, nil
}
