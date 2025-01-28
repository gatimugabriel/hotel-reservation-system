package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var (
	accessTokenSecret        = os.Getenv("ACCESS_TOKEN_SECRET")
	refreshTokenSecret       = os.Getenv("REFRESH_TOKEN_SECRET")
	passwordResetTokenSecret = os.Getenv("PASSWORD_RESET_TOKEN_SECRET")
)

// GenerateTokens : generates access & refresh token
func GenerateTokens(userID string, role any) (string, string, error) {
	userPayload := map[string]interface{}{
		"userID": userID,
		"role":   role,
	}

	// Generate access token
	accessToken, err := GenerateToken(userPayload, 15*time.Minute, accessTokenSecret)
	if err != nil {
		return "", "", fmt.Errorf("error generating access token: %w", err)
	}

	// Generate refresh token
	refreshToken, err := GenerateToken(userPayload, 30*24*time.Hour, refreshTokenSecret)
	if err != nil {
		return "", "", fmt.Errorf("error generating refresh token: %w", err)
	}

	return accessToken, refreshToken, nil
}

// ValidateToken validate any given token with its type
func ValidateToken(tokenString string, tokenType string) (map[string]interface{}, error) {
	var secret []byte
	switch tokenType {
	case "REFRESH":
		secret = []byte(refreshTokenSecret)
	case "ACCESS":
		secret = []byte(accessTokenSecret)
	case "PASSWORD_RESET":
		secret = []byte(passwordResetTokenSecret)
	default:
		return nil, fmt.Errorf("invalid token type")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	// extract claims ( gets encrypted data(payload) ) & decrypt payload
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		encryptedPayload := claims["data"].(string)

		// Decrypt payload data
		decryptedPayload, err := DecryptPayload(encryptedPayload)
		if err != nil {
			return nil, err
		}

		return decryptedPayload, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// GenerateToken : creates a JWT token on given input
// accepts flexible object input & returns an encrypted string as the token's payload
func GenerateToken(payload map[string]interface{}, tokenDuration time.Duration, secretKey string) (string, error) {
	encryptedPayload, err := EncryptPayload(payload)
	if err != nil {
		return "", fmt.Errorf("error encrypting payload: %w", err)
	}

	claims := &jwt.MapClaims{
		"data": encryptedPayload,
		"exp":  time.Now().Add(tokenDuration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}