package signup

import (
	"lsapp/auth"
	"lsapp/controller"
	"lsapp/model"
	"lsapp/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var newUser model.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUser.Mobile = "91" + newUser.Mobile
	// Create the new user
	userId, err := controller.CreateUser(&newUser)
	if err != nil || userId == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to create new user",
			"error":   err.Error(),
		})
		return
	}
	// set the cookie
	token, err := auth.GenerateJWT(newUser.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to generate token",
			"error":   err.Error(),
		})
		return
	}
	localTimeZone, err := util.GetCurrentTimeIn("Asia/Kolkata")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get current time in Asia/Kolkata timezone",
			"error":   err.Error(),
		})
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "session_token",
		Value:   token,
		Path:    "/",
		Expires: localTimeZone.Add(time.Minute * 5),
	})

	c.JSON(http.StatusOK, gin.H{
		"userId":     userId,
		"message":    "User created successfully",
		"statusCode": 200,
	})
}
