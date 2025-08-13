package user

type UserRepository interface {
	FindAll() ([]*User, error)
	FindById(id uint) (*User, error)
	FindByEmail(email string) (*User, error)
	Save(user User) (*User, error)
	Remove(id uint) error
}
