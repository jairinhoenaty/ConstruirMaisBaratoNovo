package city_usecase

import pkgcity "construir_mais_barato/app/domain/city"

type SaveCityUC struct {
	Service   pkgcity.CityService
	Assembler *CityAssembler
}

type SaveCityUCParams struct {
	Service pkgcity.CityService
}

func NewSaveCityUC(params SaveCityUCParams) SaveCityUC {
	return SaveCityUC{
		Service: params.Service,
	}
}

func (uc *SaveCityUC) Execute() (*CityPresenter, error) {
	city := GenerateCity(uc.Assembler)
	citySaved, err := uc.Service.Save(city)
	if err != nil {
		return nil, err
	}
	cityPresenter := GenerateCityPresenter(citySaved)

	return &cityPresenter, nil

}
