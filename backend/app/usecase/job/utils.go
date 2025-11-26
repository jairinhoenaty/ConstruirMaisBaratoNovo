package job_usecase

import (
	pkgjob "construir_mais_barato/app/domain/job"
)



func GenerateJob(assembler *JobAssembler) pkgjob.Job {
	job := pkgjob.Job{}

	if assembler != nil {
		job.ID = assembler.ID
		job.Title = assembler.Title
		job.HiringType = assembler.HiringType
		job.Salary = assembler.Salary
		job.SalaryType = assembler.SalaryType
		job.Location = assembler.Location
		job.Description = assembler.Description
		job.Schedule = assembler.Schedule
		job.Requirements = assembler.Requirements
		job.Benefits = assembler.Benefits
		job.ContactEmail = assembler.ContactEmail
		job.ContactPhone = assembler.ContactPhone
		job.OpeningsQuantity = assembler.OpeningsQuantity

		job.ProfessionID = assembler.ProfessionID
		job.CityID = assembler.CityID
		job.CompanyID = assembler.CompanyID


		job.Status = assembler.Status
		job.PublishedAt = assembler.PublishedAt
		job.ExpiresAt = assembler.ExpiresAt
	}

	return job
}



func GenerateJobPresenter(job *pkgjob.Job) JobPresenter {

	presenter := JobPresenter{}

	presenter.ID = job.ID
	presenter.Title = job.Title
	presenter.HiringType = job.HiringType
	presenter.Salary = job.Salary
	presenter.SalaryType = job.SalaryType
	presenter.Location = job.Location
	presenter.Description = job.Description
	presenter.Schedule = job.Schedule
	presenter.Requirements = job.Requirements
	presenter.Benefits = job.Benefits
	presenter.ContactEmail = job.ContactEmail
	presenter.ContactPhone = job.ContactPhone
	presenter.OpeningsQuantity = job.OpeningsQuantity

	presenter.ProfessionID = job.ProfessionID
	presenter.CityID = job.CityID
	presenter.CompanyID = job.CompanyID


	presenter.Status = job.Status
	presenter.PublishedAt = job.PublishedAt
	presenter.ExpiresAt = job.ExpiresAt

	return presenter
}


func GenerateJobsPresenter(jobs []*pkgjob.Job) *[]JobPresenter {
	list := make([]JobPresenter, 0)

	if jobs != nil && len(jobs) > 0 {
		for _, job := range jobs {

			presenter := JobPresenter{}

			presenter.ID = job.ID
			presenter.Title = job.Title
			presenter.HiringType = job.HiringType
			presenter.Salary = job.Salary
			presenter.SalaryType = job.SalaryType
			presenter.Location = job.Location
			presenter.Description = job.Description
			presenter.Schedule = job.Schedule
			presenter.Requirements = job.Requirements
			presenter.Benefits = job.Benefits
			presenter.ContactEmail = job.ContactEmail
			presenter.ContactPhone = job.ContactPhone
			presenter.OpeningsQuantity = job.OpeningsQuantity

			presenter.ProfessionID = job.ProfessionID
			presenter.CityID = job.CityID
			presenter.CompanyID = job.CompanyID


			presenter.Status = job.Status
			presenter.PublishedAt = job.PublishedAt
			presenter.ExpiresAt = job.ExpiresAt

			list = append(list, presenter)
		}
	}

	return &list
}
