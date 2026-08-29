package professional

import "time"

type ProfessionalRepository interface {
	// SetPremium liga/desliga o premium de um profissional sem tocar no resto
	// do cadastro. expiresAt nulo deixa a vigência em aberto.
	SetPremium(id uint, isPremium bool, expiresAt *time.Time) error
	// ExpirePremiums rebaixa profissionais cuja vigência terminou e devolve
	// quantos foram afetados. Ignora premium sem data (ativação manual).
	ExpirePremiums(now time.Time) (int64, error)
	FindAll(limit, offset int, filter string, uf string, professionalID int, order string) ([]*Professional, int64, error)
	FindById(id uint) (*Professional, error)
	FindByEmail(email string) (*Professional, error)
	FindByTelephone(telephone string) (*Professional, error)
	FindByName(name string) ([]*Professional, error)
	FindByCityAndProfession(cityID, professionID uint, limit, offset int) ([]*Professional, int64, error)
	FindByProfessionAndLocation(professionID uint, latitude float32, longitude float32, distance, limit, offset int) ([]*Professional, int64, error)
	FindByNameAndCityAndProfession(name string, cityID, professionID uint, limit, offset int) ([]*Professional, error)
	CountProfessionalsByProfession() ([]ProfessionCount, error)
	CountCityProfessionalsByState(uf string, limit, offset int) ([]CityProfessionalCount, *int64, error)
	CountProfessionalsByState(uf string, limit, offset int) ([]UFProfessionalCount, *int64, error)
	CountProfessionalsByProfessionInCity(cityID uint) ([]ProfessionCount, error)
	FindLastProfessionals(quantityRecords int) ([]Professional, error)
	Save(professional Professional) (*Professional, error)
	Remove(id uint) error
	ExportXLSX() ([]*Professional, error)
	// Retorna profissionais aleatórios por nome exato da profissão (ex.: "pedreiro")
	FindRandom(
        professionID *uint,
        professionName *string,
        verified *bool,
        online *bool,
        seed *int64,
        limit, offset int,
    ) ([]*Professional, int64, error)
}
