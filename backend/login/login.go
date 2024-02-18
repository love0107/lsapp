package login

import (
	"lsapp/auth"
	"lsapp/controller"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Email    string `orm:"column(email);size(255)"`
	Password string `orm:"column(password);size(255)"`
}

func Login(c *gin.Context) {
	var loginRequest Credentials
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Retrieve user from the database using the provided email
	user, _ := controller.GetUserByEmail(loginRequest.Email)
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "wrong email address"})
		return
	}

	// Compare hashed password from the database with the provided plaintext password
	isValidPassword := auth.ComparePasswords(loginRequest.Password, user.Password)
	if !isValidPassword {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong password"})
		return
	}

	// generate jwt token
	tokenString, err := auth.GenerateJWT(user.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":       err.Error(),
			"message":     "failed to get token",
			"tokenString": tokenString,
		})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: time.Now().Add(time.Minute * 5),
	})

	// Password is correct, return user data
	c.JSON(http.StatusOK, gin.H{"user": user})
}
