package city_usecase

type CityAssembler struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name"`
	UF   string `json:"uf"`
}

type UFCityAssembler struct {
	UF string `json:"uf"`
}
