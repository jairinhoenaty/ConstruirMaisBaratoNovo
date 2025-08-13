package city_usecase

type CityPresenter struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name"`
	UF   string `json:"uf"`
}
