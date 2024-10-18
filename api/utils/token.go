package utils

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// reference https://medium.com/@fasgolangdev/how-to-create-a-secure-authentication-api-in-golang-using-middlewares-6988632ddfd3

var secretKey = []byte("SECRET")

func GenerateToken(uid uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user_id"] = uid
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func TokenValidate(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
