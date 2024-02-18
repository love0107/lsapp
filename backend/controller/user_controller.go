package controller

import (
	"errors"
	"fmt"
	"lsapp/auth"
	"lsapp/model"

	"golang.org/x/crypto/bcrypt"
)

func AuthenticateUser(email, password string) (*model.User, error) {
	user, err := new(model.User).GetUserByEmail(email)
	if err != nil {
		fmt.Println("invalid username")
		return nil, err
	}
	//password checking...
	if user.Password != password {
		fmt.Println("invalid user password")
	}
	return user, err
}

func CreateUser(user *model.User) (id int64,err error) {
	user.Password, err = auth.HashPassword(user.Password)
	if err != nil {
		fmt.Println("something went wrong")
		return id, errors.New("error while converting hash password")
	}
	return user.CreateUser()
}

func GetUserByUserName(username string) (*model.User, error) {
	return new(model.User).GetUserByUserName(username)
}

func GetUserByMobile(moblie string)(*model.User, error){
	return new(model.User).GetUserByMobile(moblie)
}

func GetUserByEmail(email string)(*model.User, error){
	return new(model.User).GetUserByEmail(email)
}

func hashPassword(password string) (string, error) {
	// Generate a salt with a cost of 10
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hash), err
}

// Compares a hashed password with its possible plaintext equivalent
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
