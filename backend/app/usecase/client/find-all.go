package client_usecase

import (
	pkgclient "construir_mais_barato/app/domain/client"
)

type FindAllClientUC struct {
	Service   pkgclient.ClientService
	Assembler FindWithPaginationClientAssembler
}

type FindAllClientUCParams struct {
	Service pkgclient.ClientService
}

func NewFindAllClientUC(params FindAllClientUCParams) FindAllClientUC {
	return FindAllClientUC{
		Service: params.Service,
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

	return &presenters, total, nil
}
