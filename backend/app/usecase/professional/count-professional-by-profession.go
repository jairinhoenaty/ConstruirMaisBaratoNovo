package professional_usecase

import (
	pkgprofessional "construir_mais_barato/app/domain/professional"
)

type CountProfessionalsByProfessionUC struct {
	Service pkgprofessional.ProfessionalService
}

type CountProfessionalsByProfessionUCParams struct {
	Service pkgprofessional.ProfessionalService
}

func NewCountProfessionalsByProfessionUC(params CountProfessionalsByProfessionUCParams) CountProfessionalsByProfessionUC {
	return CountProfessionalsByProfessionUC{
		Service: params.Service,
	}
}

func (uc *CountProfessionalsByProfessionUC) Execute() (*[]ProfessionCountPresenter, error) {

	results, err := uc.Service.CountProfessionalsByProfession()
	if err != nil {
		return nil, err
	}

	presenters := make([]ProfessionCountPresenter, 0)
	if len(results) > 0 {
		for _, item := range results {
			presenters = append(presenters, ProfessionCountPresenter{
				ProfessionName: item.ProfessionName,
				Quantity:       item.Quantity,
			})
		}
	}
	return &presenters, nil
}
