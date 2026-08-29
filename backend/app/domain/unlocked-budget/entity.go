package unlockedbudget

import "time"

type UserType string

const (
	UserTypeProfessional UserType = "professional"
	UserTypeStore        UserType = "store"
)

// UnlockedBudget representa um orçamento desbloqueado por um profissional ou loja não-premium
type UnlockedBudget struct {
	ID             uint     `json:"id" gorm:"primaryKey"`
	UserType       UserType `json:"userType" gorm:"type:varchar(20);not null;index:idx_user_budget"`
	ProfessionalID *uint    `json:"professionalId" gorm:"index:idx_user_budget"`
	StoreID        *uint    `json:"storeId" gorm:"index:idx_user_budget"`
	BudgetID       uint     `json:"budgetId" gorm:"not null;index:idx_user_budget"`
	Status         string   `json:"status" gorm:"type:varchar(20);not null;default:'pending'"` // pending, paid, failed, cancelled
	PaymentID      string   `json:"paymentId" gorm:"type:varchar(100);index"`                  // ID do pagamento no gateway (Mercado Pago)
	// StatusToken é a chave opaca usada pelo cliente para consultar o
	// andamento do pagamento, para que o id do MercadoPago não seja exposto.
	StatusToken  string     `json:"-" gorm:"type:varchar(64);uniqueIndex"`
	Amount       float64    `json:"amount" gorm:"type:decimal(10,2);not null"` // Valor pago (ex: 30.00)
	QRCode       string     `json:"qrCode,omitempty" gorm:"type:text"`         // QR Code PIX para pagamento
	QRCodeBase64 string     `json:"qrCodeBase64,omitempty" gorm:"type:text"`   // QR Code em base64
	PaidAt       *time.Time `json:"paidAt"`                                    // Data/hora do pagamento aprovado
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
}

func (UnlockedBudget) TableName() string {
	return "unlocked_budgets"
}

func (ub *UnlockedBudget) IsPaid() bool {
	return ub.Status == "paid"
}

func (ub *UnlockedBudget) IsPending() bool {
	return ub.Status == "pending"
}

func (ub *UnlockedBudget) MarkAsPaid() {
	ub.Status = "paid"
	now := time.Now()
	ub.PaidAt = &now
}

func (ub *UnlockedBudget) MarkAsFailed() {
	ub.Status = "failed"
}
