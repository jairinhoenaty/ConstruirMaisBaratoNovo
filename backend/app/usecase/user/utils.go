package user_usecase

import (
	pkguser "construir_mais_barato/app/domain/user"
	"encoding/base64"
)

func GenerateUser(assembler *UserAssembler) pkguser.User {
	user := pkguser.User{}
	if assembler != nil {
		user.ID = assembler.ID
		user.Name = assembler.Name
		user.Email = assembler.Email
		user.Profile = assembler.Profile
		user.Password = assembler.Password
		user.GoogleToken = assembler.GoogleToken

	}
	return user
}

func GenerateUserPresenter(user *pkguser.User) UserPresenter {
	presenter := UserPresenter{}
	if user != nil {
		presenter.ID = user.ID
		presenter.Name = user.Name
		presenter.Email = user.Email
		presenter.Profile = user.Profile
		presenter.Password = user.Password
		presenter.GoogleToken = user.GoogleToken
	}
	return presenter
}

// EncryptValue codifica uma string em base64
func EncryptValue(value string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	return encoded
}

// DecryptValue decodifica uma string base64
func DecryptValue(value string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}
