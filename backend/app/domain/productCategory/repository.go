package productCategory

type ProductCategoryRepository interface {
	//	FindAll(limit, offset int) ([]*Product, int64, error)
	FindByProfession(professionID int) ([]*ProductCategory, error)
	//FindApproved() ([]*Product, error)
	//FindDayoffer() ([]*Product, error)
	Save(productCategory ProductCategory) (*ProductCategory, error)
	//FindByProfessional(professionalId uint, limit, offset int) ([]*Product, int64, error)
	FindAll() ([]*ProductCategory, error)
	FindById(id uint) (*ProductCategory, error)
	FindTopLevel() ([]*ProductCategory, error)
	FindSubcategoriesByParentID(parentID uint) ([]*ProductCategory, error)
	Remove(id uint) error
}
