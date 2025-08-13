package professional_usecase

import "time"

type ProfessionalPresenter struct {
	ID           uint                  `json:"oid,omitempty"`
	Name         string                `json:"nome"`
	Email        string                `json:"email"`
	Telephone    string                `json:"telefone"`
	LgpdAceito   string                `json:"lgpdaceito"`
	CreatedAt    time.Time             `json:"created_at"`
	Cep          string                `json:"cep"`
	Street       string                `json:"endereco"`
	Neighborhood string                `json:"bairro"`
	Cidade       CidadePresenter       `json:"cidade"`
	Professions  []ProfissionPresenter `json:"profissoes"`
	Image        []byte                `json:"image"`
	Distance     float64               `json:"distance"`
	OnLine       *bool                 `json:"online"`
	Verified     *bool                 `json:"verified"`
}

type CidadePresenter struct {
	ID   uint   `json:"oid,omitempty"`
	Name string `json:"nome"`
	UF   string `json:"uf"`
}

type ProfissionPresenter struct {
	ID          uint   `json:"oid,omitempty"`
	Name        string `json:"nome"`
	Description string `json:"descricao"`
	Icon        string `json:"icone"`
}

type ProfessionCountPresenter struct {
	ProfessionName string `json:"professionName"`
	Quantity       int    `json:"quantity"`
}

type CountProfessionalsByCity struct {
	CityName          string `json:"cidade"`
	TotalProfessional uint   `json:"totalDeProfissionais"`
}

type CityProfessionalCountPresenter struct {
	CityID            uint   `json:"cidadeId"`
	CityName          string `json:"cidade"`
	ProfessionalCount int64  `json:"quantidadeProfissionais"`
}

type UFProfessionalCountPresenter struct {
	UFName            string `json:"UF"`
	ProfessionalCount int64  `json:"quantity"`
}
