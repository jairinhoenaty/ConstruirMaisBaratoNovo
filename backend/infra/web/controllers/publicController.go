package controllers

import (
	pkgplan "construir_mais_barato/app/domain/plan"
	pkgbanneruc "construir_mais_barato/app/usecase/banner"
	pkgbudgetuc "construir_mais_barato/app/usecase/budget"
	pkgcityuc "construir_mais_barato/app/usecase/city"
	pkgclientuc "construir_mais_barato/app/usecase/client"
	pkgpcontactuc "construir_mais_barato/app/usecase/contact"
	pkgpageviewuc "construir_mais_barato/app/usecase/pageview"
	pkgplanuc "construir_mais_barato/app/usecase/plan"
	pkgpproductuc "construir_mais_barato/app/usecase/product"
	pkgproductuc "construir_mais_barato/app/usecase/product"
	pkgprofessionuc "construir_mais_barato/app/usecase/profession"
	pkgprofessionaluc "construir_mais_barato/app/usecase/professional"
	pkgregionuc "construir_mais_barato/app/usecase/region"
	pkgstoreuc "construir_mais_barato/app/usecase/store"
	pkguseruc "construir_mais_barato/app/usecase/user"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"net/http"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type PublicController struct {
	FindByEmailUCParams                                 pkguseruc.FindByEmailUCParams
	FindByPhoneUCParams                                 pkguseruc.FindByPhoneUCParams
	FindStoreByCategoryAndSubCategoryParams             pkgstoreuc.FindStoreByCategoryAndSubCategoryParams
	SaveClientUCParams                                  pkgclientuc.SaveClientUCParams
	SaveStoreUCParams                                   pkgstoreuc.SaveStoreUCParams
	FindByPageUCParams                                  pkgbanneruc.FindByPageUCParams
	SaveContactUCParams                                 pkgpcontactuc.SaveContactUCParams
	FindByCityProductUCParams                           pkgproductuc.FindByCityUCParams
	FindDayofferProductUCParams                         pkgproductuc.FindDayofferProductUCParams
	FindAllProductUCParams                              pkgproductuc.FindAllProductUCParams
	FindByUFUCParams                                    pkgcityuc.FindByUFUCParams
	SearchCityByNameUCParams                            pkgcityuc.SearchCityByNameUCParams
	FindRegionByCityIdUCParams                          pkgregionuc.FindByCityUCParams
	SaveBudgetUCParams                                  pkgbudgetuc.SaveBudgetUCParams
	UserSendEmailUCParams                               pkguseruc.UserSendEmailUCParams
	ResetPasswordUCParams                               pkguseruc.ResetPasswordUCParams
	FindProfessionUCParams                              pkgprofessionuc.FindProfessionUCParams
	FindAllProfessionUCParams                           pkgprofessionuc.FindAllProfessionUCParams
	SaveProfessionalUCParams                            pkgprofessionaluc.SaveProfessionalUCParams
	FindByCityAndProfessionUCParams                     pkgbanneruc.FindByCityAndProfessionUCParams
	FindProfessionsWithCountIdUCParams                  pkgprofessionuc.FindProfessionsWithCountIdUCParams
	FindAllWithoutPaginationProfessionParams            pkgprofessionuc.FindAllWithoutPaginationProfessionParams
	FindByProfessionalByCityAndProfessionUCParamns      pkgprofessionaluc.FindByProfessionalByCityAndProfessionUCParamns
	FindByNameProfessionalAndCityAndProfessionUCParamns pkgprofessionaluc.FindByNameProfessionalAndCityAndProfessionUCParamns
	FindAllActivePlansUCParams                          pkgplanuc.FindAllActiveUCParams
	FindPlanByUserTypeUCParams                          pkgplanuc.FindByUserTypeUCParams
	CheckoutProfessionalPremiumUCParams                 pkgprofessionaluc.CheckoutPremiumUCParams
	CheckoutStorePremiumUCParams                        pkgstoreuc.CheckoutPremiumUCParams
	IncrementPageViewUCParams                           pkgpageviewuc.IncrementPageViewUCParams
}

type PublicControllerParams struct {
	FindByEmailUCParams                                 pkguseruc.FindByEmailUCParams
	FindByPhoneUCParams                                 pkguseruc.FindByPhoneUCParams
	FindStoreByCategoryAndSubCategoryParams             pkgstoreuc.FindStoreByCategoryAndSubCategoryParams
	SaveClientUCParams                                  pkgclientuc.SaveClientUCParams
	SaveStoreUCParams                                   pkgstoreuc.SaveStoreUCParams
	FindByPageUCParams                                  pkgbanneruc.FindByPageUCParams
	SaveContactUCParams                                 pkgpcontactuc.SaveContactUCParams
	FindByCityProductUCParams                           pkgproductuc.FindByCityUCParams
	FindDayofferProductUCParams                         pkgproductuc.FindDayofferProductUCParams
	FindRegionByCityIdUCParams                          pkgregionuc.FindByCityUCParams
	FindAllProductUCParams                              pkgproductuc.FindAllProductUCParams
	FindByUFUCParams                                    pkgcityuc.FindByUFUCParams
	SearchCityByNameUCParams                            pkgcityuc.SearchCityByNameUCParams
	SaveBudgetUCParams                                  pkgbudgetuc.SaveBudgetUCParams
	UserSendEmailUCParams                               pkguseruc.UserSendEmailUCParams
	ResetPasswordUCParams                               pkguseruc.ResetPasswordUCParams
	FindProfessionUCParams                              pkgprofessionuc.FindProfessionUCParams
	FindAllProfessionUCParams                           pkgprofessionuc.FindAllProfessionUCParams
	SaveProfessionalUCParams                            pkgprofessionaluc.SaveProfessionalUCParams
	FindByCityAndProfessionUCParams                     pkgbanneruc.FindByCityAndProfessionUCParams
	FindProfessionsWithCountIdUCParams                  pkgprofessionuc.FindProfessionsWithCountIdUCParams
	FindAllWithoutPaginationProfessionParams            pkgprofessionuc.FindAllWithoutPaginationProfessionParams
	FindByProfessionalByCityAndProfessionUCParamns      pkgprofessionaluc.FindByProfessionalByCityAndProfessionUCParamns
	FindByNameProfessionalAndCityAndProfessionUCParamns pkgprofessionaluc.FindByNameProfessionalAndCityAndProfessionUCParamns
	FindAllActivePlansUCParams                          pkgplanuc.FindAllActiveUCParams
	FindPlanByUserTypeUCParams                          pkgplanuc.FindByUserTypeUCParams
	CheckoutProfessionalPremiumUCParams                 pkgprofessionaluc.CheckoutPremiumUCParams
	CheckoutStorePremiumUCParams                        pkgstoreuc.CheckoutPremiumUCParams
	IncrementPageViewUCParams                           pkgpageviewuc.IncrementPageViewUCParams
}

func NewPublicController(params *PublicControllerParams, g *echo.Group) {
	controller := PublicController{
		FindByEmailUCParams:                                 params.FindByEmailUCParams,
		FindByPhoneUCParams:                                 params.FindByPhoneUCParams,
		SaveClientUCParams:                                  params.SaveClientUCParams,
		SaveStoreUCParams:                                   params.SaveStoreUCParams,
		FindByPageUCParams:                                  params.FindByPageUCParams,
		SaveContactUCParams:                                 params.SaveContactUCParams,
		FindByCityProductUCParams:                           params.FindByCityProductUCParams,
		FindDayofferProductUCParams:                         params.FindDayofferProductUCParams,
		FindRegionByCityIdUCParams:                          params.FindRegionByCityIdUCParams,
		FindAllProductUCParams:                              params.FindAllProductUCParams,
		FindByUFUCParams:                                    params.FindByUFUCParams,
		SearchCityByNameUCParams:                            params.SearchCityByNameUCParams,
		SaveBudgetUCParams:                                  params.SaveBudgetUCParams,
		UserSendEmailUCParams:                               params.UserSendEmailUCParams,
		ResetPasswordUCParams:                               params.ResetPasswordUCParams,
		FindProfessionUCParams:                              params.FindProfessionUCParams,
		SaveProfessionalUCParams:                            params.SaveProfessionalUCParams,
		FindAllProfessionUCParams:                           params.FindAllProfessionUCParams,
		FindByCityAndProfessionUCParams:                     params.FindByCityAndProfessionUCParams,
		FindProfessionsWithCountIdUCParams:                  params.FindProfessionsWithCountIdUCParams,
		FindAllWithoutPaginationProfessionParams:            params.FindAllWithoutPaginationProfessionParams,
		FindByProfessionalByCityAndProfessionUCParamns:      params.FindByProfessionalByCityAndProfessionUCParamns,
		FindByNameProfessionalAndCityAndProfessionUCParamns: params.FindByNameProfessionalAndCityAndProfessionUCParamns,
		FindAllActivePlansUCParams:                          params.FindAllActivePlansUCParams,
		FindPlanByUserTypeUCParams:                          params.FindPlanByUserTypeUCParams,
		CheckoutProfessionalPremiumUCParams:                 params.CheckoutProfessionalPremiumUCParams,
		CheckoutStorePremiumUCParams:                        params.CheckoutStorePremiumUCParams,
		FindStoreByCategoryAndSubCategoryParams:             params.FindStoreByCategoryAndSubCategoryParams,
		IncrementPageViewUCParams:                           params.IncrementPageViewUCParams,
	}

	g.GET("/plans", controller.GetActivePlans)
	g.GET("/plans/:userType", controller.GetPlanByUserType)
	g.POST("/store/checkout/premium", controller.CheckoutStorePremium)
	g.POST("/save/budget", controller.SaveBudget)
	g.POST("/user/send-mail", controller.SendMail)
	g.POST("/user/find-by-email", controller.FindUserByEmail)
	g.POST("/user/find-by-phone", controller.FindUserByPhone)
	g.POST("/reset/password", controller.ResetPassword)
	g.GET("/professions/all", controller.FindProfessionAll)
	g.POST("/save/professional", controller.SaveProfessional)
	g.POST("/cities-by-state", controller.PublicFindCitiesByState)
	g.GET("/cities/search", controller.PublicSearchCities)
	g.GET("/professions/:quantityProfession", controller.FindProfessions)
	g.POST("/find/professions-with-count", controller.FindProfessionsWithCount)
	g.POST("/find-banner-city-and-profession", controller.FindByCityAndProfession)
	g.POST("/professional/checkout/premium", controller.CheckoutProfissionalPremium)
	g.POST("/search-all-professionals-and-city-and-profession", controller.PublicFindAllProfessionalsByCityAndProfession)
	g.POST("/search-professionals-by-name-and-city-and-profession", controller.FindByNameProfessinalsAndCityAndProfession)
	g.POST("/products/dayoffer", controller.FindProductsByDayOffer)
	g.GET("/products", controller.FindAllProducts)
	g.POST("/products/findbycity", controller.FindByCity)
	g.GET("/regions/findbycity", controller.FindRegionByCityId)

	g.POST("/contact", controller.SaveContact)

	g.POST("/banners/page", controller.FindBannerbyPage)

	g.POST("/save/store", controller.SaveStore)
	g.POST("/category-and-sub", controller.FindStoreByCategoryAndSubCategory)
	g.POST("/save/client", controller.SaveClient)
	g.POST("/upload/image", controller.uploadFile)
	g.POST("/page-view", controller.TrackPageView)

}

func (c *PublicController) GetActivePlans(ctx echo.Context) error {
	defer ctx.Request().Body.Close()

	uc := pkgplanuc.NewFindAllActiveUC(c.FindAllActivePlansUCParams)
	plans, err := uc.Execute()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, plans)
}

