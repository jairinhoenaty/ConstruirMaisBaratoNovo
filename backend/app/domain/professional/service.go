package professional

type ProfessionalService interface {
	FindAll(limit, offset int, filter string, uf string, professionID int, order string) ([]*Professional, int64, error)
	FindById(id uint) (*Professional, error)
	FindByEmail(email string) (*Professional, error)
	FindByName(email string) ([]*Professional, error)
	FindByCityAndProfession(cityID, professionID uint, limit, offset int) ([]*Professional, int64, error)
	FindByProfessionAndLocation(professionID uint, latitude float32, longitude float32, distance, limit, offset int) ([]*Professional, int64, error)
	FindByNameAndCityAndProfession(name string, cityID, professionID uint, limit, offset int) ([]*Professional, error)
	CountProfessionalsByProfession() ([]ProfessionCount, error)
	CountCityProfessionalsByState(uf string, limit, offset int) ([]CityProfessionalCount, *int64, error)
	CountProfessionalsByState(uf string, limit, offset int) ([]UFProfessionalCount, *int64, error)
	CountProfessionalsByProfessionInCity(cityID uint) ([]ProfessionCount, error)
	FindLastProfessionals(quantityRecords int) ([]Professional, error)
	Save(professional Professional) (*Professional, error)
	Remove(id uint) error
	ExportXLSX() ([]*Professional, error)
}

type professionalService struct {
	repository ProfessionalRepository
}

func NewProfessionalService(repository ProfessionalRepository) ProfessionalService {
	return &professionalService{
		repository: repository,
	}
}

func (s *professionalService) ExportXLSX() ([]*Professional, error) {
	professionals, err := s.repository.ExportXLSX()
	if err != nil {
		return nil, err
	}
	return professionals, nil
}

func (s *professionalService) FindByNameAndCityAndProfession(name string, cityID, professionID uint, limit, offset int) ([]*Professional, error) {
	professionals, err := s.repository.FindByNameAndCityAndProfession(name, cityID, professionID, limit, offset)
	if err != nil {
		return nil, err
	}
	return professionals, nil
}

func (s *professionalService) CountProfessionalsByProfessionInCity(cityID uint) ([]ProfessionCount, error) {
	results, err := s.repository.CountProfessionalsByProfessionInCity(cityID)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (s *professionalService) CountProfessionalsByState(uf string, limit, offset int) ([]UFProfessionalCount, *int64, error) {
	results, total, err := s.repository.CountProfessionalsByState(uf, limit, offset)
	if err != nil {
		return nil, nil, err
	}
	return results, total, nil
}

func (s *professionalService) CountCityProfessionalsByState(uf string, limit, offset int) ([]CityProfessionalCount, *int64, error) {
	results, total, err := s.repository.CountCityProfessionalsByState(uf, limit, offset)
	if err != nil {
		return nil, nil, err
	}
	return results, total, nil
}

func (s *professionalService) CountProfessionalsByProfession() ([]ProfessionCount, error) {
	results, err := s.repository.CountProfessionalsByProfession()
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (s *professionalService) FindLastProfessionals(quantityRecords int) ([]Professional, error) {
	professionals, err := s.repository.FindLastProfessionals(quantityRecords)
	if err != nil {
		return nil, err
	}
	return professionals, nil
}

func (s *professionalService) FindByCityAndProfession(cityID, professionID uint, limit, offset int) ([]*Professional, int64, error) {
	professionals, total, err := s.repository.FindByCityAndProfession(cityID, professionID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	return professionals, total, nil
}

func (s *professionalService) FindByProfessionAndLocation(professionID uint, latitude float32, longitude float32, distance, limit, offset int) ([]*Professional, int64, error) {
	professionals, total, err := s.repository.FindByProfessionAndLocation(professionID, latitude, longitude, distance, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	return professionals, total, nil
}

func (s *professionalService) FindAll(limit, offset int, filter string, uf string, professionID int, order string) ([]*Professional, int64, error) {
	professionals, count, err := s.repository.FindAll(limit, offset, filter, uf, professionID, order)
	if err != nil {
		return nil, 0, err
	}
	return professionals, count, nil
}

func (s *professionalService) FindById(id uint) (*Professional, error) {
	professional, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}
	return professional, nil
}

func (s *professionalService) FindByEmail(email string) (*Professional, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *professionalService) FindByName(name string) ([]*Professional, error) {
	professionals, err := s.repository.FindByName(name)
	if err != nil {
		return nil, err
	}
	return professionals, nil
}

func (s *professionalService) Save(professional Professional) (*Professional, error) {
	newprofessional, err := s.repository.Save(professional)
	if err != nil {
		return nil, err
	}
	return newprofessional, nil
}

func (s *professionalService) Remove(id uint) error {
	return s.repository.Remove(id)
}
