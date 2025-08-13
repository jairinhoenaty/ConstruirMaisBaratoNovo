package client_usecase

import (
	pkgclient "construir_mais_barato/app/domain/client"
)

type ExportXLSXClientUC struct {
	Service pkgclient.ClientService
}

type ExportXLSXClientUCParams struct {
	Service pkgclient.ClientService
}

func NewExportXLSXClientUC(params ExportXLSXClientUCParams) ExportXLSXClientUC {
	return ExportXLSXClientUC{
		Service: params.Service,
	}
}

func (uc *ExportXLSXClientUC) Execute() (*[]ClientPresenter, error) {

	clients, err := uc.Service.ExportXLSX()
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
