package controllers

import (
	pkgsolicitationuc "construir_mais_barato/app/usecase/solicitation"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SolicitationController struct {
	SaveSolicitationUCParams pkgsolicitationuc.SaveSolicitationUCParams
}

type SolicitationControllerParams struct {
	SaveSolicitationUCParams pkgsolicitationuc.SaveSolicitationUCParams
}

func (c *SolicitationController) Save(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgsolicitationuc.NewSaveSolicitationUC(c.SaveSolicitationUCParams)
	assembler := pkgsolicitationuc.SaveSolicitationAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	usecase.Assembler = &assembler

	solicitation, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, solicitation)
}

func NewSolicitationController(params *SolicitationControllerParams, g *echo.Group) {
	controller := SolicitationController{
		SaveSolicitationUCParams: params.SaveSolicitationUCParams,
	}

	g.POST("/solicitation", controller.Save)
}
