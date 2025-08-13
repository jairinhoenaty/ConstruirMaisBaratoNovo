package controllers

import (
	pkgstoreuc "construir_mais_barato/app/usecase/store"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/xuri/excelize/v2"
)

type StoreController struct {
	FindAllStoreUCParams                  pkgstoreuc.FindAllStoreUCParams
	FindByIdUCParams                             pkgstoreuc.FindByIdUCParamns
	FindByNamedUCParams                          pkgstoreuc.FindByNamedUCParams
	SaveStoreUCParams                     pkgstoreuc.SaveStoreUCParams
	DeleteStoreUCParams                   pkgstoreuc.DeleteStoreUCParams
	FindLastStoresUCParams                pkgstoreuc.FindLastStoresUCParams	
	ExportXLSXStoreUCParams               pkgstoreuc.ExportXLSXStoreUCParams
	/*
	CountStoresByStateUCParams            pkgstoreuc.CountStoresByStateUCParams
	CountStoresByProfessionUCParams       pkgstoreuc.CountStoresByProfessionUCParams
	CountStoresByProfessionInCityUCParams pkgstoreuc.CountStoresByProfessionInCityUCParams*/
}

type StoreControllerParams struct {
	FindAllStoreUCParams                  pkgstoreuc.FindAllStoreUCParams
	FindByIdUCParams                             pkgstoreuc.FindByIdUCParamns
	FindByNamedUCParams                          pkgstoreuc.FindByNamedUCParams
	SaveStoreUCParams                     pkgstoreuc.SaveStoreUCParams
	DeleteStoreUCParams                   pkgstoreuc.DeleteStoreUCParams
	FindLastStoresUCParams                pkgstoreuc.FindLastStoresUCParams
	ExportXLSXStoreUCParams               pkgstoreuc.ExportXLSXStoreUCParams
	/*CountStoresByProfessionUCParams       pkgstoreuc.CountStoresByProfessionUCParams
	CountStoresByStateUCParams            pkgstoreuc.CountStoresByStateUCParams
	CountStoresByProfessionInCityUCParams pkgstoreuc.CountStoresByProfessionInCityUCParams
	*/
}

func NewStoreController(params *StoreControllerParams, g *echo.Group) {
	controller := StoreController{
		FindAllStoreUCParams:                  params.FindAllStoreUCParams,
		FindByIdUCParams:                             params.FindByIdUCParams,
		FindByNamedUCParams:                          params.FindByNamedUCParams,
		SaveStoreUCParams:                     params.SaveStoreUCParams,
		DeleteStoreUCParams:                   params.DeleteStoreUCParams,
		ExportXLSXStoreUCParams:               params.ExportXLSXStoreUCParams,
		FindLastStoresUCParams:                params.FindLastStoresUCParams,
		/*
		CountStoresByProfessionUCParams:       params.CountStoresByProfessionUCParams,
		CountStoresByStateUCParams:            params.CountStoresByStateUCParams,
		CountStoresByProfessionInCityUCParams: params.CountStoresByProfessionInCityUCParams,
		*/
	}

	g.POST("/store", controller.Save)
	g.GET("/stores", controller.FindAll)
	g.GET("/store/:id", controller.FindById)
	g.POST("/stores/name", controller.FindByName)
	g.DELETE("/store/:id", controller.Delete)
	g.POST("/last/stores", controller.FindLastStores)
	/*
	g.POST("/stores/state", controller.CountStoresByState)
	*/
	g.POST("/export-stores-XLSX", controller.ExportStoresXLSX)
	/*g.POST("/count/store/profession", controller.CountStoreByProfession)
	g.POST("/count/stores/city", controller.CountStoresByProfessionInCity)
	*/
}

