package password

import (
	"errors"
	"lsapp/auth"
	"lsapp/controller"
	"lsapp/model"
	"lsapp/otp"
	"net/http"
	"time"

	"lsapp/communicationchannel/sms"

	"github.com/gin-gonic/gin"
)

const (
	constClickatell = "clickatell"
)

type ResetPassword struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func RestPassword(c *gin.Context) {
	var resetPassword ResetPassword
	if err := c.ShouldBindJSON(&resetPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input data",
			"error":   err.Error(),
		})
		return
	}

	// Retrieve user from the database using the provided email
	user, err := controller.GetUserByEmail(resetPassword.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get user by email",
			"error":   err.Error(),
		})
		return
	}

	token, err := auth.GenerateJWT(user.Email)
	if err != nil || token == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to generate token",
			"error":   err.Error(),
		})
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "session_token",
		Value:   token,
		Expires: time.Now().Add(time.Minute * 5),
	})

	// generate otp
	otp, err := otp.GenerateOtp(6, user)
	if err != nil || otp == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
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
	if err != nil || response == nil {
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
		// "token":   token,
	})
}

func UpdatePassword(c *gin.Context) {
	existUser, exists := c.Get("user")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
        return
    }
	 // Type assertion
	 user, ok := existUser.(*model.User)
	 if !ok {
		 c.JSON(http.StatusInternalServerError, gin.H{
			 "error": "Internal Server Error",
		 })
		 return
	 }

	// Extract the user ID from the URL parameter
	var resetPassword ResetPassword
	if err := c.ShouldBindJSON(&resetPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input data",
			"error":   err.Error(),
		})
		return
	}

	// hash the password
	password, err := auth.HashPassword(resetPassword.Password)
	if err != nil || password == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to hash password",
			"error":   err.Error(),
		})
		return
	}

	// update the password
	id, err := user.UpdateUserPassword(password)
	if id == 0 && err == nil{
        err=errors.New("failed to update password")
	}
	if err != nil || id == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to update password",
			"error":   err.Error(),
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"message": "password changed successfully",
		"status":  "success",
		"code":    http.StatusOK,
	})
	return
}
