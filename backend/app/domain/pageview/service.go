package pageview

type PageViewService interface {
	Increment(path string) (*PageView, error)
	FindAll() ([]*PageView, error)
}

type pageViewService struct {
	repository PageViewRepository
}

func NewPageViewService(repository PageViewRepository) PageViewService {
	return &pageViewService{
		repository: repository,
	}
}

func (s *pageViewService) Increment(path string) (*PageView, error) {
	return s.repository.Increment(path)
}

func (s *pageViewService) FindAll() ([]*PageView, error) {
	return s.repository.FindAll()
}
