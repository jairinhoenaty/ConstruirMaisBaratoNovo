package productCategory_usecase

import (
	pkgproductCategory "construir_mais_barato/app/domain/productCategory"
	"fmt"
)

type CreateProductCategoryUC struct {
	Service   pkgproductCategory.ProductCategoryService
	Assembler *ProductCategoryAssembler
}

type CreateProductCategoryUCParams struct {
	Service pkgproductCategory.ProductCategoryService
}

func NewCreateProductCategoryUC(params CreateProductCategoryUCParams) CreateProductCategoryUC {
	return CreateProductCategoryUC{
		Service: params.Service,
	}
}

func (uc *CreateProductCategoryUC) Execute() (*ProductCategoryPresenter, error) {
	if uc.Assembler == nil {
		return nil, fmt.Errorf("assembler cannot be nil")
	}

	// Validar nome obrigatório
	if uc.Assembler.Name == "" {
		return nil, fmt.Errorf("category name is required")
	}

	// profession_id é opcional (pode ser nil para categorias genéricas)
	// Se fornecido, qualquer validação adicional pode ser feita aqui

	// Criar entidade
	category := pkgproductCategory.ProductCategory{
		Name:         uc.Assembler.Name,
		ProfessionID: uc.Assembler.ProfessionID,
		ParentID:     uc.Assembler.ParentID,
	}

	// Salvar no banco
	savedCategory, err := uc.Service.Save(category)
	if err != nil {
		return nil, fmt.Errorf("failed to create category: %w", err)
	}

	// Retornar presenter
	presenter := GenerateProductCategoryPresenter(savedCategory)
	return &presenter, nil
}
