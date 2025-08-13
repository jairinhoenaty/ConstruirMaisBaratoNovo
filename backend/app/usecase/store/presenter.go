package store_usecase

import "time"

type StorePresenter struct {
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
	Image        []byte  			   `json:"image"`	
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

type CountStoresByCity struct {
	CityName   string `json:"cidade"`
	TotalStore uint   `json:"totalDeProfissionais"`
}

type CityStoreCountPresenter struct {
	CityID     uint   `json:"cidadeId"`
	CityName   string `json:"cidade"`
	StoreCount int64  `json:"quantidadeProfissionais"`
}
