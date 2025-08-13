package controllers

import (
	pkgregionuc "construir_mais_barato/app/usecase/region"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RegionController struct {
	FindAllRegionUCParams pkgregionuc.FindAllRegionUCParams
	FindByIdUCParams          pkgregionuc.FindByIdUCParams
	SaveRegionUCParams    pkgregionuc.SaveRegionUCParams
	DeleteRegionUCParams  pkgregionuc.DeleteRegionUCParams
}

type RegionControllerParams struct {
	FindAllRegionUCParams pkgregionuc.FindAllRegionUCParams
	FindByIdUCParams          pkgregionuc.FindByIdUCParams
	SaveRegionUCParams    pkgregionuc.SaveRegionUCParams
	DeleteRegionUCParams  pkgregionuc.DeleteRegionUCParams
}

func NewRegionController(params *RegionControllerParams, g *echo.Group) {
	controller := RegionController{
		FindAllRegionUCParams: params.FindAllRegionUCParams,
		FindByIdUCParams:          params.FindByIdUCParams,
		SaveRegionUCParams:    params.SaveRegionUCParams,
		DeleteRegionUCParams:  params.DeleteRegionUCParams,
	}

	g.POST("/region", controller.Save)
	g.GET("/regions", controller.FindAll)
	g.GET("/region/:id", controller.FindById)
	g.DELETE("/region/:id", controller.Delete)
}

func (c *RegionController) FindAll(ctx echo.Context) error {
	defer ctx.Request().Body.Close()

	assembler := pkgregionuc.FindWithPaginationRegionAssembler{}

	// Extrair parâmetros de paginação da query string
	limit := ctx.QueryParam("limit")
	offset := ctx.QueryParam("offset")
	uf := ctx.QueryParam("uf")

	// Converter os parâmetros para inteiros, com valores padrão se não forem fornecidos
	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt <= 0 {
		limitInt = 20 // valor padrão
	}

	offsetInt, err := strconv.Atoi(offset)
	if err != nil || offsetInt < 0 {
		offsetInt = 0 // valor padrão
	}

	// Definir os parâmetros de paginação no assembler
	assembler.Limit = limitInt
	assembler.Offset = offsetInt
	assembler.UF = uf

	usecase := pkgregionuc.NewFindAllRegionUC(c.FindAllRegionUCParams)
	usecase.Assembler = assembler
	regions, total, err := usecase.Execute()
	if err != nil {
		log.Printf("Erro ao executar usecase.Execute(): %v", err)
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	log.Printf("Profissões encontradas: %d", len(*regions))

	// Estrutura de resposta
	response := struct {
		Regions *[]pkgregionuc.RegionPresenter `json:"regions"`
		Total   int64                      `json:"total"` 
	}{
		Regions: regions,
		Total:       total,
	}

	return ctx.JSON(http.StatusOK, response)
}

func (c *RegionController) FindById(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgregionuc.NewFindByIdUC(c.FindByIdUCParams)
	idAssembler := ctx.Param("id")
	id, err := strconv.ParseUint(idAssembler, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	uintID := uint(id)

	usecase.ID = &uintID
	region, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, region)
}

func (c *RegionController) Save(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgregionuc.NewSaveRegionUC(c.SaveRegionUCParams)
	regionAssembler := pkgregionuc.RegionAssembler{}
	if err := ctx.Bind(&regionAssembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	usecase.Assembler = &regionAssembler

	region, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, region)

}

func (c *RegionController) Delete(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgregionuc.NewDeleteRegionUC(c.DeleteRegionUCParams)
	idAssembler := ctx.Param("id")
	id, err := strconv.ParseUint(idAssembler, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	uintID := uint(id)

	usecase.ID = &uintID
	err = usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, nil)

}