func (c *StoreController) FindByName(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	assembler := pkgstoreuc.FindByNameAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	uc := pkgstoreuc.NewFindByNamedUC(c.FindByNamedUCParams)
	uc.Assembler = &assembler
	stores, err := uc.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, stores)
}
/*
func (c *StoreController) CountStoresByProfessionInCity(ctx echo.Context) error {
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

	usecase := pkgstoreuc.NewCountStoresByProfessionInCityUC(c.CountStoresByProfessionInCityUCParams)
	usecase.CityId = &cityID
	stores, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, stores)
}

func (c *StoreController) CountStoresByState(ctx echo.Context) error {
	defer ctx.Request().Body.Close()

	assembler := pkgstoreuc.FindWithPaginationStoreByStateAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	usecase := pkgstoreuc.NewCountStoresByStateUC(c.CountStoresByStateUCParams)
	usecase.Assembler = &assembler
	stores, total, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}

	// Estrutura de resposta
	response := struct {
		Stores []*pkgstoreuc.CityStoreCountPresenter `json:"result"`
		Total         *int64                                              `json:"total"`
	}{
		Stores: stores,
		Total:         total,
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c *StoreController) CountStoreByProfession(ctx echo.Context) error {

	defer ctx.Request().Body.Close()
	usecase := pkgstoreuc.NewCountStoresByProfessionUC(c.CountStoresByProfessionUCParams)
	stores, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, stores)
}
*/
func (c *StoreController) ExportStoresXLSX(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgstoreuc.NewExportXLSXStoreUC(c.ExportXLSXStoreUCParams)
	stores, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}

	file := excelize.NewFile()
	sheet := "Logistas"
	file.NewSheet(sheet)
	headers := []string{"ID", "Nome", "Email", "Telefone", "Cidade", "Estado"}
	err = file.SetSheetRow(sheet, "A1", &headers)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to set headers in Excel sheet"})
	}
	for i, p := range *stores {

		file.SetCellValue(sheet, fmt.Sprintf("A%d", i+2), p.ID)
		file.SetCellValue(sheet, fmt.Sprintf("B%d", i+2), p.Name)
		file.SetCellValue(sheet, fmt.Sprintf("C%d", i+2), p.Email)
		file.SetCellValue(sheet, fmt.Sprintf("D%d", i+2), p.Telephone)
		file.SetCellValue(sheet, fmt.Sprintf("E%d", i+2), p.Cidade.Name)
		file.SetCellValue(sheet, fmt.Sprintf("F%d", i+2), p.Cidade.UF)
		file.SetCellValue(sheet, fmt.Sprintf("G%d", i+2), p.Cep)
		file.SetActiveSheet(i)
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

	ctx.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename=logistas.xlsx")
	ctx.Response().Header().Set(echo.HeaderContentType, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	return ctx.Blob(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", buffer.Bytes())

}

func (c *StoreController) FindLastStores(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	var assembler pkgstoreuc.FindLastStoresRequest
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	usecase := pkgstoreuc.NewFindLastStoresUC(c.FindLastStoresUCParams)
	usecase.QuantityRecords = assembler.QuantityRecords
	stores, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, stores)
}

func (c *StoreController) FindAll(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	assembler := pkgstoreuc.FindWithPaginationStoreAssembler{}

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

	usecase := pkgstoreuc.NewFindAllStoreUC(c.FindAllStoreUCParams)
	usecase.Assembler = assembler
	stores, total, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, map[string]string{"error": "Erro ao buscar profissionais"})
	}
	// Estrutura de resposta
	response := struct {
		Stores *[]pkgstoreuc.StorePresenter `json:"stores"`
		Total         int64                 `json:"total"`
	}{
		Stores: stores,
		Total:         total,
	}

	return ctx.JSON(http.StatusOK, response)
}

func (c *StoreController) FindById(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgstoreuc.NewFindByIdUC(c.FindByIdUCParams)
	idAssembler := ctx.Param("id")
	id, err := strconv.ParseUint(idAssembler, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	uintID := uint(id)

	usecase.ID = &uintID
	store, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, store)
}

func (c *StoreController) Save(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgstoreuc.NewSaveStoreUC(c.SaveStoreUCParams)
	assembler := pkgstoreuc.StoreAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	usecase.Assembler = &assembler

	store, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, store)

}


func (c *StoreController) Delete(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgstoreuc.NewDeleteStoreUC(c.DeleteStoreUCParams)
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
