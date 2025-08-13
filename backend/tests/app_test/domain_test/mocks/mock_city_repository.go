package mock_domain_test

import (
	"github.com/stretchr/testify/mock"

	pkgcity "construir_mais_barato/app/domain/city"
)

type MockCityRepository struct {
	mock.Mock
}

func (m *MockCityRepository) FindAll() ([]*pkgcity.City, error) {
	args := m.Called()
	return args.Get(0).([]*pkgcity.City), args.Error(1)
}

func (m *MockCityRepository) FindById(id uint) (*pkgcity.City, error) {
	args := m.Called(id)
	return args.Get(0).(*pkgcity.City), args.Error(1)
}

func (m *MockCityRepository) FindByUF(uf string) ([]*pkgcity.City, error) {
	args := m.Called(uf)
	return args.Get(0).([]*pkgcity.City), args.Error(1)
}

func (m *MockCityRepository) Save(city pkgcity.City) (*pkgcity.City, error) {
	args := m.Called(city)
	return args.Get(0).(*pkgcity.City), args.Error(1)
}

func (m *MockCityRepository) Remove(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
