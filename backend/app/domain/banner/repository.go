package banner

type BannerRepository interface {
	Save(banner Banner) (*Banner, error)
	FindByCityAndProfession(cityId, professionId uint) ([]*Banner, error)
	FindByPage(page string, cityId uint, regionId uint) ([]*Banner, error)
	Remove(id uint) error
}
