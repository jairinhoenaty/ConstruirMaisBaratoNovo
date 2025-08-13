package professional_usecase

import (
	pkgprofessional "construir_mais_barato/app/domain/professional"
	"fmt"
)

type FindByProfessionalByCityAndProfessionUC struct {
	Service   pkgprofessional.ProfessionalService
	Assembler *FindProfessionalByCityAndProfessionAssembler
}

type FindByProfessionalByCityAndProfessionUCParamns struct {
	Service pkgprofessional.ProfessionalService
}

func NewFindByProfessionalByCityAndProfessionUC(params FindByProfessionalByCityAndProfessionUCParamns) FindByProfessionalByCityAndProfessionUC {
	return FindByProfessionalByCityAndProfessionUC{
		Service: params.Service,
	}
}

func (uc *FindByProfessionalByCityAndProfessionUC) Execute() ([]*ProfessionalPresenter, int64, error) {
	if uc.Assembler == nil {
		return nil, 0, fmt.Errorf("invalid data")
	}

	professionals, total, err := uc.Service.FindByCityAndProfession(uc.Assembler.CityId, uc.Assembler.ProfessionId, uc.Assembler.Limit, uc.Assembler.Offset)
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
