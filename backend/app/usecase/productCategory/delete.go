package productCategory_usecase

import (
	pkgproductCategory "construir_mais_barato/app/domain/productCategory"
	pkgstore "construir_mais_barato/app/domain/store"
	"fmt"
)

type DeleteProductCategoryUC struct {
	CategoryService pkgproductCategory.ProductCategoryService
	StoreService    pkgstore.StoreService
	ID              *uint
	MigrateTo       *uint // ID da categoria para migrar lojas (opcional)
}

type DeleteProductCategoryUCParams struct {
	CategoryService pkgproductCategory.ProductCategoryService
	StoreService    pkgstore.StoreService
}

type DeleteProductCategoryOutput struct {
	Deleted        bool   `json:"deleted"`
	CategoryID     uint   `json:"categoryId"`
	MigratedStores int64  `json:"migratedStores"`
	Message        string `json:"message"`
}

func NewDeleteProductCategoryUC(params DeleteProductCategoryUCParams) DeleteProductCategoryUC {
	return DeleteProductCategoryUC{
		CategoryService: params.CategoryService,
		StoreService:    params.StoreService,
	}
}

func (uc *DeleteProductCategoryUC) Execute() (*DeleteProductCategoryOutput, error) {
	if uc.ID == nil {
		return nil, fmt.Errorf("category ID is required")
	}

	// Verificar se categoria existe
	category, err := uc.CategoryService.FindById(*uc.ID)
	if err != nil {
		return nil, fmt.Errorf("category not found: %w", err)
	}

	// Verificar se tem subcategorias (children)
	children, err := uc.CategoryService.FindSubcategoriesByParentID(*uc.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to check for subcategories: %w", err)
	}

	if len(children) > 0 {
		return nil, fmt.Errorf("cannot delete category with subcategories. Please delete subcategories first")
	}

	// Verificar se há lojas usando esta categoria
	storeCount, err := uc.StoreService.CountByCategory(*uc.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to count stores: %w", err)
	}

	// Se tem lojas e não foi fornecida categoria de migração, retornar erro
	if storeCount > 0 && uc.MigrateTo == nil {
		return nil, fmt.Errorf(
			"Categoria em uso, existem %d lojista(s) usando no momento. Faça uma migração de categoria antes de remover.",
			storeCount,
		)
	}

	var migratedStores int64

	// Se tem lojas e foi fornecida categoria de migração, migrar
	if storeCount > 0 && uc.MigrateTo != nil {
		// Verificar se categoria de destino existe
		targetCategory, err := uc.CategoryService.FindById(*uc.MigrateTo)
		if err != nil {
			return nil, fmt.Errorf("migration target category not found: %w", err)
		}

		// Não permitir migração para a própria categoria
		if *uc.MigrateTo == *uc.ID {
			return nil, fmt.Errorf("cannot migrate to the same category")
		}

		// Executar migração em massa
		migratedStores, err = uc.StoreService.MigrateCategoryBulk(*uc.ID, *uc.MigrateTo)
		if err != nil {
			return nil, fmt.Errorf("failed to migrate stores: %w", err)
		}

		if migratedStores != storeCount {
			return nil, fmt.Errorf(
				"migration incomplete: expected %d stores, migrated %d",
				storeCount,
				migratedStores,
			)
		}

		// Log ou notificação (opcional - pode ser implementado futuramente)
		// TODO: Notificar lojas sobre mudança de categoria
		_ = targetCategory // Usado para verificação, poderia ser usado em notificação
	}

	// Soft delete (GORM usa deleted_at)
	err = uc.CategoryService.Remove(*uc.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to delete category: %w", err)
	}

	message := fmt.Sprintf("Category '%s' deleted successfully", category.Name)
	if migratedStores > 0 {
		message = fmt.Sprintf(
			"Category '%s' deleted successfully. %d store(s) migrated to new category",
			category.Name,
			migratedStores,
		)
	}

	return &DeleteProductCategoryOutput{
		Deleted:        true,
		CategoryID:     *uc.ID,
		MigratedStores: migratedStores,
		Message:        message,
	}, nil
}
