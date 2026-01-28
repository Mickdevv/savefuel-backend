package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	UserId string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func CreateJWT(tokenSecret string, userId uuid.UUID) (string, error) {
	claims := Claims{
		UserId: userId.String(),
		Role:   "TestRoleValue",
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userId.String(),
			Issuer:    "savefuel",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := token.SignedString(tokenSecret)
	if err != nil {
		return "", err
	}
	fmt.Println(s)

	return s, nil
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (any, error) {
		if t.Method.Alg() != jwt.SigningMethodES256.Alg() {
			return nil, fmt.Errorf("Incorrect signing method: %v", t.Header["alg"])
		}
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	if !token.Valid {
		return uuid.Nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return uuid.Nil, err
	}

	subject, err := uuid.Parse(claims.UserId)
	if err != nil {
		return uuid.Nil, err
	}

	return subject, nil
}
