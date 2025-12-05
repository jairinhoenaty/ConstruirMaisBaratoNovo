package productCategory_usecase

import (
	pkgproductCategory "construir_mais_barato/app/domain/productCategory"
	"fmt"
)

type UpdateProductCategoryUC struct {
	Service   pkgproductCategory.ProductCategoryService
	ID        *uint
	Assembler *ProductCategoryAssembler
}

type UpdateProductCategoryUCParams struct {
	Service pkgproductCategory.ProductCategoryService
}

func NewUpdateProductCategoryUC(params UpdateProductCategoryUCParams) UpdateProductCategoryUC {
	return UpdateProductCategoryUC{
		Service: params.Service,
	}
}

func (uc *UpdateProductCategoryUC) Execute() (*ProductCategoryPresenter, error) {
	if uc.ID == nil {
		return nil, fmt.Errorf("category ID is required")
	}

	if uc.Assembler == nil {
		return nil, fmt.Errorf("assembler cannot be nil")
	}

	// Validar nome obrigatório
	if uc.Assembler.Name == "" {
		return nil, fmt.Errorf("category name is required")
	}

	// Buscar categoria existente
	existingCategory, err := uc.Service.FindById(*uc.ID)
	if err != nil {
		return nil, fmt.Errorf("category not found: %w", err)
	}

	// Atualizar campos
	existingCategory.Name = uc.Assembler.Name
	existingCategory.ProfessionID = uc.Assembler.ProfessionID
	existingCategory.ParentID = uc.Assembler.ParentID

	// Salvar alterações
	updatedCategory, err := uc.Service.Save(*existingCategory)
	if err != nil {
		return nil, fmt.Errorf("failed to update category: %w", err)
	}

	// Retornar presenter
	presenter := GenerateProductCategoryPresenter(updatedCategory)
	return &presenter, nil
}
