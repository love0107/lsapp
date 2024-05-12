package login

import (
	"lsapp/auth"
	"lsapp/controller"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Email    string `orm:"column(email);size(255)"`
	Password string `orm:"column(password);size(255)"`
}

func Login(c *gin.Context) {

 // Check for an existing cookie
    if tokenString, err := c.Cookie("session_token"); err == nil && tokenString != ""{
		// Validate the token
		isValidToken, email :=auth.ValidateJWT(tokenString)
		
       user, err:= controller.GetUserByEmail(email)
	   if err == nil &&  user != nil{
		if isValidToken {
			// If the token is valid, return a success message
			c.JSON(http.StatusOK, gin.H{
				"message": "Already logged in",
				"statusCode":200,
				"status": "success",
			})
			return
		}
	   }
	}


	var loginRequest Credentials
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input data",
			"error":   err.Error(),
		})
		return
	}

	// Retrieve user from the database using the provided email
	user, err := controller.GetUserByEmail(loginRequest.Email)
	if user == nil || err!= nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "User not found",
			"error":   err.Error(),
		})
		return
	}

	// Compare hashed password from the database with the provided plaintext password
	isValidPassword := auth.ComparePasswords(loginRequest.Password, user.Password)
	if !isValidPassword {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "wrong password",
			"error":   "invalid password",
		})
		return
	}


	// generate jwt token
	tokenString, err := auth.GenerateJWT(user.Email)
	if err != nil || tokenString == ""{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":       err.Error(),
			"message":     "failed to generate token",
			"tokenString": tokenString,
		})
		return
	}


	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "session_token",
		Value:   tokenString,
		Expires: time.Now().Add(time.Minute * 5),
	})

	// Password is correct, return user data
	c.JSON(http.StatusOK, gin.H{
		"message": "User logged in successfully",
		"statusCode":200,
		"status": "success",
	})
}
