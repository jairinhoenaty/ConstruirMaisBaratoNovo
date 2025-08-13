package store_usecase

import (
	pkgstore "construir_mais_barato/app/domain/store"
	pkguser "construir_mais_barato/app/domain/user"
	pkguseruc "construir_mais_barato/app/usecase/user"
	"fmt"
	"time"
)

type SaveStoreUC struct {
	Service     pkgstore.StoreService
	ServiceUser pkguser.UserService
	Assembler   *StoreAssembler
}

type SaveStoreUCParams struct {
	Service     pkgstore.StoreService
	ServiceUser pkguser.UserService
}

func NewSaveStoreUC(params SaveStoreUCParams) SaveStoreUC {
	return SaveStoreUC{
		Service:     params.Service,
		ServiceUser: params.ServiceUser,
	}
}

func (uc *SaveStoreUC) Execute() (*StorePresenter, error) {
	if uc.Assembler == nil {
		return nil, fmt.Errorf("invalid Data")
	}

	store := GenerateStore(uc.Assembler)

	storeSaved, err := uc.Service.Save(store)
	if err != nil {
		return nil, err
	}

	// pesquisar o usuário pelo email
	findEmailUSerUC := pkguseruc.NewFindByEmailUC(pkguseruc.FindByEmailUCParams{
		Service: uc.ServiceUser,
	})
	findEmailUSerUC.Email = &uc.Assembler.Email
	user, _ := findEmailUSerUC.Execute()
	userAssembler := pkguseruc.UserAssembler{}
	if user != nil {
		// encontrou usuário cadastrado então devo atualizar o usuário
		// criar o usuário com base nos dados que o profissional informou para fazer o login no sistema
		userAssembler.ID = user.ID
		userAssembler.Name = storeSaved.Name
		userAssembler.Email = storeSaved.Email
		//userAssembler.Password = user.Password
		userAssembler.Profile = "store"

	} else {

		// criar o usuário com base nos dados que o profissional informou para fazer o login no sistema
		userAssembler.Name = uc.Assembler.Name
		userAssembler.Email = uc.Assembler.Email
		userAssembler.Profile = "store"
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

	userPresenter := GenerateStorePresenter(storeSaved)

	return &userPresenter, nil

}
