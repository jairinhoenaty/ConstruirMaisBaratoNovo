package product_usecase

import (
	pkgproduct "construir_mais_barato/app/domain/product"
	"fmt"
)

type FindByIdUC struct {
	Service pkgproduct.ProductService
	ID      *uint
}

type FindByIdUCParams struct {
	Service pkgproduct.ProductService
}

func NewFindByIdUC(params FindByIdUCParams) FindByIdUC {
	return FindByIdUC{
		Service: params.Service,
	}
}

func (uc *FindByIdUC) Execute() (*ProductPresenter, error) {
	if uc.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	product, err := uc.Service.FindById(*uc.ID)
	if err != nil {
		return nil, err
	}

	productPresenter := GenerateProductPresenter(product)
	return &productPresenter, nil
}
