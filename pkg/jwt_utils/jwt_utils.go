package jwt_utils

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"time"
)

var SecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

const expirationTime = 1000 * 1000 * 60 * 10

func GenerateAccessToken(username string) (string, error) {
	expiration := time.Now().Add(expirationTime).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      expiration,
	})

	return token.SignedString(SecretKey)
}

func ValidateToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return SecretKey, nil
	})

	if err != nil || !token.Valid {
		log.Printf("[CustomJWT][ValidateToken] Token is invalid: %v]", err)
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
		if time.Now().After(expirationTime) {
			return false
		}
		return true
	}

	return false
}

func ExtractUsernameFromToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return SecretKey, nil
	})

	if err != nil || !token.Valid {
		log.Printf("[CustomJWT][ValidateToken] Token is invalid: %v]", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims["username"].(string), nil
	}

	return "", jwt.ErrSignatureInvalid
}
