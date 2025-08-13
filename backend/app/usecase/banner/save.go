package banner_usecase

import (
	pkgbanner "construir_mais_barato/app/domain/banner"
)

type SaveBannerUC struct {
	Service   pkgbanner.BannerService
	Assembler *BannerAssembler
}

type SaveBannerUCParams struct {
	Service pkgbanner.BannerService
}

func NewSaveBannerUC(params SaveBannerUCParams) SaveBannerUC {
	return SaveBannerUC{
		Service: params.Service,
	}
}

func (uc *SaveBannerUC) Execute() (*BannerPresenter, error) {

	banner := pkgbanner.Banner{
		CityID:        uc.Assembler.CityId,
		Link:          uc.Assembler.AccessLink,
		ProfessionIDs: uc.Assembler.Professions,
		Image:         uc.Assembler.Image,
		Page:		   uc.Assembler.Page,
		RegionID:      uc.Assembler.RegionId,
	}
	
	bannerSaved, err := uc.Service.Save(banner)
	if err != nil {
		return nil, err
	}

	professionsPresenter := make([]ProfessionPresenter, 0)
	if len(bannerSaved.Professions) > 0 {
		for _, profession := range bannerSaved.Professions {
			professionsPresenter = append(professionsPresenter, ProfessionPresenter{
				ID:   profession.ID,
				Name: profession.Name,
			})
		}
	}
	presenter := BannerPresenter{
		ID:         banner.ID,
		AccessLink: bannerSaved.Link,
		Image:      banner.Image,
		Cidade: CidadePresenter{
			ID:   banner.City.ID,
			Name: banner.City.Name,
			UF:   banner.City.UF,
		},
		Professions: &professionsPresenter,
		Page: banner.Page,
		Region: RegionPresenter{
			ID:   banner.Region.ID,
			Name: banner.Region.Name,
		},
	}

	return &presenter, nil

}
