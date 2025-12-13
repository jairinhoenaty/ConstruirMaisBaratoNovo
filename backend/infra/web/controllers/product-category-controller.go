package controllers

import (
	"net/http"
	"strconv"

	pkgproductCategoryuc "construir_mais_barato/app/usecase/productCategory"
	"github.com/labstack/echo/v4"
)

type ProductCategoryController struct {
	FindByProfessionUCParams           pkgproductCategoryuc.FindByProfessionUCParams
	SaveProductCategoryUCParams        pkgproductCategoryuc.SaveProductCategoryUCParams
	FindTopLevelUCParams               pkgproductCategoryuc.FindTopLevelUCParams
	FindSubcategoriesByParentUCParams  pkgproductCategoryuc.FindSubcategoriesByParentUCParams
	CreateProductCategoryUCParams      pkgproductCategoryuc.CreateProductCategoryUCParams
	UpdateProductCategoryUCParams      pkgproductCategoryuc.UpdateProductCategoryUCParams
	DeleteProductCategoryUCParams      pkgproductCategoryuc.DeleteProductCategoryUCParams
	CheckDependenciesUCParams          pkgproductCategoryuc.CheckDependenciesUCParams
}

type ProductCategoryControllerParams struct {
	FindByProfessionUCParams           pkgproductCategoryuc.FindByProfessionUCParams
	SaveProductCategoryUCParams        pkgproductCategoryuc.SaveProductCategoryUCParams
	FindTopLevelUCParams               pkgproductCategoryuc.FindTopLevelUCParams
	FindSubcategoriesByParentUCParams  pkgproductCategoryuc.FindSubcategoriesByParentUCParams
	CreateProductCategoryUCParams      pkgproductCategoryuc.CreateProductCategoryUCParams
	UpdateProductCategoryUCParams      pkgproductCategoryuc.UpdateProductCategoryUCParams
	DeleteProductCategoryUCParams      pkgproductCategoryuc.DeleteProductCategoryUCParams
	CheckDependenciesUCParams          pkgproductCategoryuc.CheckDependenciesUCParams
}

func NewProductCategoryController(params *ProductCategoryControllerParams, g *echo.Group) {
	controller := ProductCategoryController{
		FindByProfessionUCParams:          params.FindByProfessionUCParams,
		SaveProductCategoryUCParams:       params.SaveProductCategoryUCParams,
		FindTopLevelUCParams:              params.FindTopLevelUCParams,
		FindSubcategoriesByParentUCParams: params.FindSubcategoriesByParentUCParams,
		CreateProductCategoryUCParams:     params.CreateProductCategoryUCParams,
		UpdateProductCategoryUCParams:     params.UpdateProductCategoryUCParams,
		DeleteProductCategoryUCParams:     params.DeleteProductCategoryUCParams,
	}

	// Rotas antigas (mantidas para compatibilidade)
	g.GET("/product_category/:profession_id", controller.FindByProfession)
	g.POST("/product_category", controller.Save)

	// Rotas novas (hierarquia de categorias)
	g.GET("/categories/top-level", controller.FindTopLevel)
	g.GET("/categories/:categoryId/subcategories", controller.FindSubcategoriesByParent)
}

func NewProductCategoryAdminController(params *ProductCategoryControllerParams, g *echo.Group) {
	controller := ProductCategoryController{
		CreateProductCategoryUCParams:     params.CreateProductCategoryUCParams,
		UpdateProductCategoryUCParams:     params.UpdateProductCategoryUCParams,
		DeleteProductCategoryUCParams:     params.DeleteProductCategoryUCParams,
		CheckDependenciesUCParams:         params.CheckDependenciesUCParams,
	}

	// Rotas admin (requerem autenticação)
	g.POST("/categories", controller.CreateCategory)
	g.PUT("/categories/:id", controller.UpdateCategory)
	g.GET("/categories/:id/dependencies", controller.CheckDependencies)
	g.DELETE("/categories/:id", controller.DeleteCategory)
}

