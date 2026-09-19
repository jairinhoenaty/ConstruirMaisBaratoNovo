package accountdeletion_repository_impl

import (
	pkgaccountdeletion "construir_mais_barato/app/domain/accountDeletion"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewAccountDeletionRepositoryImpl(
	db *gorm.DB,
) pkgaccountdeletion.AccountDeletionRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) Save(
	request pkgaccountdeletion.AccountDeletionRequest,
) (*pkgaccountdeletion.AccountDeletionRequest, error) {

	if err := r.DB.Create(&request).Error; err != nil {
		return nil, err
	}

	return &request, nil
}

func (r *repository) FindAll() (
	[]*pkgaccountdeletion.AccountDeletionRequest,
	error,
) {
	var requests []*pkgaccountdeletion.AccountDeletionRequest

	err := r.DB.
		Order("created_at DESC").
		Find(&requests).
		Error

	if err != nil {
		return nil, err
	}

	return requests, nil
}

func (r *repository) FindByID(
	id uint,
) (*pkgaccountdeletion.AccountDeletionRequest, error) {

	var request pkgaccountdeletion.AccountDeletionRequest

	if err := r.DB.First(&request, id).Error; err != nil {
		return nil, err
	}

	return &request, nil
}

func (r *repository) UpdateStatus(
	id uint,
	status string,
) error {

	return r.DB.
		Model(&pkgaccountdeletion.AccountDeletionRequest{}).
		Where("id = ?", id).
		Update("status", status).
		Error
}
