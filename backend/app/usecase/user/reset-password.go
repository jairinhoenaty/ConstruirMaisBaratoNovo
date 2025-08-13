package user_usecase

import (
	pkguser "construir_mais_barato/app/domain/user"
	"fmt"
)

type ResetPasswordUC struct {
	Service   pkguser.UserService
	Assembler *ResetPasswordAssembler
}

type ResetPasswordUCParams struct {
	Service pkguser.UserService
}

func NewResetPasswordUC(params ResetPasswordUCParams) ResetPasswordUC {
	return ResetPasswordUC{
		Service: params.Service,
	}
}

func (uc *ResetPasswordUC) Execute() error {
	if uc.Assembler == nil {
		return fmt.Errorf("invalid data")
	}

	// decodificar a email
	userEmail, err := DecryptValue(uc.Assembler.Email)
	if err != nil {
		fmt.Println("Erro ao Decriptografar o email")
	}

	//pesquisar o usuário pelo email
	params := FindByEmailUCParams{
		Service: uc.Service,
	}
	userUC := NewFindByEmailUC(params)
	userUC.Email = &userEmail
	user, err := userUC.Execute()
	if err != nil {
		fmt.Println("Não encontrou o usuário com o email informado")
		return err
	}

	//salvar o usuário com a nova senha
	saveUserAssembler := UserAssembler{
		ID:       user.ID,
		Name:     user.Name,
		Email:    userEmail,
		Profile:  user.Profile,
		Password: uc.Assembler.Password,
	}
	saveParams := SaveUserUCParams{
		Service: uc.Service,
	}
	saveUserUC := NewSaveUserUC(saveParams)
	saveUserUC.Assembler = &saveUserAssembler
	_, err = saveUserUC.Execute()
	if err != nil {
		fmt.Println("Erro ao atualizar o usuário com a nova senha")
		return err
	}
	return nil
}
