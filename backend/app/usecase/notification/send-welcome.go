package notification

import (
	"context"
	"regexp"
	"strings"
)

type UserLookup interface {
	FindByID(ctx context.Context, id uint) (*UserDTO, error)
}

type UserDTO struct {
	ID       uint
	Name     string
	Whatsapp string // ex: "(24) 99309-2521" ou "24993092521"
}

type WhatsSender interface {
	SendText(toE164, body string) error
}

type SendWelcomeUC struct {
	UserSvc UserLookup
	Wpp     WhatsSender
}

type SendWelcomeInput struct {
	UserID uint
}

func normalizePhoneToE164BR(raw string) string {
	only := regexp.MustCompile(`\D+`).ReplaceAllString(raw, "")
	only = strings.TrimLeft(only, "0")
	if !strings.HasPrefix(only, "55") {
		only = "55" + only
	}
	return only
}

func (uc *SendWelcomeUC) Execute(ctx context.Context, in SendWelcomeInput) error {
	u, err := uc.UserSvc.FindByID(ctx, in.UserID)
	if err != nil || u == nil || u.Whatsapp == "" {
		return err
	}
	to := normalizePhoneToE164BR(u.Whatsapp)

	msg := "🎉 Parabéns! Sua assinatura Premium foi confirmada com sucesso.\n\n" +
		"Agora você faz parte do grupo de profissionais que têm mais visibilidade, prioridade nos orçamentos e um perfil completo com fotos, vídeos e selo de confiança.\n\n" +
		"✨ Suas novas vantagens:\n" +
		"• Seu perfil aparece primeiro nas buscas dos clientes.\n" +
		"• Você pode mostrar seu portfólio de trabalhos realizados.\n" +
		"• Recebe orçamentos em primeira mão.\n" +
		"• E ainda conta com o selo Premium que destaca sua credibilidade.\n\n" +
		"Agradecemos por confiar na Construir Mais Barato.\n" +
		"Continue oferecendo o seu melhor — agora, com muito mais destaque e oportunidades!"

	return uc.Wpp.SendText(to, msg)
}
