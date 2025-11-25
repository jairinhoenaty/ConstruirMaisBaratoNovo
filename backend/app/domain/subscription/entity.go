package subscription

import (
	"time"

	"gorm.io/gorm"
)

// PaymentStatus enum para status de pagamento
type PaymentStatus string

const (
	PaymentStatusPending  PaymentStatus = "pending"
	PaymentStatusApproved PaymentStatus = "approved"
	PaymentStatusFailed   PaymentStatus = "failed"
	PaymentStatusCanceled PaymentStatus = "canceled"
	PaymentStatusRefunded PaymentStatus = "refunded"
)

// Subscription representa uma assinatura/pagamento premium
type Subscription struct {
	ID                 uint           `gorm:"primaryKey" json:"id"`
	UserID             uint           `gorm:"not null;index" json:"user_id"`
	PlanID             uint           `gorm:"not null;index" json:"plan_id"`
	PaymentID          int64          `gorm:"index" json:"payment_id"`           // ID do pagamento no MercadoPago
	ExternalReference  string         `gorm:"type:varchar(255);index" json:"external_reference"` // Referência externa
	Amount             float64        `gorm:"type:decimal(10,2);not null" json:"amount"`
	Status             PaymentStatus  `gorm:"type:varchar(20);not null;default:'pending'" json:"status"`
	QRCode             string         `gorm:"type:text" json:"qr_code"`
	QRCodeBase64       string         `gorm:"type:text" json:"qr_code_base64"`
	IdempotencyKey     string         `gorm:"type:varchar(255);unique" json:"idempotency_key"`
	PaymentMethod      string         `gorm:"type:varchar(50)" json:"payment_method"` // pix, credit_card, etc
	PaidAt             *time.Time     `json:"paid_at,omitempty"`
	ExpiresAt          *time.Time     `json:"expires_at,omitempty"` // Para assinaturas futuras
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// TableName especifica o nome da tabela
func (Subscription) TableName() string {
	return "subscriptions"
}
