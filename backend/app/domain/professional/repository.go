package professional

type ProfessionalRepository interface {
	FindAll(limit, offset int, filter string, uf string, professionalID int, order string) ([]*Professional, int64, error)
	FindById(id uint) (*Professional, error)
	FindByEmail(email string) (*Professional, error)
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
}
