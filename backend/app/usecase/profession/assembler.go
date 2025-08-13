package profession_usecase

type ProfessionAssembler struct {
	ID          uint   `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	CityIDs     []uint `json:"city_ids"`
}

type FindWithPaginationProfessionAssembler struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
