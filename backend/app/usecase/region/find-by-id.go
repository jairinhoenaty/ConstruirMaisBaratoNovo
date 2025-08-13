package region_usecase

import (
	pkgregion "construir_mais_barato/app/domain/region"
	"fmt"
)

type FindByIdUC struct {
	Service pkgregion.RegionService
	ID      *uint
}

type FindByIdUCParams struct {
	Service pkgregion.RegionService
}

func NewFindByIdUC(params FindByIdUCParams) FindByIdUC {
	return FindByIdUC{
		Service: params.Service,
	}
}

func (uc *FindByIdUC) Execute() (*RegionPresenter, error) {
	if uc.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	profission, err := uc.Service.FindById(*uc.ID)
	if err != nil {
		return nil, err
	}

	profissionPresenter := GenerateRegionPresenter(profission)
	return &profissionPresenter, nil
}
