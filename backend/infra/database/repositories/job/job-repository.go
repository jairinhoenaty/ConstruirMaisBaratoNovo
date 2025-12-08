package job

import (
	jobdomain "construir_mais_barato/app/domain/job"
	"gorm.io/gorm"
)

type JobRepository struct {
	DB *gorm.DB
}

func NewJobRepository(db *gorm.DB) *JobRepository {
	return &JobRepository{DB: db}
}

func (r *JobRepository) Save(j jobdomain.Job) (*jobdomain.Job, error) {
	if err := r.DB.Create(&j).Error; err != nil {
		return nil, err
	}
	return &j, nil
}

func (r *JobRepository) FindAll() ([]*jobdomain.Job, error) {
	var jobs []*jobdomain.Job
	err := r.DB.
		Order("created_at desc").
		Find(&jobs).Error
	return jobs, err
}

func (r *JobRepository) FindApproved() ([]*jobdomain.Job, error) {
	var jobs []*jobdomain.Job
	err := r.DB.
		Where("approved = ?", true).
		Order("created_at desc").
		Find(&jobs).Error
	return jobs, err
}

func (r *JobRepository) Approve(id uint) error {
	return r.DB.Model(&jobdomain.Job{}).
		Where("id = ?", id).
		Update("approved", true).Error
}

func (r *JobRepository) Disapprove(id uint) error {
	return r.DB.Model(&jobdomain.Job{}).
		Where("id = ?", id).
		Update("approved", false).Error
}

func (r *JobRepository) Remove(id uint) error {
	return r.DB.Delete(&jobdomain.Job{}, id).Error
}
