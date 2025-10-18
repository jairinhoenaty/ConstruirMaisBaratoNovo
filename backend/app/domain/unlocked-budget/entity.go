package unlockedbudget

import "time"

// UnlockedBudget representa um orçamento desbloqueado por um profissional não-premium
type UnlockedBudget struct {
	ID             uint       `json:"id" gorm:"primaryKey"`
	ProfessionalID uint       `json:"professionalId" gorm:"not null;index:idx_prof_budget,unique"`
	BudgetID       uint       `json:"budgetId" gorm:"not null;index:idx_prof_budget,unique"`
	Status         string     `json:"status" gorm:"type:varchar(20);not null;default:'pending'"` // pending, paid, failed, cancelled
	PaymentID      string     `json:"paymentId" gorm:"type:varchar(100);index"`                  // ID do pagamento no gateway (Mercado Pago)
	Amount         float64    `json:"amount" gorm:"type:decimal(10,2);not null"`                 // Valor pago (ex: 30.00)
	QRCode         string     `json:"qrCode,omitempty" gorm:"type:text"`                         // QR Code PIX para pagamento
	QRCodeBase64   string     `json:"qrCodeBase64,omitempty" gorm:"type:text"`                   // QR Code em base64
	PaidAt         *time.Time `json:"paidAt"`                                                    // Data/hora do pagamento aprovado
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      time.Time  `json:"updatedAt"`
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
