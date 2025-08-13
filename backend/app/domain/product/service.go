package product

type ProductService interface {
	FindAll(limit, offset int, professionalID int, storeID int, approved string, dayoffer string) ([]*Product, int64, error)
	FindById(id uint) (*Product, error)
	FindByCity(CityID uint) ([]*Product, error)
	Save(product Product) (*Product, error)
	FindApproved() ([]*Product, error)
	FindDayoffer() ([]*Product, error)
	//FindByProfessional(professionalId uint, limit, offset int) ([]*Product, int64, error)
	Remove(id uint) error
}

type productService struct {
	repository ProductRepository
}

func NewProductService(repository ProductRepository) ProductService {
	return &productService{
		repository: repository,
	}
}

func (s *productService) Save(product Product) (*Product, error) {
	newproduct, err := s.repository.Save(product)
	if err != nil {
		return nil, err
	}
	return newproduct, nil
}

/*func (s *productService) FindByProfessional(professionalID uint, limit, offset int) ([]*Product, int64, error) {
	products, total, err := s.repository.FindByProfessional(professionalID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	return products, total, nil
}
*/

func (s *productService) FindAll(limit, offset int, professionalID int, storeID int, approved string, dayoffer string) ([]*Product, int64, error) {
	products, total, err := s.repository.FindAll(limit, offset, professionalID, storeID, approved, dayoffer)
	if err != nil {
		return nil, 0, err
	}
	return products, total, nil
}

func (s *productService) FindById(id uint) (*Product, error) {
	product, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *productService) FindApproved() ([]*Product, error) {
	products, err := s.repository.FindApproved()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *productService) FindDayoffer() ([]*Product, error) {
	products, err := s.repository.FindDayoffer()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *productService) FindByCity(CityID uint) ([]*Product, error) {
	products, err := s.repository.FindByCity(CityID)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *productService) Remove(id uint) error {
	return s.repository.Remove(id)
}
