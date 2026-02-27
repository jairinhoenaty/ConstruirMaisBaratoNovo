package controllers

import (
	pkgnotificationuc "construir_mais_barato/app/usecase/notification"
	"net/http"

	"github.com/labstack/echo/v4"
)

type NotificationController struct {
	SendAppNotificationUCParams pkgnotificationuc.SendAppNotificationUCParams
}

type NotificationControllerParams struct {
	SendAppNotificationUCParams pkgnotificationuc.SendAppNotificationUCParams
}

func NewNotificationController(params *NotificationControllerParams, g *echo.Group) {
	controller := NotificationController{
		SendAppNotificationUCParams: params.SendAppNotificationUCParams,
	}

	g.POST("/notifications/send-app", controller.SendApp)
}

func (c *NotificationController) SendApp(ctx echo.Context) error {
	defer ctx.Request().Body.Close()

	usecase := pkgnotificationuc.NewSendAppNotificationUC(c.SendAppNotificationUCParams)
	assembler := pkgnotificationuc.SendAppNotificationAssembler{}

	if err := ctx.Bind(&assembler); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Formato de requisição inválido"})
	}

	usecase.Assembler = &assembler

	if err := usecase.Execute(); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "Notificações enviadas com sucesso"})
}
