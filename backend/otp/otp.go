package otp

import (
	"fmt"
	"lsapp/auth"
	"lsapp/log"
	"lsapp/model"
	"lsapp/util"
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
	localTimeZone, err := util.GetCurrentTimeIn("Asia/Kolkata")
	if err != nil {
		log.Println("failed to get current time in Asia/Kolkata timezone")
		return "", err
	}
	expireAt := localTimeZone.Add(5 * time.Minute)
	opt := model.OTP{
		Mobile:   user.Mobile,
		Otp:      currentOTP,
		SentTime: localTimeZone,
		ExpireAt: expireAt,
		UserID:   user.Id,
		SentTo:   user.Mobile,
		Email:    user.Email,
	}
	_, err = opt.AddOtp()
	if err != nil {
		fmt.Println("failed to store OTP:", err)
		return "", err
	}

	return currentOTP, nil
}

func ValidateUserOTP(c *gin.Context) {
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
	otp, err := new(model.OTP).GetOtpByUserId(user.Id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get otp",
			"error":   err.Error(),
		})
		return

	}

	if otp.Otp != inputOTP.Otp {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "OTP is invalid",
		})
		return
	}

	// check the expiry time
	isTimeExpired := isExpired(otp.ExpireAt)
	if isTimeExpired {
		c.JSON(http.StatusRequestTimeout, gin.H{
			"message": "OTP is expired",
		})
		return
	}

	// generate jwt token
	tokenString, err := auth.GenerateJWT(otp.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "failed to get token",
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
	// set the cookie
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "session_token",
		Value:   tokenString,
		Path:    "/",
		Expires: localTimeZone.Add(time.Minute * 5),
	})

	c.JSON(http.StatusOK, gin.H{
		"message":    "OTP is validated successfully",
		"satus":      "success",
		"statusCode": 200,
	})
}

// isExpired checks
func isExpired(expTime time.Time) bool {

	// Compare the current time with the expiration time
	localTimeZone, err := util.GetCurrentTimeIn("Asia/Kolkata")
	if err != nil {
		log.Println("failed to get current time in Asia/Kolkata timezone")
		return false
	}
	return localTimeZone.After(expTime)

}
