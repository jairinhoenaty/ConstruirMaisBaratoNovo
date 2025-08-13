package region

import (
	pkgcity "construir_mais_barato/app/domain/city"

	"gorm.io/gorm"
)

type Region struct {
	gorm.Model
	Name        string
	Description string
	Icon        string
	UF  	    string 			`gorm:"index"`
	Cities      []pkgcity.City	`gorm:"many2many:regions_cities;"`
	CityIDs     []uint			`gorm:"-"`

}
