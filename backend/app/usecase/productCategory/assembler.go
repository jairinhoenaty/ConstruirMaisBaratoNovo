package productCategory_usecase

import "construir_mais_barato/app/domain/profession"

type ProductCategoryAssembler struct {
	ID           uint                   `json:"id,omitempty"`
	Name         string                 `json:"name"`
	ProfessionID *uint                  `json:"profession_id,omitempty"`
	ParentID     *uint                  `json:"parent_id,omitempty"`
	Profession   *profession.Profession `json:"profession,omitempty"`
}

type FindByProfessionAssembler struct {
	ProfessionID int `json:"profession_id"`
}
