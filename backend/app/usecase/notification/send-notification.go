package notification_usecase

import (
	"context"
	"fmt"

	pkguser "construir_mais_barato/app/domain/user"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

type SendAppNotificationUC struct {
	UserService         pkguser.UserService
	Assembler           *SendAppNotificationAssembler
	FirebaseCredentials string
}

type SendAppNotificationUCParams struct {
	UserService         pkguser.UserService
	FirebaseCredentials string
}

func NewSendAppNotificationUC(params SendAppNotificationUCParams) SendAppNotificationUC {
	return SendAppNotificationUC{
		UserService:         params.UserService,
		FirebaseCredentials: params.FirebaseCredentials,
	}
}

func (uc *SendAppNotificationUC) Execute() error {
	tokens, err := uc.UserService.FindTokensByIds(uc.Assembler.Ids, uc.Assembler.IDType)
	if err != nil {
		return fmt.Errorf("erro ao buscar tokens: %w", err)
	}

	if len(tokens) == 0 {
		return nil
	}

	ctx := context.Background()
	opt := option.WithCredentialsFile(uc.FirebaseCredentials)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return fmt.Errorf("erro ao inicializar firebase: %w", err)
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		return fmt.Errorf("erro ao criar cliente FCM: %w", err)
	}

	message := &messaging.MulticastMessage{
		Tokens: tokens,
		Android: &messaging.AndroidConfig{
			Notification: &messaging.AndroidNotification{
				// Canal com som personalizado criado no app (ver main.dart).
				ChannelID: "professional_alert_channel",
				// Nome do recurso em android/app/src/main/res/raw (sem extensão).
				Sound:    "notification_sound",
				Priority: messaging.PriorityHigh,
			},
		},
		APNS: &messaging.APNSConfig{
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					Sound: "notification_sound.aiff",
				},
			},
		},
		Notification: &messaging.Notification{
			Title: uc.Assembler.Title,
			Body:  uc.Assembler.Body,
		},
	}

	_, err = client.SendEachForMulticast(ctx, message)
	if err != nil {
		return fmt.Errorf("erro ao enviar notificações FCM: %w", err)
	}

	return nil
}
