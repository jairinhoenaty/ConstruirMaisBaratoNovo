package chat

import (
	pkgclient "construir_mais_barato/app/domain/client"
	pkgprofessional "construir_mais_barato/app/domain/professional"
	"time"

	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	Message        string
	ProfessionalID *uint
	Professional   pkgprofessional.Professional `gorm:"foreignkey:ProfessionalID"`
	ClientID       *uint
	Client         pkgclient.Client `gorm:"foreignkey:ClientID"`
	Origem         string
	CreatedAt      time.Time `gorm:"<-:create"`
}
