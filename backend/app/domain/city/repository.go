package city

type CityRepository interface {
	FindAll() ([]*City, error)
	FindById(id uint) (*City, error)
	FindByUF(uf string) ([]*City, error)
	Save(City City) (*City, error)
	Remove(id uint) error
}
