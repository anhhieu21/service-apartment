package utils

import (
	"context"
	"fmt"
	"strings"

	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/metadata"
)

const SECRET = "Ylmv9buRUxcrKVwwfMZ4KuWOZVvtjoWB"

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

	secret := []byte(SECRET)
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

func GetCustomerIdFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {

		return "", fmt.Errorf("metadata is not provided")
	}

	values := md["authorization"]

	if len(values) == 0 {

		return "", fmt.Errorf("authorization token is not provided")
	}
	accessToken := values[0]
	accessToken = strings.TrimPrefix(accessToken, "Bearer ")
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})
	if err != nil {
		return "", err
	}
	//get claims
	claims := token.Claims.(jwt.MapClaims)
	return claims["id"].(string), nil
}

func Verify(accessToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}
	//get claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
