package controllers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	notifuc "construir_mais_barato/app/usecase/notification"
	professionaldomain "construir_mais_barato/app/domain/professional"
	wpp "construir_mais_barato/infra/adapters/whatsapp"
)

type professionalLookupAdapter struct {
	svc professionaldomain.ProfessionalService
}

func (p *professionalLookupAdapter) FindByID(ctx context.Context, id uint) (*notifuc.UserDTO, error) {
	prof, err := p.svc.FindById(id)
	if err != nil || prof == nil {
		return nil, err
	}

	return &notifuc.UserDTO{
		ID:       prof.ID,
		Name:     prof.Name,
		Whatsapp: prof.Telephone,
	}, nil
}

type NotificationController struct {
	UC *notifuc.SendWelcomeUC
}

func NewNotificationController(g *echo.Group, professionalSvc professionaldomain.ProfessionalService) {
	wapi := wpp.NewCloudAPI()
	uc := &notifuc.SendWelcomeUC{
		UserSvc: &professionalLookupAdapter{svc: professionalSvc},
		Wpp:     wapi,
	}
	ctrl := &NotificationController{UC: uc}

	// Rota pública para ser chamada logo após o cadastro
	g.POST("/notifications/welcome", ctrl.SendWelcome)
}

type sendWelcomeReq struct {
	UserID uint `json:"userId"`
}

func (c *NotificationController) SendWelcome(ctx echo.Context) error {
	var req sendWelcomeReq
	if err := ctx.Bind(&req); err != nil || req.UserID == 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid userId"})
	}

	if err := c.UC.Execute(ctx.Request().Context(), notifuc.SendWelcomeInput{UserID: req.UserID}); err != nil {

		return ctx.JSON(http.StatusOK, map[string]string{"status": "queued"})
	}
	return ctx.JSON(http.StatusOK, map[string]string{"status": "sent"})
}
