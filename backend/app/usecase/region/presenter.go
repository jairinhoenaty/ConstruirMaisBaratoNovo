package region_usecase

type RegionPresenter struct {
	ID          uint              `json:"id,omitempty"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Icon        string            `json:"icon"`
	Cities      []CidadePresenter `json:"cities"`
	UF          string            `json:"uf"`
}

type CidadePresenter struct {
	ID   uint   `json:"oid,omitempty"`
	Name string `json:"nome"`
	UF   string `json:"uf"`
}

type RegionWithCountPresenter struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}
