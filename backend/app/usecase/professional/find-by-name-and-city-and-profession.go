package professional_usecase

import (
	pkgprofessional "construir_mais_barato/app/domain/professional"
	"fmt"
)

type FindByNameProfessionalAndCityAndProfessionUC struct {
	Service   pkgprofessional.ProfessionalService
	Assembler *FindProfessionalByCityAndProfessionAssembler
}

type FindByNameProfessionalAndCityAndProfessionUCParamns struct {
	Service pkgprofessional.ProfessionalService
}

func NewFindByNameProfessionalAndCityAndProfessionUC(params FindByNameProfessionalAndCityAndProfessionUCParamns) FindByNameProfessionalAndCityAndProfessionUC {
	return FindByNameProfessionalAndCityAndProfessionUC{
		Service: params.Service,
	}
}

func (uc *FindByNameProfessionalAndCityAndProfessionUC) Execute() ([]*ProfessionalPresenter, error) {
	if uc.Assembler == nil {
		return nil, fmt.Errorf("invalid data")
	}

	professionals, err := uc.Service.FindByNameAndCityAndProfession(uc.Assembler.Name, uc.Assembler.CityId, uc.Assembler.ProfessionId, uc.Assembler.Limit, uc.Assembler.Offset)
	if err != nil {
		return nil, err
	}

	presenters := make([]*ProfessionalPresenter, 0)

	for _, professional := range professionals {
		presenter := GenerateProfessionalPresenter(professional)
		presenters = append(presenters, &presenter)
	}
	return presenters, nil
}
