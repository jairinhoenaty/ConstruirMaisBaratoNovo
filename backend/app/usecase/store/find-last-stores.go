package store_usecase

import (
	pkgstore "construir_mais_barato/app/domain/store"
)

type FindLastStoresUC struct {
	Service         pkgstore.StoreService
	QuantityRecords int
}

type FindLastStoresUCParams struct {
	Service pkgstore.StoreService
}

func NewFindLastStoresUC(params FindLastStoresUCParams) FindLastStoresUC {
	return FindLastStoresUC{
		Service: params.Service,
	}
}

func (uc *FindLastStoresUC) Execute() (*[]StorePresenter, error) {

	stores, err := uc.Service.FindLastStores(uc.QuantityRecords)
	if err != nil {
		return nil, err
	}
	presenters := make([]StorePresenter, 0)
	if len(stores) > 0 {
		for _, store := range stores {
			storePresenter := GenerateStorePresenter(&store)
			presenters = append(presenters, storePresenter)
		}
	}
	return &presenters, nil
}
