package domain_test

import (
	"testing"

	pkguser "construir_mais_barato/app/domain/user"
)

func TestSetPropsByUser(t *testing.T) {
	user := pkguser.User{Name: "Usuario 1", Email: "teste@teste.com.br", Password: "123", Profile: "adm"}
	nameExpected := "Usuario 1"
	emailExpected := "teste@teste.com.br"
	passwordExpected := "123"
	profileExpected := "adm"
	if user.Name != nameExpected &&
		user.Email != emailExpected &&
		user.Password != passwordExpected &&
		user.Profile != profileExpected {
		t.Error("Dados diferente")
	}
}
