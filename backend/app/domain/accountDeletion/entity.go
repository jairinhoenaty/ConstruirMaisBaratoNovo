package accountdeletion

import (
	"time"

	"gorm.io/gorm"
)

const (
	StatusReceived  = "RECEIVED"
	StatusReviewing = "REVIEWING"
	StatusCompleted = "COMPLETED"
)

type AccountDeletionRequest struct {
	gorm.Model

	UserID    *uint  `json:"userId"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	Reason    string `json:"reason"`
	Status    string `json:"status"`

	ProcessedAt *time.Time `json:"processedAt"`
}
