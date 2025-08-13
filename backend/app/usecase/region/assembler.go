package region_usecase

type RegionAssembler struct {
	ID          uint   `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	CityIDs     []uint `json:"cityIds"`
	UF          string `json:"uf"`
}

type FindWithPaginationRegionAssembler struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	UF     string `json:"uf"`
}
