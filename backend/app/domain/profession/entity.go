package profession

import (
	pkgprofessionCategory "construir_mais_barato/app/domain/professionCategory"

	"gorm.io/gorm"
)

type Profession struct {
	gorm.Model
	Name        string
	Description string
	Icon        string
	CategoryID  *uint                                     `gorm:"index"`
	Category    *pkgprofessionCategory.ProfessionCategory `gorm:"foreignKey:CategoryID"`
	// Nulo herda o ClientTravels da categoria; preenchido sobrescreve só esta profissão.
	ClientTravels *bool
}

// ClientTravelsResolved diz se é o cliente quem se desloca até o profissional.
// Depende de Category estar carregada quando ClientTravels é nulo.
func (p Profession) ClientTravelsResolved() bool {
	if p.ClientTravels != nil {
		return *p.ClientTravels
	}
	if p.Category != nil {
		return p.Category.ClientTravels
	}
	return false
}
