package auth

import (
	pkgprofessional "construir_mais_barato/app/domain/professional"
	pkguser "construir_mais_barato/app/domain/user"
	"errors"
)

type AuthenticateUC struct {
	UserService         pkguser.UserService
	ProfessionalService pkgprofessional.ProfessionalService
	Assembler           *LoginAssembler
}

type AuthenticateUCParams struct {
	UserService         pkguser.UserService
	ProfessionalService pkgprofessional.ProfessionalService
}

func NewLoginUC(params AuthenticateUCParams) AuthenticateUC {
	return AuthenticateUC{
		UserService:         params.UserService,
		ProfessionalService: params.ProfessionalService,
	}
}

func (uc *AuthenticateUC) Execute() (*AuthenticatePresenter, error) {

	if uc.Assembler == nil {
		return nil, errors.New("Invalid data")
	}

	presenter := AuthenticatePresenter{}

	// pesquisar na tabela de usuário.
	user, _ := uc.UserService.FindByEmail(uc.Assembler.Email)
	// se não encontrar nenhum resultado, procurar na tabela de profissionais
	if user != nil {

		isValidPassword := ComparePassword(uc.Assembler.Password, user.Password)

		if isValidPassword {
			token, err := GenerateToken(UserPresenter{
				ID:   user.ID,
				Name: user.Name,
			})

			if err == nil {
				user := GenerateUserPresenter(user.ID, user.Name, user.Profile, user.Email, user.GoogleToken)
				presenter = GenerateAuthenticatePresenter(token, true, user)
			}

		}

		return &presenter, nil

	}

	return nil, errors.New("User Not found")

}
