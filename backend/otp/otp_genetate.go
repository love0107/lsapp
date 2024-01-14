package otp

import (
	"fmt"
	"lsapp/model"
	"math/rand"
	"strconv"
	"time"
)

func GenerateOtp(length int, user *model.User) string {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	otp := ""
	for i := 0; i < length; i++ {
		otp += strconv.Itoa(r.Intn(10))
	}
	fmt.Println(otp)
	// store the otp in the db
	ls_opt := model.OTP{Otp: otp}
	_, err :=ls_opt.AddOtp()
	if err!=nil{
		fmt.Println("failed to get store otp")
	}
	return otp
}
