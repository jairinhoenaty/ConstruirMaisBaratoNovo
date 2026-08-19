package store

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"

	pkgcity "construir_mais_barato/app/domain/city"
	pkgcategory "construir_mais_barato/app/domain/productCategory"
)

type Store struct {
	gorm.Model
	Name           string
	Company        string
	Email          string `gorm:"unique"`
	Telephone      string
	LgpdAceito     string
	CityID         uint
	City           pkgcity.City `gorm:"foreignKey:CityID"`
	Cep            string
	Street         string
	Neighborhood   string
	Image          []byte    `gorm:"type:longblob"`
	CreatedAt      time.Time `gorm:"<-:create"`
	IsPremiumStore *bool     `gorm:"default:false"`
	// PremiumExpiresAt é o fim da vigência do premium pago. Nulo significa
	// premium sem prazo (ativado manualmente por um administrador), que a
	// rotina de expiração não deve rebaixar.
	PremiumExpiresAt  *time.Time
	CategoryProductID uint
	Category          *pkgcategory.ProductCategory `gorm:"foreignKey:CategoryProductID"`
	SubCategories     UintSlice                    `gorm:"type:json"`
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

type UintSlice []uint

func (u UintSlice) Value() (driver.Value, error) {
	// Converte para JSON antes de salvar
	b, err := json.Marshal(u)
	return string(b), err
}

func (u *UintSlice) Scan(value interface{}) error {
	// Converte de JSON para slice quando lê do BD
	b, ok := value.([]byte)
	if !ok {
		return errors.New("invalid Scan source for UintSlice")
	}
	return json.Unmarshal(b, u)
}
