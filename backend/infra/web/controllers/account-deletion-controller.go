package controllers

import (
	pkgaccountdeletionuc "construir_mais_barato/app/usecase/accountDeletion"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AccountDeletionController struct {
	CreateAccountDeletionUCParams pkgaccountdeletionuc.CreateAccountDeletionUCParams
}

type AccountDeletionControllerParams struct {
	CreateAccountDeletionUCParams pkgaccountdeletionuc.CreateAccountDeletionUCParams
}

func NewPublicAccountDeletionController(
	params *AccountDeletionControllerParams,
	g *echo.Group,
) {
	controller := AccountDeletionController{
		CreateAccountDeletionUCParams: params.CreateAccountDeletionUCParams,
	}

	g.POST("/account-deletion-request", controller.Create)
}

func (c *AccountDeletionController) Create(ctx echo.Context) error {
	defer ctx.Request().Body.Close()

	var assembler pkgaccountdeletionuc.CreateAccountDeletionAssembler

	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			map[string]string{
				"error": "Dados da solicitação inválidos.",
			},
		)
	}

	usecase := pkgaccountdeletionuc.NewCreateAccountDeletionUC(
		c.CreateAccountDeletionUCParams,
	)

	usecase.Assembler = &assembler

	request, err := usecase.Execute()

	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			map[string]string{
				"error": err.Error(),
			},
		)
	}

	return ctx.JSON(
		http.StatusCreated,
		map[string]interface{}{
			"message": "Recebemos sua solicitação de exclusão. Nossa equipe irá processá-la.",
			"request": request,
		},
	)
}
