package mocks_test

import (
	pkgcityuc "construir_mais_barato/app/usecase/city"

	"github.com/stretchr/testify/mock"
)

// Mock dos casos de uso
type MockFindAllCityUC struct {
	mock.Mock
}

func (m *MockFindAllCityUC) Execute() (*[]pkgcityuc.CityPresenter, error) {
	args := m.Called()
	return args.Get(0).(*[]pkgcityuc.CityPresenter), args.Error(1)
}

type MockFindByIdUC struct {
	mock.Mock
	ID *uint
}

func (m *MockFindByIdUC) Execute() (*pkgcityuc.CityPresenter, error) {
	args := m.Called(m.ID)
	return args.Get(0).(*pkgcityuc.CityPresenter), args.Error(1)
}

type MockSaveCityUC struct {
	mock.Mock
}

func (m *MockSaveCityUC) Execute() (*pkgcityuc.CityPresenter, error) {
	args := m.Called()
	return args.Get(0).(*pkgcityuc.CityPresenter), args.Error(1)
}

type MockDeleteCityUC struct {
	mock.Mock
	ID *uint
}

func (m *MockDeleteCityUC) Execute() error {
	args := m.Called(m.ID)
	return args.Error(0)
}
