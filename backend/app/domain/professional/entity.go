package professional

import (
	"time"

	"gorm.io/gorm"

	pkgcity "construir_mais_barato/app/domain/city"
	pkgprofession "construir_mais_barato/app/domain/profession"
)

type Professional struct {
	gorm.Model
	Name          string
	Email         string //`gorm:"unique"`
	Company       string
	Telephone     string
	LgpdAceito    string
	CityID        uint
	City          pkgcity.City               `gorm:"foreignKey:CityID"`
	Professions   []pkgprofession.Profession `gorm:"many2many:professional_professions;"`
	ProfessionIDs []uint                     `gorm:"-"`
	Cep           string
	Street        string
	Neighborhood  string
	Image         []byte  `gorm:"type:longblob"`
	Latitude      float64 `gorm:"type:decimal(10,8)"`
	Longitude     float64 `gorm:"type:decimal(11,8)"`
	Verified      *bool   `gorm:"default:false"`
	OnLine        *bool
	CreatedAt     time.Time `gorm:"<-:create"`
	Distance      float64   `gorm:"->"` //`gorm:"-"`
}

type ProfessionCount struct {
	ProfessionName string
	Quantity       int
}

type CityProfessionalCount struct {
	CityID            uint
	CityName          string
	ProfessionalCount int64
}

type UFProfessionalCount struct {
	UFName            string
	ProfessionalCount int64
}