func (c *ProductCategoryController) FindTopLevel(ctx echo.Context) error {
	usecase := pkgproductCategoryuc.NewFindTopLevelUC(c.FindTopLevelUCParams)
	categories, err := usecase.Execute()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch top level categories",
		})
	}

	return ctx.JSON(http.StatusOK, categories)
}

func (c *ProductCategoryController) FindSubcategoriesByParent(ctx echo.Context) error {
	categoryIDStr := ctx.Param("categoryId")
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid category ID",
		})
	}

	categoryIDUint := uint(categoryID)
	usecase := pkgproductCategoryuc.NewFindSubcategoriesByParentUC(c.FindSubcategoriesByParentUCParams)
	usecase.ParentID = &categoryIDUint

	subcategories, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch subcategories",
		})
	}

	return ctx.JSON(http.StatusOK, subcategories)
}

// Métodos antigos (compatibilidade)

func (c *ProductCategoryController) FindByProfession(ctx echo.Context) error {
	defer ctx.Request().Body.Close()

	assembler := pkgproductCategoryuc.FindByProfessionAssembler{}
	usecase := pkgproductCategoryuc.NewFindByProfessionUC(c.FindByProfessionUCParams)
	idAssembler := ctx.Param("profession_id")

	id, err := strconv.Atoi(idAssembler)
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	assembler.ProfessionID = id

	usecase.Assembler = &assembler
	categories, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, categories)
}

func (c *ProductCategoryController) Save(ctx echo.Context) error {
	defer ctx.Request().Body.Close()

	usecase := pkgproductCategoryuc.NewSaveProductCategoryUC(c.SaveProductCategoryUCParams)
	productCategoryAssembler := pkgproductCategoryuc.ProductCategoryAssembler{}
	if err := ctx.Bind(&productCategoryAssembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	usecase.Assembler = &productCategoryAssembler
	productCategory, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, productCategory)
}

// Admin endpoints

func (c *ProductCategoryController) CreateCategory(ctx echo.Context) error {
	defer ctx.Request().Body.Close()

	assembler := pkgproductCategoryuc.ProductCategoryAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	usecase := pkgproductCategoryuc.NewCreateProductCategoryUC(c.CreateProductCategoryUCParams)
	usecase.Assembler = &assembler

	category, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, category)
}

func (c *ProductCategoryController) UpdateCategory(ctx echo.Context) error {
	defer ctx.Request().Body.Close()

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid category ID",
		})
	}

	assembler := pkgproductCategoryuc.ProductCategoryAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	idUint := uint(id)
	usecase := pkgproductCategoryuc.NewUpdateProductCategoryUC(c.UpdateProductCategoryUCParams)
	usecase.ID = &idUint
	usecase.Assembler = &assembler

	category, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, category)
}

func (c *ProductCategoryController) CheckDependencies(ctx echo.Context) error {
	defer ctx.Request().Body.Close()

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid category ID",
		})
	}

	idUint := uint(id)
	usecase := pkgproductCategoryuc.NewCheckDependenciesUC(c.CheckDependenciesUCParams)
	usecase.ID = &idUint

	result, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, result)
}

func (c *ProductCategoryController) DeleteCategory(ctx echo.Context) error {
	defer ctx.Request().Body.Close()

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid category ID",
		})
	}

	// Pegar query param opcional migrate_to
	migrateToStr := ctx.QueryParam("migrate_to")
	var migrateTo *uint
	if migrateToStr != "" {
		migrateToID, err := strconv.ParseUint(migrateToStr, 10, 32)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid migrate_to ID",
			})
		}
		migrateToUint := uint(migrateToID)
		migrateTo = &migrateToUint
	}

	idUint := uint(id)
	usecase := pkgproductCategoryuc.NewDeleteProductCategoryUC(c.DeleteProductCategoryUCParams)
	usecase.ID = &idUint
	usecase.MigrateTo = migrateTo

	result, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, result)
}
