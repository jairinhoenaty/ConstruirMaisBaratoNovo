package professional

import (
	"time"

	"gorm.io/gorm"

	pkgcity "construir_mais_barato/app/domain/city"
	pkgprofession "construir_mais_barato/app/domain/profession"
)

type Professional struct {
	gorm.Model
	Name          string
	Email         string //`gorm:"unique"`
	Company       string
	Telephone     string
	LgpdAceito    string
	CityID        uint
	City          pkgcity.City               `gorm:"foreignKey:CityID"`
	Professions   []pkgprofession.Profession `gorm:"many2many:professional_professions;"`
	ProfessionIDs []uint                     `gorm:"-"`
	Cep           string
	Street        string
	Neighborhood  string
	Image         []byte  `gorm:"type:longblob"`
	Latitude      float64 `gorm:"type:decimal(10,8)"`
	Longitude     float64 `gorm:"type:decimal(11,8)"`
	// Coordenadas do endereço do cadastro. Latitude/Longitude são sobrescritas
	// pelo GPS toda vez que o profissional fica online; estas não, e é delas
	// que o cliente precisa quando é ele quem se desloca.
	AddressLatitude           float64 `gorm:"type:decimal(10,8)"`
	AddressLongitude          float64 `gorm:"type:decimal(11,8)"`
	Verified                  *bool   `gorm:"default:false"`
	NegativeCertificateNumber int64
	OnLine                    *bool
	CreatedAt                 time.Time `gorm:"<-:create"`
	Distance                  float64   `gorm:"->"` //`gorm:"-"`
	IsPremium                 *bool     `gorm:"default:false"`
	// PremiumExpiresAt é o fim da vigência do premium pago. Nulo significa
	// premium sem prazo (ativado manualmente por um administrador), que a
	// rotina de expiração não deve rebaixar.
	PremiumExpiresAt *time.Time
	OnService        *bool  `gorm:"default:false"`
	CodeVerification string `json:"codeVerification"`
	DateOfBirth      string
	Experience       string
	// CodeVerification string
	MeiCnpj    string
	YoutubeUrl string
}

// ServiceLocation é o ponto de atendimento do profissional. Cai na última
// posição conhecida enquanto o cadastro não tiver coordenadas próprias, o que
// vale para todo profissional anterior aos campos de endereço.
func (p Professional) ServiceLocation() (float64, float64) {
	if p.AddressLatitude != 0 || p.AddressLongitude != 0 {
		return p.AddressLatitude, p.AddressLongitude
	}
	return p.Latitude, p.Longitude
}

type ProfessionCount struct {
	ProfessionName string
	Quantity       int
}

type CityProfessionalCount struct {
	CityID            uint
	CityName          string
	ProfessionalCount int64
}

type UFProfessionalCount struct {
	UFName            string
	ProfessionalCount int64
}
