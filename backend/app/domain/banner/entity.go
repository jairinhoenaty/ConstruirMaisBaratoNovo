package banner

import (
	pkgcity "construir_mais_barato/app/domain/city"
	pkgprofession "construir_mais_barato/app/domain/profession"
	pkgregion "construir_mais_barato/app/domain/region"

	"gorm.io/gorm"
)

type Banner struct {
	gorm.Model
	Link          string
	Image         []byte                      `gorm:"type:longblob"`
	CityID        *uint                       
	City          pkgcity.City                
	Professions   []*pkgprofession.Profession `gorm:"many2many:banner_professions;"`
	ProfessionIDs []uint                      `gorm:"-"`
	Page          string       				  
	RegionID      *uint			  
	Region        pkgregion.Region
}
