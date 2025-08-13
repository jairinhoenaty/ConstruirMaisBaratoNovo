package client_usecase

import (
	pkgclient "construir_mais_barato/app/domain/client"
	"fmt"
)

type FindByNamedUC struct {
	Service   pkgclient.ClientService
	Assembler *FindByNameAssembler
}

type FindByNamedUCParams struct {
	Service pkgclient.ClientService
}

func NewFindByNamedUC(params FindByNamedUCParams) FindByNamedUC {
	return FindByNamedUC{
		Service: params.Service,
	}
}

func (uc *FindByNamedUC) Execute() (*[]ClientPresenter, error) {
	if uc.Assembler == nil {
		return nil, fmt.Errorf("invalid data")
	}

	clients, err := uc.Service.FindByName(uc.Assembler.Name)
	if err != nil {
		return nil, err
	}

	presenters := make([]ClientPresenter, 0)
	if len(clients) > 0 {
		for _, client := range clients {
			clientPresenter := GenerateClientPresenter(client)
			presenters = append(presenters, clientPresenter)
		}
	}
	return &presenters, nil
}
