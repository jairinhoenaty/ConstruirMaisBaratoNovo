package productCategory

import (
	pkgprofession "construir_mais_barato/app/domain/profession"

	"gorm.io/gorm"
)

type ProductCategory struct {
	gorm.Model
	Name         string
	ProfessionID uint
	Profession   pkgprofession.Profession `gorm:"foreignkey:ProfessionID"`
}
