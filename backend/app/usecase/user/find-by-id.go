package user_usecase

import (
	pkguser "construir_mais_barato/app/domain/user"
	"fmt"
)

type FindByIdUC struct {
	Service pkguser.UserService
	ID      *uint
}

type FindByIdUCParamns struct {
	Service pkguser.UserService
}

func NewFindByIdUC(params FindByIdUCParamns) FindByIdUC {
	return FindByIdUC{
		Service: params.Service,
	}
}

func (uc *FindByIdUC) Execute() (*UserPresenter, error) {
	if uc.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	user, err := uc.Service.FindById(*uc.ID)
	if err != nil {
		return nil, err
	}

	userPresenter := GenerateUserPresenter(user)
	return &userPresenter, nil
}
