package productCategory_usecase

import (
	"fmt"

	pkgproductCategory "construir_mais_barato/app/domain/productCategory"
)

type FindSubcategoriesByParentUC struct {
	Service  pkgproductCategory.ProductCategoryService
	ParentID *uint
}

type FindSubcategoriesByParentUCParams struct {
	Service pkgproductCategory.ProductCategoryService
}

func NewFindSubcategoriesByParentUC(params FindSubcategoriesByParentUCParams) FindSubcategoriesByParentUC {
	return FindSubcategoriesByParentUC{Service: params.Service}
}

func (uc *FindSubcategoriesByParentUC) Execute() ([]*ProductCategoryPresenter, error) {
	if uc.ParentID == nil || *uc.ParentID == 0 {
		return nil, fmt.Errorf("parent ID is required")
	}

	subcategories, err := uc.Service.FindSubcategoriesByParentID(*uc.ParentID)
	if err != nil {
		return nil, err
	}

	presenters := make([]*ProductCategoryPresenter, 0)
	for _, subcategory := range subcategories {
		presenter := generateProductCategoryPresenter(subcategory)
		presenters = append(presenters, &presenter)
	}

	return presenters, nil
}
