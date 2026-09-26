package professionCategory_repository_impl

import (
	"gorm.io/gorm"

	pkgprofessionCategory "construir_mais_barato/app/domain/professionCategory"
)

type repository struct {
	DB *gorm.DB
}

func NewProfessionCategoryRepositoryImpl(db *gorm.DB) pkgprofessionCategory.ProfessionCategoryRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) FindAll(limit, offset int) ([]*pkgprofessionCategory.ProfessionCategory, int64, error) {
	var total int64
	if err := r.DB.Model(&pkgprofessionCategory.ProfessionCategory{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var categories []*pkgprofessionCategory.ProfessionCategory
	if err := r.DB.Limit(limit).Offset(offset).
		Order("position").Order("name").
		Find(&categories).Error; err != nil {
		return nil, 0, err
	}
	return categories, total, nil
}

func (r *repository) FindActive() ([]*pkgprofessionCategory.ProfessionCategory, error) {
	var categories []*pkgprofessionCategory.ProfessionCategory
	if err := r.DB.Where("active = ?", true).
		Order("position").Order("name").
		Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *repository) FindById(id uint) (*pkgprofessionCategory.ProfessionCategory, error) {
	category := pkgprofessionCategory.ProfessionCategory{}
	if err := r.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *repository) CountProfessions(categoryIDs []uint) (map[uint]int64, error) {
	counts := make(map[uint]int64, len(categoryIDs))
	if len(categoryIDs) == 0 {
		return counts, nil
	}

	var rows []struct {
		CategoryID uint
		Total      int64
	}
	if err := r.DB.Table("professions").
		Select("category_id, COUNT(*) as total").
		Where("category_id IN ? AND deleted_at IS NULL", categoryIDs).
		Group("category_id").
		Scan(&rows).Error; err != nil {
		return nil, err
	}

	for _, row := range rows {
		counts[row.CategoryID] = row.Total
	}
	return counts, nil
}

func (r *repository) Save(category pkgprofessionCategory.ProfessionCategory) (*pkgprofessionCategory.ProfessionCategory, error) {
	if category.ID == 0 {
		if err := r.DB.Create(&category).Error; err != nil {
			return nil, err
		}
		return &category, nil
	}

	// Select explícito: Updates com struct ignora campos zerados, e sem ele o
	// administrador nunca conseguiria desligar active ou client_travels.
	if err := r.DB.Model(&pkgprofessionCategory.ProfessionCategory{}).
		Where("id = ?", category.ID).
		Select("name", "description", "icon", "position", "active", "client_travels").
		Updates(category).Error; err != nil {
		return nil, err
	}
	return r.FindById(category.ID)
}

func (r *repository) Remove(id uint) error {
	return r.DB.Delete(&pkgprofessionCategory.ProfessionCategory{}, id).Error
}
