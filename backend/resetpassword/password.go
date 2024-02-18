package password

import (
	"lsapp/controller"
	"lsapp/otp"
	"lsapp/sms"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	constClickatell = "clickatell"
)

type ResetPassword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RestPassword(c *gin.Context) {

	var resetPassword ResetPassword
	if err := c.ShouldBindJSON(&resetPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Retrieve user from the database using the provided email
	user, err := controller.GetUserByEmail(resetPassword.Email)
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "wrong email address",
			"error":   err.Error(),
		})
		return
	}
	otp, err := otp.GenerateOtp(6, user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "failed to generate OTP",
			"error":   err.Error(),
		})
		return
	}
	dataRequest := sms.Request{
		To:      user.Mobile,
		Type:    "otp",
		Message: otp,
		Vendor:  constClickatell,
		UserId:  user.Id,
	}
	response, err := sms.SendSms(dataRequest)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "failed to send OTP",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "otp send successfully",
		"status":  response.Status,
		"code":    response.Code,
		"body":    response.Body,
	})
	// validate the opt
	// reset the password
	// log in again
}


