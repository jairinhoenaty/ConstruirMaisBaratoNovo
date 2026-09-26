package professionCategory

type ProfessionCategoryRepository interface {
	FindAll(limit, offset int) ([]*ProfessionCategory, int64, error)
	FindActive() ([]*ProfessionCategory, error)
	FindById(id uint) (*ProfessionCategory, error)
	CountProfessions(categoryIDs []uint) (map[uint]int64, error)
	Save(category ProfessionCategory) (*ProfessionCategory, error)
	Remove(id uint) error
}
