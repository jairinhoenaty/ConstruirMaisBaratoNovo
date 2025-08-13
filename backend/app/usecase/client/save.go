package client_usecase

import (
	pkgclient "construir_mais_barato/app/domain/client"
	pkguser "construir_mais_barato/app/domain/user"
	pkguseruc "construir_mais_barato/app/usecase/user"
	"fmt"
	"time"
)

type SaveClientUC struct {
	Service     pkgclient.ClientService
	ServiceUser pkguser.UserService
	Assembler   *ClientAssembler
}

type SaveClientUCParams struct {
	Service     pkgclient.ClientService
	ServiceUser pkguser.UserService
}

func NewSaveClientUC(params SaveClientUCParams) SaveClientUC {
	return SaveClientUC{
		Service:     params.Service,
		ServiceUser: params.ServiceUser,
	}
}

func (uc *SaveClientUC) Execute() (*ClientPresenter, error) {
	if uc.Assembler == nil {
		return nil, fmt.Errorf("invalid Data")
	}

	client := GenerateClient(uc.Assembler)

	clientSaved, err := uc.Service.Save(client)
	if err != nil {
		return nil, err
	}

	// pesquisar o usuário pelo email
	findEmailUSerUC := pkguseruc.NewFindByEmailUC(pkguseruc.FindByEmailUCParams{
		Service: uc.ServiceUser,
	})
	findEmailUSerUC.Email = &uc.Assembler.Email
	user, _ := findEmailUSerUC.Execute()
	// fmt.Println("Encontrou usuário pelo email ==> ", user)
	userAssembler := pkguseruc.UserAssembler{}

	if user != nil {
		//fmt.Println("-------");
		//fmt.Println(user.Password);
		//fmt.Println("-------");
		// encontrou usuário cadastrado então devo atualizar o usuário
		// criar o usuário com base nos dados que o profissional informou para fazer o login no sistema
		userAssembler.ID = user.ID
		userAssembler.Name = clientSaved.Name
		userAssembler.Email = clientSaved.Email
		//userAssembler.Password = user.Password
		userAssembler.Profile = "client"

	} else {

		// criar o usuário com base nos dados que o profissional informou para fazer o login no sistema
		userAssembler.Name = uc.Assembler.Name
		userAssembler.Email = uc.Assembler.Email
		userAssembler.Profile = "client"
		if uc.Assembler.Password == "" {

			// Obtém a data e hora atual
			now := time.Now()

			// Converte `time.Time` para string com formatação
			formattedTime := now.Format("2006-01-02 15:04:05")

			userAssembler.Password = formattedTime
		} else {
			userAssembler.Password = uc.Assembler.Password
		}
	}

	// caso de uso para salvar usuários
	ucUserParams := pkguseruc.SaveUserUCParams{
		Service: uc.ServiceUser,
	}
	ucUser := pkguseruc.NewSaveUserUC(ucUserParams)
	ucUser.Assembler = &userAssembler
	_, err = ucUser.Execute()
	if err != nil {
		fmt.Println("Erro ao salvar o usuário criado para o profissional =>  " + err.Error())
		return nil, err
	}

	userPresenter := GenerateClientPresenter(clientSaved)

	return &userPresenter, nil

}
