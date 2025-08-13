package region_usecase

import (
	pkgregion "construir_mais_barato/app/domain/region"
)

type FindRegionsWithCountIdUC struct {
	Service pkgregion.RegionService
}

type FindRegionsWithCountIdUCParams struct {
	Service pkgregion.RegionService
}

func NewFindRegionsWithCountIdUC(params FindRegionsWithCountIdUCParams) FindRegionsWithCountIdUC {
	return FindRegionsWithCountIdUC{
		Service: params.Service,
	}
}

func (uc *FindRegionsWithCountIdUC) Execute() (*[]RegionWithCountPresenter, error) {

	result, err := uc.Service.FindRegionsWithCount()
	if err != nil {
		return nil, err
	}

	presenters := make([]RegionWithCountPresenter, 0)
	for _, res := range result {

		presenter := RegionWithCountPresenter{
			Name:  res["region"].(string),
			Count: int(res["count"].(int64)),
		}

		presenters = append(presenters, presenter)
	}

	return &presenters, nil
}
