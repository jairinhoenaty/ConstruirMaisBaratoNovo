package mock_domain_test

import (
	"github.com/stretchr/testify/mock"

	pkgprofession "construir_mais_barato/app/domain/profession"
)

type MockProfessionRepository struct {
	mock.Mock
}

func (m *MockProfessionRepository) FindProfessionsWithCount() ([]map[string]interface{}, error) {
	// Argumentos passados ​​para a função MockProfessionRepository.Called()
	args := m.Called()

	// Verifica se o primeiro argumento é um slice de mapas
	if args.Get(0) == nil {
		return nil, args.Error(1) // Retorna um erro, pois não há dados para retornar
	}

	// Converte e retorna o primeiro argumento como []map[string]interface{}
	return args.Get(0).([]map[string]interface{}), args.Error(1)
}

func (m *MockProfessionRepository) Find(quantityProfessions uint) ([]*pkgprofession.Profession, error) {
	args := m.Called()
	return args.Get(0).([]*pkgprofession.Profession), args.Error(1)
}

func (m *MockProfessionRepository) FindAll() ([]*pkgprofession.Profession, error) {
	args := m.Called()
	return args.Get(0).([]*pkgprofession.Profession), args.Error(1)
}

func (m *MockProfessionRepository) FindById(id uint) (*pkgprofession.Profession, error) {
	args := m.Called(id)
	return args.Get(0).(*pkgprofession.Profession), args.Error(1)
}

func (m *MockProfessionRepository) Save(city pkgprofession.Profession) (*pkgprofession.Profession, error) {
	args := m.Called(city)
	return args.Get(0).(*pkgprofession.Profession), args.Error(1)
}

func (m *MockProfessionRepository) Remove(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
