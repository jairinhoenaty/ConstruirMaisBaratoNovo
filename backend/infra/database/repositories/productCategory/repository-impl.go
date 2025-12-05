package productCategory_repository_impl

import (
	"gorm.io/gorm"

	pkgproductCategory "construir_mais_barato/app/domain/productCategory"
)

type repository struct {
	DB *gorm.DB
}

func NewProductCategoryRepositoryImpl(db *gorm.DB) pkgproductCategory.ProductCategoryRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) FindByProfession(professionID int) ([]*pkgproductCategory.ProductCategory, error) {
	productCategories := make([]*pkgproductCategory.ProductCategory, 0)
	if err := r.DB.Where("profession_id = ?", professionID).
		Preload("Profession").
		Order("name asc").
		Find(&productCategories).Error; err != nil {
		return nil, err
	}
	return productCategories, nil
}

func (r *repository) Save(productCategory pkgproductCategory.ProductCategory) (*pkgproductCategory.ProductCategory, error) {
	if err := r.DB.Save(&productCategory).Error; err != nil {
		return nil, err
	}
	return &productCategory, nil
}

func (r *repository) FindAll() ([]*pkgproductCategory.ProductCategory, error) {
	categories := make([]*pkgproductCategory.ProductCategory, 0)
	if err := r.DB.Preload("Parent").
		Preload("Children").
		Order("name asc").
		Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *repository) FindById(id uint) (*pkgproductCategory.ProductCategory, error) {
	category := &pkgproductCategory.ProductCategory{}
	if err := r.DB.Preload("Parent").
		Preload("Children").
		First(category, id).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *repository) FindTopLevel() ([]*pkgproductCategory.ProductCategory, error) {
	categories := make([]*pkgproductCategory.ProductCategory, 0)
	if err := r.DB.Where("parent_id IS NULL").
		Preload("Children").
		Order("name asc").
		Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *repository) FindSubcategoriesByParentID(parentID uint) ([]*pkgproductCategory.ProductCategory, error) {
	subcategories := make([]*pkgproductCategory.ProductCategory, 0)
	if err := r.DB.Where("parent_id = ?", parentID).
		Order("name asc").
		Find(&subcategories).Error; err != nil {
		return nil, err
	}
	return subcategories, nil
}

func (r *repository) Remove(id uint) error {
	if err := r.DB.Delete(&pkgproductCategory.ProductCategory{}, id).Error; err != nil {
		return err
	}
	return nil
}
