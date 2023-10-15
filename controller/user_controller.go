package controller

import (
	"errors"
	"fmt"
	"lsapp/auth"
	"lsapp/model"

	"golang.org/x/crypto/bcrypt"
)

func AuthenticateUser(username, password string) (*model.User, error) {
	user, err := new(model.User).GetUserByUserName(username)
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
func CreateUser(username, password string) (*model.User, error) {
	password, err := auth.HashPassword(password)
	if err != nil {
		fmt.Println("something went wrong")
		return nil, errors.New("error while converting hash password")
	}
	return new(model.User).CreateUser(username, password)
}

func GetUserByUserName(username string) (*model.User, error) {
	return new(model.User).GetUserByUserName(username)
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
