package job_usecase


type JobAssembler struct {
    Cargo             string  `json:"cargo"`
    Contratacao       string  `json:"contratacao"`
    Salario           float64 `json:"salario"`
    TipoSalario       string  `json:"tipoSalario"`
    Local             string  `json:"local"`
    Descricao         string  `json:"descricao"`
    Horario           string  `json:"horario"`
    Requisitos        string  `json:"requisitos"`
    Beneficios        string  `json:"beneficios"`
    ContatoEmail      string  `json:"contatoEmail"`
    ContatoTelefone   string  `json:"contatoTelefone"`
    QuantidadeVagas   int     `json:"quantidadeVagas"`
    EmpresaID         uint    `json:"empresaId"`
    ProfissaoID       *uint   `json:"profissaoId"`
    CidadeID          uint    `json:"cidadeId"`
}
