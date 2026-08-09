package plan

import (
	"time"

	"gorm.io/gorm"
)

// UserType enum para tipos de usuário
type UserType string

const (
	UserTypeProfessional UserType = "professional"
	UserTypeStore        UserType = "store"
	UserTypeSolicitation UserType = "solicitation"
)

// Plan representa um plano de assinatura premium
type Plan struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	UserType    UserType       `gorm:"type:varchar(20);not null;index" json:"user_type"`
	Name        string         `gorm:"type:varchar(100);not null" json:"name"`
	Price       float64        `gorm:"type:decimal(10,2);not null" json:"price"`
	Description string         `gorm:"type:text" json:"description"`
	Features    string         `gorm:"type:json" json:"features"` // JSON array de features
	IsActive    bool           `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// TableName especifica o nome da tabela
func (Plan) TableName() string {
	return "plans"
}
