package signup

import (
	"lsapp/controller"
	"lsapp/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var newUser model.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// did not need to check for email , mobile and userName if we make in table it as unique

	// Check if the user already exists
	// existUser, _ := controller.GetUserByUserName(newUser.UserName)
	// if existUser != nil && existUser.UserName == newUser.UserName {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "user name already exist", "user": existUser})
	// 	return
	// }

	// check for exist email
	// existUser, _ = controller.GetUserByEmail(newUser.Email)
	// if existUser != nil && existUser.Email == newUser.Email {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "user email already exist","user": existUser})
	// 	return
	// }

	// check for exist mobile
	// existUser, _ = controller.GetUserByMobile(newUser.Mobile)
	// if existUser != nil && existUser.Mobile == newUser.Mobile {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "mobile already exist", "user": existUser})
	// 	return
	// }

	// Create the new user
	userId, err := controller.CreateUser(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to create new user",
			"error":   err.Error(),
		})
		return
	}

	if userId == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to create new user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"userId":  userId,
		"message": "User created successfully",
	})
}
