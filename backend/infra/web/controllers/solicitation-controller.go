package controllers

import (
	pkgsolicitationuc "construir_mais_barato/app/usecase/solicitation"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SolicitationController struct {
	SaveSolicitationUCParams pkgsolicitationuc.SaveSolicitationUCParams
	UpdateFeedbackUCParams   pkgsolicitationuc.UpdateFeedbackUCParams
}

type SolicitationControllerParams struct {
	SaveSolicitationUCParams pkgsolicitationuc.SaveSolicitationUCParams
	UpdateFeedbackUCParams   pkgsolicitationuc.UpdateFeedbackUCParams
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

func (c *SolicitationController) Feedback(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgsolicitationuc.NewUpdateFeedbackUC(c.UpdateFeedbackUCParams)
	assembler := pkgsolicitationuc.UpdateFeedbackAssembler{}
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
		UpdateFeedbackUCParams:   params.UpdateFeedbackUCParams,
	}

	g.POST("/solicitation", controller.Save)
	g.PATCH("/solicitation/feedback", controller.Feedback)
}
