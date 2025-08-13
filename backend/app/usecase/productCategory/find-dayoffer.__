package product_usecase

import (
	pkgproduct "construir_mais_barato/app/domain/product"
)

type FindDayofferProductUC struct {
	Service   pkgproduct.ProductService
	Assembler FindWithPaginationProductAssembler
}

type FindDayofferProductUCParams struct {
	Service pkgproduct.ProductService
}

func NewFindDayofferProductUC(params FindDayofferProductUCParams) FindDayofferProductUC {
	return FindDayofferProductUC{
		Service: params.Service,
	}
}

func (uc *FindDayofferProductUC) Execute() (*[]ProductPresenter, error) {

	products, err := uc.Service.FindDayoffer()
	if err != nil {
		return nil, err
	}

	presenters := make([]ProductPresenter, 0)
	if len(products) > 0 {
		//presenters = GenerateProductPresenter(products)
		for _, product := range products {
			productPresenter := GenerateProductPresenter(product)
			presenters = append(presenters, productPresenter)
			//append(presenters, productPresenter)
		}
	}

	return &presenters, nil
}
