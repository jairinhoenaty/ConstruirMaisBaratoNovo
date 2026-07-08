package controllers

import (
	pkgpageviewuc "construir_mais_barato/app/usecase/pageview"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PageViewController struct {
	FindAllPageViewsUCParams pkgpageviewuc.FindAllPageViewsUCParams
}

type PageViewControllerParams struct {
	FindAllPageViewsUCParams pkgpageviewuc.FindAllPageViewsUCParams
}

func NewPageViewController(params *PageViewControllerParams, g *echo.Group) {
	controller := PageViewController{
		FindAllPageViewsUCParams: params.FindAllPageViewsUCParams,
	}

	g.GET("/page-views", controller.FindAll)
}

func (c *PageViewController) FindAll(ctx echo.Context) error {
	usecase := pkgpageviewuc.NewFindAllPageViewsUC(c.FindAllPageViewsUCParams)
	pageViews, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, pageViews)
}
