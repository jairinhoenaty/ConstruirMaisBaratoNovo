package product_usecase

import (
	pkgproduct "construir_mais_barato/app/domain/product"
)

type FindAllProductUC struct {
	Service   pkgproduct.ProductService
	Assembler FindWithPaginationProductAssembler
}

type FindAllProductUCParams struct {
	Service pkgproduct.ProductService
}

func NewFindAllProductUC(params FindAllProductUCParams) FindAllProductUC {
	return FindAllProductUC{
		Service: params.Service,
	}
}

func (uc *FindAllProductUC) Execute() (*[]ProductPresenter, int64, error) {

	products, total, err := uc.Service.FindAll(uc.Assembler.Limit, uc.Assembler.Offset,uc.Assembler.ProfessionalID,uc.Assembler.StoreID,uc.Assembler.Approved,uc.Assembler.DayOffer)
	if err != nil {
		return nil, 0, err
	}

	presenters := make([]ProductPresenter, 0)
	if len(products) > 0 {
		//presenters = GenerateProductPresenter(products)
		for _, product := range products {
			//productPresenter := GenerateProductPresenter(product)
			presenters = append(presenters,GenerateProductPresenter(product))
			//append(presenters, productPresenter)
		}
	}

	return &presenters, total, nil
}
