package user_usecase

import (
	pkguser "construir_mais_barato/app/domain/user"
	"fmt"
)

type FindByEmailAndPasswordUC struct {
	Service   pkguser.UserService
	Assembler *LoginAssembler
}

type FindByEmailAndPasswordUCParamns struct {
	Service pkguser.UserService
}

func NewFindByEmailAndPasswordUC(params FindByEmailAndPasswordUCParamns) FindByEmailAndPasswordUC {
	return FindByEmailAndPasswordUC{
		Service: params.Service,
	}
}

func (uc *FindByEmailAndPasswordUC) Execute() (*UserPresenter, error) {
	if uc.Assembler == nil {
		return nil, fmt.Errorf("invalid data")
	}

	user, err := uc.Service.FindByEmail(uc.Assembler.Email)
	if err != nil {
		return nil, err
	}

	userPresenter := GenerateUserPresenter(user)
	return &userPresenter, nil
}
