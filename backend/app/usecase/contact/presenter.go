package contact_usecase

import (
	"construir_mais_barato/app/domain/city"
	"construir_mais_barato/app/domain/client"
	"construir_mais_barato/app/domain/product"
	"construir_mais_barato/app/domain/professional"
	"construir_mais_barato/app/domain/store"
	"time"
)

type ContactPresenter struct {
	ID             uint                      `json:"id,omitempty"`
	Name           string                    `json:"name"`
	Telephone      string                    `json:"telefone"`
	Email          string                    `json:"email"`
	Message        string                    `json:"mensagem"`
	Status         string                    `json:"status"`
	CityID         *uint                     `json:"city_id"`
	City           city.City                 `json:"city"`
	ProfessionalID *uint                     `json:"professional_id"`
	Professional   professional.Professional `json:"professional"`
	ClientID       *uint                     `json:"client_id"`
	Client         client.Client             `json:"client"`
	StoreID        *uint                     `json:"store_id"`
	Store          store.Store               `json:"store"`
	ProductID      *uint                     `json:"product_id"`
	Product        product.Product           `json:"product"`
	CreatedAt      time.Time                 `json:"created_at"`
	Approved       bool                      `json:"approved"`
}
