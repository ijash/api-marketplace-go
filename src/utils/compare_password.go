package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// ComparePassword compares a plain text password with a hashed password.
func ComparePassword(plainPassword, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err
}
