package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	pkgauthuc "construir_mais_barato/app/usecase/auth"
)

type AuthController struct {
	AuthenticateUCParams pkgauthuc.AuthenticateUCParams
}

type AuthControllerParams struct {
	AuthenticateUCParams pkgauthuc.AuthenticateUCParams
}

func NewAuthController(params AuthControllerParams, g *echo.Group) {
	controller := AuthController{
		AuthenticateUCParams: params.AuthenticateUCParams,
	}

	g.POST("/login", controller.Login)
	g.POST("/validalogin", controller.ValidaLogin)
}

func (c *AuthController) Login(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	// parse do corpo da requisição para obter as credenciais
	fmt.Println("parse do corpo da requisição para obter as credenciais")
	//fmt.Println(ctx.Request().Body);
	assembler := pkgauthuc.LoginAssembler{}

	if err := ctx.Bind(&assembler); err != nil {
		fmt.Println("Erro no parse do assembler => ", err)
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	usecase := pkgauthuc.NewLoginUC(c.AuthenticateUCParams)

	usecase.Assembler = &assembler

	result, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	return ctx.JSON(http.StatusOK, result)

}

func (c *AuthController) ValidaLogin(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	// parse do corpo da requisição para obter as credenciais
	fmt.Println("parse do corpo da requisição para obter as credenciais")
	//fmt.Println(ctx.Request().Body);
	assembler := pkgauthuc.ValidaLoginAssembler{}

	if err := ctx.Bind(&assembler); err != nil {
		fmt.Println("Erro no parse do assembler => ", err)
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	/* usecase := pkgauthuc.NewLoginUC(c.AuthenticateUCParams)

	usecase.Assembler = &assembler

	result, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}*

	return ctx.JSON(http.StatusOK, result)
	*/
	return ctx.JSON(http.StatusOK, nil)
}
