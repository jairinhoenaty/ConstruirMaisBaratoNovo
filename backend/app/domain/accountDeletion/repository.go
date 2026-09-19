package accountdeletion

type AccountDeletionRepository interface {
	Save(request AccountDeletionRequest) (*AccountDeletionRequest, error)
	FindAll() ([]*AccountDeletionRequest, error)
	FindByID(id uint) (*AccountDeletionRequest, error)
	UpdateStatus(id uint, status string) error
}
