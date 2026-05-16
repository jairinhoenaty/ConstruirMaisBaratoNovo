package user_usecase

import (
	pkgclient "construir_mais_barato/app/domain/client"
	pkgprofessional "construir_mais_barato/app/domain/professional"
	"fmt"
)

type FindByPhoneUC struct {
	ClientService       pkgclient.ClientService
	ProfessionalService pkgprofessional.ProfessionalService
	Telephone           *string
}

type FindByPhoneUCParams struct {
	ClientService       pkgclient.ClientService
	ProfessionalService pkgprofessional.ProfessionalService
}

func NewFindByPhoneUC(params FindByPhoneUCParams) FindByPhoneUC {
	return FindByPhoneUC{
		ClientService:       params.ClientService,
		ProfessionalService: params.ProfessionalService,
	}
}

func (uc *FindByPhoneUC) Execute() error {
	if uc.Telephone == nil {
		return fmt.Errorf("invalid data")
	}

	_, errClient := uc.ClientService.FindByTelephone(*uc.Telephone)
	if errClient == nil {
		return nil
	}

	_, errProfessional := uc.ProfessionalService.FindByTelephone(*uc.Telephone)
	if errProfessional == nil {
		return nil
	}

	return fmt.Errorf("telefone não encontrado")
}
