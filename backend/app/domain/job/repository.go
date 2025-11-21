package job

type JobRepository interface {
	Create(job *Job) error
	Update(job *Job) error
	Delete(id uint) error
	FindByID(id uint) (*Job, error)
	FindAll() ([]*Job, error)
	FindByCompanyID(companyID uint) ([]*Job, error)
	FindByProfessionID(professionID uint) ([]*Job, error)
	FindByCityID(cityID uint) ([]*Job, error)
	FindActiveJobs() ([]*Job, error)
	FindByStatus(status string) ([]*Job, error)
	FindByMultipleCriteria(companyID, cityID, professionID *uint, status string) ([]*Job, error)
}
