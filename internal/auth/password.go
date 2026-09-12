package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plain string) (string, error) {
	if plain == "" {
		return "", errors.New("password is required")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func CheckPassword(hash, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
}
