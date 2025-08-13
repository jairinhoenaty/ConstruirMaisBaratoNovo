package controllers

import (
	pkgprofessionuc "construir_mais_barato/app/usecase/profession"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProfessionController struct {
	FindAllProfessionUCParams pkgprofessionuc.FindAllProfessionUCParams
	FindByIdUCParams          pkgprofessionuc.FindByIdUCParams
	SaveProfessionUCParams    pkgprofessionuc.SaveProfessionUCParams
	DeleteProfessionUCParams  pkgprofessionuc.DeleteProfessionUCParams
}

type ProfessionControllerParams struct {
	FindAllProfessionUCParams pkgprofessionuc.FindAllProfessionUCParams
	FindByIdUCParams          pkgprofessionuc.FindByIdUCParams
	SaveProfessionUCParams    pkgprofessionuc.SaveProfessionUCParams
	DeleteProfessionUCParams  pkgprofessionuc.DeleteProfessionUCParams
}

func NewProfessionController(params *ProfessionControllerParams, g *echo.Group) {
	controller := ProfessionController{
		FindAllProfessionUCParams: params.FindAllProfessionUCParams,
		FindByIdUCParams:          params.FindByIdUCParams,
		SaveProfessionUCParams:    params.SaveProfessionUCParams,
		DeleteProfessionUCParams:  params.DeleteProfessionUCParams,
	}

	g.POST("/profession", controller.Save)
	g.GET("/professions", controller.FindAll)
	g.GET("/profession/:id", controller.FindById)
	g.DELETE("/profession/:id", controller.Delete)
}

func (c *ProfessionController) FindAll(ctx echo.Context) error {
	defer ctx.Request().Body.Close()

	assembler := pkgprofessionuc.FindWithPaginationProfessionAssembler{}

	// Extrair parâmetros de paginação da query string
	limit := ctx.QueryParam("limit")
	offset := ctx.QueryParam("offset")

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

	usecase := pkgprofessionuc.NewFindAllProfessionUC(c.FindAllProfessionUCParams)
	usecase.Assembler = assembler
	professions, total, err := usecase.Execute()
	if err != nil {
		log.Printf("Erro ao executar usecase.Execute(): %v", err)
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	log.Printf("Profissões encontradas: %d", len(*professions))

	// Estrutura de resposta
	response := struct {
		Professions *[]pkgprofessionuc.ProfessionPresenter `json:"profissoes"`
		Total       int64                                  `json:"total"`
	}{
		Professions: professions,
		Total:       total,
	}

	return ctx.JSON(http.StatusOK, response)
}

func (c *ProfessionController) FindById(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgprofessionuc.NewFindByIdUC(c.FindByIdUCParams)
	idAssembler := ctx.Param("id")
	id, err := strconv.ParseUint(idAssembler, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	uintID := uint(id)

	usecase.ID = &uintID
	profession, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, profession)
}

func (c *ProfessionController) Save(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgprofessionuc.NewSaveProfessionUC(c.SaveProfessionUCParams)
	professionAssembler := pkgprofessionuc.ProfessionAssembler{}
	if err := ctx.Bind(&professionAssembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	usecase.Assembler = &professionAssembler

	profession, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, profession)

}

func (c *ProfessionController) Delete(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgprofessionuc.NewDeleteProfessionUC(c.DeleteProfessionUCParams)
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
