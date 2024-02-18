package otp

import (
	"fmt"
	"lsapp/model"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

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
		SentTo: user.Mobile,
	}
	_, err := opt.AddOtp()
	if err != nil {
		fmt.Println("failed to store OTP:", err)
		return "", err
	}

	return currentOTP, nil
}

func ValidateOTP(c *gin.Context) {
    
}
