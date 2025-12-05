package store_usecase

import (
	pkgstore "construir_mais_barato/app/domain/store"
	"fmt"
)

type FindStoreByCategoryAndSubCategory struct {
	Service   pkgstore.StoreService
	Assembler *FindByCategoryAndSubCategory
}

type FindStoreByCategoryAndSubCategoryParams struct {
	Service pkgstore.StoreService
}

func NewFindStoreByCategoryAndSubCategory(params FindStoreByCategoryAndSubCategoryParams) FindStoreByCategoryAndSubCategory {
	return FindStoreByCategoryAndSubCategory{
		Service: params.Service,
	}
}

func (uc *FindStoreByCategoryAndSubCategory) Execute() (*[]StorePresenter, error) {
	if uc.Assembler == nil {
		return nil, fmt.Errorf("invalid data")
	}

	stores, err := uc.Service.FindByCategoryAndSubCategory(uc.Assembler.CategoryID, uc.Assembler.SubCategoriesID)
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
