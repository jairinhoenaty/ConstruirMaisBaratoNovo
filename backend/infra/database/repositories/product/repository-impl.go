package product_repository_impl

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"

	pkgproduct "construir_mais_barato/app/domain/product"
)

type repository struct {
	DB *gorm.DB
}

func NewProductRepositoryImpl(db *gorm.DB) pkgproduct.ProductRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) FindAll(limit, offset int, professionalID int, storeID int, approved string, dayoffer string) ([]*pkgproduct.Product, int64, error) {
	var total int64
	fmt.Println("PROCURANDO PRODUTOS")

	var where string = "0=0"
	if professionalID != 0 {
		where = where + " and professional_id=" + strconv.Itoa(professionalID)
	}
	if storeID != 0 {
		where = where + " and store_id=" + strconv.Itoa(storeID)
	}
	if storeID != 0 {
		where = where + " and store_id=" + strconv.Itoa(storeID)
	}
	if dayoffer != "" {
		var dayofferStr = "false"
		if dayoffer == "S" {
			dayofferStr = "true"
		}
		where = where + " and dayoffer=" + dayofferStr
	}

	if approved != "" {
		var approvedStr = "false"
		if approved == "S" {
			approvedStr = "true"
		}
		where = where + " and approved=" + approvedStr
	}

	// Contagem total de orçamentos
	if err := r.DB.Model(&pkgproduct.Product{}).Where(where).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if limit == 0 {
		limit = 10
		offset = 0
	}

	var products []*pkgproduct.Product

	if err := r.DB.
		Preload("Category").
		Preload("Professional").
		Preload("Store").
		Limit(limit).
		Offset(offset).
		Where(where).
		Order("dayoffer DESC,created_at DESC,id DESC").
		Find(&products).Error; err != nil {
		return nil, 0, err
	}
	return products, total, nil

}

func (r *repository) FindById(id uint) (*pkgproduct.Product, error) {
	product := pkgproduct.Product{}
	if err := r.DB.Preload("Category").
		First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

/*func (r *repository) FindByProfessional(uf string) ([]*pkgproduct.Product, error) {
	cities := make([]*pkgproduct.Product, 0)
	if err := r.DB.Where("uf = ?", uf).Find(&cities).Error; err != nil {
		return nil, err
	}
	return cities, nil
}*/

func (r *repository) Save(product pkgproduct.Product) (*pkgproduct.Product, error) {
	if err := r.DB.Save(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *repository) Remove(id uint) error {

	if err := r.DB.Delete(&pkgproduct.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) FindApproved() ([]*pkgproduct.Product, error) {
	var products []*pkgproduct.Product
	if err := r.DB.
		Preload("productCategory.ProductCategory").
		Where("Approved = true").
		Where("deleted_at IS NULL").
		Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *repository) FindDayoffer() ([]*pkgproduct.Product, error) {
	var products []*pkgproduct.Product
	if err := r.DB.
		Preload("Category").
		Preload("Professional").
		Where("dayoffer = true").
		Where("approved = true").
		Where("deleted_at IS NULL").
		Limit(5).
		Order("created_at desc").
		Find(&products).
		Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *repository) FindByCity(CityID uint) ([]*pkgproduct.Product, error) {
	var products []*pkgproduct.Product
	if err := r.DB.
		Joins("LEFT OUTER JOIN professionals ON products.professional_id = professionals.id").
		Joins("LEFT OUTER JOIN stores ON products.store_id = stores.id").
		Where("(professionals.city_id = "+strconv.FormatUint(uint64(CityID), 10)+" or stores.city_id = "+strconv.FormatUint(uint64(CityID), 10)+")").
		Where("products.deleted_at IS NULL").
		Where("products.approved", true).
		Order("products.dayoffer DESC,products.created_at desc").
		Find(&products).
		Error; err != nil {
		return nil, err
	}
	return products, nil
}
