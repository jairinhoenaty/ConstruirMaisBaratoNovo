package client_usecase

import (
	pkgclient "construir_mais_barato/app/domain/client"
	pkguser "construir_mais_barato/app/domain/user"
	pkguseruc "construir_mais_barato/app/usecase/user"
	"fmt"
)

type DeleteClientUC struct {
	Service     pkgclient.ClientService
	ServiceUser pkguser.UserService
	ID          *uint
}

type DeleteClientUCParams struct {
	Service     pkgclient.ClientService
	ServiceUser pkguser.UserService
}

func NewDeleteClientUC(params DeleteClientUCParams) DeleteClientUC {
	return DeleteClientUC{
		Service:     params.Service,
		ServiceUser: params.ServiceUser,
	}
}

func (uc *DeleteClientUC) Execute() error {
	if uc.ID == nil {
		return fmt.Errorf("invalid id")
	}

	oldClient, err := uc.Service.FindById(*uc.ID)

	err = uc.Service.Remove(*uc.ID)
	if err != nil {
		return err
	}

	// deve remover o profissional da tabela de usuário.
	paramsUser := pkguseruc.DeleteUserWithClientUCParams{
		Service: uc.ServiceUser,
	}
	ucUser := pkguseruc.NewDeleteUserWithClientUC(paramsUser)
	ucUser.Name = oldClient.Name
	ucUser.Email = oldClient.Email
	err = ucUser.Execute()
	if err != nil {
		return err
	}

	return nil
}
