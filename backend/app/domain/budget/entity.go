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
	Approved  bool
	CreatedAt time.Time `gorm:"<-:create"`
}
