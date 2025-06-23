package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecretKey = []byte("your-secret-key") // replace with env var in production

type JWTClaims struct {
	UserID int64 `json:"user_id"`
	Role   UserModel `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID int64, role UserModel) (string, error) {
	claims := JWTClaims{
		UserID: userID,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}
