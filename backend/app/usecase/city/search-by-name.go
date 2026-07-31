package city_usecase

import (
	pkgcity "construir_mais_barato/app/domain/city"
	"fmt"
	"strings"
)

const (
	// minSearchTermLength evita que uma busca com 1 caractere retorne
	// praticamente toda a base de municipios.
	minSearchTermLength = 2
	// defaultSearchLimit e maxSearchLimit delimitam o tamanho da lista de
	// sugestoes devolvida ao autocomplete.
	defaultSearchLimit = 8
	maxSearchLimit     = 50
)

type SearchCityByNameUC struct {
	Service   pkgcity.CityService
	Assembler *SearchCityAssembler
}

type SearchCityByNameUCParams struct {
	Service pkgcity.CityService
}

func NewSearchCityByNameUC(params SearchCityByNameUCParams) SearchCityByNameUC {
	return SearchCityByNameUC{
		Service: params.Service,
	}
}

func (uc *SearchCityByNameUC) Execute() ([]*CityPresenter, error) {
	if uc.Assembler == nil {
		return nil, fmt.Errorf("invalid search term")
	}

	// Termo curto nao e erro: o autocomplete consulta a cada tecla digitada,
	// entao devolvemos lista vazia ate haver caracteres suficientes.
	term := strings.TrimSpace(uc.Assembler.Term)
	if len([]rune(term)) < minSearchTermLength {
		return []*CityPresenter{}, nil
	}

	limit := uc.Assembler.Limit
	if limit <= 0 {
		limit = defaultSearchLimit
	}
	if limit > maxSearchLimit {
		limit = maxSearchLimit
	}

	cities, err := uc.Service.SearchByName(term, limit)
	if err != nil {
		return nil, err
	}

	citiesPresenter := make([]*CityPresenter, 0)
	for _, city := range cities {
		cityPresenter := GenerateCityPresenter(city)

		citiesPresenter = append(citiesPresenter, &cityPresenter)
	}

	return citiesPresenter, nil
}
