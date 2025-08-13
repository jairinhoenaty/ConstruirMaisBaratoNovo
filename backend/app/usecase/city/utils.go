package city_usecase

import pkgcity "construir_mais_barato/app/domain/city"

func GenerateCity(assembler *CityAssembler) pkgcity.City {
	city := pkgcity.City{}
	if assembler != nil {
		city.ID = assembler.ID
		city.Name = assembler.Name
		city.UF = assembler.UF
	}
	return city
}

func GenerateCityPresenter(city *pkgcity.City) CityPresenter {
	presenter := CityPresenter{}
	if city != nil {
		presenter.ID = city.ID
		presenter.Name = city.Name
		presenter.UF = city.UF
	}
	return presenter
}
