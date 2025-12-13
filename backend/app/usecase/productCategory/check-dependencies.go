package productCategory_usecase

import (
	pkgproductCategory "construir_mais_barato/app/domain/productCategory"
	pkgstore "construir_mais_barato/app/domain/store"
	"fmt"
)

type CheckDependenciesUC struct {
	CategoryService pkgproductCategory.ProductCategoryService
	StoreService    pkgstore.StoreService
	ID              *uint
}

type CheckDependenciesUCParams struct {
	CategoryService pkgproductCategory.ProductCategoryService
	StoreService    pkgstore.StoreService
}

type CheckDependenciesOutput struct {
	CategoryID              uint                           `json:"categoryId"`
	CategoryName            string                         `json:"categoryName"`
	HasDependencies         bool                           `json:"hasDependencies"`
	StoreCount              int64                          `json:"storeCount"`
	SubcategoryCount        int                            `json:"subcategoryCount"`
	CanDelete               bool                           `json:"canDelete"`
	RequiresMigration       bool                           `json:"requiresMigration"`
	SuggestedMigrationPaths []SuggestedMigrationPath       `json:"suggestedMigrationPaths,omitempty"`
	Message                 string                         `json:"message"`
}

type SuggestedMigrationPath struct {
	CategoryID   uint   `json:"categoryId"`
	CategoryName string `json:"categoryName"`
	ParentID     *uint  `json:"parentId,omitempty"`
	Level        string `json:"level"` // "same_level", "parent", "sibling"
}

func NewCheckDependenciesUC(params CheckDependenciesUCParams) CheckDependenciesUC {
	return CheckDependenciesUC{
		CategoryService: params.CategoryService,
		StoreService:    params.StoreService,
	}
}

func (uc *CheckDependenciesUC) Execute() (*CheckDependenciesOutput, error) {
	if uc.ID == nil {
		return nil, fmt.Errorf("category ID is required")
	}

	// Buscar categoria
	category, err := uc.CategoryService.FindById(*uc.ID)
	if err != nil {
		return nil, fmt.Errorf("category not found: %w", err)
	}

	output := &CheckDependenciesOutput{
		CategoryID:   category.ID,
		CategoryName: category.Name,
	}

	// Verificar subcategorias
	subcategories, err := uc.CategoryService.FindSubcategoriesByParentID(*uc.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to check subcategories: %w", err)
	}
	output.SubcategoryCount = len(subcategories)

	// Verificar lojas que usam esta categoria ou subcategoria
	storeCount, err := uc.StoreService.CountByCategory(*uc.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to count stores: %w", err)
	}
	output.StoreCount = storeCount

	// Determinar se tem dependências
	output.HasDependencies = output.SubcategoryCount > 0 || output.StoreCount > 0

	// Determinar se pode deletar
	output.CanDelete = !output.HasDependencies
	output.RequiresMigration = output.StoreCount > 0

	// Gerar mensagem
	if !output.HasDependencies {
		output.Message = "Esta categoria pode ser removida sem problemas"
	} else if output.SubcategoryCount > 0 && output.StoreCount > 0 {
		output.Message = fmt.Sprintf(
			"Esta categoria possui %d subcategoria(s) e %d loja(s) vinculada(s). "+
				"Remova as subcategorias primeiro e migre as lojas para outra categoria",
			output.SubcategoryCount,
			output.StoreCount,
		)
	} else if output.SubcategoryCount > 0 {
		output.Message = fmt.Sprintf(
			"Esta categoria possui %d subcategoria(s). Remova as subcategorias primeiro",
			output.SubcategoryCount,
		)
	} else if output.StoreCount > 0 {
		output.Message = fmt.Sprintf(
			"Esta categoria possui %d loja(s) vinculada(s). Escolha uma categoria para migrar as lojas",
			output.StoreCount,
		)
		// Buscar sugestões de migração
		output.SuggestedMigrationPaths = uc.getSuggestedMigrationPaths(category)
	}

	return output, nil
}

func (uc *CheckDependenciesUC) getSuggestedMigrationPaths(category *pkgproductCategory.ProductCategory) []SuggestedMigrationPath {
	suggestions := []SuggestedMigrationPath{}

	// Sugestão 1: Categoria pai (se existir)
	if category.ParentID != nil {
		parent, err := uc.CategoryService.FindById(*category.ParentID)
		if err == nil {
			suggestions = append(suggestions, SuggestedMigrationPath{
				CategoryID:   parent.ID,
				CategoryName: parent.Name,
				ParentID:     parent.ParentID,
				Level:        "parent",
			})
		}
	}

	// Sugestão 2: Categorias irmãs (mesmo parent_id)
	var siblings []*pkgproductCategory.ProductCategory
	var err error

	if category.ParentID != nil {
		// Buscar subcategorias do mesmo pai
		siblings, err = uc.CategoryService.FindSubcategoriesByParentID(*category.ParentID)
	} else {
		// Buscar outras categorias de nível raiz
		allTopLevel, err2 := uc.CategoryService.FindTopLevel()
		if err2 == nil {
			siblings = allTopLevel
		}
		err = err2
	}

	if err == nil {
		for _, sibling := range siblings {
			// Não sugerir a própria categoria
			if sibling.ID == category.ID {
				continue
			}

			suggestions = append(suggestions, SuggestedMigrationPath{
				CategoryID:   sibling.ID,
				CategoryName: sibling.Name,
				ParentID:     sibling.ParentID,
				Level:        "sibling",
			})
		}
	}

	return suggestions
}
