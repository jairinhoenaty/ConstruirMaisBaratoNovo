package city

type CityService interface {
	FindAll() ([]*City, error)
	FindById(id uint) (*City, error)
	FindByUF(uf string) ([]*City, error)
	Save(City City) (*City, error)
	Remove(id uint) error
}

type cityService struct {
	repository CityRepository
}

func NewCityService(repository CityRepository) CityService {
	return &cityService{
		repository: repository,
	}
}

func (s *cityService) FindAll() ([]*City, error) {
	cities, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return cities, nil
}
func (s *cityService) FindById(id uint) (*City, error) {
	city, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}
	return city, nil
}

func (s *cityService) FindByUF(uf string) ([]*City, error) {
	cities, err := s.repository.FindByUF(uf)
	if err != nil {
		return nil, err
	}
	return cities, nil
}

func (s *cityService) Save(City City) (*City, error) {
	newCity, err := s.repository.Save(City)
	if err != nil {
		return nil, err
	}
	return newCity, nil
}
func (s *cityService) Remove(id uint) error {
	return s.repository.Remove(id)
}
