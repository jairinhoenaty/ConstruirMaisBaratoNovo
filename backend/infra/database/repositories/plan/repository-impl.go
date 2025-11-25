package plan_repository_impl

import (
	pkgplan "construir_mais_barato/app/domain/plan"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewPlanRepositoryImpl(db *gorm.DB) pkgplan.PlanRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) FindAll() ([]*pkgplan.Plan, error) {
	var plans []*pkgplan.Plan
	if err := r.DB.Order("user_type ASC, price ASC").Find(&plans).Error; err != nil {
		return nil, err
	}
	return plans, nil
}

func (r *repository) FindByID(id uint) (*pkgplan.Plan, error) {
	var plan pkgplan.Plan
	if err := r.DB.First(&plan, id).Error; err != nil {
		return nil, err
	}
	return &plan, nil
}

func (r *repository) FindByUserType(userType pkgplan.UserType) (*pkgplan.Plan, error) {
	var plan pkgplan.Plan
	if err := r.DB.Where("user_type = ? AND is_active = ?", userType, true).First(&plan).Error; err != nil {
		return nil, err
	}
	return &plan, nil
}

func (r *repository) FindAllActive() ([]*pkgplan.Plan, error) {
	var plans []*pkgplan.Plan
	if err := r.DB.Where("is_active = ?", true).Order("user_type ASC, price ASC").Find(&plans).Error; err != nil {
		return nil, err
	}
	return plans, nil
}