func (c *PublicController) GetPlanByUserType(ctx echo.Context) error {
	defer ctx.Request().Body.Close()

	userTypeStr := ctx.Param("userType")
	if userTypeStr != "professional" && userTypeStr != "store" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user type. Use 'professional' or 'store'"})
	}

	uc := pkgplanuc.NewFindByUserTypeUC(c.FindPlanByUserTypeUCParams)
	userTypeEnum := pkgplan.UserType(userTypeStr)
	uc.UserType = &userTypeEnum

	plan, err := uc.Execute()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, plan)
}

func (c *PublicController) CheckoutStorePremium(ctx echo.Context) error {
	defer ctx.Request().Body.Close()

	var assembler pkgstoreuc.PayerAssembler
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	checkoutUC := pkgstoreuc.NewCheckoutPremiumUC(c.CheckoutStorePremiumUCParams)
	checkoutUC.Assembler = assembler
	result, err := checkoutUC.Execute()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, result)
}

func (c *PublicController) CheckoutProfissionalPremium(ctx echo.Context) error {
	defer ctx.Request().Body.Close()

	var assembler pkgprofessionaluc.PayerAssembler
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	checkoutUC := pkgprofessionaluc.NewCheckoutPremiumUC(c.CheckoutProfessionalPremiumUCParams)
	checkoutUC.Assembler = assembler
	result, err := checkoutUC.Execute()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, result)
}

