package budget

import (
	"time"

	"gorm.io/gorm"

	pkgcity "construir_mais_barato/app/domain/city"
	pkgclient "construir_mais_barato/app/domain/client"
	pkgprofessional "construir_mais_barato/app/domain/professional"
)

type Budget struct {
	gorm.Model
	Name                       string
	Email                      string
	Telephone                  string
	Description                string
	ProfessionalIDs            *[]uint                        `gorm:"-"`
	Professionals              []pkgprofessional.Professional `gorm:"many2many:budgets_professionals;"`
	CityID                     *uint
	City                       pkgcity.City `gorm:"foreignKey:CityID"`
	TermResponsabilityAccepted bool
	ClientID                   *uint
	Client                     pkgclient.Client `gorm:"foreignKey:ClientID"`
	Approved                   bool
	CreatedAt                  time.Time `gorm:"<-:create"`
}
