package user_usecase

import (
	pkguser "construir_mais_barato/app/domain/user"
	"fmt"
)

type DeleteUserWithProfessionalUC struct {
	Service pkguser.UserService
	Name    string
	Email   string
}

type DeleteUserWithProfessionalUCParams struct {
	Service pkguser.UserService
}

func NewDeleteUserWithProfessionalUC(params DeleteUserWithProfessionalUCParams) DeleteUserWithProfessionalUC {
	return DeleteUserWithProfessionalUC{
		Service: params.Service,
	}
}

func (uc *DeleteUserWithProfessionalUC) Execute() error {
	if uc.Name == "" && uc.Email == "" {
		return fmt.Errorf("invalid id")
	}

	user, err := uc.Service.FindByEmail(uc.Email)
	if err != nil {
		return fmt.Errorf("invalid user")
	}
	if user.Name == uc.Name {
		err := uc.Service.Remove(user.ID)
		if err != nil {
			return err
		}
	}
	return nil
}
