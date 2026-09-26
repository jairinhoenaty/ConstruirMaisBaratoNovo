package professionCategory_usecase

type ProfessionCategoryAssembler struct {
	ID            uint   `json:"id,omitempty"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Icon          string `json:"icon"`
	Position      int    `json:"position"`
	Active        bool   `json:"active"`
	ClientTravels bool   `json:"client_travels"`
}

type FindWithPaginationAssembler struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
