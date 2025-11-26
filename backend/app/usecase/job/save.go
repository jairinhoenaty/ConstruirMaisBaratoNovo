package job_usecase

import (
    "fmt"

    pkgjob "construir_mais_barato/app/domain/job"
)

type SaveJobUC struct {
    Service   pkgjob.JobService
    Assembler *JobAssembler
}

type SaveJobUCParams struct {
    Service pkgjob.JobService
}

func NewSaveJobUC(params SaveJobUCParams) SaveJobUC {
    return SaveJobUC{
        Service: params.Service,
    }
}


func (uc *SaveJobUC) Execute() (*JobPresenter, error) {

    if uc.Assembler == nil {
        return nil, fmt.Errorf("invalid data: assembler is nil")
    }

    jobEntity := GenerateJob(uc.Assembler)

 
    savedJob, err := uc.Service.CreateJob(&jobEntity)
    if err != nil {
        return nil, err
    }

  
    presenter := GenerateJobPresenter(*savedJob)

    return &presenter, nil
}
