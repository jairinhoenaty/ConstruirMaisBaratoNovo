package product

type ProductRepository interface {
	FindAll(limit, offset int, professionalID int, storeID int, approved string, dayoffer string) ([]*Product, int64, error)
	FindById(id uint) (*Product, error)
	FindByCity(CityID uint) ([]*Product, error)
	FindApproved() ([]*Product, error)
	FindDayoffer() ([]*Product, error)
	Save(product Product) (*Product, error)
	//FindByProfessional(professionalId uint, limit, offset int) ([]*Product, int64, error)
	Remove(id uint) error
}
