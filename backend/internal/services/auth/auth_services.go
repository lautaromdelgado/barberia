package auth_services

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTService es una estructura que representa un servicio de JWT
type JWTService struct {
	secretKey string
}

// NewJWTService es un constructor de JWTService
func NewJWTService() *JWTService {
	secretJWT := os.Getenv("JWT_SECRET")
	if secretJWT == "" {
		log.Panicln("jwt secret is not ser: &v", secretJWT)
	}
	return &JWTService{secretKey: secretJWT}
}

// GenerateToken genera un token JWT
func (j *JWTService) GenerateToken(id uint, nombre, apellido, correo, rol string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  id,
		"name":     nombre,
		"apellido": apellido,
		"correo":   correo,
		"rol":      rol,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

// ValidateToken valida un token JWT
func (j *JWTService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.secretKey), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token, nil
}