func (c *PublicController) FindByCityAndProfession(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	assembler := pkgbanneruc.FindByCityIdAndProfessionIDAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	usecase := pkgbanneruc.NewFindByCityAndProfessionUC(c.FindByCityAndProfessionUCParams)
	usecase.Assembler = &assembler
	banner, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, banner)
}

func (c *PublicController) FindByNameProfessinalsAndCityAndProfession(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	assembler := pkgprofessionaluc.FindProfessionalByCityAndProfessionAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	uc := pkgprofessionaluc.NewFindByNameProfessionalAndCityAndProfessionUC(c.FindByNameProfessionalAndCityAndProfessionUCParamns)
	uc.Assembler = &assembler
	professionals, err := uc.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, professionals)
}

func (c *PublicController) ResetPassword(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	assembler := pkguseruc.ResetPasswordAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request format"})
	}

	usecase := pkguseruc.NewResetPasswordUC(c.ResetPasswordUCParams)
	usecase.Assembler = &assembler

	err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, map[string]string{"error": "Erro ao alterar a senha"})
	}
	return ctx.JSON(http.StatusOK, nil)
}

func (c *PublicController) SendMail(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	var request struct {
		Email string `json:"email"`
	}
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request format"})
	}

	usecase := pkguseruc.NewSendEmailUC(c.UserSendEmailUCParams)
	usecase.Email = request.Email

	err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, map[string]string{"error": "Erro ao enviar o email"})
	}
	return ctx.JSON(http.StatusOK, nil)
}

