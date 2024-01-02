package otp

import (
	"fmt"
	"lsapp/model"
	"math/rand"
	"strconv"
	"time"
)

func GenerateOtp(length, userId int) string {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	otp := ""
	for i := 0; i < length; i++ {
		otp += strconv.Itoa(r.Intn(10))
	}
	fmt.Println(otp)
	// store the otp in the db
	model.StoreOtp(otp, userId)
	return otp
}
