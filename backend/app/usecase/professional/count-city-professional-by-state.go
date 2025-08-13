package professional_usecase

import (
	pkgprofessional "construir_mais_barato/app/domain/professional"
	"fmt"
)

type CountCityProfessionalsByStateUC struct {
	Service   pkgprofessional.ProfessionalService
	Assembler *FindWithPaginationProfessionalByStateAssembler
}

type CountCityProfessionalsByStateUCParams struct {
	Service pkgprofessional.ProfessionalService
}

func NewCountCityProfessionalsByStateUC(params CountCityProfessionalsByStateUCParams) CountCityProfessionalsByStateUC {
	return CountCityProfessionalsByStateUC{
		Service: params.Service,
	}
}

func (uc *CountCityProfessionalsByStateUC) Execute() ([]*CityProfessionalCountPresenter, *int64, error) {
	if uc.Assembler == nil {
		return nil, nil, fmt.Errorf("invalid data")
	}

	results, total, err := uc.Service.CountCityProfessionalsByState(uc.Assembler.UF, uc.Assembler.Limit, uc.Assembler.Offset)
	if err != nil {
		return nil, nil, err
	}

	presenters := make([]*CityProfessionalCountPresenter, 0)
	if len(results) > 0 {
		for _, item := range results {
			presenters = append(presenters, &CityProfessionalCountPresenter{
				CityID:            item.CityID,
				CityName:          item.CityName,
				ProfessionalCount: item.ProfessionalCount,
			})
		}
	}
	return presenters, total, nil
}
