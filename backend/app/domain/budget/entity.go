package budget

import (
	"time"

	"gorm.io/gorm"

	pkgcity "construir_mais_barato/app/domain/city"
	// pkgclient "construir_mais_barato/app/domain/client"
	pkgprofessional "construir_mais_barato/app/domain/professional"
	pkgstore "construir_mais_barato/app/domain/store"
)

type Budget struct {
	gorm.Model
	Name                       string
	Email                      string
	Telephone                  string
	Description                string
	ProfessionalIDs            *[]uint                        `gorm:"-"`
	Professionals              []pkgprofessional.Professional `gorm:"many2many:budgets_professionals;"`
	Stores                     []pkgstore.Store               `gorm:"many2many:budgets_stores;"`
	StoresIDs                  *[]uint                        `gorm:"-"`
	CityID                     *uint
	City                       pkgcity.City `gorm:"foreignKey:CityID"`
	TermResponsabilityAccepted bool
	// ClientID                   *uint `gorm:"default:null"`
	// Client                     pkgclient.Client `gorm:"-"`
	Approved bool
	// Recusas individuais por destinatário (cada um só perde o que recusou).
	Refusals  []BudgetRefusal `gorm:"foreignKey:BudgetID"`
	CreatedAt time.Time       `gorm:"<-:create"`
}

type RecipientType string

const (
	RecipientTypeProfessional RecipientType = "professional"
	RecipientTypeStore        RecipientType = "store"
)

// Recusa de um orçamento por um destinatário específico.
type BudgetRefusal struct {
	gorm.Model
	BudgetID      uint          `gorm:"index;uniqueIndex:idx_budget_recipient"`
	RecipientID   uint          `gorm:"uniqueIndex:idx_budget_recipient"`
	RecipientType RecipientType `gorm:"size:20;uniqueIndex:idx_budget_recipient"`
}
