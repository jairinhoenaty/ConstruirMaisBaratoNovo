package professional_usecase

import (
	pkgprofessional "construir_mais_barato/app/domain/professional"
	"fmt"
)

type FindByNamedUC struct {
	Service   pkgprofessional.ProfessionalService
	Assembler *FindByNameAssembler
}

type FindByNamedUCParams struct {
	Service pkgprofessional.ProfessionalService
}

func NewFindByNamedUC(params FindByNamedUCParams) FindByNamedUC {
	return FindByNamedUC{
		Service: params.Service,
	}
}

func (uc *FindByNamedUC) Execute() (*[]ProfessionalPresenter, error) {
	if uc.Assembler == nil {
		return nil, fmt.Errorf("invalid data")
	}

	professionals, err := uc.Service.FindByName(uc.Assembler.Name)
	if err != nil {
		return nil, err
	}

	presenters := make([]ProfessionalPresenter, 0)
	if len(professionals) > 0 {
		for _, professional := range professionals {
			professionalPresenter := GenerateProfessionalPresenter(professional)
			presenters = append(presenters, professionalPresenter)
		}
	}
	return &presenters, nil
}
