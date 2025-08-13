package productCategory_usecase

import "construir_mais_barato/app/domain/profession"

type ProductCategoryAssembler struct {
	ID           uint                  `json:"id,omitempty"`
	Name         string                `json:"name"`
	ProfessionID uint                  `json:"profession_id"`
	Profession   profession.Profession `json:"profession"`
}

type FindByProfessionAssembler struct {
	ProfessionID int `json:"profession_id"`
}
