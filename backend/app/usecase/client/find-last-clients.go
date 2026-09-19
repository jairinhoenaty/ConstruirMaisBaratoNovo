package client_usecase

import (
	pkgclient "construir_mais_barato/app/domain/client"
	pkguser "construir_mais_barato/app/domain/user"
	"fmt"
)

type FindLastClientsUC struct {
	Service         pkgclient.ClientService
	ServiceUser     pkguser.UserService
	QuantityRecords int
}

type FindLastClientsUCParams struct {
	Service     pkgclient.ClientService
	ServiceUser pkguser.UserService
}

func NewFindLastClientsUC(params FindLastClientsUCParams) FindLastClientsUC {
	return FindLastClientsUC{
		Service:     params.Service,
		ServiceUser: params.ServiceUser,
	}
}

func (uc *FindLastClientsUC) Execute() (*[]ClientPresenter, error) {

	clients, err := uc.Service.FindLastClients(uc.QuantityRecords)
	if err != nil {
		return nil, err
	}
	presenters := make([]ClientPresenter, 0)
	if len(clients) > 0 {
		for _, client := range clients {
			clientPresenter := GenerateClientPresenter(&client)
			presenters = append(presenters, clientPresenter)
		}
	}

	// A tag de síndico é informativa: sem ela a lista ainda deve sair.
	if err := markCondoManagers(uc.ServiceUser, presenters); err != nil {
		fmt.Println("Erro ao marcar síndicos nos últimos clientes => " + err.Error())
	}

	return &presenters, nil
}
