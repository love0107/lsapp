package main

import (
	"lsapp/login"
	"lsapp/persistance"
	"lsapp/signup"

	"github.com/gin-gonic/gin"
)

func main() {
	persistance.Init()
	router := gin.Default()

	router.POST("/signup", signup.SignUp)
	router.POST("/login", login.Login)
	router.Run(":8080")
}

// func getUser(c *gin.Context) {
// 	// Extract the user ID from the URL parameter
// 	idStr := c.Param("id")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
// 		return
// 	}

// 	// Get the user by ID
// 	user, err := new(model.User).GetUserById(id)
// 	if err != nil {
// 		// Handle the error, perhaps by sending an error response
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
// 		return
// 	}

// 	// Check if the user is found
// 	if user == nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 		return
// 	}

// 	// Return the user data as JSON using c.JSON
// 	c.JSON(http.StatusOK, gin.H{"user": user})
// }
