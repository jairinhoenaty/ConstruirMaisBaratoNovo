package chat_usecase

import (
	"construir_mais_barato/app/domain/client"
	"construir_mais_barato/app/domain/professional"
	"time"
)

type ChatPresenter struct {
	ID             uint                      `json:"id,omitempty"`
	Message        string                    `json:"message"`
	ProfessionalID *uint                     `json:"professional_id"`
	Professional   professional.Professional `json:"professional"`
	ClientID       *uint                     `json:"client_id"`
	Client         client.Client             `json:"client"`
	CreatedAt      time.Time                 `json:"created_at"`
	Origem         string                    `json:"origem"`
}
