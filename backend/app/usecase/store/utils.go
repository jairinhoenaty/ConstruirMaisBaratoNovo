package store_usecase

import (
	// pkgprofession "construir_mais_barato/app/domain/profession"
	pkgstore "construir_mais_barato/app/domain/store"
)

func GenerateStore(assembler *StoreAssembler) pkgstore.Store {
	store := pkgstore.Store{}
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

		store.ID = assembler.ID
		store.Name = assembler.Name
		store.Email = assembler.Email
		store.Telephone = assembler.Telephone
		store.LgpdAceito = assembler.LgpdAceito
		store.CityID = assembler.CityID
		store.Cep = assembler.Cep
		store.Street = assembler.Street
		store.Neighborhood = assembler.Neighborhood
		store.Image = assembler.Image		


	}
	return store
}

func GenerateStorePresenter(store *pkgstore.Store) StorePresenter {
	presenter := StorePresenter{}
	if store != nil {


		cidadePresenter := CidadePresenter{
			ID:   store.CityID,
			Name: store.City.Name,
			UF:   store.City.UF,
		}

		presenter.ID = store.ID
		presenter.Name = store.Name
		presenter.Email = store.Email
		presenter.Telephone = store.Telephone
		presenter.LgpdAceito = store.LgpdAceito
		presenter.CreatedAt = store.CreatedAt;
		presenter.Cidade = cidadePresenter
		presenter.Cep = store.Cep
		presenter.Street = store.Street		
		presenter.Neighborhood = store.Neighborhood
		presenter.Image = store.Image


	}
	return presenter
}
