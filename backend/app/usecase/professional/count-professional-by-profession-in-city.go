package professional_usecase

import (
	pkgprofessional "construir_mais_barato/app/domain/professional"
	"fmt"
)

type CountProfessionalsByProfessionInCityUC struct {
	Service pkgprofessional.ProfessionalService
	CityId  *uint
}

type CountProfessionalsByProfessionInCityUCParams struct {
	Service pkgprofessional.ProfessionalService
}

func NewCountProfessionalsByProfessionInCityUC(params CountProfessionalsByProfessionInCityUCParams) CountProfessionalsByProfessionInCityUC {
	return CountProfessionalsByProfessionInCityUC{
		Service: params.Service,
	}
}

func (uc *CountProfessionalsByProfessionInCityUC) Execute() ([]*ProfessionCountPresenter, error) {
	if uc.CityId == nil {
		return nil, fmt.Errorf("invalid data")
	}

	results, err := uc.Service.CountProfessionalsByProfessionInCity(*uc.CityId)
	if err != nil {
		return nil, err
	}

	presenters := make([]*ProfessionCountPresenter, 0)
	if len(results) > 0 {
		for _, item := range results {
			presenters = append(presenters, &ProfessionCountPresenter{
				ProfessionName: item.ProfessionName,
				Quantity:       item.Quantity,
			})
		}
	}
	return presenters, nil
}
