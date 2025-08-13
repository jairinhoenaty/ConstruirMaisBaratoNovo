package domain_test

import (
	"testing"

	pkgcity "construir_mais_barato/app/domain/city"
)

func TestSetPropsByCity(t *testing.T) {
	city := pkgcity.City{Name: "Lins", UF: "SP"}
	nameExpected := "Lins"
	ufExpected := "SP"
	if city.Name != nameExpected && city.UF != ufExpected {
		t.Errorf("Expected city name to be %s, uf to be %s, got %s instead %s intead ", nameExpected, ufExpected, city.Name, city.UF)
	}
}
