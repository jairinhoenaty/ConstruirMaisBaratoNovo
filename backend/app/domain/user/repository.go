package user

type IDType string

const (
	IDTypeUser         IDType = "user"
	IDTypeProfessional IDType = "professional"
)

type UserRepository interface {
	FindAll() ([]*User, error)
	FindById(id uint) (*User, error)
	FindByEmail(email string) (*User, error)
	FindTokensByIds(ids []uint, idType IDType) ([]string, error)
	Save(user User) (*User, error)
	Remove(id uint) error
}
