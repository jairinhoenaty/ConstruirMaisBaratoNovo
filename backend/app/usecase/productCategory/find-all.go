package productCategory_usecase

import (
	pkgproductCategory "construir_mais_barato/app/domain/productCategory"
)

type FindAllUC struct {
	Service pkgproductCategory.ProductCategoryService
}

type FindAllUCParams struct {
	Service pkgproductCategory.ProductCategoryService
}

func NewFindAllUC(params FindAllUCParams) FindAllUC {
	return FindAllUC{Service: params.Service}
}

func (uc *FindAllUC) Execute() ([]*ProductCategoryPresenter, error) {
	categories, err := uc.Service.FindAll()
	if err != nil {
		return nil, err
	}

	presenters := make([]*ProductCategoryPresenter, 0)
	for _, category := range categories {
		presenter := generateProductCategoryPresenter(category)
		presenters = append(presenters, &presenter)
	}

	return presenters, nil
}
