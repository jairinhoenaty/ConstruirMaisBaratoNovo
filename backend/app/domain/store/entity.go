package store

import (
	"time"

	"gorm.io/gorm"

	pkgcity "construir_mais_barato/app/domain/city"
)

type Store struct {
	gorm.Model
	Name         string
	Company      string
	Email        string `gorm:"unique"`
	Telephone    string
	LgpdAceito   string
	CityID       uint
	City         pkgcity.City `gorm:"foreignKey:CityID"`
	Cep          string
	Street       string
	Neighborhood string
	Image        []byte    `gorm:"type:longblob"`
	CreatedAt    time.Time `gorm:"<-:create"`
}

type StoreCount struct {
	StoreName string
	Quantity  int
}

type CityStoreCount struct {
	CityID     uint
	CityName   string
	StoreCount int64
}
