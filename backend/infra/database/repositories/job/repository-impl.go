package job

import (
	"errors"
	pkgjob "construir_mais_barato/app/domain/job"

	"gorm.io/gorm"
)

type JobRepository struct {
	DB *gorm.DB
}

func NewJobRepository(db *gorm.DB) pkgjob.JobRepository {
	return &JobRepository{DB: db}
}

func (r *JobRepository) Create(job *pkgjob.Job) error {
	if job == nil {
		return errors.New("job cannot be nil")
	}
	return r.DB.Create(job).Error
}

func (r *JobRepository) Update(job *pkgjob.Job) error {
	if job == nil {
		return errors.New("job cannot be nil")
	}
	return r.DB.Save(job).Error
}

func (r *JobRepository) Delete(id uint) error {
	if id == 0 {
		return errors.New("invalid job id")
	}
	return r.DB.Delete(&pkgjob.Job{}, id).Error
}

func (r *JobRepository) FindByID(id uint) (*pkgjob.Job, error) {
	if id == 0 {
		return nil, errors.New("invalid job id")
	}
	var job *pkgjob.Job
	err := r.DB.
		Preload("Company").
		Preload("Profession").
		Preload("City").
		First(&job, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("job not found")
		}
		return nil, err
	}
	return job, nil
}

func (r *JobRepository) FindAll() ([]*pkgjob.Job, error) {
	var jobs []*pkgjob.Job
	err := r.DB.
		Preload("Company").
		Preload("Profession").
		Preload("City").
		Find(&jobs).Error
	return jobs, err
}

func (r *JobRepository) FindByCompanyID(companyID uint) ([]*pkgjob.Job, error) {
	if companyID == 0 {
		return nil, errors.New("invalid company id")
	}
	var jobs []*pkgjob.Job
	err := r.DB.
		Where("company_id = ?", companyID).
		Preload("Company").
		Preload("Profession").
		Preload("City").
		Find(&jobs).Error
	return jobs, err
}

func (r *JobRepository) FindByProfessionID(professionID uint) ([]*pkgjob.Job, error) {
	if professionID == 0 {
		return nil, errors.New("invalid profession id")
	}
	var jobs []*pkgjob.Job
	err := r.DB.
		Where("profession_id = ?", professionID).
		Preload("Company").
		Preload("Profession").
		Preload("City").
		Find(&jobs).Error
	return jobs, err
}

func (r *JobRepository) FindByCityID(cityID uint) ([]*pkgjob.Job, error) {
	if cityID == 0 {
		return nil, errors.New("invalid city id")
	}
	var jobs []*pkgjob.Job
	err := r.DB.
		Where("city_id = ?", cityID).
		Preload("Company").
		Preload("Profession").
		Preload("City").
		Find(&jobs).Error
	return jobs, err
}

func (r *JobRepository) FindActiveJobs() ([]*pkgjob.Job, error) {
	var jobs []*pkgjob.Job
	err := r.DB.
		Where("status = ? AND (expires_at IS NULL OR expires_at > NOW())", "active").
		Preload("Company").
		Preload("Profession").
		Preload("City").
		Find(&jobs).Error
	return jobs, err
}

func (r *JobRepository) FindByStatus(status string) ([]*pkgjob.Job, error) {
	if status == "" {
		return nil, errors.New("status cannot be empty")
	}
	var jobs []*pkgjob.Job
	err := r.DB.
		Where("status = ?", status).
		Preload("Company").
		Preload("Profession").
		Preload("City").
		Find(&jobs).Error
	return jobs, err
}

func (r *JobRepository) FindByMultipleCriteria(companyID, cityID, professionID *uint, status string) ([]*pkgjob.Job, error) {
	query := r.DB

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if companyID != nil && *companyID != 0 {
		query = query.Where("company_id = ?", *companyID)
	}

	if cityID != nil && *cityID != 0 {
		query = query.Where("city_id = ?", *cityID)
	}

	if professionID != nil && *professionID != 0 {
		query = query.Where("profession_id = ?", *professionID)
	}

	var jobs []*pkgjob.Job
	err := query.
		Preload("Company").
		Preload("Profession").
		Preload("City").
		Find(&jobs).Error
	return jobs, err
}
