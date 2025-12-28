package services

import (
	"belajargolang/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = utils.GetJWTSecret()

func GenerateToken(userID int, email string) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // expired 1 hari
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}
