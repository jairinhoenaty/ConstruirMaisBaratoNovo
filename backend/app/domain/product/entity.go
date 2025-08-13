package product

import (
	pkgproductCategory "construir_mais_barato/app/domain/productCategory"
	pkgprofessional "construir_mais_barato/app/domain/professional"
	pkgstore "construir_mais_barato/app/domain/store"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name          string
	Image         []byte `gorm:"type:longblob"`
	Price         float64
	OriginalPrice float64
	Discount      float64
	Approved      bool
	Dayoffer      bool
	//	Professions   []*pkgprofession.Profession `gorm:"many2many:banner_professions;" json:"professions"`
	ProfessionID   uint
	ProfessionalID *uint
	Professional   pkgprofessional.Professional `gorm:"foreignkey:ProfessionalID"`
	Description    string
	CategoryID     uint
	Category       pkgproductCategory.ProductCategory `gorm:"foreignkey:CategoryID"`
	CreatedAt      time.Time                          `gorm:"<-:create"`
	StoreID        *uint
	Store          pkgstore.Store `gorm:"foreignkey:StoreID"`
}
