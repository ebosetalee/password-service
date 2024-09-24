package auth

import (
	"fmt"
	"time"

	"github.com/ebosetalee/password-service.git/config"
	"github.com/golang-jwt/jwt/v5"
)

var env = config.Env

func GenerateJWT(userID string, email string) (string, error) {

	expiration := time.Second * time.Duration(env.JWTExpiration)
	secret := []byte(env.JWTSecret)

	fmt.Println("jwt-exp", expiration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    userID,
		"email":     email,
		"expiresAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func verifyJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(env.JWTSecret), nil
	})
}
