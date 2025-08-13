package productCategory_usecase

import pkgproductCategory "construir_mais_barato/app/domain/productCategory"

type SaveProductCategoryUC struct {
	Service   pkgproductCategory.ProductCategoryService
	Assembler *ProductCategoryAssembler
}

type SaveProductCategoryUCParams struct {
	Service pkgproductCategory.ProductCategoryService
}

func NewSaveProductCategoryUC(params SaveProductCategoryUCParams) SaveProductCategoryUC {
	return SaveProductCategoryUC{
		Service: params.Service,
	}
}

func (uc *SaveProductCategoryUC) Execute() (*ProductCategoryPresenter, error) {
	productCategory := GenerateProductCategory(uc.Assembler)
	productCategorySaved, err := uc.Service.Save(productCategory)
	if err != nil {
		println("Error saving product category:", err.Error())
	}

	if err != nil {
		return nil, err
	}
	productCategoryPresenter := GenerateProductCategoryPresenter(productCategorySaved)

	return &productCategoryPresenter, nil

}
