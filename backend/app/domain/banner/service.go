package banner

type BannerService interface {
	Save(banner Banner) (*Banner, error)
	FindByCityAndProfession(cityId, professionId uint) ([]*Banner, error)
	FindByPage(page string, cityId uint, regionId uint) ([]*Banner, error)
	Remove(id uint) error
}

type bannerService struct {
	repository BannerRepository
}

func NewBannerService(repository BannerRepository) BannerService {
	return &bannerService{
		repository: repository,
	}
}

func (s *bannerService) Save(banner Banner) (*Banner, error) {
	newbanner, err := s.repository.Save(banner)
	if err != nil {
		return nil, err
	}
	return newbanner, nil
}

func (s *bannerService) FindByCityAndProfession(cityId, professionId uint) ([]*Banner, error) {
	banners, err := s.repository.FindByCityAndProfession(cityId, professionId)
	if err != nil {
		return nil, err
	}
	return banners, nil
}

func (s *bannerService) FindByPage(page string, cityId uint, regionId uint) ([]*Banner, error) {
	banners, err := s.repository.FindByPage(page, cityId, regionId)
	if err != nil {
		return nil, err
	}
	return banners, nil
}

func (s *bannerService) Remove(id uint) error {
	return s.repository.Remove(id)
}
