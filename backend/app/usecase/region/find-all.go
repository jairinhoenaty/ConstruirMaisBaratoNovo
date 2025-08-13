package region_usecase

import pkgregion "construir_mais_barato/app/domain/region"

type FindAllRegionUC struct {
	Service   pkgregion.RegionService
	Assembler FindWithPaginationRegionAssembler
}

type FindAllRegionUCParams struct {
	Service pkgregion.RegionService
}

func NewFindAllRegionUC(params FindAllRegionUCParams) FindAllRegionUC {
	return FindAllRegionUC{
		Service: params.Service,
	}
}

func (uc *FindAllRegionUC) Execute() (*[]RegionPresenter, int64, error) {

	regions, total, err := uc.Service.FindAll(uc.Assembler.Limit, uc.Assembler.Offset,uc.Assembler.UF)
	if err != nil {
		return nil, 0, err
	}
	presenters := make([]RegionPresenter, 0)
	if len(regions) > 0 {
		for _, region := range regions {
			regionPresenter := GenerateRegionPresenter(region)
			presenters = append(presenters, regionPresenter)
		}
	}

	return &presenters, total, nil
}
