package region

type RegionService interface {
	Find(quantityRegions uint) ([]*Region, error)
	FindAll(limit, offset int, uf string) ([]*Region, int64, error)
	FindAllWithoutPagination() ([]*Region, error)
	FindByCity(cityId uint) (*Region, error)
	FindById(id uint) (*Region, error)
	FindRegionsWithCount() ([]map[string]interface{}, error)
	Save(region Region) (*Region, error)
	Remove(id uint) error
}

type regionService struct {
	repository RegionRepository
}

func NewRegionService(repository RegionRepository) RegionService {
	return &regionService{
		repository: repository,
	}
}

func (s *regionService) FindAllWithoutPagination() ([]*Region, error) {
	regions, err := s.repository.FindAllWithoutPagination()
	if err != nil {
		return nil, err
	}
	return regions, nil
}

func (s *regionService) FindRegionsWithCount() ([]map[string]interface{}, error) {
	regionals, err := s.repository.FindRegionsWithCount()
	if err != nil {
		return nil, err
	}
	return regionals, nil
}

func (s *regionService) FindByCity(cityId uint) (*Region, error) {
	region, err := s.repository.FindByCity(cityId)
	if err != nil {
		return nil, err
	}
	return region, nil
}

func (s *regionService) Find(quantityRegions uint) ([]*Region, error) {
	regions, err := s.repository.Find(quantityRegions)
	if err != nil {
		return nil, err
	}
	return regions, nil
}

func (s *regionService) FindAll(limit, offset int, uf string) ([]*Region, int64, error) {
	regions, total, err := s.repository.FindAll(limit, offset, uf)
	if err != nil {
		return nil, 0, err
	}
	return regions, total, nil
}

func (s *regionService) FindById(id uint) (*Region, error) {
	region, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}
	return region, nil
}

func (s *regionService) Save(region Region) (*Region, error) {
	newregion, err := s.repository.Save(region)
	if err != nil {
		return nil, err
	}
	return newregion, nil
}

func (s *regionService) Remove(id uint) error {
	return s.repository.Remove(id)
}
