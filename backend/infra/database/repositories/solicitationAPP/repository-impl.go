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

func (r *repository) UpdateFeedback(idFirebase string, rating int, feedback string) (*pkgsolicitation.SolicitationApp, error) {
	var solicitation pkgsolicitation.SolicitationApp
	if err := r.DB.
		Where("id_firebase = ?", idFirebase).
		First(&solicitation).Error; err != nil {
		return nil, err
	}

	if err := r.DB.
		Model(&solicitation).
		Updates(map[string]interface{}{
			"rating":   rating,
			"feedback": feedback,
		}).Error; err != nil {
		return nil, err
	}

	return &solicitation, nil
}
