package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	pkgauthuc "construir_mais_barato/app/usecase/auth"
)

type AuthController struct {
	AuthenticateUCParams pkgauthuc.AuthenticateUCParams
	ExchangeCodeUCParams pkgauthuc.ExchangeCodeUCParams
	RedeemCodeUCParams   pkgauthuc.RedeemCodeUCParams
}

type AuthControllerParams struct {
	AuthenticateUCParams pkgauthuc.AuthenticateUCParams
	ExchangeCodeUCParams pkgauthuc.ExchangeCodeUCParams
	RedeemCodeUCParams   pkgauthuc.RedeemCodeUCParams
}

func NewAuthController(params AuthControllerParams, g *echo.Group) {
	controller := AuthController{
		AuthenticateUCParams: params.AuthenticateUCParams,
		RedeemCodeUCParams:   params.RedeemCodeUCParams,
	}

	g.POST("/login", controller.Login)
	g.POST("/validalogin", controller.ValidaLogin)
	g.POST("/redeem-code", controller.RedeemCode)
}

func NewAuthAuthenticatedController(params AuthControllerParams, g *echo.Group) {
	controller := AuthController{
		ExchangeCodeUCParams: params.ExchangeCodeUCParams,
	}

	g.POST("/exchange-code", controller.ExchangeCode)
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

func (c *AuthController) ExchangeCode(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	// parse do corpo da requisição para obter as credenciais
	fmt.Println("parse do corpo da requisição para obter as credenciais")
	//fmt.Println(ctx.Request().Body);
	assembler := pkgauthuc.ExchangeCodeAssembler{}

	if err := ctx.Bind(&assembler); err != nil {
		fmt.Println("Erro no parse do assembler => ", err)
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	usecase := pkgauthuc.NewExchangeCode(c.ExchangeCodeUCParams)

	usecase.Assembler = &assembler

	result, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	return ctx.JSON(http.StatusOK, result)

}
func (c *AuthController) RedeemCode(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	// parse do corpo da requisição para obter as credenciais
	fmt.Println("parse do corpo da requisição para obter as credenciais")
	//fmt.Println(ctx.Request().Body);
	assembler := pkgauthuc.RedeemCodeAssembler{}

	if err := ctx.Bind(&assembler); err != nil {
		fmt.Println("Erro no parse do assembler => ", err)
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	usecase := pkgauthuc.NewRedeemCode(c.RedeemCodeUCParams)

	usecase.Assembler = &assembler

	result, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	return ctx.JSON(http.StatusOK, result)

}
