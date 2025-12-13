package productCategory_usecase

import pkgproductCategory "construir_mais_barato/app/domain/productCategory"

func GenerateProductCategory(assembler *ProductCategoryAssembler) pkgproductCategory.ProductCategory {
	productCategory := pkgproductCategory.ProductCategory{}
	if assembler != nil {
		productCategory.ID = assembler.ID
		productCategory.Name = assembler.Name
		productCategory.ProfessionID = assembler.ProfessionID
		productCategory.Profession = assembler.Profession
		productCategory.ParentID = assembler.ParentID
	}
	return productCategory
}

func GenerateProductCategoryPresenter(productCategory *pkgproductCategory.ProductCategory) ProductCategoryPresenter {

	presenter := ProductCategoryPresenter{}

	presenter.ID = productCategory.ID
	presenter.Name = productCategory.Name
	presenter.ProfessionID = productCategory.ProfessionID
	presenter.Profession = productCategory.Profession
	presenter.ParentID = productCategory.ParentID

	// Map children if present
	if len(productCategory.Children) > 0 {
		children := make([]ProductCategoryPresenter, len(productCategory.Children))
		for i, child := range productCategory.Children {
			children[i] = GenerateProductCategoryPresenter(&child)
		}
		presenter.Children = children
	}

	return presenter
}

func GenerateProductCategorysPresenter(productCategorys []*pkgproductCategory.ProductCategory) *[]ProductCategoryPresenter {
	list := make([]ProductCategoryPresenter, 0)
	if productCategorys != nil && len(productCategorys) > 0 {
		for _, productCategory := range productCategorys {
			presenter := GenerateProductCategoryPresenter(productCategory)
			list = append(list, presenter)
		}
	}

	return &list
}
