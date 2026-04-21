package exchangecode

import (
	"errors"
	"time"

	"gorm.io/gorm"

	pkgexchangecode "construir_mais_barato/app/domain/exchange-codes"
)

type repository struct {
	DB *gorm.DB
}

func NewExchangeCodeRepositoryImpl(db *gorm.DB) pkgexchangecode.ExchangeCodeRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) Create(ec *pkgexchangecode.ExchangeCode) error {
	return r.DB.Create(ec).Error
}

func (r *repository) Redeem(code string) (*pkgexchangecode.ExchangeCode, error) {
	now := time.Now()

	result := r.DB.Model(&pkgexchangecode.ExchangeCode{}).
		Where("code = ? AND used_at IS NULL AND expires_at > ?", code, now).
		Update("used_at", now)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("invalid, expired or already used code")
	}

	var ec pkgexchangecode.ExchangeCode
	if err := r.DB.Preload("User").First(&ec, "code = ?", code).Error; err != nil {
		return nil, err
	}
	// Estudar alternativa para usar apenas colunas utilizáveis no futuro:
	// r.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
	// 	return db.Select("id, name, profile, email")
	// }).First(&ec, "code = ?", code)

	return &ec, nil
}
