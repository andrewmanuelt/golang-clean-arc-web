package helper

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.StandardClaims
}

func Token(secret string) string {
	claims := &Claims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))

	ShowError(err)

	return t
}

func ValidateToken(stringToken string, secret string) bool {
	token, _ := jwt.Parse(stringToken, func(stringToken *jwt.Token) (interface{}, error) {
		if _, ok := stringToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", stringToken.Header["alg"])
		}

		return []byte(secret), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	}

	return false
}
