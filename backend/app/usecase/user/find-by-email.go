package user_usecase

import (
	pkguser "construir_mais_barato/app/domain/user"
	"fmt"
)

type FindByEmailUC struct {
	Service pkguser.UserService
	Email   *string
}

type FindByEmailUCParams struct {
	Service pkguser.UserService
}

func NewFindByEmailUC(params FindByEmailUCParams) FindByEmailUC {
	return FindByEmailUC{
		Service: params.Service,
	}
}

func (uc *FindByEmailUC) Execute() (*UserPresenter, error) {
	if uc.Email == nil {
		return nil, fmt.Errorf("invalid data")
	}

	user, err := uc.Service.FindByEmail(*uc.Email)
	if err != nil {
		return nil, err
	}

	userPresenter := GenerateUserPresenter(user)
	return &userPresenter, nil
}
