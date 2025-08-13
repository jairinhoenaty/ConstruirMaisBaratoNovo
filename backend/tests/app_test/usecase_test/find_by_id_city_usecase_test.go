package usecase_teste

import (
	"testing"

	pkgcity "construir_mais_barato/app/domain/city"
	pkgcityuc "construir_mais_barato/app/usecase/city"
	pkgcityrepository "construir_mais_barato/infra/database/repositories/city"

	"github.com/stretchr/testify/assert"
)

func TestFindByIdUC(t *testing.T) {
	//configurar o banco de dados
	db := setupTestDB()
	defer db.Migrator().DropTable(&pkgcity.City{})

	// iniciar o serviço do caso de uso
	cityService := pkgcity.NewCityService(pkgcityrepository.NewCityRepositoryImpl(db))

	// cria uma cidade para ser cadastrada
	city := pkgcity.City{
		Name: "City A ",
		UF:   "UF ",
	}
	savedCity, err := cityService.Save(city)
	assert.NoError(t, err)

	//criar o caso de uso e buscar os resultados
	uc := pkgcityuc.NewFindByIdUC(pkgcityuc.FindByIdUCParams{
		Service: cityService,
	})

	//adicionar o id no caso de uso
	idTesting := uint(1)
	uc.ID = &idTesting

	presenter, err := uc.Execute()
	assert.NoError(t, err)
	assert.NotNil(t, presenter)
	assert.Equal(t, savedCity.ID, presenter.ID)
	assert.Equal(t, savedCity.Name, presenter.Name)
	assert.Equal(t, savedCity.UF, presenter.UF)

}
