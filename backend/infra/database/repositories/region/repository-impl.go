package region_repository_impl

import (
	"gorm.io/gorm"

	pkgcity "construir_mais_barato/app/domain/city"
	pkgregion "construir_mais_barato/app/domain/region"
)

type repository struct {
	DB *gorm.DB
}

func NewRegionRepositoryImpl(db *gorm.DB) pkgregion.RegionRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) FindAllWithoutPagination() ([]*pkgregion.Region, error) {
	var regions []*pkgregion.Region
	if err := r.DB.Order("name").Find(&regions).Error; err != nil {
		return nil, err
	}
	return regions, nil
}

func (r *repository) Find(quantityRegions uint) ([]*pkgregion.Region, error) {
	var regions []*pkgregion.Region
	if err := r.DB.Preload("Cities").Limit(int(quantityRegions)).Find(&regions).Error; err != nil {
		return nil, err
	}
	return regions, nil
}

func (r *repository) FindAll(limit, offset int,uf string) ([]*pkgregion.Region, int64, error) {

	var regions []*pkgregion.Region
	var where = "";
	var total int64 =0;

	if err := r.DB.Model(&pkgregion.Region{}).Where(where).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if uf != "" {
		where = "uf = '"+ uf+"'";
	}	
	if err := r.DB.Preload("Cities").Where(where).Limit(limit).Offset(offset).Order("name").Find(&regions).Error; err != nil {
		return nil, 0, err
	}
	return regions, total, nil
}

func (r *repository) FindRegionsWithCount() ([]map[string]interface{}, error) {
	var result []map[string]interface{}

	// Realiza a consulta usando uma junção entre as tabelas Regional e regional_regions
	// e depois agrupa os resultados pela profissão para contar quantos profissionais estão associados a cada uma
	if err := r.DB.Table("regional_regions").
		Select("regions.name AS region, COUNT(*) AS count").
		Joins("JOIN regions ON regional_regions.region_id = regions.id").
		Group("regions.name").
		Order("regions.name").
		Scan(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (r *repository) FindById(id uint) (*pkgregion.Region, error) {
	region := pkgregion.Region{}
	if err := r.DB.Preload("Cities").First(&region, id).Error; err != nil {
		return nil, err
	}
	return &region, nil
}

func (r *repository) FindByCity(cityid uint) (*pkgregion.Region, error) {
	region := pkgregion.Region{}
	if err := r.DB.
	Joins("JOIN regions_cities ON regions.id = regions_cities.region_id").
	Preload("Cities").
	Where("regions_cities.city_id",cityid).First(&region).Error; err != nil {
		return nil, err
	}
	return &region, nil
}


func (r *repository) Save(region pkgregion.Region) (*pkgregion.Region, error) {

	// Iniciar uma transação para garantir a atomicidade das operações
	err := r.DB.Transaction(func(tx *gorm.DB) error {

		if err := r.DB.Save(&region).Error; err != nil {
			return nil
		}
		// Limpar associações existentes (opcional, dependendo dos requisitos)
		if err := tx.Model(&region).Association("Cities").Clear(); err != nil {
			return err
		}

		// Adicionar novas associações
		var cities []pkgcity.City
		if len(region.CityIDs) > 0 {
			if err := tx.Where("id IN ?", region.CityIDs).Find(&cities).Error; err != nil {
				return err
			}
			if err := tx.Model(&region).Association("Cities").Replace(&cities); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}


	return &region, nil
}

func (r *repository) Remove(id uint) error {

	if err := r.DB.Delete(&pkgregion.Region{}, id).Error; err != nil {
		return err
	}
	return nil
}
