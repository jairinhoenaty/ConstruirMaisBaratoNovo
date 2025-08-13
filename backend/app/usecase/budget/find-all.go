package budget_usecase

import pkgbudget "construir_mais_barato/app/domain/budget"

type FindAllBudgetUC struct {
	Service   pkgbudget.BudgetService
	Assembler FindWithPaginationBudgetAssembler
}

type FindAllBudgetUCParams struct {
	Service pkgbudget.BudgetService
}

func NewFindAllBudgetUC(params FindAllBudgetUCParams) FindAllBudgetUC {
	return FindAllBudgetUC{
		Service: params.Service,
	}
}

func (uc *FindAllBudgetUC) Execute() (*[]BudgetPresenter, int64, error) {

	budgets, total, err := uc.Service.FindAll(uc.Assembler.Limit, uc.Assembler.Offset)
	if err != nil {
		return nil, 0, err
	}
	presenters := make([]BudgetPresenter, 0)
	if len(budgets) > 0 {
		for _, budget := range budgets {
			budgetPresenter := GenerateBudgetPresenter(budget)
			presenters = append(presenters, budgetPresenter)
		}
	}
	return &presenters, total, nil
}