func (c *PublicController) SaveBudget(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgbudgetuc.NewSaveBudgetUC(c.SaveBudgetUCParams)
	budgetAssembler := pkgbudgetuc.BudgetAssembler{}
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

func (c *PublicController) TrackPageView(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgpageviewuc.NewIncrementPageViewUC(c.IncrementPageViewUCParams)
	assembler := pkgpageviewuc.IncrementPageViewAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	if assembler.Path == "" {
		assembler.Path = "/"
	}
	usecase.Assembler = &assembler
	pageView, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, pageView)
}

func (c *PublicController) FindProfessionsWithCount(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgprofessionuc.NewFindProfessionsWithCountIdUC(c.FindProfessionsWithCountIdUCParams)
	professions, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, professions)
}

func (c *PublicController) SaveProfessional(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgprofessionaluc.NewSaveProfessionalUC(c.SaveProfessionalUCParams)
	assembler := pkgprofessionaluc.ProfessionalAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, map[string]string{
			"error":   "Failed to bind request data",
			"details": err.Error(),
		})
	}
	usecase.Assembler = &assembler
	professional, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, professional)

}

func (c *PublicController) FindStoreByCategoryAndSubCategory(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	assembler := pkgstoreuc.FindByCategoryAndSubCategory{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	uc := pkgstoreuc.NewFindStoreByCategoryAndSubCategory(c.FindStoreByCategoryAndSubCategoryParams)
	uc.Assembler = &assembler
	stores, err := uc.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, stores)
}

func (c *PublicController) SaveStore(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgstoreuc.NewSaveStoreUC(c.SaveStoreUCParams)
	assembler := pkgstoreuc.StoreAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, map[string]string{
			"error":   "Failed to bind request data",
			"details": err.Error(),
		})
	}
	usecase.Assembler = &assembler
	store, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, store)

}

func (c *PublicController) SaveClient(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgclientuc.NewSaveClientUC(c.SaveClientUCParams)
	assembler := pkgclientuc.ClientAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, map[string]string{
			"error":   "Failed to bind request data",
			"details": err.Error(),
		})
	}
	usecase.Assembler = &assembler
	store, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, store)

}

