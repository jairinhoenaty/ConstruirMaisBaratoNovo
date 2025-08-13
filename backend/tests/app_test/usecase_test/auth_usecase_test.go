package usecase_teste

import (
	"fmt"
	"testing"

	pkgauthuc "construir_mais_barato/app/usecase/auth"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {

	userPresenter := pkgauthuc.UserPresenter{
		ID:   1,
		Name: "Mariazinha",
	}

	token, err := pkgauthuc.GenerateToken(userPresenter)

	fmt.Println("Token => ", token)

	assert.NoError(t, err)
	assert.NotNil(t, token)

}

func TestValidateToken(t *testing.T) {
	userPresenter := pkgauthuc.UserPresenter{
		ID:   1,
		Name: "Joaozinho",
	}

	// Generate a valid token
	validToken, err := pkgauthuc.GenerateToken(userPresenter)
	assert.NoError(t, err)
	assert.NotNil(t, validToken)

	// test valid Token
	isValidToken := pkgauthuc.ValidateToken(validToken)
	assert.NotNil(t, isValidToken, "Token should not be nil")

	// test invalid token
	invalidToken := "invalid.token.string"
	resultInvalidToken := pkgauthuc.ValidateToken(invalidToken)
	assert.False(t, resultInvalidToken)

}

func TestGenerateHashPassword(t *testing.T) {
	senha, err := pkgauthuc.GenerateHashPassword("123")
	assert.NoError(t, err)
	//fmt.Println(senha)
	assert.NotEmpty(t, senha)
}

func TestCompareValidPassaword(t *testing.T) {

	password := "123"

	senha, err := pkgauthuc.GenerateHashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, senha)

	ok := pkgauthuc.ComparePassword(password, senha)
	assert.True(t, ok)
}

func TestCompareInvalidPassaword(t *testing.T) {

	errorPassword := "1234"
	senha, err := pkgauthuc.GenerateHashPassword(errorPassword)
	assert.NoError(t, err)
	assert.NotEmpty(t, senha)

	password := "123"
	ok := pkgauthuc.ComparePassword(password, senha)
	assert.False(t, ok)
}
