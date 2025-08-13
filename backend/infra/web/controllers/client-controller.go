package controllers

import (
	pkgclientuc "construir_mais_barato/app/usecase/client"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/xuri/excelize/v2"
)

type ClientController struct {
	FindAllClientUCParams                  pkgclientuc.FindAllClientUCParams
	FindByIdUCParams                             pkgclientuc.FindByIdUCParamns
	FindByNamedUCParams                          pkgclientuc.FindByNamedUCParams
	SaveClientUCParams                     pkgclientuc.SaveClientUCParams
	DeleteClientUCParams                   pkgclientuc.DeleteClientUCParams
	FindLastClientsUCParams                pkgclientuc.FindLastClientsUCParams
	ExportXLSXClientUCParams               pkgclientuc.ExportXLSXClientUCParams
	/*CountClientsByStateUCParams            pkgclientuc.CountClientsByStateUCParams
	CountClientsByProfessionUCParams       pkgclientuc.CountClientsByProfessionUCParams
	CountClientsByProfessionInCityUCParams pkgclientuc.CountClientsByProfessionInCityUCParams*/
}

type ClientControllerParams struct {
	FindAllClientUCParams                  pkgclientuc.FindAllClientUCParams
	FindByIdUCParams                             pkgclientuc.FindByIdUCParamns
	FindByNamedUCParams                          pkgclientuc.FindByNamedUCParams
	SaveClientUCParams                     pkgclientuc.SaveClientUCParams
	DeleteClientUCParams                   pkgclientuc.DeleteClientUCParams
	FindLastClientsUCParams                pkgclientuc.FindLastClientsUCParams
	
	ExportXLSXClientUCParams               pkgclientuc.ExportXLSXClientUCParams
	/*CountClientsByProfessionUCParams       pkgclientuc.CountClientsByProfessionUCParams
	CountClientsByStateUCParams            pkgclientuc.CountClientsByStateUCParams
	CountClientsByProfessionInCityUCParams pkgclientuc.CountClientsByProfessionInCityUCParams
	*/
}

func NewClientController(params *ClientControllerParams, g *echo.Group) {
	controller := ClientController{
		FindAllClientUCParams:                  params.FindAllClientUCParams,
		FindByIdUCParams:                             params.FindByIdUCParams,
		FindByNamedUCParams:                          params.FindByNamedUCParams,
		SaveClientUCParams:                     params.SaveClientUCParams,
		DeleteClientUCParams:                   params.DeleteClientUCParams,
		ExportXLSXClientUCParams:               params.ExportXLSXClientUCParams,
		FindLastClientsUCParams:                params.FindLastClientsUCParams,
		/*
		CountClientsByProfessionUCParams:       params.CountClientsByProfessionUCParams,
		CountClientsByStateUCParams:            params.CountClientsByStateUCParams,
		CountClientsByProfessionInCityUCParams: params.CountClientsByProfessionInCityUCParams,
		*/
	}

	g.POST("/client", controller.Save)
	g.GET("/clients", controller.FindAll)
	g.GET("/client/:id", controller.FindById)
	g.POST("/clients/name", controller.FindByName)
	g.DELETE("/client/:id", controller.Delete)
	g.POST("/last/clients", controller.FindLastClients)
	/*g.POST("/clients/state", controller.CountClientsByState)*/
	g.POST("/export-clients-XLSX", controller.ExportClientsXLSX)
	/*
	g.POST("/count/client/profession", controller.CountClientByProfession)
	g.POST("/count/clients/city", controller.CountClientsByProfessionInCity)
	*/
}