func (c *PublicController) FindProfessions(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	quantityProfessionAssembler := ctx.Param("quantityProfession")

	usecase := pkgprofessionuc.NewFindProfessionUC(c.FindProfessionUCParams)
	quantProf, err := strconv.ParseUint(quantityProfessionAssembler, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	uintQuantProf := uint(quantProf)
	usecase.QuantityProfessions = uintQuantProf

	professions, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, professions)
}

func (c *PublicController) PublicFindCitiesByState(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	assembler := pkgcityuc.UFCityAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	usecase := pkgcityuc.NewFindByUFUC(c.FindByUFUCParams)
	usecase.Assembler = &assembler
	cities, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, cities)
}

// PublicSearchCities alimenta o autocomplete de cidades do site.
// Ex.: GET /publica/cities/search?q=sao&limit=8
func (c *PublicController) PublicSearchCities(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	assembler := pkgcityuc.SearchCityAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	usecase := pkgcityuc.NewSearchCityByNameUC(c.SearchCityByNameUCParams)
	usecase.Assembler = &assembler
	cities, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, cities)
}

func (c *PublicController) PublicFindAllProfessionalsByCityAndProfession(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	assembler := pkgprofessionaluc.FindProfessionalByCityAndProfessionAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	usecase := pkgprofessionaluc.NewFindByProfessionalByCityAndProfessionUC(c.FindByProfessionalByCityAndProfessionUCParamns)
	usecase.Assembler = &assembler
	professionals, total, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	// Estrutura de resposta
	response := struct {
		Professionals []*pkgprofessionaluc.ProfessionalPresenter `json:"profissionais"`
		Total         int64                                      `json:"total"`
	}{
		Professionals: professionals,
		Total:         total,
	}

	return ctx.JSON(http.StatusOK, response)
}

func (c *PublicController) FindProfessionAll(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgprofessionuc.NewFindAllWithoutPaginationProfessionUC(c.FindAllWithoutPaginationProfessionParams)
	professions, err := usecase.Execute()

	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, professions)
}

func (c *PublicController) FindProductsByDayOffer(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgproductuc.NewFindDayofferProductUC(c.FindDayofferProductUCParams)

	product, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, product)
}

func (c *PublicController) FindAllProducts(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	assembler := pkgpproductuc.FindWithPaginationProductAssembler{}

	limit := ctx.QueryParam("limit")
	offset := ctx.QueryParam("offset")
	professionalID := ctx.QueryParam("professional_id")
	professionalIDInt, err := strconv.Atoi(professionalID)
	storeID := ctx.QueryParam("store_id")
	storeIDInt, err := strconv.Atoi(storeID)

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
	usecase := pkgproductuc.NewFindAllProductUC(c.FindAllProductUCParams)
	usecase.Assembler = assembler

	product, total, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	response := struct {
		Products *[]pkgproductuc.ProductPresenter `json:"products"`
		Total    int64                            `json:"total"`
	}{
		Products: product,
		Total:    total,
	}
	return ctx.JSON(http.StatusOK, response)

}

func (c *PublicController) FindByCity(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	assembler := pkgproductuc.FindByCityAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}
	usecase := pkgproductuc.NewFindByCityUC(c.FindByCityProductUCParams)
	usecase.Assembler = &assembler

	product, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, product)
}

func (c *PublicController) FindRegionByCityId(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgregionuc.NewFindByCityUC(c.FindRegionByCityIdUCParams)
	idAssembler := ctx.QueryParam("cityId")
	id, err := strconv.ParseUint(idAssembler, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	uintID := uint(id)

	usecase.ID = &uintID
	region, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, region)
}

func (c *PublicController) SaveContact(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgpcontactuc.NewSaveContactUC(c.SaveContactUCParams)
	contactAssembler := pkgpcontactuc.ContactAssembler{}
	if err := ctx.Bind(&contactAssembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	usecase.Assembler = &contactAssembler

	contact, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, contact)

}

func (c *PublicController) FindBannerbyPage(ctx echo.Context) error {

	defer ctx.Request().Body.Close()

	assembler := pkgbanneruc.FindByPageAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	uc := pkgbanneruc.NewFindByPageUC(c.FindByPageUCParams)
	uc.Assembler = &assembler
	banners, err := uc.Execute()

	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, banners)
}

