package controllers

import (
	pkgprofessionaluc "construir_mais_barato/app/usecase/professional"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/xuri/excelize/v2"
)

type ProfessionalController struct {
	FindAllProfessionalUCParams                  pkgprofessionaluc.FindAllProfessionalUCParams
	FindByIdUCParams                             pkgprofessionaluc.FindByIdUCParamns
	FindByNamedUCParams                          pkgprofessionaluc.FindByNamedUCParams
	FindByProfessionAndLocationUCParams          pkgprofessionaluc.FindByProfessionAndLocationUCParams
	SaveProfessionalUCParams                     pkgprofessionaluc.SaveProfessionalUCParams
	DeleteProfessionalUCParams                   pkgprofessionaluc.DeleteProfessionalUCParams
	FindLastProfessionalsUCParams                pkgprofessionaluc.FindLastProfessionalsUCParams
	ExportXLSXProfessionalUCParams               pkgprofessionaluc.ExportXLSXProfessionalUCParams
	CountProfessionalsByStateUCParams            pkgprofessionaluc.CountProfessionalsByStateUCParams
	CountProfessionalsByProfessionUCParams       pkgprofessionaluc.CountProfessionalsByProfessionUCParams
	CountProfessionalsByProfessionInCityUCParams pkgprofessionaluc.CountProfessionalsByProfessionInCityUCParams
	CountCityProfessionalsByStateUCParams        pkgprofessionaluc.CountCityProfessionalsByStateUCParams
}

type ProfessionalControllerParams struct {
	FindAllProfessionalUCParams                  pkgprofessionaluc.FindAllProfessionalUCParams
	FindByIdUCParams                             pkgprofessionaluc.FindByIdUCParamns
	FindByNamedUCParams                          pkgprofessionaluc.FindByNamedUCParams
	FindByProfessionAndLocationUCParams          pkgprofessionaluc.FindByProfessionAndLocationUCParams
	SaveProfessionalUCParams                     pkgprofessionaluc.SaveProfessionalUCParams
	DeleteProfessionalUCParams                   pkgprofessionaluc.DeleteProfessionalUCParams
	FindLastProfessionalsUCParams                pkgprofessionaluc.FindLastProfessionalsUCParams
	ExportXLSXProfessionalUCParams               pkgprofessionaluc.ExportXLSXProfessionalUCParams
	CountProfessionalsByProfessionUCParams       pkgprofessionaluc.CountProfessionalsByProfessionUCParams
	CountProfessionalsByStateUCParams            pkgprofessionaluc.CountProfessionalsByStateUCParams
	CountProfessionalsByProfessionInCityUCParams pkgprofessionaluc.CountProfessionalsByProfessionInCityUCParams
	CountCityProfessionalsByStateUCParams        pkgprofessionaluc.CountCityProfessionalsByStateUCParams
}

func NewProfessionalController(params *ProfessionalControllerParams, g *echo.Group) {
	controller := ProfessionalController{
		FindAllProfessionalUCParams:                  params.FindAllProfessionalUCParams,
		FindByIdUCParams:                             params.FindByIdUCParams,
		FindByNamedUCParams:                          params.FindByNamedUCParams,
		FindByProfessionAndLocationUCParams:          params.FindByProfessionAndLocationUCParams,
		SaveProfessionalUCParams:                     params.SaveProfessionalUCParams,
		DeleteProfessionalUCParams:                   params.DeleteProfessionalUCParams,
		ExportXLSXProfessionalUCParams:               params.ExportXLSXProfessionalUCParams,
		FindLastProfessionalsUCParams:                params.FindLastProfessionalsUCParams,
		CountProfessionalsByProfessionUCParams:       params.CountProfessionalsByProfessionUCParams,
		CountProfessionalsByStateUCParams:            params.CountProfessionalsByStateUCParams,
		CountProfessionalsByProfessionInCityUCParams: params.CountProfessionalsByProfessionInCityUCParams, CountCityProfessionalsByStateUCParams: params.CountCityProfessionalsByStateUCParams,
	}

	g.POST("/professional", controller.Save)
	g.GET("/professionals", controller.FindAll)
	g.GET("/professional/:id", controller.FindById)
	g.POST("/professionals/name", controller.FindByName)
	g.POST("/professionals/profession-location", controller.PublicFindByProfessionAndLocation)
	g.DELETE("/professional/:id", controller.Delete)
	g.POST("/last/professionals", controller.FindLastProfessionals)
	g.POST("/professionals/state", controller.CountCityProfessionalsByState)
	g.POST("/count/professionals/state", controller.CountProfessionalsByState)
	g.POST("/export-professionals-XLSX", controller.ExportProfessionalsXLSX)
	g.POST("/count/professional/profession", controller.CountProfessionalByProfession)
	g.POST("/count/professionals/city", controller.CountProfessionalsByProfessionInCity)

}

