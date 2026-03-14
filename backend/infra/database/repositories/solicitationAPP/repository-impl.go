package solicitationapp_repository_impl

import (
	"gorm.io/gorm"

	pkgsolicitation "construir_mais_barato/app/domain/solicitationAPP"
)

type repository struct {
	DB *gorm.DB
}

func NewSolicitationAppRepositoryImpl(db *gorm.DB) pkgsolicitation.SolicitationAppRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) Save(solicitation pkgsolicitation.SolicitationApp) (*pkgsolicitation.SolicitationApp, error) {
	if err := r.DB.Create(&solicitation).Error; err != nil {
		return nil, err
	}
	return &solicitation, nil
}
