package store_usecase

import (
	pkgstore "construir_mais_barato/app/domain/store"
)

type FindAllStoreUC struct {
	Service   pkgstore.StoreService
	Assembler FindWithPaginationStoreAssembler
}

type FindAllStoreUCParams struct {
	Service pkgstore.StoreService
}

func NewFindAllStoreUC(params FindAllStoreUCParams) FindAllStoreUC {
	return FindAllStoreUC{
		Service: params.Service,
	}
}

func (uc *FindAllStoreUC) Execute() (*[]StorePresenter, int64, error) {

	stores, total, err := uc.Service.FindAll(uc.Assembler.Limit, uc.Assembler.Offset)
	if err != nil {
		return nil, 0, err
	}

	presenters := make([]StorePresenter, 0)
	if len(stores) > 0 {
		for _, store := range stores {
			storePresenter := GenerateStorePresenter(store)
			presenters = append(presenters, storePresenter)
		}
	}

	return &presenters, total, nil
}
