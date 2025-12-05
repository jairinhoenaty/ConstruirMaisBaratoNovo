package budget_usecase

type BudgetAssembler struct {
	ID                         uint    `json:"id,omitempty"`
	Name                       string  `json:"name"`
	Email                      string  `json:"email"`
	Telephone                  string  `json:"telephone"`
	Description                string  `json:"description"`
	ProfessionalsId            *[]uint `json:"professionalsId"`
	StoresId                   *[]uint `json:"storesId"`
	CityID                     *uint   `json:"cityId"`
	TermResponsabilityAccepted bool    `json:"termResponsabilityAccepted"`
	CategoryId                 int     `json:"categoryId"`
	// ClientID                   *uint   `json:"clientId"`
	Approved bool `json:"approved"`
}

type ProfessionalAssembler struct {
	ID          uint                   `json:"id,omitempty"`
	Name        string                 `json:"name"`
	Email       string                 `json:"email"`
	Telephone   string                 `json:"telephone"`
	Professions *[]ProfessionAssembler `json:"professions"`
}

type ProfessionAssembler struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name"`
}

type FindBudgetByMontAndProfessionalIDAssembler struct {
	Page           uint   `json:"page"`
	Month          string `json:"month"`
	PageSize       uint   `json:"pagesize"`
	ProfessionalID *uint  `json:"professionalID"`
	StoreID        *uint  `json:"storeID"`
	ClientID       int    `json:"clientID"`
}

type FindWithPaginationBudgetAssembler struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type FindBudgetByEmailAssembler struct {
	Email string `json:"email"`
}
