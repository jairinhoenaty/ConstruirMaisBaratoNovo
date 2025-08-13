package city_usecase

import pkgcity "construir_mais_barato/app/domain/city"

type FindAllCityUC struct {
	Service pkgcity.CityService
}

type FindAllCityUCParams struct {
	Service pkgcity.CityService
}

func NewFindAllCityUC(params FindAllCityUCParams) FindAllCityUC {
	return FindAllCityUC{
		Service: params.Service,
	}
}

func (uc *FindAllCityUC) Execute() (*[]CityPresenter, error) {

	city, err := uc.Service.FindAll()
	if err != nil {
		return nil, err
	}
	presenters := make([]CityPresenter, 0)
	if len(city) > 0 {
		for _, City := range city {
			presenters = append(presenters, CityPresenter{
				ID:   City.ID,
				Name: City.Name,
				UF:   City.UF,
			})
		}
	}
	return &presenters, nil
}
