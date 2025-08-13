package user

type UserService interface {
	FindAll() ([]*User, error)
	FindById(id uint) (*User, error)
	FindByEmail(email string) (*User, error)
	Save(user User) (*User, error)
	Remove(id uint) error
}

type userService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

func (s *userService) FindAll() ([]*User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) FindById(id uint) (*User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) FindByEmail(email string) (*User, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) Save(user User) (*User, error) {
	newUser, err := s.repository.Save(user)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (s *userService) Remove(id uint) error {
	return s.repository.Remove(id)
}
