package hash

import (
	"golang.org/x/crypto/bcrypt"
)

// hashString hashes given string
func hashString(str string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), 14)
	return string(bytes), err
}

// checkStash compares raw string with it's hashed values
func checkStringHash(str, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	return err == nil
}
