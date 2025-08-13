package productCategory_usecase

import "construir_mais_barato/app/domain/profession"

type ProductCategoryPresenter struct {
	ID           uint                  `json:"id,omitempty"`
	Name         string                `json:"name"`
	ProfessionID uint                  `json:"professional_id"`
	Profession   profession.Profession `json:"profession"`
}
