package contact

import (
	pkgcity "construir_mais_barato/app/domain/city"
	pkgclient "construir_mais_barato/app/domain/client"
	pkgproduct "construir_mais_barato/app/domain/product"
	pkgprofessional "construir_mais_barato/app/domain/professional"
	pkgstore "construir_mais_barato/app/domain/store"
	"time"

	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	Name           string
	Telephone      string
	Email          string
	Message        string
	Status         string
	CityID         *uint
	City           pkgcity.City `gorm:"foreignkey:CityID"`
	ProfessionalID *uint
	Professional   pkgprofessional.Professional `gorm:"foreignkey:ProfessionalID"`
	ClientID       *uint
	Client         pkgclient.Client `gorm:"foreignkey:ClientID"`
	StoreID        *uint
	Store          pkgstore.Store `gorm:"foreignkey:StoreID"`
	ProductID      *uint
	Product        pkgproduct.Product `gorm:"foreignkey:ProductID"`
	CreatedAt      time.Time          `gorm:"<-:create"`
	Approved       bool
}
