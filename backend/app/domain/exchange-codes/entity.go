package exchangecodes

import (
	"time"

	pkguser "construir_mais_barato/app/domain/user"
)

type ExchangeCode struct {
	Code      string    `gorm:"primaryKey;size:128"`
	UserID    uint      `gorm:"not null;index"`
	ExpiresAt time.Time `gorm:"not null;index"`
	UsedAt    *time.Time
	CreatedAt time.Time `gorm:"<-:create"`

	User pkguser.User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
