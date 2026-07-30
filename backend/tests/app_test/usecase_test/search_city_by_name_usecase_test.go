package usecase_teste

import (
	"testing"

	pkgcity "construir_mais_barato/app/domain/city"
	pkgcityuc "construir_mais_barato/app/usecase/city"
	pkgcityrepository "construir_mais_barato/infra/database/repositories/city"

	"github.com/stretchr/testify/assert"
)

func buildSearchCityUC(t *testing.T, cities []pkgcity.City) pkgcityuc.SearchCityByNameUC {
	t.Helper()

	db := setupTestDB()
	t.Cleanup(func() {
		db.Migrator().DropTable(&pkgcity.City{})
	})

	cityService := pkgcity.NewCityService(pkgcityrepository.NewCityRepositoryImpl(db))
	for _, city := range cities {
		_, err := cityService.Save(city)
		assert.NoError(t, err)
	}

	return pkgcityuc.NewSearchCityByNameUC(pkgcityuc.SearchCityByNameUCParams{
		Service: cityService,
	})
}

func TestSearchCityByNameUC_PrioritizaNomesQueComecamComOTermo(t *testing.T) {
	uc := buildSearchCityUC(t, []pkgcity.City{
		{Name: "Ribeirao Bonito", UF: "SP"},
		{Name: "Bonito", UF: "MS"},
	})

	uc.Assembler = &pkgcityuc.SearchCityAssembler{Term: "bonito"}
	presenters, err := uc.Execute()

	assert.NoError(t, err)
	assert.Len(t, presenters, 2)
	// "Bonito" comeca com o termo, entao vem antes de "Ribeirao Bonito"
	assert.Equal(t, "Bonito", presenters[0].Name)
	assert.Equal(t, "MS", presenters[0].UF)
	assert.Equal(t, "Ribeirao Bonito", presenters[1].Name)
}

func TestSearchCityByNameUC_RespeitaOLimite(t *testing.T) {
	uc := buildSearchCityUC(t, []pkgcity.City{
		{Name: "Santa Barbara", UF: "MG"},
		{Name: "Santa Clara", UF: "RJ"},
		{Name: "Santa Rita", UF: "PB"},
	})

	uc.Assembler = &pkgcityuc.SearchCityAssembler{Term: "santa", Limit: 2}
	presenters, err := uc.Execute()

	assert.NoError(t, err)
	assert.Len(t, presenters, 2)
}

func TestSearchCityByNameUC_TermoCurtoRetornaListaVazia(t *testing.T) {
	uc := buildSearchCityUC(t, []pkgcity.City{
		{Name: "Salvador", UF: "BA"},
	})

	uc.Assembler = &pkgcityuc.SearchCityAssembler{Term: "s"}
	presenters, err := uc.Execute()

	assert.NoError(t, err)
	assert.Empty(t, presenters)
}

func TestSearchCityByNameUC_TermoApenasComCuringasNaoRetornaTudo(t *testing.T) {
	uc := buildSearchCityUC(t, []pkgcity.City{
		{Name: "Salvador", UF: "BA"},
		{Name: "Recife", UF: "PE"},
	})

	// Sem sanitizacao, "%%" casaria com todos os municipios da base.
	uc.Assembler = &pkgcityuc.SearchCityAssembler{Term: "%%"}
	presenters, err := uc.Execute()

	assert.NoError(t, err)
	assert.Empty(t, presenters)
}

func TestSearchCityByNameUC_SemAssemblerRetornaErro(t *testing.T) {
	uc := buildSearchCityUC(t, nil)

	presenters, err := uc.Execute()

	assert.Error(t, err)
	assert.Nil(t, presenters)
}
