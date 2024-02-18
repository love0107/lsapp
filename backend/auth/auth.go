package auth

import (
	"lsapp/model"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func HashPassword(password string) (string, error) {
	// Generate a salt with a cost of 10
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hash), err
}

// Compares a hashed password with its possible plaintext equivalent
func ComparePasswords(plainPassword, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plainPassword))
	return err == nil
}

func GenerateJWT(email string) (string, error) {
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	crd, err := model.GetConfigByType("jwt")
	if err != nil {
		return "", err
	}
	tokenString, err := token.SignedString([]byte(crd["key"]))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
