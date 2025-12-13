package unlockprice

import (
	"time"

	"gorm.io/gorm"
)

type UserType string

const (
	UserTypeProfessional UserType = "professional"
	UserTypeStore        UserType = "store"
)

type UnlockPrice struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	UserType    UserType       `gorm:"type:varchar(20);not null;uniqueIndex:idx_user_type_active" json:"user_type"`
	Price       float64        `gorm:"type:decimal(10,2);not null" json:"price"`
	Description string         `gorm:"type:varchar(255)" json:"description"`
	IsActive    bool           `gorm:"default:true;uniqueIndex:idx_user_type_active" json:"is_active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (UnlockPrice) TableName() string {
	return "unlock_prices"
}
