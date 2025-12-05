package productCategory_usecase

import "construir_mais_barato/app/domain/profession"

type ProductCategoryPresenter struct {
	ID           uint                        `json:"id,omitempty"`
	Name         string                     `json:"name"`
	ProfessionID *uint                      `json:"professional_id,omitempty"`
	Profession   *profession.Profession     `json:"profession,omitempty"`
	ParentID     *uint                      `json:"parent_id,omitempty"`
	Children     []ProductCategoryPresenter `json:"children,omitempty"`
}
