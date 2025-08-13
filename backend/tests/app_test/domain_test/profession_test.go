package domain_test

import (
	"testing"

	pkgprofession "construir_mais_barato/app/domain/profession"
)

func TestSetPropsByProfession(t *testing.T) {
	profession := pkgprofession.Profession{Name: "Arquiteto", Description: "Descrição da profissão de arquiteto", Icon: "icon.png"}
	nameExpected := "Arquiteto"
	descriptionExpected := "Descrição da profissão de arquiteto"
	iconExpected := "icon.png"
	if profession.Name != nameExpected &&
		profession.Description != descriptionExpected &&
		profession.Icon != iconExpected {
		t.Error("Dados diferente")
	}
}