func (c *ProfessionalController) FindByName(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	assembler := pkgprofessionaluc.FindByNameAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	uc := pkgprofessionaluc.NewFindByNamedUC(c.FindByNamedUCParams)
	uc.Assembler = &assembler
	professionals, err := uc.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, professionals)
}

func (c *ProfessionalController) CountProfessionalsByProfessionInCity(ctx echo.Context) error {
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

	usecase := pkgprofessionaluc.NewCountProfessionalsByProfessionInCityUC(c.CountProfessionalsByProfessionInCityUCParams)
	usecase.CityId = &cityID
	professionals, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, professionals)
}

func (c *ProfessionalController) CountProfessionalsByState(ctx echo.Context) error {
	defer ctx.Request().Body.Close()

	assembler := pkgprofessionaluc.FindWithPaginationProfessionalByStateAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	usecase := pkgprofessionaluc.NewCountProfessionalsByStateUC(c.CountProfessionalsByStateUCParams)
	usecase.Assembler = &assembler
	professionals, total, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}

	// Estrutura de resposta
	response := struct {
		Professionals []*pkgprofessionaluc.UFProfessionalCountPresenter `json:"result"`
		Total         *int64                                            `json:"total"`
	}{
		Professionals: professionals,
		Total:         total,
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c *ProfessionalController) CountCityProfessionalsByState(ctx echo.Context) error {
	defer ctx.Request().Body.Close()

	assembler := pkgprofessionaluc.FindWithPaginationProfessionalByStateAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	usecase := pkgprofessionaluc.NewCountCityProfessionalsByStateUC(c.CountCityProfessionalsByStateUCParams)
	usecase.Assembler = &assembler
	professionals, total, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}

	// Estrutura de resposta
	response := struct {
		Professionals []*pkgprofessionaluc.CityProfessionalCountPresenter `json:"result"`
		Total         *int64                                              `json:"total"`
	}{
		Professionals: professionals,
		Total:         total,
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c *ProfessionalController) CountProfessionalByProfession(ctx echo.Context) error {

	defer ctx.Request().Body.Close()
	usecase := pkgprofessionaluc.NewCountProfessionalsByProfessionUC(c.CountProfessionalsByProfessionUCParams)
	professionals, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, professionals)
}

