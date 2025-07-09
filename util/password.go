package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword handles the password hashing using bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed  to hash password: %w", err)

	}
	return string(hashedPassword), nil
}

// CheckPassword check if provided password is correct
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
