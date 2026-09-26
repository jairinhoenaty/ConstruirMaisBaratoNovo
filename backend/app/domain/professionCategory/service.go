package professionCategory

type ProfessionCategoryService interface {
	FindAll(limit, offset int) ([]*ProfessionCategory, int64, error)
	FindActive() ([]*ProfessionCategory, error)
	FindById(id uint) (*ProfessionCategory, error)
	CountProfessions(categoryIDs []uint) (map[uint]int64, error)
	Save(category ProfessionCategory) (*ProfessionCategory, error)
	Remove(id uint) error
}

type professionCategoryService struct {
	repository ProfessionCategoryRepository
}

func NewProfessionCategoryService(repository ProfessionCategoryRepository) ProfessionCategoryService {
	return &professionCategoryService{
		repository: repository,
	}
}

func (s *professionCategoryService) FindAll(limit, offset int) ([]*ProfessionCategory, int64, error) {
	return s.repository.FindAll(limit, offset)
}

func (s *professionCategoryService) FindActive() ([]*ProfessionCategory, error) {
	return s.repository.FindActive()
}

func (s *professionCategoryService) FindById(id uint) (*ProfessionCategory, error) {
	return s.repository.FindById(id)
}

func (s *professionCategoryService) CountProfessions(categoryIDs []uint) (map[uint]int64, error) {
	return s.repository.CountProfessions(categoryIDs)
}

func (s *professionCategoryService) Save(category ProfessionCategory) (*ProfessionCategory, error) {
	return s.repository.Save(category)
}

func (s *professionCategoryService) Remove(id uint) error {
	return s.repository.Remove(id)
}
