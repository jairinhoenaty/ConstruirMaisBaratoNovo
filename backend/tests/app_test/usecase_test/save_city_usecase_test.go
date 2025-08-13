package usecase_teste

import (
	"testing"

	"github.com/stretchr/testify/assert"

	pkgcity "construir_mais_barato/app/domain/city"
	pkgcityuc "construir_mais_barato/app/usecase/city"
	pkgcityrepository "construir_mais_barato/infra/database/repositories/city"
)

func TestSaveCityUC_Execute(t *testing.T) {
	//configurar banco de dados
	db := setupTestDB()
	defer db.Migrator().DropTable(&pkgcity.City{})

	//iniciar o serviço do caso de uso
	cityService := pkgcity.NewCityService(pkgcityrepository.NewCityRepositoryImpl(db))
	saveCityUC := pkgcityuc.NewSaveCityUC(pkgcityuc.SaveCityUCParams{
		Service: cityService,
	})

	//assembler para o caso de uso
	cityAssembler := pkgcityuc.CityAssembler{
		Name: "Lins",
		UF:   "SP",
	}

	//adicona o assembler no caso de uso
	saveCityUC.Assembler = &cityAssembler

	cityPresenter, err := saveCityUC.Execute()
	assert.NoError(t, err)
	assert.NotNil(t, cityPresenter)
	assert.Equal(t, cityAssembler.Name, cityPresenter.Name)
	assert.Equal(t, cityAssembler.UF, cityPresenter.UF)

}
