package domain_test

import (
	"testing"

	pkgcity "construir_mais_barato/app/domain/city"
	pkgmock "construir_mais_barato/tests/app_test/domain_test/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCityService(t *testing.T) {

	mockRepo := new(pkgmock.MockCityRepository)
	cityService := pkgcity.NewCityService(mockRepo)

	t.Run("Test method findAll", func(t *testing.T) {
		mockCities := []*pkgcity.City{
			{Name: "Lins", UF: "SP"},
			{Name: "Manaus", UF: "AM"},
		}

		mockRepo.On("FindAll").Return(mockCities, nil)

		cities, err := cityService.FindAll()

		assert.NoError(t, err)
		assert.Equal(t, mockCities, cities)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Test method findById", func(t *testing.T) {
		mockCity := &pkgcity.City{
			Name: "Lins", UF: "SP",
		}

		mockRepo.On("FindById", uint(1)).Return(mockCity, nil)

		city, err := cityService.FindById(1)

		assert.NoError(t, err)
		assert.Equal(t, mockCity, city)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Test method findByUF", func(t *testing.T) {
		mockCities := []*pkgcity.City{
			{Name: "Lins", UF: "SP"},
			{Name: "Manaus", UF: "AM"},
		}

		mockRepo.On("FindById", uint(1)).Return(mockCities, nil)

		city, err := cityService.FindByUF("SP")

		assert.NoError(t, err)
		assert.Equal(t, mockCities, city)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Test method Save (insert)", func(t *testing.T) {
		newCity := pkgcity.City{Name: "Bauru", UF: "SP"}
		savedCity := &pkgcity.City{Name: "Cafelandia", UF: "SP"}

		mockRepo.On("Save", newCity).Return(savedCity, nil)

		city, err := cityService.Save(newCity)

		assert.NoError(t, err)
		assert.Equal(t, savedCity, city)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Test method Remove", func(t *testing.T) {
		mockRepo.On("Remove", uint(1)).Return(nil)

		err := cityService.Remove(1)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

}
