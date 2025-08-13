package budget_usecase

import (
	pkgbudget "construir_mais_barato/app/domain/budget"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type CheckExpiredBudgetUC struct {
	Service pkgbudget.BudgetService
}

type CheckExpiredBudgetUCParams struct {
	Service pkgbudget.BudgetService
}

func NewCheckExpiredBudgetUC(params CheckExpiredBudgetUCParams) CheckExpiredBudgetUC {
	return CheckExpiredBudgetUC{
		Service: params.Service,
	}
}

func (uc *CheckExpiredBudgetUC) Execute() error {

	now := time.Now()
	// buscar orçamentos com 14 dias vencidos
	expiredBudgets, err := uc.Service.FindExpiredBudgets(now.AddDate(0, 0, -15))
	if err != nil {
		fmt.Println("Error finding expired budgets:", err)
		return err
	}

	for _, budget := range expiredBudgets {
		budget.DeletedAt = gorm.DeletedAt{Time: now}

		// chamar o caso de uso para salvar o oeçamento com a desativação
		professionalIDs := make([]uint, 0)
		if len(*budget.ProfessionalIDs) > 0 {
			for _, id := range *budget.ProfessionalIDs {
				professionalIDs = append(professionalIDs, id)

			}
		}
		saveAssembler := BudgetAssembler{
			ID:              budget.ID,
			Name:            budget.Name,
			Email:           budget.Email,
			Telephone:       budget.Telephone,
			Description:     budget.Description,
			ProfessionalsId: &professionalIDs,
		}
		saveParams := SaveBudgetUCParams{
			Service: uc.Service,
		}
		saveUC := NewSaveBudgetUC(saveParams)
		saveUC.Assembler = &saveAssembler
		_, err := saveUC.Execute()
		if err != nil {
			fmt.Println("Erro ao desativar o orçamento vencido!. ID => ", budget.ID)
			return err
		}
	}
	return nil
}
