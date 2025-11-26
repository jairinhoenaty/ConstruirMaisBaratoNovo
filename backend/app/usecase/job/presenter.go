package job_usecase

import (
    pkgjob "construir_mais_barato/app/domain/job"
)

type JobPresenter struct {
    ID               uint     `json:"id"`
    Cargo            string   `json:"cargo"`
    Contratacao      string   `json:"contratacao"`
    Salario          *float64 `json:"salario"`
    TipoSalario      string   `json:"tipoSalario"`
    Local            string   `json:"local"`
    Descricao        string   `json:"descricao"`
    Horario          string   `json:"horario"`
    Requisitos       string   `json:"requisitos"`
    Beneficios       string   `json:"beneficios"`
    ContatoEmail     string   `json:"contatoEmail"`
    ContatoTelefone  string   `json:"contatoTelefone"`
    QuantidadeVagas  int      `json:"quantidadeVagas"`
    EmpresaID        uint     `json:"empresaId"`
    ProfissaoID      *uint    `json:"profissaoId"`
    CidadeID         uint     `json:"cidadeId"`
    Status           string   `json:"status"`
}


func GenerateJobPresenter(job pkgjob.Job) JobPresenter {

    return JobPresenter{
        ID:              job.ID,
        Cargo:           job.Title,
        Contratacao:     job.HiringType,
        Salario:         job.Salary,
        TipoSalario:     job.SalaryType,
        Local:           job.Location,
        Descricao:       job.Description,
        Horario:         job.Schedule,
        Requisitos:      job.Requirements,
        Beneficios:      job.Benefits,
        ContatoEmail:    job.ContactEmail,
        ContatoTelefone: job.ContactPhone,
        QuantidadeVagas: job.OpeningsQuantity,
        EmpresaID:       job.CompanyID,
        ProfissaoID:     job.ProfessionID,
        CidadeID:        job.CityID,
        Status:          job.Status,
    }
}
