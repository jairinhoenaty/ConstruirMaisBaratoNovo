package budget_usecase

import (
	"fmt"
	"strings"
	"time"
)

type BudgetPresenter struct {
	ID                         uint                     `json:"id,omitempty"`
	Name                       string                   `json:"name"`
	Email                      string                   `json:"email"`
	Telephone                  string                   `json:"telephone"`
	Description                string                   `json:"description"`
	CreatedAt                  time.Time                `json:"created_at"`
	Professionals              *[]ProfessionalPresenter `json:"professionals"`
	Stores                     *[]StorePresenter        `json:"stores"`
	CityID                     *uint                    `json:"cityId"`
	City                       CityPresenter            `json:"city"`
	TermResponsabilityAccepted bool                     `json:"termResponsabilityAccepted"`
	// ClientID                   *uint                    `json:"clientId"`
	// Client                     ClientPresenter          `json:"client"`
	Approved bool `json:"approved"`
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
type StorePresenter struct {
	ID                uint                   `json:"id,omitempty"`
	Name              string                 `json:"name"`
	Email             string                 `json:"email"`
	Telephone         string                 `json:"telephone"`
	Professions       *[]ProfessionPresenter `json:"professions"`
	City              CityPresenter          `json:"city"`
	CategoryProductID uint
	// SubCategories     UintSlice `gorm:"type:json"`
}

// Name              string
// 	Company           string
// 	Email             string `gorm:"unique"`
// 	Telephone         string
// 	LgpdAceito        string
// 	CityID            uint
// 	City              pkgcity.City `gorm:"foreignKey:CityID"`
// 	Cep               string
// 	Street            string
// 	Neighborhood      string
// 	Image             []byte    `gorm:"type:longblob"`
// 	CreatedAt         time.Time `gorm:"<-:create"`
// 	IsPremiumStore    *bool     `gorm:"default:false"`
// 	CategoryProductID uint
// 	SubCategories     UintSlice `gorm:"type:json"`

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

// MaskedBudgetPresenter é usado para retornar dados de orçamento mascarados
// para usuários sem acesso (não-premium, não-desbloqueado)
type MaskedBudgetPresenter struct {
	ID                         uint                     `json:"id,omitempty"`
	Name                       string                   `json:"name"`          // Mascarado
	Email                      string                   `json:"email"`         // Mascarado
	Telephone                  string                   `json:"telephone"`     // Mascarado
	Description                string                   `json:"description"`   // Completo
	CreatedAt                  time.Time                `json:"created_at"`    // Completo
	Professionals              *[]ProfessionalPresenter `json:"professionals"` // Pode mostrar
	CityID                     *uint                    `json:"cityId"`        // Completo
	City                       CityPresenter            `json:"city"`          // Completo
	TermResponsabilityAccepted bool                     `json:"termResponsabilityAccepted"`
	Approved                   bool                     `json:"approved"`
	IsLocked                   bool                     `json:"isLocked"` // Indica que dados estão mascarados
}

// MaskString mascara uma string deixando apenas os primeiros visibleChars caracteres
func MaskString(s string, visibleChars int) string {
	if len(s) <= visibleChars {
		return strings.Repeat("*", len(s))
	}
	return s[:visibleChars] + strings.Repeat("*", len(s)-visibleChars)
}

// MaskEmail mascara um email (ex: jo***@gmail.com)
func MaskEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return "***@***.com"
	}
	username := parts[0]
	domain := parts[1]

	if len(username) <= 2 {
		return strings.Repeat("*", len(username)) + "@" + domain
	}
	return username[:2] + strings.Repeat("*", len(username)-2) + "@" + domain
}

// MaskPhone mascara um telefone (ex: (XX) *****-1234)
func MaskPhone(phone string) string {
	// Remove non-numeric characters
	cleaned := strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, phone)

	if len(cleaned) < 4 {
		return strings.Repeat("*", len(phone))
	}

	// Show last 4 digits
	lastFour := cleaned[len(cleaned)-4:]

	// Format back
	if len(cleaned) >= 10 {
		// (XX) XXXXX-1234 -> (XX) *****-1234
		return fmt.Sprintf("(XX) *****-%s", lastFour)
	}

	// Fallback format
	return strings.Repeat("*", len(cleaned)-4) + lastFour
}

// GenerateMaskedBudgetPresenter cria um presenter com dados mascarados
func GenerateMaskedBudgetPresenter(budget interface{}) MaskedBudgetPresenter {
	// Converter interface{} para *BudgetPresenter
	budgetPresenter, ok := budget.(*BudgetPresenter)
	if !ok {
		return MaskedBudgetPresenter{}
	}

	return MaskedBudgetPresenter{
		ID:                         budgetPresenter.ID,
		Name:                       MaskString(budgetPresenter.Name, 3),
		Email:                      MaskEmail(budgetPresenter.Email),
		Telephone:                  MaskPhone(budgetPresenter.Telephone),
		Description:                budgetPresenter.Description,
		CreatedAt:                  budgetPresenter.CreatedAt,
		Professionals:              budgetPresenter.Professionals,
		CityID:                     budgetPresenter.CityID,
		City:                       budgetPresenter.City,
		TermResponsabilityAccepted: budgetPresenter.TermResponsabilityAccepted,
		Approved:                   budgetPresenter.Approved,
		IsLocked:                   true,
	}
}
