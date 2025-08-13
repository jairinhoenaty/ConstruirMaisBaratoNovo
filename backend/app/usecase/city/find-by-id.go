package city_usecase

import (
	pkgcity "construir_mais_barato/app/domain/city"
	"fmt"
)

type FindByIdUC struct {
	Service pkgcity.CityService
	ID      *uint
}

type FindByIdUCParams struct {
	Service pkgcity.CityService
}

func NewFindByIdUC(params FindByIdUCParams) FindByIdUC {
	return FindByIdUC{
		Service: params.Service,
	}
}

func (uc *FindByIdUC) Execute() (*CityPresenter, error) {
	if uc.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	city, err := uc.Service.FindById(*uc.ID)
	if err != nil {
		return nil, err
	}

	cityPresenter := GenerateCityPresenter(city)
	return &cityPresenter, nil
}
