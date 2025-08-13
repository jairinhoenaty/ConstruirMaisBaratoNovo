package profession_usecase

type ProfessionPresenter struct {
	ID          uint   `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type ProfessionWithCountPresenter struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}
