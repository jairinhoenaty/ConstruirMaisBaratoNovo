package budget_usecase

import (
	pkgbudget "construir_mais_barato/app/domain/budget"
	"fmt"
)

type RefuseBudgetUC struct {
	Service       pkgbudget.BudgetService
	ID            *uint
	RecipientID   uint
	RecipientType string
}

type RefuseBudgetUCParams struct {
	Service pkgbudget.BudgetService
}

func NewRefuseBudgetUC(params RefuseBudgetUCParams) RefuseBudgetUC {
	return RefuseBudgetUC{
		Service: params.Service,
	}
}

func (uc *RefuseBudgetUC) Execute() error {
	if uc.ID == nil {
		return fmt.Errorf("invalid id")
	}
	if uc.RecipientID == 0 {
		return fmt.Errorf("invalid recipient id")
	}
	recipientType := pkgbudget.RecipientType(uc.RecipientType)
	if recipientType != pkgbudget.RecipientTypeProfessional &&
		recipientType != pkgbudget.RecipientTypeStore {
		return fmt.Errorf("invalid recipient type")
	}

	return uc.Service.Refuse(*uc.ID, uc.RecipientID, recipientType)
}
