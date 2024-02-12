package login

import (
	"lsapp/auth"
	"lsapp/controller"
	"lsapp/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginRequest model.User
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Retrieve user from the database using the provided email
	user, err := controller.GetUserByEmail(loginRequest.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Compare hashed password from the database with the provided plaintext password
	isValidPassword := auth.ComparePasswords(loginRequest.Password, user.Password)
	if !isValidPassword {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
		return
	}
	
	// Password is correct, return user data
	c.JSON(http.StatusOK, gin.H{"user": user})
}
