package product_usecase

import pkgproduct "construir_mais_barato/app/domain/product"

type SaveProductUC struct {
	Service   pkgproduct.ProductService
	Assembler *ProductAssembler
}

type SaveProductUCParams struct {
	Service pkgproduct.ProductService
}

func NewSaveProductUC(params SaveProductUCParams) SaveProductUC {
	return SaveProductUC{
		Service: params.Service,
	}
}

func (uc *SaveProductUC) Execute() (*ProductPresenter, error) {
	product := GenerateProduct(uc.Assembler)
	productSaved, err := uc.Service.Save(product)
	if err != nil {
		return nil, err
	}
	productPresenter := GenerateProductPresenter(productSaved)

	return &productPresenter, nil

}
