package client_usecase

import (
	pkgclient "construir_mais_barato/app/domain/client"
)

type FindLastClientsUC struct {
	Service         pkgclient.ClientService
	QuantityRecords int
}

type FindLastClientsUCParams struct {
	Service pkgclient.ClientService
}

func NewFindLastClientsUC(params FindLastClientsUCParams) FindLastClientsUC {
	return FindLastClientsUC{
		Service: params.Service,
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
	return &presenters, nil
}
