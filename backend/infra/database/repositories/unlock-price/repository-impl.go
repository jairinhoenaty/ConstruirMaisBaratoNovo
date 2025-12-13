package unlock_price_repository_impl

import (
	unlockprice "construir_mais_barato/app/domain/unlock-price"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewUnlockPriceRepositoryImpl(db *gorm.DB) unlockprice.UnlockPriceRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) FindByUserType(userType unlockprice.UserType) (*unlockprice.UnlockPrice, error) {
	var price unlockprice.UnlockPrice
	err := r.DB.Where("user_type = ?", userType).First(&price).Error
	if err != nil {
		return nil, err
	}
	return &price, nil
}

func (r *repository) FindActiveByUserType(userType unlockprice.UserType) (*unlockprice.UnlockPrice, error) {
	var price unlockprice.UnlockPrice
	err := r.DB.Where("user_type = ? AND is_active = ?", userType, true).First(&price).Error
	if err != nil {
		return nil, err
	}
	return &price, nil
}

func (r *repository) FindAll() ([]*unlockprice.UnlockPrice, error) {
	var prices []*unlockprice.UnlockPrice
	err := r.DB.Order("user_type asc").Find(&prices).Error
	if err != nil {
		return nil, err
	}
	return prices, nil
}

func (r *repository) Save(unlockPrice unlockprice.UnlockPrice) (*unlockprice.UnlockPrice, error) {
	err := r.DB.Save(&unlockPrice).Error
	if err != nil {
		return nil, err
	}
	return &unlockPrice, nil
}

func (r *repository) Update(unlockPrice *unlockprice.UnlockPrice) error {
	return r.DB.Save(unlockPrice).Error
}
