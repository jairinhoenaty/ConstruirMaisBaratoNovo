package product_usecase

import (
	pkgcity "construir_mais_barato/app/domain/product"
	"fmt"
)

type DeleteProductUC struct {
	Service pkgcity.ProductService
	ID      *uint
}

type DeleteProductUCParams struct {
	Service pkgcity.ProductService
}

func NewDeleteProductUC(params DeleteProductUCParams) DeleteProductUC {
	return DeleteProductUC{
		Service: params.Service,
	}
}

func (uc *DeleteProductUC) Execute() error {
	if uc.ID == nil {
		return fmt.Errorf("invalid id")
	}

	err := uc.Service.Remove(*uc.ID)
	if err != nil {
		return err
	}
	return nil
}
