package store_usecase

import (
	pkgstore "construir_mais_barato/app/domain/store"
	pkguser "construir_mais_barato/app/domain/user"
	pkguseruc "construir_mais_barato/app/usecase/user"
	"fmt"
)

type FindByIdUC struct {
	Service     pkgstore.StoreService
	ServiceUser pkguser.UserService
	ID          *uint
}

type FindByIdUCParamns struct {
	Service     pkgstore.StoreService
	ServiceUser pkguser.UserService
}

func NewFindByIdUC(params FindByIdUCParamns) FindByIdUC {
	return FindByIdUC{
		Service:     params.Service,
		ServiceUser: params.ServiceUser,
	}
}

func (uc *FindByIdUC) Execute() (*StorePresenter, error) {
	if uc.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	//o id do profissional recebido é o id do profissional da tabela de usuário
	userParams := pkguseruc.FindByIdUCParamns{
		Service: uc.ServiceUser,
	}
	userUC := pkguseruc.NewFindByIdUC(userParams)
	userUC.ID = uc.ID
	user, err := userUC.Execute()
	if err != nil {
		return nil, fmt.Errorf("Usuário não encontrado com o id informado")
	}

	//pesquisar pelo nome do usuário na tabela de profissionais
	foundStore, err := uc.Service.FindByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	// Verifica se a lista de profissionais não está vazia
	if foundStore == nil {
		return nil, fmt.Errorf("Nenhum profissional encontrado")
	}

	// Obtém o ID do primeiro profissional
	store := foundStore

	storePresenter := GenerateStorePresenter(store)
	return &storePresenter, nil
}
