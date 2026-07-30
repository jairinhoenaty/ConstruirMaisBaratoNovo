package city_usecase

type CityAssembler struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name"`
	UF   string `json:"uf"`
}

type UFCityAssembler struct {
	UF string `json:"uf"`
}

// A tag query permite o bind a partir da querystring (GET).
type SearchCityAssembler struct {
	Term  string `json:"term" query:"q"`
	Limit int    `json:"limit" query:"limit"`
}
