package signup

import (
	"fmt"
	"lsapp/controller"
	"lsapp/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// get the user by user name to check weather user already exist or not
	_, err := controller.GetUserByUserName(user.UserName)

	if err == nil { // retuning from here as username already exist
		fmt.Println("username already exist")
		return
	}
	_, err = controller.CreateUser(user.UserName, user.Password)
	if err != nil {
		fmt.Println("unable to create new user")
	}
}
