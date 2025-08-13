package professional_usecase

import (
	pkgprofessional "construir_mais_barato/app/domain/professional"
	"fmt"
)

type FindByProfessionAndLocationUC struct {
	Service   pkgprofessional.ProfessionalService
	Assembler *FindByProfessionAndLocationAssembler
}

type FindByProfessionAndLocationUCParams struct {
	Service pkgprofessional.ProfessionalService
}

func NewFindByProfessionAndLocationUC(params FindByProfessionAndLocationUCParams) FindByProfessionAndLocationUC {
	return FindByProfessionAndLocationUC{
		Service: params.Service,
	}
}

func (uc *FindByProfessionAndLocationUC) Execute() ([]*ProfessionalPresenter, int64, error) {
	if uc.Assembler == nil {
		return nil, 0, fmt.Errorf("invalid data")
	}

	professionals, total, err :=
		uc.Service.FindByProfessionAndLocation(uc.Assembler.ProfessionId,
			uc.Assembler.Latitude, uc.Assembler.Longitude,
			uc.Assembler.Distance, uc.Assembler.Limit, uc.Assembler.Offset)
	if err != nil {
		return nil, 0, err
	}

	presenters := make([]*ProfessionalPresenter, 0)

	for _, professional := range professionals {
		presenter := GenerateProfessionalPresenter(professional)
		presenters = append(presenters, &presenter)
	}
	return presenters, total, nil
}
