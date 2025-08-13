package user_usecase

import (
	pkguser "construir_mais_barato/app/domain/user"
	pkgauthuc "construir_mais_barato/app/usecase/auth"
	"fmt"
)

type SaveUserUC struct {
	Service   pkguser.UserService
	Assembler *UserAssembler
}

type SaveUserUCParams struct {
	Service pkguser.UserService
}

func NewSaveUserUC(params SaveUserUCParams) SaveUserUC {
	return SaveUserUC{
		Service: params.Service,
	}
}

func (uc *SaveUserUC) Execute() (*UserPresenter, error) {
	user := GenerateUser(uc.Assembler)
	if user.Password != "" {
		password, err := pkgauthuc.GenerateHashPassword(user.Password)
		if err != nil {
			return nil, err
		}
		user.Password = password
	} else {
		//pesquisar na base de dados o usuário e adicionar a senha

		// pesquisar o usuário pelo email
		findEmailUSerUC := NewFindByEmailUC(FindByEmailUCParams{
			Service: uc.Service,
		})
		findEmailUSerUC.Email = &uc.Assembler.Email
		userByEmail, _ := findEmailUSerUC.Execute()
		fmt.Println("Encontrou usuário pelo email ==> ", user)
		// adicionar a senha cadastrada na base ao usuario
		user.Password = userByEmail.Password

	}

	userSaved, err := uc.Service.Save(user)
	if err != nil {
		return nil, err
	}
	userPresenter := GenerateUserPresenter(userSaved)

	return &userPresenter, nil

}
