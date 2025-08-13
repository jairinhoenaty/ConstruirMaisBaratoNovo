package controllers

import (
	pkgpchatuc "construir_mais_barato/app/usecase/chat"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ChatController struct {
	FindAllChatUCParams pkgpchatuc.FindAllChatUCParams
	FindByUserChatUCParams pkgpchatuc.FindByUserChatUCParams	
	//FindByIdUCParams       pkgpchatuc.FindByIdUCParams
	SaveChatUCParams    pkgpchatuc.SaveChatUCParams
	DeleteChatUCParams  pkgpchatuc.DeleteChatUCParams
}

type ChatControllerParams struct {
	FindAllChatUCParams pkgpchatuc.FindAllChatUCParams
	FindByUserChatUCParams pkgpchatuc.FindByUserChatUCParams	
	//FindByIdUCParams       pkgpchatuc.FindByIdUCParams
	SaveChatUCParams    pkgpchatuc.SaveChatUCParams
	DeleteChatUCParams  pkgpchatuc.DeleteChatUCParams
}

func NewChatController(params *ChatControllerParams, g *echo.Group) {
	controller := ChatController{
		FindAllChatUCParams: params.FindAllChatUCParams,
		FindByUserChatUCParams: params.FindByUserChatUCParams,		
	//	FindByIdUCParams:       params.FindByIdUCParams,
		SaveChatUCParams:    params.SaveChatUCParams,
		DeleteChatUCParams:  params.DeleteChatUCParams,
	}

	g.POST("/chat", controller.Save)
	g.GET("/chats", controller.FindAll)
	g.POST("/chats", controller.FindByUser)
	//g.GET("/chat/:id", controller.FindById)
	g.DELETE("/chat/:id", controller.Delete)
}

func (c *ChatController) FindAll(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	
	limit,err := strconv.Atoi(ctx.QueryParam("limit"))
	if (err !=nil && limit == 0) {
		limit=20
	}
	offset,err := strconv.Atoi(ctx.QueryParam("offset"))
	if (err !=nil && offset == 0) {
		offset=0
	}

	usecase := pkgpchatuc.NewFindAllChatUC(c.FindAllChatUCParams)
	chats, total, err := usecase.Execute(limit,offset)
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}

		// Estrutura de resposta
	response := struct {
		Chats *[]pkgpchatuc.ChatPresenter `json:"chats"`
		Total       int64                          `json:"total"`
	}{
		Chats: chats,
		Total:       total,
	}
	return ctx.JSON(http.StatusOK, response)
}


func (c *ChatController) FindByUser(ctx echo.Context) error {
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
	limit,ok4 :=  requestBody["limit"].(int);
	if (!ok4 || limit == 0) {
		limit=20
	}

	offset,ok5 :=  requestBody["offset"].(int);
	if (!ok5 || offset == 0) {
		offset=0
	}
	

	if (!ok1 || professionalID == 0) && (!ok2 || clientID == 0) {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "faltam parametros"})
	}

	usecase := pkgpchatuc.NewFindByUserUC(c.FindByUserChatUCParams)
	chats, total, err := usecase.Execute(limit,offset, professionalID,clientID)
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	// Estrutura de resposta
	response := struct {
		Chats *[]pkgpchatuc.ChatPresenter `json:"chats"`
		Total       int64                                  `json:"total"`
	}{
		Chats: chats,
		Total:       total,
	}

	return ctx.JSON(http.StatusOK, response)	
}


/*func (c *ChatController) FindById(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgpchatuc.NewFindByIdUC(c.FindByIdUCParams)
	
	idAssembler := ctx.Param("id")
	id, err := strconv.ParseUint(idAssembler, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	uintID := uint(id)

	usecase.ID = &uintID
	chat, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, chat)
}
*/

func (c *ChatController) Save(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgpchatuc.NewSaveChatUC(c.SaveChatUCParams)
	chatAssembler := pkgpchatuc.ChatAssembler{}
	if err := ctx.Bind(&chatAssembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	usecase.Assembler = &chatAssembler

	chat, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, chat)

}

func (c *ChatController) Delete(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgpchatuc.NewDeleteChatUC(c.DeleteChatUCParams)
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
