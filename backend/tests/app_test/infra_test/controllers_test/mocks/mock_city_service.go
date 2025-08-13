package mocks_test

import (
	pkgcity "construir_mais_barato/app/domain/city"

	"github.com/stretchr/testify/mock"
)

// Mock do serviço CityService
type MockCityService struct {
	mock.Mock
}

func (m *MockCityService) FindAll() ([]*pkgcity.City, error) {
	args := m.Called()
	return args.Get(0).([]*pkgcity.City), args.Error(1)
}

func (m *MockCityService) FindById(id uint) (*pkgcity.City, error) {
	args := m.Called(id)
	return args.Get(0).(*pkgcity.City), args.Error(1)
}

func (m *MockCityService) FindByUF(uf string) ([]*pkgcity.City, error) {
	args := m.Called(uf)
	return args.Get(0).([]*pkgcity.City), args.Error(1)
}

func (m *MockCityService) Save(city pkgcity.City) (*pkgcity.City, error) {
	args := m.Called(city)
	//fmt.Println(args...)
	return args.Get(0).(*pkgcity.City), args.Error(1)
}

func (m *MockCityService) Remove(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
