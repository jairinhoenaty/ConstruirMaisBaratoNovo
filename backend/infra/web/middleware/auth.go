package middleware

import (
	"net/http"
	"strings"

	pkgauthuc "construir_mais_barato/app/usecase/auth"

	"github.com/labstack/echo/v4"
)


func VerifyAndValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := getToken(c)
		
		if token != "" {
			if ok := pkgauthuc.ValidateToken(token); !ok {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token informado é inválido"})
			}
			// Se o token for válido, continue com o próximo manipulador
			return next(c)
		} else {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Token de autorização ausente ou inválido"})
		}

	}
}

func getToken(c echo.Context) string {

	tokenString := ""
	
	// Obter o token do cabeçalho de autorização
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return tokenString
	}

	// Verificar o formato do token (Bearer <token>)
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return tokenString
	}

	// Obter o token JWT
	tokenString = tokenParts[1]

	return tokenString
}
