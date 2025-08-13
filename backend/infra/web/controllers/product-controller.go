package controllers

import (
	pkgpproductuc "construir_mais_barato/app/usecase/product"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	FindAllProductUCParams pkgpproductuc.FindAllProductUCParams
	SaveProductUCParams    pkgpproductuc.SaveProductUCParams
	FindDayOfferUCParams   pkgpproductuc.FindDayofferProductUCParams
	FindByIdUCParams       pkgpproductuc.FindByIdUCParams
	DeleteProductUCParams  pkgpproductuc.DeleteProductUCParams
	/*
		FindByEmailUCParams                  pkgpproductuc.FindByEmailUCParams
		FindByMonthAndProfessionalIDUCParams pkgpproductuc.FindByMonthAndProfessionalIDUCParams*/
}

type ProductControllerParams struct {
	FindAllProductUCParams pkgpproductuc.FindAllProductUCParams
	FindByIdUCParams       pkgpproductuc.FindByIdUCParams
	SaveProductUCParams    pkgpproductuc.SaveProductUCParams
	FindDayOfferUCParams   pkgpproductuc.FindDayofferProductUCParams
	DeleteProductUCParams  pkgpproductuc.DeleteProductUCParams
	/*
		FindByEmailUCParams                  pkgpproductuc.FindByEmailUCParams
		FindByMonthAndProfessionalIDUCParams pkgpproductuc.FindByMonthAndProfessionalIDUCParams*/
}

func NewProductController(params *ProductControllerParams, g *echo.Group) {
	controller := ProductController{
		FindAllProductUCParams: params.FindAllProductUCParams,
		FindByIdUCParams:       params.FindByIdUCParams,
		SaveProductUCParams:    params.SaveProductUCParams,
		DeleteProductUCParams:  params.DeleteProductUCParams,
		/*
			FindByEmailUCParams:                  params.FindByEmailUCParams,
			FindByMonthAndProfessionalIDUCParams: params.FindByMonthAndProfessionalIDUCParams,*/
	}

	g.POST("/product", controller.Save)
	g.GET("/products", controller.FindAll)
	//g.GET("/products", controller.Find)
	g.GET("/products/dayoffer", controller.FindByDayOffer)
	g.GET("/product/:id", controller.FindById)
	g.DELETE("/product/:id", controller.Delete)
	/*
		g.POST("/products/month", controller.FindProductByMonthAndProfessionalId)
		g.POST("/product/email", controller.FindProductByEmail)*/
}

/*
	func (c *ProductController) FindProductByMonthAndProfessionalId(ctx echo.Context) error {
		defer ctx.Request().Body.Close()
		usecase := pkgpproductuc.NewFindByMonthAndProfessionalIDUC(c.FindByMonthAndProfessionalIDUCParams)
		assembler := pkgpproductuc.FindProductByMontAndProfessionalIDAssembler{}
		if err := ctx.Bind(&assembler); err != nil {
			return ctx.JSON(http.StatusPreconditionFailed, err)
		}

		usecase.Assembler = &assembler
		products, err := usecase.Execute()
		if err != nil {
			return ctx.JSON(http.StatusPreconditionFailed, nil)
		}
		fmt.Println(" ===== ProductController == Orçamentos => ", products)
		return ctx.JSON(http.StatusOK, products)
	}
*/
func (c *ProductController) FindAll(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	assembler := pkgpproductuc.FindWithPaginationProductAssembler{}

	// Extrair parâmetros de paginação da query string
	limit := ctx.QueryParam("limit")
	offset := ctx.QueryParam("offset")

	professionalID := ctx.QueryParam("professional_id")
	professionalIDInt, err := strconv.Atoi(professionalID)

	storeID := ctx.QueryParam("store_id")
	storeIDInt, err := strconv.Atoi(storeID)

	dayoffer := ctx.QueryParam("dayoffer")
	approved := ctx.QueryParam("approved")
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
	assembler.ProfessionalID = professionalIDInt
	assembler.StoreID = storeIDInt
	assembler.DayOffer = dayoffer
	assembler.Approved = approved

	usecase := pkgpproductuc.NewFindAllProductUC(c.FindAllProductUCParams)
	usecase.Assembler = assembler
	products, total, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, map[string]string{"error": "Erro ao buscar Produtos"})
	}

	// Estrutura de resposta
	response := struct {
		Products *[]pkgpproductuc.ProductPresenter `json:"products"`
		Total    int64                             `json:"total"`
	}{
		Products: products,
		Total:    total,
	}

	return ctx.JSON(http.StatusOK, response)
}

func (c *ProductController) FindByDayOffer(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgpproductuc.NewFindDayofferProductUC(c.FindDayOfferUCParams)

	product, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, product)
}

func (c *ProductController) FindById(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgpproductuc.NewFindByIdUC(c.FindByIdUCParams)
	idAssembler := ctx.Param("id")
	id, err := strconv.ParseUint(idAssembler, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	uintID := uint(id)

	usecase.ID = &uintID
	product, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, product)
}

/*
func (c *ProductController) FindProductByEmail(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgpproductuc.NewFindByEmailUC(c.FindByEmailUCParams)

	assembler := pkgpproductuc.FindProductByEmailAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	usecase.Assembler = assembler
	product, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, product)
}
*/

func (c *ProductController) Save(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgpproductuc.NewSaveProductUC(c.SaveProductUCParams)
	productAssembler := pkgpproductuc.ProductAssembler{}
	if err := ctx.Bind(&productAssembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	usecase.Assembler = &productAssembler

	product, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, product)

}

func (c *ProductController) Delete(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgpproductuc.NewDeleteProductUC(c.DeleteProductUCParams)
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
