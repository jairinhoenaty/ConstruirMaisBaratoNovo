package subscription

import "time"

type SubscriptionService interface {
	FindByID(id uint) (*Subscription, error)
	FindByPaymentID(paymentID int64) (*Subscription, error)
	FindByStatusToken(statusToken string) (*Subscription, error)
	FindByExternalReference(externalReference string) (*Subscription, error)
	FindPendingByUserAndPlan(userID, planID uint) (*Subscription, error)
	FindExpired(now time.Time) ([]*Subscription, error)
	Save(subscription Subscription) (*Subscription, error)
	Update(subscription *Subscription) error
}

type subscriptionService struct {
	repository SubscriptionRepository
}

func NewSubscriptionService(repository SubscriptionRepository) SubscriptionService {
	return &subscriptionService{
		repository: repository,
	}
}

func (s *subscriptionService) FindByID(id uint) (*Subscription, error) {
	return s.repository.FindByID(id)
}

func (s *subscriptionService) FindByPaymentID(paymentID int64) (*Subscription, error) {
	return s.repository.FindByPaymentID(paymentID)
}

func (s *subscriptionService) FindByStatusToken(statusToken string) (*Subscription, error) {
	return s.repository.FindByStatusToken(statusToken)
}

func (s *subscriptionService) FindByExternalReference(externalReference string) (*Subscription, error) {
	return s.repository.FindByExternalReference(externalReference)
}

func (s *subscriptionService) FindPendingByUserAndPlan(userID, planID uint) (*Subscription, error) {
	return s.repository.FindPendingByUserAndPlan(userID, planID)
}

func (s *subscriptionService) FindExpired(now time.Time) ([]*Subscription, error) {
	return s.repository.FindExpired(now)
}

func (s *subscriptionService) Save(subscription Subscription) (*Subscription, error) {
	return s.repository.Save(subscription)
}

func (s *subscriptionService) Update(subscription *Subscription) error {
	return s.repository.Update(subscription)
}
