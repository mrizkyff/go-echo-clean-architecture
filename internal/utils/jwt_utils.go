package utils

import (
	"errors"
	"fmt"
	"go-echo-clean-architecture/internal/models"
	"go-echo-clean-architecture/internal/models/config"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// JWTUtil JWT utils struct
type JWTUtil struct {
	config config.JWTConfig
}

// NewJWTUtil membuat instance baru dari JWTUtil
func NewJWTUtil(config config.JWTConfig) *JWTUtil {
	return &JWTUtil{
		config: config,
	}
}

// GenerateToken menghasilkan token JWT baru
func (j *JWTUtil) GenerateToken(userID, username, email string) (string, error) {
	// Buat claims dengan info user dan expiration time
	claims := models.JWTClaims{
		ID:       userID,
		Username: username,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(j.config.ExpirationHours))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    j.config.Issuer,
		},
	}

	// Buat token dengan signing method HS256 dan claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token dengan secret key
	tokenString, err := token.SignedString([]byte(j.config.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken memvalidasi dan mengekstrak claims dari token
func (j *JWTUtil) ValidateToken(tokenString string) (*models.JWTClaims, error) {
	// Parse token dengan metode validasi kustom
	token, err := jwt.ParseWithClaims(tokenString, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validasi algoritma signing
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method not valid: %v", token.Header["alg"])
		}
		return []byte(j.config.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// Validasi token dan ekstrak claims
	if claims, ok := token.Claims.(*models.JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// ExtractTokenFromHeader mengekstrak token dari header Authorization
func ExtractTokenFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header not found")
	}

	// Biasanya format header: "Bearer {token}"
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("format header Authorization tidak ditemukan")
	}

	return parts[1], nil
}

// RefreshToken memperpanjang token dengan membuat token baru
func (j *JWTUtil) RefreshToken(tokenString string) (string, error) {
	// Validasi token lama
	claims, err := j.ValidateToken(tokenString)
	if err != nil {
		return "", err
	}

	// Buat token baru dengan claims yang sama tetapi expiration time baru
	return j.GenerateToken(claims.ID, claims.Username, claims.Email)
}

// GetUserFromContext mengambil data user dari context Echo
func (j *JWTUtil) GetUserFromContext(c echo.Context) (*models.JWTClaimsTyping, error) {
	claimsInterface := c.Get("claims")
	if claimsInterface == nil {
		return nil, errors.New("claims not found in context")
	}

	claims, ok := claimsInterface.(*models.JWTClaims)
	if !ok {
		return nil, errors.New("cannot convert claims")
	}

	loggedInUserId, _ := uuid.Parse(claims.ID)

	return &models.JWTClaimsTyping{
		ID:               loggedInUserId,
		Username:         claims.Username,
		Email:            claims.Email,
		RegisteredClaims: jwt.RegisteredClaims{},
	}, nil
}
