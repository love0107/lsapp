package auth

import (
	"lsapp/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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

func ValidateJWT(tokenString string) (bool, string) {
	// Get the JWT secret key
	crd, err := model.GetConfigByType("jwt")
	if err != nil || crd["key"] == "" {
		return false, ""
	}
	secretKey := []byte(crd["key"])

	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	// Check if the token is valid
	if err != nil || !token.Valid {
		return false, ""
	}

	// Return the email from the claims
	return true, claims.Email
}

func AuthRequired() (gin.HandlerFunc) {
	return func(c *gin.Context) {
		// Get the JWT string from the cookie
		tokenString, err := c.Cookie("session_token")
		if err != nil {
			// If the cookie is not set, return an unauthorized status
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Validate the token
		isValidToken, email := ValidateJWT(tokenString)
		user, err := new(model.User).GetUserByEmail(email)
		if err != nil || user == nil || !isValidToken {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		// Store the user information in the context.
        c.Set("user", user)
		// If the token is valid, continue to the route handler
		c.Next()
	}
}
