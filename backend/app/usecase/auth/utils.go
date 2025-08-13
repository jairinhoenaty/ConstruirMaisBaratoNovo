package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("asdlkjqwoeiurtghrtghlkp")

func GenerateToken(user UserPresenter) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"name": user.Name,
		"exp":  time.Now().Add(time.Hour * 4).Unix(), // token terá validade por 4 horas
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func ValidateToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return false
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	} else {
		return false
	}
}

func GenerateAuthenticatePresenter(token string, isLoged bool, data UserPresenter) AuthenticatePresenter {
	dataProfile := ""
	if data.Profile != nil {
		dataProfile = *data.Profile
	}

	return AuthenticatePresenter{
		Token:   token,
		IsLoged: isLoged,
		User: UserPresenter{
			ID:          data.ID,
			Name:        data.Name,
			Profile:     &dataProfile,
			Email:       data.Email,
			GoogleToken: data.GoogleToken,
		},
	}
}

func GenerateUserPresenter(id uint, name, profile string, email string, google_token string) UserPresenter {
	return UserPresenter{
		ID:          id,
		Name:        name,
		Profile:     &profile,
		Email:       email,
		GoogleToken: google_token,
	}
}

func GenerateHashPassword(senha string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CompararSenha compara a senha informada com a senha criptografada
func ComparePassword(senhaInformada, senhaCriptografada string) bool {
	/*fmt.Println("Senha==========");
	fmt.Println(senhaInformada);
	fmt.Println("Senha Cript==========");
	fmt.Println(senhaCriptografada);
	fmt.Println("==========");
	fmt.Println(GenerateHashPassword("123456"));
	fmt.Println("==========");*/
	err := bcrypt.CompareHashAndPassword([]byte(senhaCriptografada), []byte(senhaInformada))
	return err == nil
}
