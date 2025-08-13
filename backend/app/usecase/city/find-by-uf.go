package city_usecase

import (
	pkgcity "construir_mais_barato/app/domain/city"
	"fmt"
)

type FindByUFUC struct {
	Service   pkgcity.CityService
	Assembler *UFCityAssembler
}

type FindByUFUCParams struct {
	Service pkgcity.CityService
}

func NewFindByUFUC(params FindByUFUCParams) FindByUFUC {
	return FindByUFUC{
		Service: params.Service,
	}
}

func (uc *FindByUFUC) Execute() ([]*CityPresenter, error) {
	if &uc.Assembler.UF == nil {
		return nil, fmt.Errorf("invalid UF")
	}

	cities, err := uc.Service.FindByUF(*&uc.Assembler.UF)
	if err != nil {
		return nil, err
	}

	citiesPresenter := make([]*CityPresenter, 0)
	for _, city := range cities {
		cityPresenter := GenerateCityPresenter(city)

		citiesPresenter = append(citiesPresenter, &cityPresenter)
	}

	return citiesPresenter, nil
}
