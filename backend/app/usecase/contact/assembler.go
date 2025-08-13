package contact_usecase

type ContactAssembler struct {
	ID             uint   `json:"id,omitempty"`
	Name           string `json:"name"`
	Telephone      string `json:"telefone"`
	Email          string `json:"email"`
	Mensagem       string `json:"mensagem"`
	Status         string `json:"status"`
	CityID         *uint  `json:"city_id"`
	ProfessionalID *uint  `json:"professional_id"`
	ClientID       *uint  `json:"client_id"`
	StoreID        *uint  `json:"store_id"`
	ProductID      *uint  `json:"product_id"`
	Approved       bool   `json:"approved"`
}

type FindWithPaginationUserAssembler struct {
	Limit          int  `json:"limit"`
	Offset         int  `json:"offset"`
	ProfessionalID uint `json:"professional_id"`
	ClientID       uint `json:"client_id"`
	StoreID        uint `json:"store_id"`
}

type FindWithPaginationContactAssembler struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
