package region_usecase

import pkgregion "construir_mais_barato/app/domain/region"

type SaveRegionUC struct {
	Service   pkgregion.RegionService
	Assembler *RegionAssembler
}

type SaveRegionUCParams struct {
	Service pkgregion.RegionService
}

func NewSaveRegionUC(params SaveRegionUCParams) SaveRegionUC {
	return SaveRegionUC{
		Service: params.Service,
	}
}

func (uc *SaveRegionUC) Execute() (*RegionPresenter, error) {
	profission := GenerateRegion(uc.Assembler)
	profissionSaved, err := uc.Service.Save(profission)
	if err != nil {
		return nil, err
	}
	profissionPresenter := GenerateRegionPresenter(profissionSaved)

	return &profissionPresenter, nil

}
