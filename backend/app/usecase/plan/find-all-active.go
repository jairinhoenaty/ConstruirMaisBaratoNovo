package plan_usecase

import (
	pkgplan "construir_mais_barato/app/domain/plan"
)

type FindAllActiveUC struct {
	Service pkgplan.PlanService
}

type FindAllActiveUCParams struct {
	Service pkgplan.PlanService
}

func NewFindAllActiveUC(params FindAllActiveUCParams) FindAllActiveUC {
	return FindAllActiveUC{
		Service: params.Service,
	}
}

func (uc *FindAllActiveUC) Execute() ([]*pkgplan.Plan, error) {
	plans, err := uc.Service.FindAllActive()
	if err != nil {
		return nil, err
	}
	return plans, nil
}
