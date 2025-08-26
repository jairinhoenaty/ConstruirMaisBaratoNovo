package budget_usecase

import "time"

type BudgetPresenter struct {
	ID                         uint                     `json:"id,omitempty"`
	Name                       string                   `json:"name"`
	Email                      string                   `json:"email"`
	Telephone                  string                   `json:"telephone"`
	Description                string                   `json:"description"`
	CreatedAt                  time.Time                `json:"created_at"`
	Professionals              *[]ProfessionalPresenter `json:"professionals"`
	CityID                     *uint                    `json:"cityId"`
	City                       CityPresenter            `json:"city"`
	TermResponsabilityAccepted bool                     `json:"termResponsabilityAccepted"`
	// ClientID                   *uint                    `json:"clientId"`
	// Client                     ClientPresenter          `json:"client"`
	Approved                   bool                     `json:"approved"`
}

type CityPresenter struct {
	Name string `json:"name"`
	UF   string `json:"uf"`
}

type ProfessionalPresenter struct {
	ID          uint                   `json:"id,omitempty"`
	Name        string                 `json:"name"`
	Email       string                 `json:"email"`
	Telephone   string                 `json:"telephone"`
	Professions *[]ProfessionPresenter `json:"professions"`
	City        CityPresenter          `json:"city"`
}

type ProfessionPresenter struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name"`
}

// type ClientPresenter struct {
// 	ID        uint          `json:"id,omitempty"`
// 	Name      string        `json:"name"`
// 	Email     string        `json:"email"`
// 	Telephone string        `json:"telephone"`
// 	City      CityPresenter `json:"city"`
// }
