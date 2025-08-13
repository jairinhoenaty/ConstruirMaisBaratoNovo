package budget_usecase

import (
	pkgbudget "construir_mais_barato/app/domain/budget"
	"fmt"
)

type DeleteBudgetUC struct {
	Service pkgbudget.BudgetService
	ID      *uint
}

type DeleteBudgetUCParams struct {
	Service pkgbudget.BudgetService
}

func NewDeleteBudgetUC(params DeleteBudgetUCParams) DeleteBudgetUC {
	return DeleteBudgetUC{
		Service: params.Service,
	}
}

func (uc *DeleteBudgetUC) Execute() error {
	if uc.ID == nil {
		return fmt.Errorf("invalid id")
	}

	err := uc.Service.Remove(*uc.ID)
	if err != nil {
		return err
	}
	return nil
}
