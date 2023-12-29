package login

import (
	"fmt"
	"lsapp/controller"
	"lsapp/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ls_user, err := controller.GetUserByUserName(user.UserName)
	if err != nil {
		fmt.Println("wrong username and password")
		return
	}
	if user.Password != ls_user.Password {
		fmt.Println("wrong username and password")
		return
	}
	// still not completed yet
	fmt.Printf("log in success!!")
}
