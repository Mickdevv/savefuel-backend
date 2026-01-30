package auth

import (
	"fmt"
	"net/mail"
	"strings"
	"time"

	"github.com/alexedwards/argon2id"
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
	s, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", err
	}

	return s, nil
}

func ValidateJWT(tokenString string, tokenSecret string) (Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (any, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("Incorrect signing method: %v", t.Header["alg"])
		}
		return []byte(tokenSecret), nil
	})
	if err != nil {
		fmt.Println("Parsing failed")
		return Claims{}, err
	}

	if !token.Valid {
		fmt.Println("Token invalid")
		return Claims{}, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		fmt.Println("Error casting to claims")
		return Claims{}, err
	}

	_, err = uuid.Parse(claims.UserId)
	if err != nil {
		return Claims{}, err
	}

	return *claims, nil
}

func CheckPassword(password, hash string) (bool, error) {
	ok, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		return false, err
	}
	return ok, nil

}

func HashPassword(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func ValidatePassword(password string) error {
	var validationErrors []string
	if len(password) < 5 {
		validationErrors = append(validationErrors, "Password must be at least 5 chartacters in length")
	}

	if len(validationErrors) > 0 {
		return fmt.Errorf(strings.Join(validationErrors, ", "))
	}
	return nil
}

func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
