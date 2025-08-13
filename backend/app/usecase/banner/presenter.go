package banner_usecase

type BannerPresenter struct {
	ID             uint                   `json:"id,omitempty"`
	AccessLink     string                 `json:"link"`
	Cidade         CidadePresenter        `json:"cidade"`
	Professions    *[]ProfessionPresenter `json:"professions"`
	ProfessionsIds *[]uint                `json:"professionsIds"`
	Page           string                 `json:"page"`
	Image          []byte                 `json:"image"`
	Region         RegionPresenter        `json:"region"`
}

type ProfessionPresenter struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name"`
}

type CidadePresenter struct {
	ID   uint   `json:"oid,omitempty"`
	Name string `json:"nome"`
	UF   string `json:"uf"`
}

type RegionPresenter struct {
	ID   uint   `json:"oid,omitempty"`
	Name string `json:"nome"`
	UF   string `json:"uf"`
}
