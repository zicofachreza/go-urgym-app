package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	accessSecret  = []byte(os.Getenv("JWT_ACCESS_SECRET"))
	refreshSecret = []byte(os.Getenv("JWT_REFRESH_SECRET"))
)

// Claims bisa disesuaikan dengan struktur payload kamu
type Claims struct {
	Payload map[string]any `json:"payload"`
	jwt.RegisteredClaims
}

// --- Sign Token (Access) ---
func SignToken(payload map[string]any) (string, error) {
	claims := Claims{
		Payload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(accessSecret)
}

// --- Sign Refresh Token ---
func SignRefreshToken(payload map[string]any) (string, error) {
	claims := Claims{
		Payload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(refreshSecret)
}

// --- Verify Access Token ---
func VerifyToken(tokenString string) (map[string]any, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return accessSecret, nil
	})

	if err != nil {
		// jwt/v5: gunakan IsError untuk deteksi expired
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, NewCodeError("AccessTokenExpired")
		}
		return nil, NewCodeError("InvalidAccessToken")
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.Payload, nil
	}

	return nil, NewCodeError("InvalidAccessToken")
}

// --- Verify Refresh Token ---
func VerifyRefreshToken(tokenString string) (map[string]any, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return refreshSecret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, NewCodeError("RefreshTokenExpired")
		}
		return nil, NewCodeError("InvalidRefreshToken")
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.Payload, nil
	}

	return nil, NewCodeError("InvalidRefreshToken")
}
