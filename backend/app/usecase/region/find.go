package region_usecase

import pkgregion "construir_mais_barato/app/domain/region"

type FindRegionUC struct {
	Service             pkgregion.RegionService
	QuantityRegions uint
}

type FindRegionUCParams struct {
	Service pkgregion.RegionService
}

func NewFindRegionUC(params FindRegionUCParams) FindRegionUC {
	return FindRegionUC{
		Service: params.Service,
	}
}

func (uc *FindRegionUC) Execute() (*[]RegionPresenter, error) {

	regions, err := uc.Service.Find(uc.QuantityRegions)
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
