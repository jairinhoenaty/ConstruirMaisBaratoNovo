package professionCategory

import (
	"gorm.io/gorm"
)

// ProfessionCategory é a categoria master que agrupa profissões ("Casa e Lar",
// "Automotivo", "Beleza"). ClientTravels vale para todas as profissões da
// categoria, e cada profissão pode sobrescrevê-lo individualmente.
type ProfessionCategory struct {
	gorm.Model
	Name          string `gorm:"index"`
	Description   string
	Icon          string
	Position      int
	Active        bool
	ClientTravels bool
}
