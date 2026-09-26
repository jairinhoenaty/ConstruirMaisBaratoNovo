package professionCategory_usecase

import pkgprofessionuc "construir_mais_barato/app/usecase/profession"

type ProfessionCategoryPresenter struct {
	ID              uint   `json:"id,omitempty"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Icon            string `json:"icon"`
	Position        int    `json:"position"`
	Active          bool   `json:"active"`
	ClientTravels   bool   `json:"client_travels"`
	ProfessionCount int64  `json:"profession_count"`
}

type CategoryWithProfessionsPresenter struct {
	ProfessionCategoryPresenter
	Professions []pkgprofessionuc.ProfessionPresenter `json:"professions"`
}
