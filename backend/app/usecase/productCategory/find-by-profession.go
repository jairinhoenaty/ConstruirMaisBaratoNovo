package productCategory_usecase

import (
	pkgproductCategory "construir_mais_barato/app/domain/productCategory"
	"fmt"
)

type FindByProfessionUC struct {
	Service   pkgproductCategory.ProductCategoryService
	Assembler *FindByProfessionAssembler
}

type FindByProfessionUCParams struct {
	Service pkgproductCategory.ProductCategoryService
}

func NewFindByProfessionUC(params FindByProfessionUCParams) FindByProfessionUC {
	return FindByProfessionUC{
		Service: params.Service,
	}
}

func (uc *FindByProfessionUC) Execute() ([]*ProductCategoryPresenter, error) {

	if uc.Assembler.ProfessionID == 0 {
		return nil, fmt.Errorf("invalid Profession")
	}

	productCategorys, err := uc.Service.FindByProfession(uc.Assembler.ProfessionID)
	if err != nil {
		return nil, err
	}

	productCategorysPresenter := make([]*ProductCategoryPresenter, 0)
	for _, productCategory := range productCategorys {
		productCategoryPresenter := GenerateProductCategoryPresenter(productCategory)

		productCategorysPresenter = append(productCategorysPresenter, &productCategoryPresenter)
	}

	return productCategorysPresenter, nil
}
