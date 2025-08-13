package banner_repository_impl

import (
	pkgbanner "construir_mais_barato/app/domain/banner"
	pkgprofession "construir_mais_barato/app/domain/profession"
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}


func NewBannerRepositoryImpl(db *gorm.DB) pkgbanner.BannerRepository {
	return &repository{
		DB: db,
	}
}

// FindByPage implements banner.BannerRepository.
func (r *repository) FindByPage(page string,cityId uint, regionId uint) ([]*pkgbanner.Banner, error) {

	var where string = "banners.page = '"+ page+"'"

	if (cityId!=0) {
		where = where+ " and banners.city_id="+strconv.Itoa(int(cityId))
	}

	if (regionId!=0 && page!="B") {
		where = "("+where+ ") or (banners.region_id="+strconv.Itoa(int(regionId))+" and banners.page='U')"
	}
	if (regionId!=0 && page=="B") {
		where = where+ " and banners.region_id="+strconv.Itoa(int(regionId))	
	}
	//println("where: ", where)
	banners := make([]*pkgbanner.Banner, 0)
		err := r.DB.Where("deleted_at IS NULL").
				Preload("City").
				Preload("Professions").
				Preload("Region").
				Where(where).
			//Preload("page").
			Find(&banners).Error
	
		if err != nil {
			return nil, err
		}
	
		return banners, nil
	}
	

func (r *repository) Save(banner pkgbanner.Banner) (*pkgbanner.Banner, error) {
	// Buscar a cidade correspondente ao CityID
	//var city pkgcity.City
	
/*	if err := r.DB.First(&city, banner.CityID).Error; err != nil {
		return nil, err
	}
		*/

	// Buscar as profissões correspondentes aos IDs fornecidos
	var professions []*pkgprofession.Profession
	/*if err := r.DB.Where("id IN ?", banner.ProfessionIDs).Find(&professions).Error; err != nil {
		return nil, err
	}*/

	fmt.Println( banner.CityID)

	newBanner := pkgbanner.Banner{
		Link:        banner.Link,
		Image:       banner.Image,
		CityID:      banner.CityID,
		//City:        city,
		Professions: professions,
		Page:        banner.Page,
		RegionID:    banner.RegionID,
	}
	if err := r.DB.Save(&newBanner).Error; err != nil {
		return nil, err
	}
	return &banner, nil
}

func (r *repository) FindByCityAndProfession(cityId, professionId uint) ([]*pkgbanner.Banner, error) {
	banners := make([]*pkgbanner.Banner, 0)
	err := r.DB.Joins("JOIN banner_professions bp ON bp.banner_id = banners.id").
		Where("deleted_at IS NULL").
		//Where("banners.city_id = ? AND bp.profession_id = ?", cityId, professionId).
		Where("banners.city_id = ? ", cityId).		
		Preload("City").
		Preload("Professions").
		Preload("Regions").
		Find(&banners).Error

	if err != nil {
		return nil, err
	}

	return banners, nil
}

func (r *repository) Remove(id uint) error {

	if err := r.DB.Delete(&pkgbanner.Banner{}, id).Error; err != nil {
		return err
	}
	return nil
}

