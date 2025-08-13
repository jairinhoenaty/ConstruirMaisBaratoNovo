package professional_usecase

import (
	pkgprofessional "construir_mais_barato/app/domain/professional"
	pkguser "construir_mais_barato/app/domain/user"
	pkguseruc "construir_mais_barato/app/usecase/user"
	"fmt"
)

type DeleteProfessionalUC struct {
	Service     pkgprofessional.ProfessionalService
	ServiceUser pkguser.UserService
	ID          *uint
}

type DeleteProfessionalUCParams struct {
	Service     pkgprofessional.ProfessionalService
	ServiceUser pkguser.UserService
}

func NewDeleteProfessionalUC(params DeleteProfessionalUCParams) DeleteProfessionalUC {
	return DeleteProfessionalUC{
		Service:     params.Service,
		ServiceUser: params.ServiceUser,
	}
}

func (uc *DeleteProfessionalUC) Execute() error {
	if uc.ID == nil {
		return fmt.Errorf("invalid id")
	}

	oldProfessional, err := uc.Service.FindById(*uc.ID)

	err = uc.Service.Remove(*uc.ID)
	if err != nil {
		return err
	}

	// deve remover o profissional da tabela de usuário.
	paramsUser := pkguseruc.DeleteUserWithProfessionalUCParams{
		Service: uc.ServiceUser,
	}
	ucUser := pkguseruc.NewDeleteUserWithProfessionalUC(paramsUser)
	ucUser.Name = oldProfessional.Name
	ucUser.Email = oldProfessional.Email
	err = ucUser.Execute()
	if err != nil {
		return err
	}

	return nil
}
