package auth

import (
	"time"

	"github.com/ByChanderZap/api-basics/config"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(secret string, userID string) (string, error) {
	expiration := time.Hour * time.Duration(config.Envs.JWTExpirationInHours)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
