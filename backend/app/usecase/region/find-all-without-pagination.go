package region_usecase

import pkgregion "construir_mais_barato/app/domain/region"

type FindAllWithoutPaginationRegionUC struct {
	Service pkgregion.RegionService
}

type FindAllWithoutPaginationRegionParams struct {
	Service pkgregion.RegionService
}

func NewFindAllWithoutPaginationRegionUC(params FindAllWithoutPaginationRegionParams) FindAllWithoutPaginationRegionUC {
	return FindAllWithoutPaginationRegionUC{
		Service: params.Service,
	}
}

func (uc *FindAllWithoutPaginationRegionUC) Execute() (*[]RegionPresenter, error) {

	regions, err := uc.Service.FindAllWithoutPagination()
	if err != nil {
		return nil, err
	}
	presenters := make([]RegionPresenter, 0)
	if len(regions) > 0 {
		for _, region := range regions {
			presenters = append(presenters, RegionPresenter{
				ID:          region.ID,
				Name:        region.Name,
				Description: region.Description,
				Icon:        region.Icon,
			})
		}
	}
	return &presenters, nil
}
