package pkg

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID string, rol string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = userID
	claims["rol"] = rol
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expira en 24 horas

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
