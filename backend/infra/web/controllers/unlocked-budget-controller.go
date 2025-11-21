package controllers

import (
	pkguser "construir_mais_barato/app/domain/user"
	pkgauthuc "construir_mais_barato/app/usecase/auth"
	pkgunlockedbudgetuc "construir_mais_barato/app/usecase/unlocked-budget"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type UnlockedBudgetController struct {
	CheckBudgetAccessUCParams   pkgunlockedbudgetuc.CheckBudgetAccessParams
	UnlockBudgetPaymentUCParams pkgunlockedbudgetuc.UnlockBudgetPaymentParams
	UserService                 pkguser.UserService
}

type UnlockedBudgetControllerParams struct {
	CheckBudgetAccessUCParams   pkgunlockedbudgetuc.CheckBudgetAccessParams
	UnlockBudgetPaymentUCParams pkgunlockedbudgetuc.UnlockBudgetPaymentParams
	UserService                 pkguser.UserService
}

func NewUnlockedBudgetController(params *UnlockedBudgetControllerParams, g *echo.Group) {
	controller := UnlockedBudgetController{
		CheckBudgetAccessUCParams:   params.CheckBudgetAccessUCParams,
		UnlockBudgetPaymentUCParams: params.UnlockBudgetPaymentUCParams,
		UserService:                 params.UserService,
	}

	// Rotas autenticadas
	g.GET("/budget/:id/check-access", controller.CheckAccess)
	g.POST("/budget/:id/unlock", controller.UnlockBudget)
}

// CheckAccess verifica se o profissional tem acesso ao orçamento
func (c *UnlockedBudgetController) CheckAccess(ctx echo.Context) error {
	defer ctx.Request().Body.Close()

	token := getTokenFromHeader(ctx)
	if token == "" {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Token não fornecido"})
	}

	userID, err := pkgauthuc.GetUserIDFromToken(token)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Token inválido"})
	}

	user, err := c.UserService.FindById(userID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Usuário não encontrado"})
	}

	professional, err := c.CheckBudgetAccessUCParams.ProfessionalService.FindByEmail(user.Email)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Profissional não encontrado"})
	}

	budgetIDParam := ctx.Param("id")
	budgetID, err := strconv.ParseUint(budgetIDParam, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID do orçamento inválido"})
	}

	usecase := pkgunlockedbudgetuc.NewCheckBudgetAccessUC(c.CheckBudgetAccessUCParams)
	result, err := usecase.Execute(pkgunlockedbudgetuc.CheckBudgetAccessInput{
		ProfessionalID: professional.ID,
		BudgetID:       uint(budgetID),
	})

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, result)
}

// UnlockBudget cria um pagamento PIX para desbloquear um orçamento
func (c *UnlockedBudgetController) UnlockBudget(ctx echo.Context) error {
	defer ctx.Request().Body.Close()

	token := getTokenFromHeader(ctx)
	if token == "" {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Token não fornecido"})
	}

	userID, err := pkgauthuc.GetUserIDFromToken(token)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Token inválido"})
	}

	user, err := c.UserService.FindById(userID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Usuário não encontrado"})
	}

	professional, err := c.UnlockBudgetPaymentUCParams.ProfessionalService.FindByEmail(user.Email)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Profissional não encontrado"})
	}

	budgetIDParam := ctx.Param("id")
	budgetID, err := strconv.ParseUint(budgetIDParam, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "ID do orçamento inválido"})
	}

	var input pkgunlockedbudgetuc.UnlockBudgetPaymentInput
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Dados inválidos"})
	}

	input.ProfessionalID = professional.ID
	input.BudgetID = uint(budgetID)

	usecase := pkgunlockedbudgetuc.NewUnlockBudgetPaymentUC(c.UnlockBudgetPaymentUCParams)
	result, err := usecase.Execute(input)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, result)
}

// getTokenFromHeader extrai o token JWT do header Authorization
func getTokenFromHeader(c echo.Context) string {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return ""
	}

	return tokenParts[1]
}
