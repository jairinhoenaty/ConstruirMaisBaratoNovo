package city_repository_impl

import (
	"gorm.io/gorm"

	pkgcity "construir_mais_barato/app/domain/city"
)

type repository struct {
	DB *gorm.DB
}

func NewCityRepositoryImpl(db *gorm.DB) pkgcity.CityRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) FindAll() ([]*pkgcity.City, error) {
	var citys []*pkgcity.City
	if err := r.DB.Find(&citys).Error; err != nil {
		return nil, err
	}
	return citys, nil
}

func (r *repository) FindById(id uint) (*pkgcity.City, error) {
	city := pkgcity.City{}
	if err := r.DB.First(&city, id).Error; err != nil {
		return nil, err
	}
	return &city, nil
}

func (r *repository) FindByUF(uf string) ([]*pkgcity.City, error) {
	cities := make([]*pkgcity.City, 0)
	if err := r.DB.Where("uf = ?", uf).Find(&cities).Error; err != nil {
		return nil, err
	}
	return cities, nil
}

func (r *repository) Save(city pkgcity.City) (*pkgcity.City, error) {
	if err := r.DB.Save(&city).Error; err != nil {
		return nil, err
	}
	return &city, nil
}

func (r *repository) Remove(id uint) error {

	if err := r.DB.Delete(&pkgcity.City{}, id).Error; err != nil {
		return err
	}
	return nil
}
