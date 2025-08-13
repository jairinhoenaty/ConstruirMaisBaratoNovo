package region_usecase

import pkgregion "construir_mais_barato/app/domain/region"

func GenerateRegion(assembler *RegionAssembler) pkgregion.Region {
	region := pkgregion.Region{}
	if assembler != nil {
		region.Name = assembler.Name
		region.Description = assembler.Description
		region.Icon = assembler.Icon
		region.CityIDs = assembler.CityIDs
		region.UF = assembler.UF
		
	}
	if assembler.ID > 0 {
		region.ID = assembler.ID
	}

	return region
}

func GenerateRegionPresenter(region *pkgregion.Region) RegionPresenter {
	presenter := RegionPresenter{}
	if region != nil {
		cidadePresenter := make([]CidadePresenter, 0)
		if len(region.Cities) > 0 {
			for _, city := range region.Cities {
				cidadePresenter = append(cidadePresenter, CidadePresenter{
					ID:          city.ID,
					Name:        city.Name,
					UF: city.UF,
				})

			}
		}

		presenter.ID = region.ID
		presenter.Name = region.Name
		presenter.Description = region.Description
		presenter.Icon = region.Icon
		presenter.Cities = cidadePresenter
		presenter.UF = region.UF
	}
	return presenter
}
