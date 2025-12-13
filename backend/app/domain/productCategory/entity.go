package productCategory

import (
	pkgprofession "construir_mais_barato/app/domain/profession"

	"gorm.io/gorm"
)

type ProductCategory struct {
	gorm.Model
	Name       string
	ProfessionID *uint                    `gorm:"index"` // NULL para categorias genéricas
	Profession *pkgprofession.Profession `gorm:"foreignkey:ProfessionID"`
	ParentID   *uint                     `gorm:"index"` // NULL para categorias pai
	Parent     *ProductCategory          `gorm:"foreignkey:ParentID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Children   []ProductCategory         `gorm:"foreignkey:ParentID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}
