package productCategory_usecase

import (
	pkgproductCategory "construir_mais_barato/app/domain/productCategory"
)

type FindTopLevelUC struct {
	Service pkgproductCategory.ProductCategoryService
}

type FindTopLevelUCParams struct {
	Service pkgproductCategory.ProductCategoryService
}

func NewFindTopLevelUC(params FindTopLevelUCParams) FindTopLevelUC {
	return FindTopLevelUC{Service: params.Service}
}

func (uc *FindTopLevelUC) Execute() ([]*ProductCategoryPresenter, error) {
	categories, err := uc.Service.FindTopLevel()
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

func generateProductCategoryPresenter(category *pkgproductCategory.ProductCategory) ProductCategoryPresenter {
	presenter := ProductCategoryPresenter{
		ID:           category.ID,
		Name:         category.Name,
		ProfessionID: category.ProfessionID,
		ParentID:     category.ParentID,
	}

	if category.Profession != nil && category.Profession.ID > 0 {
		presenter.Profession = category.Profession
	}

	if len(category.Children) > 0 {
		presenter.Children = make([]ProductCategoryPresenter, 0)
		for _, child := range category.Children {
			childPresenter := generateProductCategoryPresenter(&child)
			presenter.Children = append(presenter.Children, childPresenter)
		}
	}

	return presenter
}
