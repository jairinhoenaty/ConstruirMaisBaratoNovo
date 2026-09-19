package client_usecase

import (
	pkgclient "construir_mais_barato/app/domain/client"
	pkguser "construir_mais_barato/app/domain/user"
	"fmt"
)

type FindAllClientUC struct {
	Service     pkgclient.ClientService
	ServiceUser pkguser.UserService
	Assembler   FindWithPaginationClientAssembler
}

type FindAllClientUCParams struct {
	Service     pkgclient.ClientService
	ServiceUser pkguser.UserService
}

func NewFindAllClientUC(params FindAllClientUCParams) FindAllClientUC {
	return FindAllClientUC{
		Service:     params.Service,
		ServiceUser: params.ServiceUser,
	}
}

func (uc *FindAllClientUC) Execute() (*[]ClientPresenter, int64, error) {

	clients, total, err := uc.Service.FindAll(uc.Assembler.Limit, uc.Assembler.Offset)
	if err != nil {
		return nil, 0, err
	}

	presenters := make([]ClientPresenter, 0)
	if len(clients) > 0 {
		for _, client := range clients {
			clientPresenter := GenerateClientPresenter(client)
			presenters = append(presenters, clientPresenter)
		}
	}

	// A tag de síndico é informativa: sem ela a lista ainda deve sair.
	if err := markCondoManagers(uc.ServiceUser, presenters); err != nil {
		fmt.Println("Erro ao marcar síndicos na lista de clientes => " + err.Error())
	}

	return &presenters, total, nil
}
