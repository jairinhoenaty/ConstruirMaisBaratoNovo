package controllers

import (
	pkguseruc "construir_mais_barato/app/usecase/user"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	FindAllUserUCParams pkguseruc.FindAllUserUCParams
	FindByIdUCParams    pkguseruc.FindByIdUCParamns
	FindByEmailUCParams pkguseruc.FindByEmailUCParams
	SaveUserUCParams    pkguseruc.SaveUserUCParams
	DeleteUserUCParams  pkguseruc.DeleteUserUCParams
}

type UserControllerParams struct {
	FindAllUserUCParams pkguseruc.FindAllUserUCParams
	FindByIdUCParams    pkguseruc.FindByIdUCParamns
	FindByEmailUCParams pkguseruc.FindByEmailUCParams
	SaveUserUCParams    pkguseruc.SaveUserUCParams
	DeleteUserUCParams  pkguseruc.DeleteUserUCParams
}

func NewUserController(params *UserControllerParams, g *echo.Group) {

	controller := UserController{
		FindByIdUCParams:    params.FindByIdUCParams,
		SaveUserUCParams:    params.SaveUserUCParams,
		DeleteUserUCParams:  params.DeleteUserUCParams,
		FindAllUserUCParams: params.FindAllUserUCParams,
		FindByEmailUCParams: params.FindByEmailUCParams,
	}

	g.GET("/users", controller.FindAll)
	g.GET("/user/:id", controller.FindById)
	g.POST("/user", controller.Save)
	g.POST("/find-by-email", controller.FindByEmail)
	g.DELETE("/user/:id", controller.Delete)

}

func (c *UserController) FindByEmail(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	var request struct {
		Email string `json:"email"`
	}
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request format"})
	}

	usecase := pkguseruc.NewFindByEmailUC(c.FindByEmailUCParams)
	usecase.Email = &request.Email

	user, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, map[string]string{"error": "E-mail não encontrado"})
	}
	return ctx.JSON(http.StatusOK, user)

}

func (c *UserController) FindAll(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkguseruc.NewFindAllUserUC(c.FindAllUserUCParams)
	users, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, users)
}

func (c *UserController) FindById(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkguseruc.NewFindByIdUC(c.FindByIdUCParams)
	idAssembler := ctx.Param("id")
	id, err := strconv.ParseUint(idAssembler, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	uintID := uint(id)

	usecase.ID = &uintID
	user, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, user)
}

func (c *UserController) Save(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkguseruc.NewSaveUserUC(c.SaveUserUCParams)
	userAssembler := pkguseruc.UserAssembler{}
	if err := ctx.Bind(&userAssembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	usecase.Assembler = &userAssembler

	user, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, user)

}

func (c *UserController) Delete(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkguseruc.NewDeleteUserUC(c.DeleteUserUCParams)
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
