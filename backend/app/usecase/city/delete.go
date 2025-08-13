package city_usecase

import (
	pkgcity "construir_mais_barato/app/domain/city"
	"fmt"
)

type DeleteCityUC struct {
	Service pkgcity.CityService
	ID      *uint
}

type DeleteCityUCParams struct {
	Service pkgcity.CityService
}

func NewDeleteCityUC(params DeleteCityUCParams) DeleteCityUC {
	return DeleteCityUC{
		Service: params.Service,
	}
}

func (uc *DeleteCityUC) Execute() error {
	if uc.ID == nil {
		return fmt.Errorf("invalid id")
	}

	err := uc.Service.Remove(*uc.ID)
	if err != nil {
		return err
	}
	return nil
}
