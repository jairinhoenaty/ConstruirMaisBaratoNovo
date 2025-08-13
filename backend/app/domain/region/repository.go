package region

type RegionRepository interface {
	Find(quantityRegions uint) ([]*Region, error)
	FindAll(limit, offset int, uf string) ([]*Region, int64, error)
	FindAllWithoutPagination() ([]*Region, error)
	FindByCity(cityId uint) (*Region, error)
	FindById(id uint) (*Region, error)
	FindRegionsWithCount() ([]map[string]interface{}, error)
	Save(Region Region) (*Region, error)
	Remove(id uint) error
}
