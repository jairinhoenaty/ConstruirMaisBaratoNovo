package store

type StoreRepository interface {
	FindAll(limit, offset int) ([]*Store, int64, error)
	FindById(id uint) (*Store, error)
	FindByEmail(email string) (*Store, error)
	FindByName(name string) ([]*Store, error)
	/*
		FindByCityAndProfession(cityID, professionID uint, limit, offset int) ([]*Professional, int64, error)
		FindByNameAndCityAndProfession(name string, cityID, professionID uint, limit, offset int) ([]*Professional, error)
		CountProfessionalsByProfession() ([]ProfessionCount, error)
		CountProfessionalsByState(uf string, limit, offset int) ([]CityProfessionalCount, *int64, error)
		CountProfessionalsByProfessionInCity(cityID uint) ([]ProfessionCount, error)
	*/
	FindLastStores(quantityRecords int) ([]Store, error)
	FindByCategoryAndSubCategory(categoryID int, subCategories []int) ([]*Store, error)
	CountByCategory(categoryID uint) (int64, error)
	CountBySubCategory(subCategoryID uint) (int64, error)
	MigrateCategoryBulk(fromCategoryID, toCategoryID uint) (int64, error)
	MigrateSubCategoryBulk(fromSubCategoryID, toSubCategoryID uint) (int64, error)
	Save(store Store) (*Store, error)
	Remove(id uint) error
	ExportXLSX() ([]*Store, error)
}
