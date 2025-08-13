package controllers

import (
	pkgcityuc "construir_mais_barato/app/usecase/city"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CityController struct {
	FindAllCityUCParams pkgcityuc.FindAllCityUCParams
	FindByIdUCParams    pkgcityuc.FindByIdUCParams
	FindByUFUCParams    pkgcityuc.FindByUFUCParams
	SaveCityUCParams    pkgcityuc.SaveCityUCParams
	DeleteCityUCParams  pkgcityuc.DeleteCityUCParams
}

type CityControllerParams struct {
	FindAllCityUCParams pkgcityuc.FindAllCityUCParams
	FindByIdUCParams    pkgcityuc.FindByIdUCParams
	FindByUFUCParams    pkgcityuc.FindByUFUCParams
	SaveCityUCParams    pkgcityuc.SaveCityUCParams
	DeleteCityUCParams  pkgcityuc.DeleteCityUCParams
}

func NewCityController(params *CityControllerParams, g *echo.Group) {
	controller := CityController{
		FindAllCityUCParams: params.FindAllCityUCParams,
		FindByIdUCParams:    params.FindByIdUCParams,
		FindByUFUCParams:    params.FindByUFUCParams,
		SaveCityUCParams:    params.SaveCityUCParams,
		DeleteCityUCParams:  params.DeleteCityUCParams,
	}

	g.POST("/city", controller.Save)
	g.GET("/cities", controller.FindAll)
	g.GET("/city/:id", controller.FindById)
	g.POST("/city-uf", controller.FindByUF)
	g.DELETE("/city/:id", controller.Delete)
}

func (c *CityController) FindAll(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgcityuc.NewFindAllCityUC(c.FindAllCityUCParams)
	cities, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, cities)
}

func (c *CityController) FindById(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgcityuc.NewFindByIdUC(c.FindByIdUCParams)
	idAssembler := ctx.Param("id")
	id, err := strconv.ParseUint(idAssembler, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	uintID := uint(id)

	usecase.ID = &uintID
	city, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, city)
}

func (c *CityController) FindByUF(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgcityuc.NewFindByUFUC(c.FindByUFUCParams)

	ufAssembler := pkgcityuc.UFCityAssembler{}
	if err := ctx.Bind(&ufAssembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	usecase.Assembler = &ufAssembler
	cities, err := usecase.Execute()

	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, cities)
}

func (c *CityController) Save(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgcityuc.NewSaveCityUC(c.SaveCityUCParams)
	cityAssembler := pkgcityuc.CityAssembler{}
	if err := ctx.Bind(&cityAssembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	usecase.Assembler = &cityAssembler

	city, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, city)

}

func (c *CityController) Delete(ctx echo.Context) error {

	defer ctx.Request().Body.Close()
	usecase := pkgcityuc.NewDeleteCityUC(c.DeleteCityUCParams)
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
