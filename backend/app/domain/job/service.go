package job

type JobService interface {
	CreateJob(job *Job) error
	UpdateJob(job *Job) error
	CloseJob(jobID uint) error
	ActivateJob(jobID uint) error
	GetJobDetails(jobID uint) (*Job, error)
	SearchJobs(filters map[string]interface{}) ([]*Job, error)
}
