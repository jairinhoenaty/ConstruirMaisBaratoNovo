package professional_usecase

import (
	pkgprofessional "construir_mais_barato/app/domain/professional"
)

type FindAllProfessionalUC struct {
	Service   pkgprofessional.ProfessionalService
	Assembler FindWithPaginationProfessionalAssembler
}

type FindAllProfessionalUCParams struct {
	Service pkgprofessional.ProfessionalService
}

func NewFindAllProfessionalUC(params FindAllProfessionalUCParams) FindAllProfessionalUC {
	return FindAllProfessionalUC{
		Service: params.Service,
	}
}

func (uc *FindAllProfessionalUC) Execute() (*[]ProfessionalPresenter, int64, error) {

	professionals, total, err := uc.Service.FindAll(uc.Assembler.Limit, uc.Assembler.Offset, uc.Assembler.Filter, uc.Assembler.Uf, uc.Assembler.ProfessionId, uc.Assembler.Order)
	if err != nil {
		return nil, 0, err
	}

	presenters := make([]ProfessionalPresenter, 0)
	if len(professionals) > 0 {
		for _, professional := range professionals {
			professionalPresenter := GenerateProfessionalPresenter(professional)
			presenters = append(presenters, professionalPresenter)
		}
	}

	return &presenters, total, nil
}
