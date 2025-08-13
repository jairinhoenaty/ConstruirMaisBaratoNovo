package store_usecase

import (
	pkgstore "construir_mais_barato/app/domain/store"
	"fmt"
)

type FindByNamedUC struct {
	Service   pkgstore.StoreService
	Assembler *FindByNameAssembler
}

type FindByNamedUCParams struct {
	Service pkgstore.StoreService
}

func NewFindByNamedUC(params FindByNamedUCParams) FindByNamedUC {
	return FindByNamedUC{
		Service: params.Service,
	}
}

func (uc *FindByNamedUC) Execute() (*[]StorePresenter, error) {
	if uc.Assembler == nil {
		return nil, fmt.Errorf("invalid data")
	}

	stores, err := uc.Service.FindByName(uc.Assembler.Name)
	if err != nil {
		return nil, err
	}

	presenters := make([]StorePresenter, 0)
	if len(stores) > 0 {
		for _, store := range stores {
			storePresenter := GenerateStorePresenter(store)
			presenters = append(presenters, storePresenter)
		}
	}
	return &presenters, nil
}
