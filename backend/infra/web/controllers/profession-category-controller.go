package controllers

import (
	"net/http"
	"strconv"

	pkgprofessionCategoryuc "construir_mais_barato/app/usecase/professionCategory"

	"github.com/labstack/echo/v4"
)

type ProfessionCategoryController struct {
	FindAllUCParams         pkgprofessionCategoryuc.FindAllUCParams
	FindActiveUCParams      pkgprofessionCategoryuc.FindActiveUCParams
	FindByIdUCParams        pkgprofessionCategoryuc.FindByIdUCParams
	FindProfessionsUCParams pkgprofessionCategoryuc.FindProfessionsUCParams
	SaveUCParams            pkgprofessionCategoryuc.SaveUCParams
	DeleteUCParams          pkgprofessionCategoryuc.DeleteUCParams
}

type ProfessionCategoryControllerParams struct {
	FindAllUCParams         pkgprofessionCategoryuc.FindAllUCParams
	FindActiveUCParams      pkgprofessionCategoryuc.FindActiveUCParams
	FindByIdUCParams        pkgprofessionCategoryuc.FindByIdUCParams
	FindProfessionsUCParams pkgprofessionCategoryuc.FindProfessionsUCParams
	SaveUCParams            pkgprofessionCategoryuc.SaveUCParams
	DeleteUCParams          pkgprofessionCategoryuc.DeleteUCParams
}

func newProfessionCategoryController(params *ProfessionCategoryControllerParams) ProfessionCategoryController {
	return ProfessionCategoryController{
		FindAllUCParams:         params.FindAllUCParams,
		FindActiveUCParams:      params.FindActiveUCParams,
		FindByIdUCParams:        params.FindByIdUCParams,
		FindProfessionsUCParams: params.FindProfessionsUCParams,
		SaveUCParams:            params.SaveUCParams,
		DeleteUCParams:          params.DeleteUCParams,
	}
}

func NewProfessionCategoryPublicController(params *ProfessionCategoryControllerParams, g *echo.Group) {
	controller := newProfessionCategoryController(params)

	g.GET("/profession-categories", controller.FindActive)
	g.GET("/profession-categories/:id/professions", controller.FindProfessions)
}

func NewProfessionCategoryController(params *ProfessionCategoryControllerParams, g *echo.Group) {
	controller := newProfessionCategoryController(params)

	g.GET("/profession-categories", controller.FindAll)
	g.GET("/profession-category/:id", controller.FindById)
	g.GET("/profession-category/:id/professions", controller.FindProfessions)
	g.POST("/profession-category", controller.Save)
	g.DELETE("/profession-category/:id", controller.Delete)
}

func (c *ProfessionCategoryController) FindActive(ctx echo.Context) error {
	defer ctx.Request().Body.Close()

	usecase := pkgprofessionCategoryuc.NewFindActiveUC(c.FindActiveUCParams)
	categories, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	return ctx.JSON(http.StatusOK, categories)
}

func (c *ProfessionCategoryController) FindAll(ctx echo.Context) error {
	defer ctx.Request().Body.Close()

	limit, err := strconv.Atoi(ctx.QueryParam("limit"))
	if err != nil || limit <= 0 {
		limit = 20
	}
	offset, err := strconv.Atoi(ctx.QueryParam("offset"))
	if err != nil || offset < 0 {
		offset = 0
	}

	usecase := pkgprofessionCategoryuc.NewFindAllUC(c.FindAllUCParams)
	usecase.Assembler = pkgprofessionCategoryuc.FindWithPaginationAssembler{
		Limit:  limit,
		Offset: offset,
	}

	categories, total, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}

	return ctx.JSON(http.StatusOK, struct {
		Categories []pkgprofessionCategoryuc.ProfessionCategoryPresenter `json:"categorias"`
		Total      int64                                                 `json:"total"`
	}{
		Categories: categories,
		Total:      total,
	})
}

func (c *ProfessionCategoryController) FindById(ctx echo.Context) error {
	defer ctx.Request().Body.Close()

	id, err := parseUintParam(ctx, "id")
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	usecase := pkgprofessionCategoryuc.NewFindByIdUC(c.FindByIdUCParams)
	usecase.ID = &id

	category, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusNotFound, nil)
	}

	return ctx.JSON(http.StatusOK, category)
}

func (c *ProfessionCategoryController) FindProfessions(ctx echo.Context) error {
	defer ctx.Request().Body.Close()

	id, err := parseUintParam(ctx, "id")
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	usecase := pkgprofessionCategoryuc.NewFindProfessionsUC(c.FindProfessionsUCParams)
	usecase.CategoryID = &id

	category, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusNotFound, nil)
	}

	return ctx.JSON(http.StatusOK, category)
}

func (c *ProfessionCategoryController) Save(ctx echo.Context) error {
	defer ctx.Request().Body.Close()

	assembler := pkgprofessionCategoryuc.ProfessionCategoryAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	usecase := pkgprofessionCategoryuc.NewSaveUC(c.SaveUCParams)
	usecase.Assembler = &assembler

	category, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}

	return ctx.JSON(http.StatusOK, category)
}

func (c *ProfessionCategoryController) Delete(ctx echo.Context) error {
	defer ctx.Request().Body.Close()

	id, err := parseUintParam(ctx, "id")
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	usecase := pkgprofessionCategoryuc.NewDeleteUC(c.DeleteUCParams)
	usecase.ID = &id

	if err := usecase.Execute(); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, nil)
}

func parseUintParam(ctx echo.Context, name string) (uint, error) {
	value, err := strconv.ParseUint(ctx.Param(name), 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(value), nil
}
