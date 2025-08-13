package user_usecase

import (
	pkguser "construir_mais_barato/app/domain/user"
	"fmt"
)

type DeleteUserUC struct {
	Service pkguser.UserService
	ID      *uint
}

type DeleteUserUCParams struct {
	Service pkguser.UserService
}

func NewDeleteUserUC(params DeleteUserUCParams) DeleteUserUC {
	return DeleteUserUC{
		Service: params.Service,
	}
}

func (uc *DeleteUserUC) Execute() error {
	if uc.ID == nil {
		return fmt.Errorf("invalid id")
	}

	err := uc.Service.Remove(*uc.ID)
	if err != nil {
		return err
	}
	return nil
}
