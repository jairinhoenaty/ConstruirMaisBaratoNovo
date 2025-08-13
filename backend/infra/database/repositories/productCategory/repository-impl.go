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

/*
func (r *repository) Remove(id uint) error {

	if err := r.DB.Delete(&pkgproduct.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}
*/
