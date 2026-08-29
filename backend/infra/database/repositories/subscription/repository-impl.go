package subscription_repository_impl

import (
	"time"

	pkgsubscription "construir_mais_barato/app/domain/subscription"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewSubscriptionRepositoryImpl(db *gorm.DB) pkgsubscription.SubscriptionRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) FindByID(id uint) (*pkgsubscription.Subscription, error) {
	var subscription pkgsubscription.Subscription
	if err := r.DB.Where("id = ?", id).First(&subscription).Error; err != nil {
		return nil, err
	}
	return &subscription, nil
}

func (r *repository) FindByPaymentID(paymentID int64) (*pkgsubscription.Subscription, error) {
	var subscription pkgsubscription.Subscription
	if err := r.DB.Where("payment_id = ?", paymentID).First(&subscription).Error; err != nil {
		return nil, err
	}
	return &subscription, nil
}

func (r *repository) FindByStatusToken(statusToken string) (*pkgsubscription.Subscription, error) {
	var subscription pkgsubscription.Subscription
	if err := r.DB.Where("status_token = ?", statusToken).First(&subscription).Error; err != nil {
		return nil, err
	}
	return &subscription, nil
}

func (r *repository) FindByExternalReference(externalReference string) (*pkgsubscription.Subscription, error) {
	var subscription pkgsubscription.Subscription
	if err := r.DB.Where("external_reference = ?", externalReference).First(&subscription).Error; err != nil {
		return nil, err
	}
	return &subscription, nil
}

func (r *repository) FindPendingByUserAndPlan(userID, planID uint) (*pkgsubscription.Subscription, error) {
	var subscription pkgsubscription.Subscription
	err := r.DB.
		Where("user_id = ? AND plan_id = ? AND status = ?", userID, planID, pkgsubscription.PaymentStatusPending).
		Order("created_at DESC").
		First(&subscription).Error
	if err != nil {
		return nil, err
	}
	return &subscription, nil
}

func (r *repository) FindExpired(now time.Time) ([]*pkgsubscription.Subscription, error) {
	var subscriptions []*pkgsubscription.Subscription
	err := r.DB.
		Where("status = ? AND expires_at IS NOT NULL AND expires_at < ?", pkgsubscription.PaymentStatusApproved, now).
		Find(&subscriptions).Error
	if err != nil {
		return nil, err
	}
	return subscriptions, nil
}

func (r *repository) Save(subscription pkgsubscription.Subscription) (*pkgsubscription.Subscription, error) {
	if err := r.DB.Create(&subscription).Error; err != nil {
		return nil, err
	}
	return &subscription, nil
}

func (r *repository) Update(subscription *pkgsubscription.Subscription) error {
	return r.DB.Save(subscription).Error
}
