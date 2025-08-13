package client_usecase

import (
	// pkgprofession "construir_mais_barato/app/domain/profession"
	pkgclient "construir_mais_barato/app/domain/client"
)

func GenerateClient(assembler *ClientAssembler) pkgclient.Client {
	client := pkgclient.Client{}
	if assembler != nil {

		// professions := make([]pkgprofession.Profession, 0)
		// if len(assembler.Professions) > 0 {
		// 	for _, profession := range assembler.Professions {
		// 		professions = append(professions, pkgprofession.Profession{
		// 			Name:        profession.Name,
		// 			Description: profession.Description,
		// 			Icon:        profession.Icon,
		// 		})

		// 	}
		// }

		client.ID = assembler.ID
		client.Name = assembler.Name
		client.Email = assembler.Email
		client.Telephone = assembler.Telephone
		client.LgpdAceito = assembler.LgpdAceito
		client.CityID = assembler.CityID
		client.Cep = assembler.Cep
		client.Street = assembler.Street
		client.Neighborhood = assembler.Neighborhood
		client.Image = assembler.Image		


	}
	return client
}

func GenerateClientPresenter(client *pkgclient.Client) ClientPresenter {
	presenter := ClientPresenter{}
	if client != nil {


		cidadePresenter := CidadePresenter{
			ID:   client.CityID,
			Name: client.City.Name,
			UF:   client.City.UF,
		}

		presenter.ID = client.ID
		presenter.Name = client.Name
		presenter.Email = client.Email
		presenter.Telephone = client.Telephone
		presenter.LgpdAceito = client.LgpdAceito
		presenter.CreatedAt = client.CreatedAt;
		presenter.Cidade = cidadePresenter
		presenter.Cep = client.Cep
		presenter.Street = client.Street		
		presenter.Neighborhood = client.Neighborhood
		presenter.Image = client.Image


	}
	return presenter
}
