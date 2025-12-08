package controllers

import (
	jobuc "construir_mais_barato/app/usecase/job"
	jobmodel "construir_mais_barato/app/domain/job"
	"github.com/labstack/echo/v4"
	"strconv"
)

type JobController struct {
	UC *jobuc.JobUseCase
}

func NewJobController(uc *jobuc.JobUseCase, g *echo.Group) {
	controller := &JobController{UC: uc}

	g.POST("/jobs", controller.Save)
	g.GET("/jobs", controller.ListAll)
	g.GET("/jobs/approved", controller.ListApproved)
	g.PUT("/jobs/approve/:id", controller.Approve)
	g.PUT("/jobs/disapprove/:id", controller.Disapprove)
	g.DELETE("/jobs/:id", controller.Delete)
}

func (c *JobController) Save(ctx echo.Context) error {
	var req jobmodel.Job
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(400, err)
	}
	result, err := c.UC.Save(req)
	if err != nil {
		return ctx.JSON(500, err)
	}
	return ctx.JSON(201, result)
}

func (c *JobController) ListAll(ctx echo.Context) error {
	r, err := c.UC.ListAll()
	if err != nil {
		return ctx.JSON(500, err)
	}
	return ctx.JSON(200, r)
}

func (c *JobController) ListApproved(ctx echo.Context) error {
	r, err := c.UC.ListApproved()
	if err != nil {
		return ctx.JSON(500, err)
	}
	return ctx.JSON(200, r)
}

func (c *JobController) Approve(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	return useJSON(ctx, c.UC.Approve(uint(id)))
}

func (c *JobController) Disapprove(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	return useJSON(ctx, c.UC.Disapprove(uint(id)))
}

func (c *JobController) Delete(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	return useJSON(ctx, c.UC.Delete(uint(id)))
}

func useJSON(ctx echo.Context, err error) error {
	if err != nil {
		return ctx.JSON(500, err)
	}
	return ctx.JSON(200, "ok")
}
