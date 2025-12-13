package plan_usecase

import (
	pkgplan "construir_mais_barato/app/domain/plan"
	"fmt"
)

type FindByUserTypeUC struct {
	Service  pkgplan.PlanService
	UserType *pkgplan.UserType
}

type FindByUserTypeUCParams struct {
	Service pkgplan.PlanService
}

func NewFindByUserTypeUC(params FindByUserTypeUCParams) FindByUserTypeUC {
	return FindByUserTypeUC{
		Service: params.Service,
	}
}

func (uc *FindByUserTypeUC) Execute() (*pkgplan.Plan, error) {
	if uc.UserType == nil {
		return nil, fmt.Errorf("user type is required")
	}

	plan, err := uc.Service.FindByUserType(*uc.UserType)
	if err != nil {
		return nil, fmt.Errorf("plan not found for user type: %s", *uc.UserType)
	}

	return plan, nil
}
