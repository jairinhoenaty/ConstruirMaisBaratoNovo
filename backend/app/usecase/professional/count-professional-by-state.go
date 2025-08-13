package professional_usecase

import (
	pkgprofessional "construir_mais_barato/app/domain/professional"
	"fmt"
)

type CountProfessionalsByStateUC struct {
	Service   pkgprofessional.ProfessionalService
	Assembler *FindWithPaginationProfessionalByStateAssembler
}

type CountProfessionalsByStateUCParams struct {
	Service pkgprofessional.ProfessionalService
}

func NewCountProfessionalsByStateUC(params CountProfessionalsByStateUCParams) CountProfessionalsByStateUC {
	return CountProfessionalsByStateUC{
		Service: params.Service,
	}
}

func (uc *CountProfessionalsByStateUC) Execute() ([]*UFProfessionalCountPresenter, *int64, error) {
	if uc.Assembler == nil {
		return nil, nil, fmt.Errorf("invalid data")
	}

	results, total, err := uc.Service.CountProfessionalsByState(uc.Assembler.UF, uc.Assembler.Limit, uc.Assembler.Offset)
	if err != nil {
		return nil, nil, err
	}

	presenters := make([]*UFProfessionalCountPresenter, 0)
	if len(results) > 0 {
		for _, item := range results {
			presenters = append(presenters, &UFProfessionalCountPresenter{
				UFName:          item.UFName,
				ProfessionalCount: item.ProfessionalCount,
			})
		}
	}
	return presenters, total, nil
}
