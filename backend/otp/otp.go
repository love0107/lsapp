package otp

import (
	"fmt"
	"lsapp/auth"
	"lsapp/model"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ValidateOTP struct {
	Otp string `json:"otp"`
}

func GenerateOtp(length int, user *model.User) (string, error) {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	currentOTP := ""
	for i := 0; i < length; i++ {
		currentOTP += strconv.Itoa(r.Intn(10))
	}

	// Store the OTP in the database
	expireAt := time.Now().Add(5 * time.Minute)
	opt := model.OTP{
		Mobile:   user.Mobile,
		Otp:      currentOTP,
		SentTime: time.Now(),
		ExpireAt: expireAt,
		UserID:   user.Id,
		SentTo:   user.Mobile,
		Email:    user.Email,
	}
	_, err := opt.AddOtp()
	if err != nil {
		fmt.Println("failed to store OTP:", err)
		return "", err
	}

	return currentOTP, nil
}

func ValidateUserOTP(c *gin.Context) {
	// Check for an existing cookie

	var inputOTP ValidateOTP
	if err := c.ShouldBindJSON(&inputOTP); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input data",
			"error":   err.Error(),
		})
		return
	}

	// get the otp by otp
	otp, err := new(model.OTP).GetOtpByOTP(inputOTP.Otp)

	if err != nil || otp == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get otp",
			"error":   err.Error(),
		})
		return

	}
	// check the expiry time
	isTimeExpired := isExpired(otp.ExpireAt)
	if isTimeExpired {
		c.JSON(http.StatusRequestTimeout, gin.H{
			"message": "OTP is expired",
			"error":   err.Error(),
		})
		return
	}

	// generate jwt token
	tokenString, err := auth.GenerateJWT(otp.Email)
	if err != nil || tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":       err.Error(),
			"message":     "failed to get token",
			"tokenString": tokenString,
		})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "session_token",
		Value:   tokenString,
		Expires: time.Now().Add(time.Minute * 5),
	})

	c.JSON(http.StatusRequestTimeout, gin.H{
		"message":    "OTP is validated successfully",
		"satus":      "success",
		"statusCode": 200,
	})
}

// isExpired checks
func isExpired(expTime time.Time) bool {

	// Compare the current time with the expiration time
	return time.Now().After(expTime)

}
