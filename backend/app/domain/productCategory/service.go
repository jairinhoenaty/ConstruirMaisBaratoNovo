package productCategory

type ProductCategoryService interface {
	/*FindAll(limit, offset int) ([]*Product, int64, error)
	FindById(id uint) (*Product, error)
	Save(product Product) (*Product, error)
	FindApproved() ([]*Product, error)
	FindDayoffer() ([]*Product, error)
	*/
	//FindByProfessional(professionalId uint, limit, offset int) ([]*Product, int64, error)
	FindByProfession(professionID int) ([]*ProductCategory, error)
	Save(productCategory ProductCategory) (*ProductCategory, error)
	FindAll() ([]*ProductCategory, error)
	FindById(id uint) (*ProductCategory, error)
	FindTopLevel() ([]*ProductCategory, error)
	FindSubcategoriesByParentID(parentID uint) ([]*ProductCategory, error)
	Remove(id uint) error
}

type productCategoryService struct {
	repository ProductCategoryRepository
}

func NewProductCategoryService(repository ProductCategoryRepository) ProductCategoryService {
	return &productCategoryService{
		repository: repository,
	}
}

func (s *productCategoryService) FindByProfession(professionID int) ([]*ProductCategory, error) {
	productCategory, err := s.repository.FindByProfession(professionID)
	if err != nil {
		return nil, err
	}
	return productCategory, nil
}

func (s *productCategoryService) Save(productCategory ProductCategory) (*ProductCategory, error) {
	newproductCategory, err := s.repository.Save(productCategory)

	if err != nil {
		return nil, err
	}
	return newproductCategory, nil

}

func (s *productCategoryService) FindAll() ([]*ProductCategory, error) {
	categories, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (s *productCategoryService) FindById(id uint) (*ProductCategory, error) {
	category, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *productCategoryService) FindTopLevel() ([]*ProductCategory, error) {
	categories, err := s.repository.FindTopLevel()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (s *productCategoryService) FindSubcategoriesByParentID(parentID uint) ([]*ProductCategory, error) {
	subcategories, err := s.repository.FindSubcategoriesByParentID(parentID)
	if err != nil {
		return nil, err
	}
	return subcategories, nil
}

func (s *productCategoryService) Remove(id uint) error {
	err := s.repository.Remove(id)
	if err != nil {
		return err
	}
	return nil
}

/*
func (s *productService) FindByProfessional(professionalID uint, limit, offset int) ([]*Product, int64, error) {
	products, total, err := s.repository.FindByProfessional(professionalID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	return products, total, nil
}


func (s *productService) FindAll(limit, offset int) ([]*Product, int64, error) {
	products, total, err := s.repository.FindAll(limit, offset)
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
*/