func (c *PublicController) FindUserByEmail(ctx echo.Context) error {
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
		//return ctx.JSON(http.StatusPreconditionFailed, map[string]string{"error": "E-mail não encontrado"})
		return ctx.JSON(http.StatusOK, false)
	}
	if user != nil {
		return ctx.JSON(http.StatusOK, true)
	}

	return ctx.JSON(http.StatusOK, false)

}

func (c *PublicController) FindUserByPhone(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	var request struct {
		Telephone string `json:"telephone"`
	}
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request format"})
	}

	usecase := pkguseruc.NewFindByPhoneUC(c.FindByPhoneUCParams)
	usecase.Telephone = &request.Telephone

	err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusOK, false)
	}
	return ctx.JSON(http.StatusOK, true)
}

/*
func (c *PublicController) uploadImage(ctx echo.Context) error {
 file, err := ctx.FormFile("image")

  if err != nil {
   log.Println("Error in uploading Image : ", err)
   return ctx.JSON(http.StatusInternalServerError,"Server error")

  }

  uniqueId := uuid.New()

  filename := strings.Replace(uniqueId.String(), "-", "", -1)

  fileExt := strings.Split(file.Filename, ".")[1]

  image := fmt.Sprintf("%s.%s", filename, fileExt)

  // Generate a unique filename
  newFile, err := os.Create("./images/" + image)

  //err = ctx.SaveFile(file, fmt.Sprintf("./images/%s", image))

  if err != nil {
   log.Println("Error in saving Image :", err)
//   return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	return ctx.JSON(http.StatusInternalServerError,"Server error")

  }
  defer newFile.Close()

    _, err = io.Copy(newFile, file)
    if err != nil {
        http.Error(w, "Error copying the file", http.StatusInternalServerError)
        return
    }

  imageUrl := fmt.Sprintf("http://localhost:3000/images/%s", image)

  data := map[string]interface{}{

   "imageName": image,
   "imageUrl":  imageUrl,
   "header":    file.Header,
   "size":      file.Size,
  }

//  return c.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully", "data":
// data})
	return ctx.JSON(http.StatusCreated,data)

}
*/

func (c *PublicController) uploadFile(ctx echo.Context) error {
	fmt.Println("File Upload Endpoint Hit")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	dir_upload := os.Getenv("DIR_UPLOAD")
	if dir_upload == "" {
		dir_upload = "../files/images/upload"
	}

	if _, err := os.Stat(dir_upload); os.IsNotExist(err) {
		fmt.Println("Dir not exist " + dir_upload + ", creating")
		err := os.MkdirAll(dir_upload, 0755)
		if err != nil {
			fmt.Printf("Error creating upload directory: %v\n", err)
			return ctx.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to create upload directory",
			})
		}
	}

	file, err := ctx.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve uploaded file",
		})
	}
	fmt.Printf("Uploaded File: %+v\n", file.Filename)
	fmt.Printf("File Size: %+v\n", file.Size)
	fmt.Printf("MIME Header: %+v\n", file.Header)

	ext := filepath.Ext(file.Filename)
	if ext == "" {
		ext = ".png" // fallback extension
	}

	tempFile, err := os.CreateTemp(dir_upload, "upload-*"+ext)
	if err != nil {
		fmt.Printf("Error creating temporary file: %v\n", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create temporary file",
		})
	}
	defer tempFile.Close()

	fileData, err := file.Open()
	if err != nil {
		fmt.Printf("Error opening uploaded file: %v\n", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to open uploaded file",
		})
	}
	defer fileData.Close()

	fileBytes, err := io.ReadAll(fileData)
	if err != nil {
		fmt.Printf("Error reading file contents: %v\n", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to read file contents",
		})
	}

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to write file",
		})
	}

	err = os.Chmod(tempFile.Name(), 0755)
	if err != nil {
		fmt.Printf("Error setting file permissions: %v\n", err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to set file permissions",
		})
	}

	uri := "/images/upload" + strings.Replace(tempFile.Name(), dir_upload, "", -1)
	return ctx.JSON(http.StatusCreated, map[string]string{
		"message": "Successfully Uploaded File",
		"uri":     uri})

}
