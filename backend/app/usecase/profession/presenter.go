package profession_usecase

type ProfessionPresenter struct {
	ID          uint   `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	CategoryID  *uint  `json:"category_id,omitempty"`
	// ClientTravels já vem resolvido (override da profissão ou padrão da
	// categoria): app e plataforma consomem este campo e não recalculam nada.
	ClientTravels bool `json:"client_travels"`
	// Só o administrador precisa distinguir "herda" (nulo) de "sobrescreve".
	ClientTravelsOverride *bool `json:"client_travels_override"`
}

type ProfessionWithCountPresenter struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}
