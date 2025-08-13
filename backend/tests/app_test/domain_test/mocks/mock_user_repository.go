package mock_domain_test

import (
	"github.com/stretchr/testify/mock"

	pkguser "construir_mais_barato/app/domain/user"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindAll() ([]*pkguser.User, error) {
	args := m.Called()
	return args.Get(0).([]*pkguser.User), args.Error(1)
}

func (m *MockUserRepository) FindById(id uint) (*pkguser.User, error) {
	args := m.Called(id)
	return args.Get(0).(*pkguser.User), args.Error(1)
}

func (m *MockUserRepository) Save(city pkguser.User) (*pkguser.User, error) {
	args := m.Called(city)
	return args.Get(0).(*pkguser.User), args.Error(1)
}

func (m *MockUserRepository) Remove(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
