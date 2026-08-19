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
	// PaymentStatusExpired marca assinaturas premium cujo prazo terminou
	PaymentStatusExpired PaymentStatus = "expired"
	// PaymentStatusBypassed registra que o usuário seguiu o fluxo no app sem
	// que o pagamento tenha sido confirmado. Serve apenas para auditoria.
	PaymentStatusBypassed PaymentStatus = "bypassed"
)

// Subscription representa uma assinatura/pagamento premium
type Subscription struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UserID    uint   `gorm:"not null;index" json:"user_id"`
	PlanID    uint   `gorm:"not null;index" json:"plan_id"`
	UserType  string `gorm:"type:varchar(20);index" json:"user_type"` // professional, store, solicitation
	PaymentID int64  `gorm:"index" json:"payment_id"`                 // ID do pagamento no MercadoPago
	// StatusToken é a chave opaca usada pelo site e pelo app para consultar o
	// andamento do pagamento. Existe para que o id do MercadoPago nunca chegue
	// ao cliente: os endpoints de status e de bypass são públicos, e um id
	// adivinhável permitiria ler pagamentos alheios e sujar a auditoria.
	StatusToken       string         `gorm:"type:varchar(64);uniqueIndex" json:"-"`
	ExternalReference string         `gorm:"type:varchar(255);index" json:"external_reference"` // Referência externa
	ReferenceID       string         `gorm:"type:varchar(100);index" json:"reference_id"`       // Contexto extra (ex: id da solicitação no Firebase)
	Amount            float64        `gorm:"type:decimal(10,2);not null" json:"amount"`
	Status            PaymentStatus  `gorm:"type:varchar(20);not null;default:'pending'" json:"status"`
	QRCode            string         `gorm:"type:text" json:"qr_code"`
	QRCodeBase64      string         `gorm:"type:text" json:"qr_code_base64"`
	IdempotencyKey    string         `gorm:"type:varchar(255);unique" json:"idempotency_key"`
	PaymentMethod     string         `gorm:"type:varchar(50)" json:"payment_method"` // pix, credit_card, etc
	PaidAt            *time.Time     `json:"paid_at,omitempty"`
	ExpiresAt         *time.Time     `json:"expires_at,omitempty"` // Fim da vigência do premium
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// TableName especifica o nome da tabela
func (Subscription) TableName() string {
	return "subscriptions"
}

func (s *Subscription) IsApproved() bool {
	return s.Status == PaymentStatusApproved
}

func (s *Subscription) IsPending() bool {
	return s.Status == PaymentStatusPending
}

// MarkAsApproved registra o pagamento confirmado. duration zero significa que a
// assinatura não expira (é o caso da taxa por solicitação, que é avulsa).
func (s *Subscription) MarkAsApproved(paidAt time.Time, duration time.Duration) {
	s.Status = PaymentStatusApproved
	s.PaidAt = &paidAt
	if duration > 0 {
		expiresAt := paidAt.Add(duration)
		s.ExpiresAt = &expiresAt
	}
}

func (s *Subscription) MarkAs(status PaymentStatus) {
	s.Status = status
}
