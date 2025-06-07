package jwt

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	log.Println([]byte(os.Getenv("JWT_SECRET")))
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}