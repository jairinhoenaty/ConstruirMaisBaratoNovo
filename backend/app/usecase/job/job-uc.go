package jobuc

import jobdomain "construir_mais_barato/app/domain/job"

type JobUseCase struct {
	service jobdomain.JobService
}

func NewJobUseCase(s jobdomain.JobService) *JobUseCase {
	return &JobUseCase{service: s}
}

func (uc *JobUseCase) Save(job jobdomain.Job) (*jobdomain.Job, error) {
	return uc.service.Save(job)
}
func (uc *JobUseCase) ListAll() ([]*jobdomain.Job, error) {
	return uc.service.FindAll()
}
func (uc *JobUseCase) ListApproved() ([]*jobdomain.Job, error) {
	return uc.service.FindApproved()
}
func (uc *JobUseCase) Approve(id uint) error {
	return uc.service.Approve(id)
}
func (uc *JobUseCase) Disapprove(id uint) error {
	return uc.service.Disapprove(id)
}
func (uc *JobUseCase) Delete(id uint) error {
	return uc.service.Remove(id)
}
