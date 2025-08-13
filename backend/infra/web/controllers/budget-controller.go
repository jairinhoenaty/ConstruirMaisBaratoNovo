package controllers

import (
	pkgpbudgetuc "construir_mais_barato/app/usecase/budget"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BudgetController struct {
	FindAllBudgetUCParams                pkgpbudgetuc.FindAllBudgetUCParams
	FindByIdUCParams                     pkgpbudgetuc.FindByIdUCParams
	SaveBudgetUCParams                   pkgpbudgetuc.SaveBudgetUCParams
	DeleteBudgetUCParams                 pkgpbudgetuc.DeleteBudgetUCParams
	FindByEmailUCParams                  pkgpbudgetuc.FindByEmailUCParams
	FindByMonthAndProfessionalIDUCParams pkgpbudgetuc.FindByMonthAndProfessionalIDUCParams
}

type BudgetControllerParams struct {
	FindAllBudgetUCParams                pkgpbudgetuc.FindAllBudgetUCParams
	FindByIdUCParams                     pkgpbudgetuc.FindByIdUCParams
	SaveBudgetUCParams                   pkgpbudgetuc.SaveBudgetUCParams
	DeleteBudgetUCParams                 pkgpbudgetuc.DeleteBudgetUCParams
	FindByEmailUCParams                  pkgpbudgetuc.FindByEmailUCParams
	FindByMonthAndProfessionalIDUCParams pkgpbudgetuc.FindByMonthAndProfessionalIDUCParams
}

func NewBudgetController(params *BudgetControllerParams, g *echo.Group) {
	controller := BudgetController{
		FindAllBudgetUCParams:                params.FindAllBudgetUCParams,
		FindByIdUCParams:                     params.FindByIdUCParams,
		SaveBudgetUCParams:                   params.SaveBudgetUCParams,
		DeleteBudgetUCParams:                 params.DeleteBudgetUCParams,
		FindByEmailUCParams:                  params.FindByEmailUCParams,
		FindByMonthAndProfessionalIDUCParams: params.FindByMonthAndProfessionalIDUCParams,
	}

	g.POST("/budget", controller.Save)
	g.GET("/budgets", controller.FindAll)
	g.GET("/budget/:id", controller.FindById)
	g.DELETE("/budget/:id", controller.Delete)
	g.POST("/budgets/month", controller.FindBudgetByMonthAndProfessionalId)
	g.POST("/budget/email", controller.FindBudgetByEmail)
}

func (c *BudgetController) FindBudgetByMonthAndProfessionalId(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgpbudgetuc.NewFindByMonthAndProfessionalIDUC(c.FindByMonthAndProfessionalIDUCParams)
	assembler := pkgpbudgetuc.FindBudgetByMontAndProfessionalIDAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	usecase.Assembler = &assembler
	fmt.Println("JSON REQUEST")
	fmt.Println("assembler", usecase.Assembler)
	budgets, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	fmt.Println(" ===== BudgetController == Orçamentos => ", budgets)
	return ctx.JSON(http.StatusOK, budgets)
}

func (c *BudgetController) FindAll(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	assembler := pkgpbudgetuc.FindWithPaginationBudgetAssembler{}

	// Extrair parâmetros de paginação da query string
	limit := ctx.QueryParam("limit")
	offset := ctx.QueryParam("offset")

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

	usecase := pkgpbudgetuc.NewFindAllBudgetUC(c.FindAllBudgetUCParams)
	usecase.Assembler = assembler
	budgets, total, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, map[string]string{"error": "Erro ao buscar orçamentos"})
	}

	// Estrutura de resposta
	response := struct {
		Budgets *[]pkgpbudgetuc.BudgetPresenter `json:"budgets"`
		Total   int64                           `json:"total"`
	}{
		Budgets: budgets,
		Total:   total,
	}

	return ctx.JSON(http.StatusOK, response)
}

// func (c *BudgetController) FindAll(ctx echo.Context) error {

// 	// Extrair page e pageSize da query string
// 	page, err := strconv.Atoi(ctx.QueryParam("page"))
// 	if err != nil || page < 1 {
// 		page = 1 // Default para página 1 se o parâmetro não for válido
// 	}

// 	pageSize, err := strconv.Atoi(ctx.QueryParam("pageSize"))
// 	if err != nil || pageSize < 1 {
// 		pageSize = 10 // Default para 10 registros por página se o parâmetro não for válido
// 	}

// 	usecase := pkgpbudgetuc.NewFindAllBudgetUC(c.FindAllBudgetUCParams)
// 	usecase.Page = page
// 	usecase.PageSize = pageSize
// 	budgets, err := usecase.Execute()
// 	if err != nil {
// 		return ctx.JSON(http.StatusPreconditionFailed, nil)
// 	}
// 	return ctx.JSON(http.StatusOK, budgets)
// }

func (c *BudgetController) FindById(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgpbudgetuc.NewFindByIdUC(c.FindByIdUCParams)
	idAssembler := ctx.Param("id")
	id, err := strconv.ParseUint(idAssembler, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	uintID := uint(id)

	usecase.ID = &uintID
	budget, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, budget)
}

func (c *BudgetController) FindBudgetByEmail(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgpbudgetuc.NewFindByEmailUC(c.FindByEmailUCParams)

	assembler := pkgpbudgetuc.FindBudgetByEmailAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	usecase.Assembler = assembler
	budget, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, budget)
}

func (c *BudgetController) Save(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgpbudgetuc.NewSaveBudgetUC(c.SaveBudgetUCParams)
	budgetAssembler := pkgpbudgetuc.BudgetAssembler{}
	if err := ctx.Bind(&budgetAssembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	usecase.Assembler = &budgetAssembler

	budget, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, budget)

}

func (c *BudgetController) Delete(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgpbudgetuc.NewDeleteBudgetUC(c.DeleteBudgetUCParams)
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
