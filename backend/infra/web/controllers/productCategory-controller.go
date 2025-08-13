package controllers

import (
	pkgpproductCategoryuc "construir_mais_barato/app/usecase/productCategory"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductCategoryController struct {
	FindByProfessionUCParams    pkgpproductCategoryuc.FindByProfessionUCParams
	SaveProductCategoryUCParams pkgpproductCategoryuc.SaveProductCategoryUCParams
}

type ProductCategoryControllerParams struct {
	FindByProfessionUCParams    pkgpproductCategoryuc.FindByProfessionUCParams
	SaveProductCategoryUCParams pkgpproductCategoryuc.SaveProductCategoryUCParams
}

func NewProductCategoryController(params *ProductCategoryControllerParams, g *echo.Group) {
	controller := ProductCategoryController{
		FindByProfessionUCParams:    params.FindByProfessionUCParams,
		SaveProductCategoryUCParams: params.SaveProductCategoryUCParams,
	}

	g.GET("/product_category/:profession_id", controller.FindByProfession)
	g.POST("/product_category", controller.Save)
}

func (c *ProductCategoryController) FindByProfession(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	fmt.Println("Controller")
	assembler := pkgpproductCategoryuc.FindByProfessionAssembler{}
	usecase := pkgpproductCategoryuc.NewFindByProfessionUC(c.FindByProfessionUCParams)
	idAssembler := ctx.Param("profession_id")
	
	id, err := strconv.Atoi(idAssembler)
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	assembler.ProfessionID = id

	usecase.Assembler = &assembler
	product, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, product)
}

func (c *ProductCategoryController) Save(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgpproductCategoryuc.NewSaveProductCategoryUC(c.SaveProductCategoryUCParams)
	productCategoryAssembler := pkgpproductCategoryuc.ProductCategoryAssembler{}
	if err := ctx.Bind(&productCategoryAssembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	usecase.Assembler = &productCategoryAssembler
	productCategory, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, productCategory)

}
