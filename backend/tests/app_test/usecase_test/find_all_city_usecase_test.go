package usecase_teste

import (
	"fmt"
	"testing"

	pkgcity "construir_mais_barato/app/domain/city"
	pkgcityuc "construir_mais_barato/app/usecase/city"
	pkgcityrepository "construir_mais_barato/infra/database/repositories/city"

	"github.com/stretchr/testify/assert"
)

func TestFindAllCitiesUC_Execute(t *testing.T) {

	//configurar o banco de dados
	db := setupTestDB()
	defer db.Migrator().DropTable(&pkgcity.City{})

	// iniciar o serviço do caso de uso
	cityService := pkgcity.NewCityService(pkgcityrepository.NewCityRepositoryImpl(db))

	// salvar varias as cidades
	for i := 0; i < 5; i++ {
		numString := fmt.Sprint(i)
		mockcity := &pkgcity.City{
			Name: "City A " + numString,
			UF:   "UF " + numString,
		}
		savedCity, err := cityService.Save(*mockcity)
		assert.NoError(t, err)
		assert.Equal(t, savedCity.Name, mockcity.Name)
		assert.Equal(t, savedCity.UF, mockcity.UF)
	}

	//criar o caso de uso e buscar os resultados
	uc := pkgcityuc.NewFindAllCityUC(pkgcityuc.FindAllCityUCParams{
		Service: cityService,
	})

	presenters, err := uc.Execute()
	assert.NoError(t, err)
	assert.NotNil(t, presenters)

}
