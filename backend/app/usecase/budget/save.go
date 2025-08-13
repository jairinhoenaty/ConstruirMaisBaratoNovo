package budget_usecase

import pkgbudget "construir_mais_barato/app/domain/budget"

type SaveBudgetUC struct {
	Service   pkgbudget.BudgetService
	Assembler *BudgetAssembler
}

type SaveBudgetUCParams struct {
	Service pkgbudget.BudgetService
}

func NewSaveBudgetUC(params SaveBudgetUCParams) SaveBudgetUC {
	return SaveBudgetUC{
		Service: params.Service,
	}
}

func (uc *SaveBudgetUC) Execute() (*BudgetPresenter, error) {

	budget := GenerateBudget(uc.Assembler)
	budgetSaved, err := uc.Service.Save(budget)
	if err != nil {
		return nil, err
	}
	budgetPresenter := GenerateBudgetPresenter(budgetSaved)

	return &budgetPresenter, nil

}
