package region_usecase

import (
	pkgregion "construir_mais_barato/app/domain/region"
	"fmt"
)

type FindByCityUC struct {
	Service pkgregion.RegionService
	ID  *uint
}

type FindByCityUCParams struct {
	Service pkgregion.RegionService
}

func NewFindByCityUC(params FindByCityUCParams) FindByCityUC {
	return FindByCityUC{
		Service: params.Service,
	}
}

func (uc *FindByCityUC) Execute() (*RegionPresenter, error) {

	if uc.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	region, err := uc.Service.FindByCity(*uc.ID)
	if err != nil {
		return nil, err
	}

	regionPresenter := GenerateRegionPresenter(region)
	return &regionPresenter, nil

}
