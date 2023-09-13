package jwt

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Secret key being used to sign tokens
var (
	issKey    = os.Getenv("ISS_KEY")
	secretKey = []byte(os.Getenv("JWT_SECRET"))
)

// Data we save in each token
type Claims struct {
	id uint
	jwt.RegisteredClaims
}

// Generate a JWT token and assign an ID to its claims and return it
func GenerateToken(id uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	claims["id"] = id
	claims["iss"] = issKey
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		log.Fatal("Error in generating key")
		return "", err
	}

	return tokenString, nil
}

// Parse a JWT token and returns the ID in its claims
func ParseToken(tokenStr string) (float64, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := claims["id"].(float64)
		return id, nil
	} else {
		return 0, err
	}
}
