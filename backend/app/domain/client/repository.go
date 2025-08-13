package client

type ClientRepository interface {
	FindAll(limit, offset int) ([]*Client, int64, error)
	FindById(id uint) (*Client, error)
	FindByEmail(email string) (*Client, error)
	FindByName(name string) ([]*Client, error)
	/*
		FindByCityAndProfession(cityID, professionID uint, limit, offset int) ([]*Professional, int64, error)
		FindByNameAndCityAndProfession(name string, cityID, professionID uint, limit, offset int) ([]*Professional, error)
		CountProfessionalsByProfession() ([]ProfessionCount, error)
		CountProfessionalsByState(uf string, limit, offset int) ([]CityProfessionalCount, *int64, error)
		CountProfessionalsByProfessionInCity(cityID uint) ([]ProfessionCount, error)
	*/
	FindLastClients(quantityRecords int) ([]Client, error)

	Save(client Client) (*Client, error)

	Remove(id uint) error
	ExportXLSX() ([]*Client, error)
}
