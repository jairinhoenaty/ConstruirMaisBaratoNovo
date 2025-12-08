package job

type JobService interface {
	Save(job Job) (*Job, error)
	FindAll() ([]*Job, error)
	FindApproved() ([]*Job, error)
	Approve(id uint) error
	Disapprove(id uint) error
	Remove(id uint) error
}

type jobService struct {
	repository JobRepository
}


type JobRepository interface {
	Save(job Job) (*Job, error)
	FindAll() ([]*Job, error)
	FindApproved() ([]*Job, error)
	Approve(id uint) error
	Disapprove(id uint) error
	Remove(id uint) error
}

func NewJobService(repository JobRepository) JobService {
	return &jobService{repository}
}

func (s *jobService) Save(job Job) (*Job, error) {
	return s.repository.Save(job)
}

func (s *jobService) FindAll() ([]*Job, error) {
	return s.repository.FindAll()
}

func (s *jobService) FindApproved() ([]*Job, error) {
	return s.repository.FindApproved()
}

func (s *jobService) Approve(id uint) error {
	return s.repository.Approve(id)
}

func (s *jobService) Disapprove(id uint) error {
	return s.repository.Disapprove(id)
}

func (s *jobService) Remove(id uint) error {
	return s.repository.Remove(id)
}
