package banner_usecase

import (
	pkgbanner "construir_mais_barato/app/domain/banner"
	"fmt"
)

type FindByCityAndProfessionUC struct {
	Service   pkgbanner.BannerService
	Assembler *FindByCityIdAndProfessionIDAssembler
}

type FindByCityAndProfessionUCParams struct {
	Service pkgbanner.BannerService
}

func NewFindByCityAndProfessionUC(params FindByCityAndProfessionUCParams) FindByCityAndProfessionUC {
	return FindByCityAndProfessionUC{
		Service: params.Service,
	}
}

func (uc FindByCityAndProfessionUC) Execute() ([]*BannerPresenter, error) {
	if uc.Assembler == nil {
		return nil, fmt.Errorf("invalid data")
	}

	banners, err := uc.Service.FindByCityAndProfession(uc.Assembler.CityId, uc.Assembler.ProfessionId)
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
