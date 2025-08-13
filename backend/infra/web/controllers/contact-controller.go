package controllers

import (
	pkgpcontactuc "construir_mais_barato/app/usecase/contact"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ContactController struct {
	FindAllContactUCParams pkgpcontactuc.FindAllContactUCParams
	FindByUserContactUCParams pkgpcontactuc.FindByUserContactUCParams	
	FindByIdUCParams       pkgpcontactuc.FindByIdUCParams
	SaveContactUCParams    pkgpcontactuc.SaveContactUCParams
	DeleteContactUCParams  pkgpcontactuc.DeleteContactUCParams
}

type ContactControllerParams struct {
	FindAllContactUCParams pkgpcontactuc.FindAllContactUCParams
	FindByUserContactUCParams pkgpcontactuc.FindByUserContactUCParams	
	FindByIdUCParams       pkgpcontactuc.FindByIdUCParams
	SaveContactUCParams    pkgpcontactuc.SaveContactUCParams
	DeleteContactUCParams  pkgpcontactuc.DeleteContactUCParams
}

func NewContactController(params *ContactControllerParams, g *echo.Group) {
	controller := ContactController{
		FindAllContactUCParams: params.FindAllContactUCParams,
		FindByUserContactUCParams: params.FindByUserContactUCParams,		
		FindByIdUCParams:       params.FindByIdUCParams,
		SaveContactUCParams:    params.SaveContactUCParams,
		DeleteContactUCParams:  params.DeleteContactUCParams,
	}

	g.POST("/contact", controller.Save)
	g.GET("/contacts", controller.FindAll)
	g.POST("/contacts", controller.FindByUser)
	g.GET("/contact/:id", controller.FindById)
	g.DELETE("/contact/:id", controller.Delete)
}

func (c *ContactController) FindAll(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	
	limit,err := strconv.Atoi(ctx.QueryParam("limit"))
	if (err !=nil && limit == 0) {
		limit=20
	}
	offset,err := strconv.Atoi(ctx.QueryParam("offset"))
	if (err !=nil && offset == 0) {
		offset=0
	}

	usecase := pkgpcontactuc.NewFindAllContactUC(c.FindAllContactUCParams)
	contacts, total, err := usecase.Execute(limit,offset)
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}

		// Estrutura de resposta
	response := struct {
		Contacts *[]pkgpcontactuc.ContactPresenter `json:"contacts"`
		Total       int64                          `json:"total"`
	}{
		Contacts: contacts,
		Total:       total,
	}
	return ctx.JSON(http.StatusOK, response)
}


func (c *ContactController) FindByUser(ctx echo.Context) error {
	defer ctx.Request().Body.Close()

	var requestBody map[string]interface{}

	// Bind the request body to a map
	if err := ctx.Bind(&requestBody); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// Extrai o valor do requestBody como float64 (considerando que pode ser um tipo numérico compatível)
	/*cityIDFloat, ok := requestBody["cityID"].(float64)
	if !ok || cityIDFloat == 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "City Id is required"})
	}

	// Converte o float64 para uint
	cityID := uint(cityIDFloat)	
	*/

	professionalIDStr,ok1 :=  requestBody["professional_id"].(float64);
	professionalID := int(professionalIDStr);
	clientIDStr,ok2 :=  requestBody["client_id"].(float64);
	clientID := int(clientIDStr);
	storeIDStr,ok3 :=  requestBody["store_id"].(float64);
	storeID := int(storeIDStr);
	limit,ok4 :=  requestBody["limit"].(int);
	if (!ok4 || limit == 0) {
		limit=20
	}

	offset,ok5 :=  requestBody["offset"].(int);
	if (!ok5 || offset == 0) {
		offset=0
	}
	

	if (!ok1 || professionalID == 0) && (!ok2 || clientID == 0) && (!ok3 || storeID == 0){
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "faltam parametros"})
	}

	usecase := pkgpcontactuc.NewFindByUserUC(c.FindByUserContactUCParams)
	contacts, total, err := usecase.Execute(limit,offset, professionalID,clientID,storeID)
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	// Estrutura de resposta
	response := struct {
		Contacts *[]pkgpcontactuc.ContactPresenter `json:"contacts"`
		Total       int64                                  `json:"total"`
	}{
		Contacts: contacts,
		Total:       total,
	}

	return ctx.JSON(http.StatusOK, response)	
}


func (c *ContactController) FindById(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgpcontactuc.NewFindByIdUC(c.FindByIdUCParams)
	
	idAssembler := ctx.Param("id")
	id, err := strconv.ParseUint(idAssembler, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	uintID := uint(id)

	usecase.ID = &uintID
	contact, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, contact)
}

func (c *ContactController) Save(ctx echo.Context) error {
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

func (c *ContactController) Delete(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgpcontactuc.NewDeleteContactUC(c.DeleteContactUCParams)
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