func (c *ClientController) FindByName(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	assembler := pkgclientuc.FindByNameAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	uc := pkgclientuc.NewFindByNamedUC(c.FindByNamedUCParams)
	uc.Assembler = &assembler
	clients, err := uc.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, clients)
}
/*
func (c *ClientController) CountClientsByProfessionInCity(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	var requestBody map[string]interface{}

	// Bind the request body to a map
	if err := ctx.Bind(&requestBody); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// Extrai o valor do requestBody como float64 (considerando que pode ser um tipo numérico compatível)
	cityIDFloat, ok := requestBody["cityID"].(float64)
	if !ok || cityIDFloat == 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "City Id is required"})
	}

	// Converte o float64 para uint
	cityID := uint(cityIDFloat)

	usecase := pkgclientuc.NewCountClientsByProfessionInCityUC(c.CountClientsByProfessionInCityUCParams)
	usecase.CityId = &cityID
	clients, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, clients)
}

func (c *ClientController) CountClientsByState(ctx echo.Context) error {
	defer ctx.Request().Body.Close()

	assembler := pkgclientuc.FindWithPaginationClientByStateAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	usecase := pkgclientuc.NewCountClientsByStateUC(c.CountClientsByStateUCParams)
	usecase.Assembler = &assembler
	clients, total, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}

	// Estrutura de resposta
	response := struct {
		Clients []*pkgclientuc.CityClientCountPresenter `json:"result"`
		Total         *int64                                              `json:"total"`
	}{
		Clients: clients,
		Total:         total,
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c *ClientController) CountClientByProfession(ctx echo.Context) error {

	defer ctx.Request().Body.Close()
	usecase := pkgclientuc.NewCountClientsByProfessionUC(c.CountClientsByProfessionUCParams)
	clients, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, clients)
}
*/
func (c *ClientController) ExportClientsXLSX(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgclientuc.NewExportXLSXClientUC(c.ExportXLSXClientUCParams)
	clients, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	file := excelize.NewFile()
	sheet := "Clientes"
	file.NewSheet(sheet)
	headers := []string{"ID", "Nome", "Email", "Telefone", "Cidade", "Estado"}
	err = file.SetSheetRow(sheet, "A1", &headers)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to set headers in Excel sheet"})
	}
	for i, p := range *clients {

		file.SetCellValue(sheet, fmt.Sprintf("A%d", i+2), p.ID)
		file.SetCellValue(sheet, fmt.Sprintf("B%d", i+2), p.Name)
		file.SetCellValue(sheet, fmt.Sprintf("C%d", i+2), p.Email)
		file.SetCellValue(sheet, fmt.Sprintf("D%d", i+2), p.Telephone)
		file.SetCellValue(sheet, fmt.Sprintf("E%d", i+2), p.Cidade.Name)
		file.SetCellValue(sheet, fmt.Sprintf("F%d", i+2), p.Cidade.UF)
		file.SetCellValue(sheet, fmt.Sprintf("G%d", i+2), p.Cep)
		file.SetActiveSheet(i)
		fmt.Println(p.Name)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	buffer, err := file.WriteToBuffer()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate Excel file"})
	}

	ctx.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename=clientes.xlsx")
	ctx.Response().Header().Set(echo.HeaderContentType, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	return ctx.Blob(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", buffer.Bytes())

}

func (c *ClientController) FindLastClients(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	var assembler pkgclientuc.FindLastClientsRequest
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	usecase := pkgclientuc.NewFindLastClientsUC(c.FindLastClientsUCParams)
	usecase.QuantityRecords = assembler.QuantityRecords
	clients, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, clients)
}

func (c *ClientController) FindAll(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	assembler := pkgclientuc.FindWithPaginationClientAssembler{}

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

	usecase := pkgclientuc.NewFindAllClientUC(c.FindAllClientUCParams)
	usecase.Assembler = assembler
	clients, total, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, map[string]string{"error": "Erro ao buscar clientes"})
	}
	// Estrutura de resposta
	response := struct {
		Clients *[]pkgclientuc.ClientPresenter `json:"clients"`
		Total         int64                    `json:"total"`
	}{
		Clients: clients,
		Total:         total,
	}

	return ctx.JSON(http.StatusOK, response)
}

func (c *ClientController) FindById(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgclientuc.NewFindByIdUC(c.FindByIdUCParams)
	idAssembler := ctx.Param("id")
	id, err := strconv.ParseUint(idAssembler, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	uintID := uint(id)

	usecase.ID = &uintID
	client, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, client)
}

func (c *ClientController) Save(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgclientuc.NewSaveClientUC(c.SaveClientUCParams)
	assembler := pkgclientuc.ClientAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	usecase.Assembler = &assembler

	client, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, client)

}


func (c *ClientController) Delete(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgclientuc.NewDeleteClientUC(c.DeleteClientUCParams)
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
