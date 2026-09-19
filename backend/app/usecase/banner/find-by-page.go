package banner_usecase

import (
	pkgbanner "construir_mais_barato/app/domain/banner"
	"fmt"
)

type FindByPageUC struct {
	Service   pkgbanner.BannerService
	Assembler *FindByPageAssembler
}

type FindByPageUCParams struct {
	Service pkgbanner.BannerService
}

func NewFindByPageUC(params FindByPageUCParams) FindByPageUC {
	return FindByPageUC{
		Service: params.Service,
	}
}

func (uc FindByPageUC) Execute() ([]*BannerPresenter, error) {
	if uc.Assembler.Page == "" {
		return nil, fmt.Errorf("invalid data")
	}
	
	// assemblerJson, _ := json.Marshal(uc.Assembler)
	// fmt.Println("uc.Assembler ===> ", string(assemblerJson))

	// CityId e RegionId sao ponteiros opcionais: o app manda apenas a pagina.
	// Sem esta guarda um payload sem esses campos derrubava o handler com nil
	// pointer dereference. Zero e o valor que o repositorio trata como "sem
	// filtro", que e o mesmo que o painel web ja envia.
	var cityId, regionId uint
	if uc.Assembler.CityId != nil {
		cityId = *uc.Assembler.CityId
	}
	if uc.Assembler.RegionId != nil {
		regionId = *uc.Assembler.RegionId
	}

	banners, err := uc.Service.FindByPage(uc.Assembler.Page, cityId, regionId)

	if err != nil {
		return nil, err
	}

	presenters := make([]*BannerPresenter, 0)
	if len(banners) > 0 {
		for _, banner := range banners {
			presenter := GenerateBannerPresenter(*banner)
			presenters = append(presenters, &presenter)
		}
	}
	return presenters, nil

}