func (c *ProfessionalController) ExportProfessionalsXLSX(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgprofessionaluc.NewExportXLSXProfessionalUC(c.ExportXLSXProfessionalUCParams)
	professionals, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}

	file := excelize.NewFile()
	sheet := "Profissionais"
	file.NewSheet(sheet)
	headers := []string{"ID", "Nome", "Email", "Telefone", "Profissões", "Cidade", "Estado"}
	err = file.SetSheetRow(sheet, "A1", &headers)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to set headers in Excel sheet"})
	}
	for i, p := range *professionals {
		var professions []string
		for _, profession := range p.Professions {
			professions = append(professions, profession.Name)
		}
		professionNames := strings.Join(professions, ", ")

		file.SetCellValue(sheet, fmt.Sprintf("A%d", i+2), p.ID)
		file.SetCellValue(sheet, fmt.Sprintf("B%d", i+2), p.Name)
		file.SetCellValue(sheet, fmt.Sprintf("C%d", i+2), p.Email)
		file.SetCellValue(sheet, fmt.Sprintf("D%d", i+2), p.Telephone)
		file.SetCellValue(sheet, fmt.Sprintf("E%d", i+2), professionNames)
		file.SetCellValue(sheet, fmt.Sprintf("F%d", i+2), p.Cidade.Name)
		file.SetCellValue(sheet, fmt.Sprintf("G%d", i+2), p.Cidade.UF)
		file.SetCellValue(sheet, fmt.Sprintf("H%d", i+2), p.Cep)
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

	ctx.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename=profissionais.xlsx")
	ctx.Response().Header().Set(echo.HeaderContentType, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	return ctx.Blob(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", buffer.Bytes())

}

func (c *ProfessionalController) FindLastProfessionals(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	var assembler pkgprofessionaluc.FindLastProfessionalsRequest
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	usecase := pkgprofessionaluc.NewFindLastProfessionalsUC(c.FindLastProfessionalsUCParams)
	usecase.QuantityRecords = assembler.QuantityRecords
	professionals, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, professionals)
}

func (c *ProfessionalController) FindAll(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	assembler := pkgprofessionaluc.FindWithPaginationProfessionalAssembler{}

	// Extrair parâmetros de paginação da query string
	limit := ctx.QueryParam("limit")
	offset := ctx.QueryParam("offset")
	filter := ctx.QueryParam("filter")
	order := ctx.QueryParam("order")
	uf := ctx.QueryParam("uf")
	profession_id := ctx.QueryParam("profession_id")
	profession_idInt, err := strconv.Atoi(profession_id)
	if err != nil || profession_idInt <= 0 {
		profession_idInt = 0 // valor padrão
	}

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
	assembler.Filter = filter
	assembler.Order = order
	assembler.Uf = uf
	assembler.ProfessionId = profession_idInt

	usecase := pkgprofessionaluc.NewFindAllProfessionalUC(c.FindAllProfessionalUCParams)
	usecase.Assembler = assembler
	professionals, total, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, map[string]string{"error": "Erro ao buscar profissionais"})
	}
	// Estrutura de resposta
	response := struct {
		Professionals *[]pkgprofessionaluc.ProfessionalPresenter `json:"profissionais"`
		Total         int64                                      `json:"total"`
	}{
		Professionals: professionals,
		Total:         total,
	}

	return ctx.JSON(http.StatusOK, response)
}

func (c *ProfessionalController) FindById(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgprofessionaluc.NewFindByIdUC(c.FindByIdUCParams)
	idAssembler := ctx.Param("id")
	id, err := strconv.ParseUint(idAssembler, 10, 32)

	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	uintID := uint(id)

	usecase.ID = &uintID
	professional, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, professional)
}

func (c *ProfessionalController) Save(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgprofessionaluc.NewSaveProfessionalUC(c.SaveProfessionalUCParams)
	assembler := pkgprofessionaluc.ProfessionalAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}

	usecase.Assembler = &assembler

	professional, err := usecase.Execute()
	if err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
	return ctx.JSON(http.StatusOK, professional)

}

func (c *ProfessionalController) Delete(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	usecase := pkgprofessionaluc.NewDeleteProfessionalUC(c.DeleteProfessionalUCParams)
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

func (c *ProfessionalController) PublicFindByProfessionAndLocation(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	assembler := pkgprofessionaluc.FindByProfessionAndLocationAssembler{}
	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusPreconditionFailed, err)
	}
	
	usecase := pkgprofessionaluc.NewFindByProfessionAndLocationUC(c.FindByProfessionAndLocationUCParams)
	usecase.Assembler = &assembler

	//print(usecase.Assembler.Distance)

	if usecase.Assembler.Distance == 0 {
		return ctx.JSON(http.StatusPreconditionFailed, nil)
	}
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
