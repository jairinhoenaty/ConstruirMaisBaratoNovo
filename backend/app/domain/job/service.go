package job

type JobService interface {
	CreateJob(job *Job) error
	UpdateJob(job *Job) error
	CloseJob(jobID uint) error
	ActivateJob(jobID uint) error
	GetJobDetails(jobID uint) (*Job, error)
	SearchJobs(filters map[string]interface{}) ([]*Job, error)
}


type jobService struct {
    repository JobRepository
}

func NewJobService(repo JobRepository) JobService {
    return &jobService{repository: repo}
}

func (s *jobService) CreateJob(job *Job) (*Job, error) {
    return s.repository.Save(job)
}

func (s *jobService) ListJobs() ([]Job, error) {
    return s.repository.FindAll()
}
