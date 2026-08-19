package subscription

import "time"

type SubscriptionRepository interface {
	FindByID(id uint) (*Subscription, error)
	FindByPaymentID(paymentID int64) (*Subscription, error)
	FindByStatusToken(statusToken string) (*Subscription, error)
	FindByExternalReference(externalReference string) (*Subscription, error)
	// FindPendingByUserAndPlan devolve o pagamento pendente mais recente do
	// usuário para o plano, usado para reaproveitar o QR Code ainda válido.
	FindPendingByUserAndPlan(userID, planID uint) (*Subscription, error)
	// FindExpired lista assinaturas aprovadas cuja vigência terminou.
	FindExpired(now time.Time) ([]*Subscription, error)
	Save(subscription Subscription) (*Subscription, error)
	Update(subscription *Subscription) error
}
