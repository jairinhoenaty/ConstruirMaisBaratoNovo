package professional_usecase

import (
	pkgprofessional "construir_mais_barato/app/domain/professional"
	pkguseruc "construir_mais_barato/app/usecase/user"
	"fmt"
)

type FindByNameAndPasswordUC struct {
	Service   pkgprofessional.ProfessionalService
	Assembler *pkguseruc.LoginAssembler
}

type FindByNameAndPasswordUCParamns struct {
	Service pkgprofessional.ProfessionalService
}

func NewFindByNameAndPasswordUC(params FindByNameAndPasswordUCParamns) FindByNameAndPasswordUC {
	return FindByNameAndPasswordUC{
		Service: params.Service,
	}
}

func (uc *FindByNameAndPasswordUC) Execute() (*ProfessionalPresenter, error) {
	if uc.Assembler == nil {
		return nil, fmt.Errorf("invalid data")
	}

	professional, err := uc.Service.FindByEmail(uc.Assembler.Email)
	if err != nil {
		return nil, err
	}

	professionalPresenter := GenerateProfessionalPresenter(professional)
	return &professionalPresenter, nil
}
