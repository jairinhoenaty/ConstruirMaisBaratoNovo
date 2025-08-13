package professional_usecase

import (
	pkgprofessional "construir_mais_barato/app/domain/professional"
)

type FindLastProfessionalsUC struct {
	Service         pkgprofessional.ProfessionalService
	QuantityRecords int
}

type FindLastProfessionalsUCParams struct {
	Service pkgprofessional.ProfessionalService
}

func NewFindLastProfessionalsUC(params FindLastProfessionalsUCParams) FindLastProfessionalsUC {
	return FindLastProfessionalsUC{
		Service: params.Service,
	}
}

func (uc *FindLastProfessionalsUC) Execute() (*[]ProfessionalPresenter, error) {

	professionals, err := uc.Service.FindLastProfessionals(uc.QuantityRecords)
	if err != nil {
		return nil, err
	}
	presenters := make([]ProfessionalPresenter, 0)
	if len(professionals) > 0 {
		for _, professional := range professionals {
			professionalPresenter := GenerateProfessionalPresenter(&professional)
			presenters = append(presenters, professionalPresenter)
		}
	}
	return &presenters, nil
}
