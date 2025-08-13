package product_usecase

import (
	"construir_mais_barato/app/domain/productCategory"
	"construir_mais_barato/app/domain/professional"
	"construir_mais_barato/app/domain/store"
)

type ProductPresenter struct {
	ID            uint    `json:"id,omitempty"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Image         []byte  `json:"image"`
	Price         float64 `json:"price"`
	OriginalPrice float64 `json:"originalprice"`
	Discount      float64 `json:"discount"`
	Approved      bool    `json:"approved"`
	Dayoffer      bool    `json:"dayoffer"`
	//	Professions   []*pkgprofession.Profession `gorm:"many2many:banner_professions;" json:"professions"`
	ProfessionID   uint                            `json:"professionID"`
	CategoryID     uint                            `json:"categoryID"`
	Category       productCategory.ProductCategory `json:"category"`
	ProfessionalID *uint                           `json:"professionalID"`
	Professional   professional.Professional       `json:"professional"`
	StoreID        *uint                           `json:"storeID"`
	Store          store.Store                     `json:"store"`
}
