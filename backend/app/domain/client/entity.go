package client

import (
	"time"

	"gorm.io/gorm"

	pkgcity "construir_mais_barato/app/domain/city"
)

type Client struct {
	gorm.Model
	Name         string
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

type ClientCount struct {
	Quantity int
}

type CityClientCount struct {
	CityID      uint
	CityName    string
	ClientCount int64
}
