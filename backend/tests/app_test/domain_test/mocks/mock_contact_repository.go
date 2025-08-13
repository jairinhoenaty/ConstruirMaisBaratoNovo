package mock_domain_test

import (
	"github.com/stretchr/testify/mock"

	pkgcontact "construir_mais_barato/app/domain/contact"
)

type MockContactRepository struct {
	mock.Mock
}

func (m *MockContactRepository) FindAll() ([]*pkgcontact.Contact, error) {
	args := m.Called()
	return args.Get(0).([]*pkgcontact.Contact), args.Error(1)
}

func (m *MockContactRepository) FindById(id uint) (*pkgcontact.Contact, error) {
	args := m.Called(id)
	return args.Get(0).(*pkgcontact.Contact), args.Error(1)
}

func (m *MockContactRepository) Save(city pkgcontact.Contact) (*pkgcontact.Contact, error) {
	args := m.Called(city)
	return args.Get(0).(*pkgcontact.Contact), args.Error(1)
}

func (m *MockContactRepository) Remove(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
