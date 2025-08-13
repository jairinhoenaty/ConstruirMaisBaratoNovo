package product_usecase

import (
	pkgproduct "construir_mais_barato/app/domain/product"
	"fmt"
)

type FindByCityUC struct {
	Service   pkgproduct.ProductService
	Assembler *FindByCityAssembler
}

type FindByCityUCParams struct {
	Service pkgproduct.ProductService
}

func NewFindByCityUC(params FindByCityUCParams) FindByCityUC {
	return FindByCityUC{
		Service: params.Service,
	}
}

func (uc *FindByCityUC) Execute() ([]*ProductPresenter, error) {

	if uc.Assembler.CityID == 0 {
		return nil, fmt.Errorf("invalid City")
	}

	products, err := uc.Service.FindByCity(uc.Assembler.CityID)
	if err != nil {
		return nil, err
	}

	productsPresenter := make([]*ProductPresenter, 0)
	for _, product := range products {
		productPresenter := GenerateProductPresenter(product)

		productsPresenter = append(productsPresenter, &productPresenter)
	}

	return productsPresenter, nil
}
