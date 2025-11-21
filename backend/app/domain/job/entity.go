package job

import (
	pkgcity "construir_mais_barato/app/domain/city"
	pkgprofession "construir_mais_barato/app/domain/profession"
	pkgprofessional "construir_mais_barato/app/domain/professional"
	"time"

	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	Title             string
	HiringType        string                           `gorm:"type:varchar(50);index"`
	Salary            *float64                         `gorm:"type:decimal(10,2)"`
	SalaryType        string                           `gorm:"type:varchar(50)"` // mensal, por hora, etc
	Location          string
	Description       string                           `gorm:"type:longtext"`
	Schedule          string
	Requirements      string                           `gorm:"type:longtext"`
	Benefits          string                           `gorm:"type:longtext"`
	ContactEmail      string
	ContactPhone      string
	OpeningsQuantity  int                              `gorm:"default:1"`
	Status            string                           `gorm:"type:varchar(20);default:'active';index"`
	CompanyID         uint
	Company           pkgprofessional.Professional     `gorm:"foreignKey:CompanyID"`
	ProfessionID      *uint
	Profession        *pkgprofession.Profession        `gorm:"foreignKey:ProfessionID"`
	CityID            uint
	City              pkgcity.City                     `gorm:"foreignKey:CityID"`
	PublishedAt       *time.Time
	ExpiresAt         *time.Time
	CreatedAt         time.Time `gorm:"<-:create"`
	UpdatedAt         time.Time
}

// IsActive verifica se a vaga ainda está ativa
func (j *Job) IsActive() bool {
	if j.Status != "active" {
		return false
	}
	if j.ExpiresAt != nil && j.ExpiresAt.Before(time.Now()) {
		return false
	}
	return true
}

// IsExpired verifica se a vaga expirou
func (j *Job) IsExpired() bool {
	return !j.IsActive()
}
