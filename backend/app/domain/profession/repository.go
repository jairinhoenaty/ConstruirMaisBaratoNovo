package profession

type ProfessionRepository interface {
	Find(quantityProfessions uint) ([]*Profession, error)
	FindAll(limit, offset int) ([]*Profession, int64, error)
	FindAllWithoutPagination() ([]*Profession, error)
	FindById(id uint) (*Profession, error)
	FindProfessionsWithCount() ([]map[string]interface{}, error)
	Save(Profession Profession) (*Profession, error)
	Remove(id uint) error
}
