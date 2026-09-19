package accountdeletion

type AccountDeletionService interface {
	Save(request AccountDeletionRequest) (*AccountDeletionRequest, error)
	FindAll() ([]*AccountDeletionRequest, error)
	FindByID(id uint) (*AccountDeletionRequest, error)
	UpdateStatus(id uint, status string) error
}

type accountDeletionService struct {
	repository AccountDeletionRepository
}

func NewAccountDeletionService(
	repository AccountDeletionRepository,
) AccountDeletionService {
	return &accountDeletionService{
		repository: repository,
	}
}

func (s *accountDeletionService) Save(
	request AccountDeletionRequest,
) (*AccountDeletionRequest, error) {
	return s.repository.Save(request)
}

func (s *accountDeletionService) FindAll() ([]*AccountDeletionRequest, error) {
	return s.repository.FindAll()
}

func (s *accountDeletionService) FindByID(
	id uint,
) (*AccountDeletionRequest, error) {
	return s.repository.FindByID(id)
}

func (s *accountDeletionService) UpdateStatus(
	id uint,
	status string,
) error {
	return s.repository.UpdateStatus(id, status)
}
