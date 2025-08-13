package banner_usecase

import pkgbanner "construir_mais_barato/app/domain/banner"

func GenerateBannerPresenter(banner pkgbanner.Banner) BannerPresenter {
	return BannerPresenter{
		ID:         banner.ID,
		AccessLink: banner.Link,
		Image:      banner.Image,
		Cidade: CidadePresenter{
			ID:   banner.City.ID,
			Name: banner.City.Name,
			UF:   banner.City.UF,
		},
		ProfessionsIds: &banner.ProfessionIDs,
		Page: 			banner.Page,
		Region: RegionPresenter{
			ID:   banner.Region.ID,
			Name: banner.Region.Name,
		},
	}

}
