package user_usecase

import (
	pkguser "construir_mais_barato/app/domain/user"
	"fmt"
)

type DeleteUserWithClientUC struct {
	Service pkguser.UserService
	Name    string
	Email   string
}

type DeleteUserWithClientUCParams struct {
	Service pkguser.UserService
}

func NewDeleteUserWithClientUC(params DeleteUserWithClientUCParams) DeleteUserWithClientUC {
	return DeleteUserWithClientUC{
		Service: params.Service,
	}
}

func (uc *DeleteUserWithClientUC) Execute() error {
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
