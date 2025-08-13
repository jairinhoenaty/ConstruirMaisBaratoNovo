package infra_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	pkgcity "construir_mais_barato/app/domain/city"
	pkgcityuc "construir_mais_barato/app/usecase/city"
	pkgcitycontroller "construir_mais_barato/infra/web/controllers"
	pkgmockcontroller "construir_mais_barato/tests/app_test/infra_test/controllers_test/mocks"
)

func TestCityController_FindAll(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/cities", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	//criado o mock do serviço
	mockService := new(pkgmockcontroller.MockCityService)
	mockService.On("FindAll").Return([]*pkgcity.City{
		{Model: gorm.Model{ID: 1}, Name: "Lins", UF: "SP"},
	}, nil)

	// criado o mock do caso de uso
	findAllCityUC := pkgcityuc.NewFindAllCityUC(pkgcityuc.FindAllCityUCParams{
		Service: mockService,
	})

	params := pkgcitycontroller.CityControllerParams{
		FindAllCityUCParams: pkgcityuc.FindAllCityUCParams{
			Service: findAllCityUC.Service,
		},
	}

	controller := pkgcitycontroller.CityController{
		FindAllCityUCParams: params.FindAllCityUCParams,
	}

	if assert.NoError(t, controller.FindAll(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Contains(t, res.Body.String(), "Lins")
	}

}

func TestCityController_FindById(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/city/1", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	//setando os parametros na requisição
	c.SetPath("/city/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	//criado o mock do serviço
	mockService := new(pkgmockcontroller.MockCityService)
	mockService.On("FindById", uint(1)).Return(&pkgcity.City{Model: gorm.Model{ID: 1}, Name: "Lins", UF: "SP"}, nil)

	// criado o mock do caso de uso
	findByIdCityUC := pkgcityuc.NewFindByIdUC(pkgcityuc.FindByIdUCParams{
		Service: mockService,
	})

	params := pkgcitycontroller.CityControllerParams{
		FindByIdUCParams: pkgcityuc.FindByIdUCParams{
			Service: findByIdCityUC.Service,
		},
	}

	controller := pkgcitycontroller.CityController{
		FindByIdUCParams: params.FindByIdUCParams,
	}

	if assert.NoError(t, controller.FindById(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Contains(t, res.Body.String(), "Lins")
	}
}

func TestCityController_Save(t *testing.T) {
	e := echo.New()
	cityJSON := `{"name":"Cafelandia", "uf":"SP"}`
	req := httptest.NewRequest(http.MethodPost, "/city", strings.NewReader(cityJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	mockService := new(pkgmockcontroller.MockCityService)
	expectedCity := &pkgcity.City{
		Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Name:  "Cafelandia",
		UF:    "SP",
	}

	mockService.On("Save", mock.MatchedBy(func(city pkgcity.City) bool {
		return city.Name == "Cafelandia" && city.UF == "SP"
	})).Return(expectedCity, nil)

	saveCityUC := pkgcityuc.NewSaveCityUC(pkgcityuc.SaveCityUCParams{Service: mockService})

	params := pkgcitycontroller.CityControllerParams{
		SaveCityUCParams: pkgcityuc.SaveCityUCParams{
			Service: saveCityUC.Service,
		},
	}

	controller := pkgcitycontroller.CityController{
		SaveCityUCParams: params.SaveCityUCParams,
	}

	if assert.NoError(t, controller.Save(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Contains(t, res.Body.String(), "Cafelandia")
	}

}

func TestCityController_Delete(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/city/1", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	//setando os parametros na requisição
	c.SetPath("/city/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	//criado o mock do serviço
	mockService := new(pkgmockcontroller.MockCityService)
	mockService.On("Remove", uint(1)).Return(nil)

	deleteCityUC := pkgcityuc.NewDeleteCityUC(pkgcityuc.DeleteCityUCParams{Service: mockService})

	params := pkgcitycontroller.CityControllerParams{
		DeleteCityUCParams: pkgcityuc.DeleteCityUCParams{
			Service: deleteCityUC.Service,
		},
	}

	controller := pkgcitycontroller.CityController{
		DeleteCityUCParams: params.DeleteCityUCParams,
	}

	if assert.NoError(t, controller.Delete(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
	}

}
