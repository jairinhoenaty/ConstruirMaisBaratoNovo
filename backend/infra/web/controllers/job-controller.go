package controllers

import (
    pkgjob "construir_mais_barato/app/domain/job"
    pkgjobuc "construir_mais_barato/app/usecase/job"

    "github.com/labstack/echo/v4"
    "net/http"
)

type JobControllerParams struct {
    SaveJobUCParams   pkgjobuc.SaveJobUCParams
    FindAllJobUCParams pkgjobuc.FindAllJobUCParams
    DeleteJobUCParams pkgjobuc.DeleteJobUCParams
}

type JobController struct {
    saveUC   pkgjobuc.SaveJobUC
    listUC   pkgjobuc.FindAllJobUC
    deleteUC pkgjobuc.DeleteJobUC
}

func NewJobController(params *JobControllerParams, privateGroup, publicGroup *echo.Group) {

    controller := &JobController{
        saveUC:   pkgjobuc.NewSaveJobUC(params.SaveJobUCParams),
        listUC:   pkgjobuc.NewFindAllJobUC(params.FindAllJobUCParams),
        deleteUC: pkgjobuc.NewDeleteJobUC(params.DeleteJobUCParams),
    }

    // rota pública
    publicGroup.POST("/job", controller.Save)

    // rotas privadas
    privateGroup.GET("/jobs", controller.FindAll)
    privateGroup.DELETE("/job/:id", controller.Delete)
}

func (c *JobController) Save(ctx echo.Context) error {
    var dto pkgjob.JobDTO
    if err := ctx.Bind(&dto); err != nil {
        return ctx.JSON(http.StatusBadRequest, err)
    }

    result, err := c.saveUC.Execute(&dto)
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err)
    }

    return ctx.JSON(http.StatusOK, result)
}

func (c *JobController) FindAll(ctx echo.Context) error {
    result, err := c.listUC.Execute()
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err)
    }
    return ctx.JSON(http.StatusOK, result)
}

func (c *JobController) Delete(ctx echo.Context) error {
    id := ctx.Param("id")
    err := c.deleteUC.Execute(id)
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err)
    }
    return ctx.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
