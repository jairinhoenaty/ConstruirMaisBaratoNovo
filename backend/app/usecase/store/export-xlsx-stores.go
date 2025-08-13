package store_usecase

import (
	pkgstore "construir_mais_barato/app/domain/store"
)

type ExportXLSXStoreUC struct {
	Service pkgstore.StoreService
}

type ExportXLSXStoreUCParams struct {
	Service pkgstore.StoreService
}

func NewExportXLSXStoreUC(params ExportXLSXStoreUCParams) ExportXLSXStoreUC {
	return ExportXLSXStoreUC{
		Service: params.Service,
	}
}

func (uc *ExportXLSXStoreUC) Execute() (*[]StorePresenter, error) {

	stores, err := uc.Service.ExportXLSX()
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
