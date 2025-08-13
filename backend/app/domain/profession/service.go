package profession

type ProfessionService interface {
	Find(quantityProfessions uint) ([]*Profession, error)
	FindAll(limit, offset int) ([]*Profession, int64, error)
	FindAllWithoutPagination() ([]*Profession, error)
	FindById(id uint) (*Profession, error)
	FindProfessionsWithCount() ([]map[string]interface{}, error)
	Save(profession Profession) (*Profession, error)
	Remove(id uint) error
}

type professionService struct {
	repository ProfessionRepository
}

func NewProfessionService(repository ProfessionRepository) ProfessionService {
	return &professionService{
		repository: repository,
	}
}

func (s *professionService) FindAllWithoutPagination() ([]*Profession, error) {
	professions, err := s.repository.FindAllWithoutPagination()
	if err != nil {
		return nil, err
	}
	return professions, nil
}

func (s *professionService) FindProfessionsWithCount() ([]map[string]interface{}, error) {
	professionals, err := s.repository.FindProfessionsWithCount()
	if err != nil {
		return nil, err
	}
	return professionals, nil
}

func (s *professionService) Find(quantityProfessions uint) ([]*Profession, error) {
	professions, err := s.repository.Find(quantityProfessions)
	if err != nil {
		return nil, err
	}
	return professions, nil
}

func (s *professionService) FindAll(limit, offset int) ([]*Profession, int64, error) {
	professions, total, err := s.repository.FindAll(limit, offset)
	if err != nil {
		return nil, 0, err
	}
	return professions, total, nil
}

func (s *professionService) FindById(id uint) (*Profession, error) {
	profession, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}
	return profession, nil
}

func (s *professionService) Save(profession Profession) (*Profession, error) {
	newprofession, err := s.repository.Save(profession)
	if err != nil {
		return nil, err
	}
	return newprofession, nil
}

func (s *professionService) Remove(id uint) error {
	return s.repository.Remove(id)
}
