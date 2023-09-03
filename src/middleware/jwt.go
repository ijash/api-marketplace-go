package middleware

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type jwtCustomClaims struct {
	Name string `json:"name"`
	Id   string `json:"id"`
	jwt.RegisteredClaims
}

func GenerateJwt(name string, id string) (string, error) {
	log.Println("GenerateJwt accessed")
	claims := &jwtCustomClaims{
		name,
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("PRIVATE_KEY_JWT")))
	if err != nil {
		return "", err
	}
	return t, nil
}

func ExtractUserIDFromJWT(tokenString string) (string, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &jwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("PRIVATE_KEY_JWT")), nil
	})

	if err != nil {
		log.Printf("Error parsing JWT: %v", err)
		return "", err
	}

	if !token.Valid {
		log.Printf("Invalid JWT token")
		return "", fmt.Errorf("invalid token")
	}

	// Type-assert to access custom claims
	claims, ok := token.Claims.(*jwtCustomClaims)
	if !ok {
		log.Printf("Invalid custom claims in JWT")
		return "", fmt.Errorf("invalid claims")
	}

	return claims.Id, nil
}
