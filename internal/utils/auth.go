package utils

import (
	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func GenerateAccessToken(id string, expiresAt int64) (string, error) {
	// header := map[string]interface{}{
	// 	"alg": "HS256",
	// 	"typ": "JWT",
	// }
	payload := map[string]interface{}{
		"id":  id,
		"exp": expiresAt,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(payload))

	secret := []byte("Ylmv9buRUxcrKVwwfMZ4KuWOZVvtjoWB")
	tokenString, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateRefreshToken(id string) (string, error) {
	// headerRefresh := map[string]interface{}{
	// 	"alg": "HS256",
	// 	"typ": "JWT",
	// }
	payloadRefresh := map[string]interface{}{
		"id": id,
	}

	tokenRefresh := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(payloadRefresh))

	secret := []byte("secret")
	tokenRefreshString, err := tokenRefresh.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenRefreshString, nil
}

func HashPassword(password string) string {

	// use algorithm bcrypt to encode password with salt
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash)
}

func CompareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
