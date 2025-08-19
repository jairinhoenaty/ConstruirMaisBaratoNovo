package banner_usecase

import (
	pkgbanner "construir_mais_barato/app/domain/banner"
	"encoding/json"
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
	
	assemblerJson, _ := json.Marshal(uc.Assembler)
	fmt.Println("uc.Assembler ===> ", string(assemblerJson))

	banners, err := uc.Service.FindByPage(uc.Assembler.Page,*uc.Assembler.CityId,*uc.Assembler.RegionId)

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
