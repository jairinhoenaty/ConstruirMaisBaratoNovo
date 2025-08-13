package domain_test

import (
	"testing"

	pkgcontact "construir_mais_barato/app/domain/contact"
)

func TestSetPropsByContact(t *testing.T) {
	contact := pkgcontact.Contact{
		Name:      "Alfredinho da Silva",
		Telephone: "12999987878",
		Email:     "alfredo@teste.com.br",
		Message:   "Testando a entidade de contato",
	}

	nameExpected := "Alfredinho da Silva"
	telephoneExpected := "12999987878"
	emailExpected := "alfredo@teste.com.br"

	if contact.Name != nameExpected &&
		contact.Telephone != telephoneExpected &&
		contact.Email != emailExpected {
		t.Error("Dados diferente")
	}

}
