package helpers

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type Claims struct {
	UserID   int
	FullName string
	Email    string
	Role     string
	jwt.RegisteredClaims
}

var MapTypeToken = map[string]time.Duration{
	"token":         time.Hour * 3,
	"refresh_token": time.Hour * 24,
}

var secretKey = []byte(os.Getenv("UMS_JWT_SECRET_KEY"))

func GenerateToken(ctx context.Context, userID int, userFullName, userEmail, userRole, tokenType string) (string, error) {
	claims := &Claims{
		UserID:   userID,
		FullName: userFullName,
		Email:    userEmail,
		Role:     userRole,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    os.Getenv("APP_NAME"),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(MapTypeToken[tokenType])),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func ValidateToken(ctx context.Context, token string) (*Claims, error) {
	var (
		claims = &Claims{}
		ok     bool
	)

	jwtToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("failed to validate method jwt: %v", t.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok = jwtToken.Claims.(*Claims); !ok || !jwtToken.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
