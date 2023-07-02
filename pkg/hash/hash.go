package hash

import (
	"golang.org/x/crypto/bcrypt"
)

// Hashes given string
func HashString(str string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), 14)
	return string(bytes), err
}

// Compares raw string with it's hashed values
func CheckStringHash(str, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	return err == nil
}
