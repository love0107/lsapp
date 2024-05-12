package main

import (
	"lsapp/auth"
	"lsapp/login"
	"lsapp/otp"
	"lsapp/password"
	"lsapp/persistance"
	"lsapp/signup"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	persistance.Init()

	router := setupRouter()

	// router.POST("/password/validate", otp.ValidateOTP)
	router.Run(":8080")
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	// CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}  // Update with your frontend URL
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"} // Allow the methods needed

	router.Use(cors.New(config))

	// Define routes
	// Public routes
	public := router.Group("/")
	{
		public.POST("/signup", signup.SignUp)
		public.POST("/login", login.Login)
		public.POST("/password/reset", password.RestPassword)
	}

	// Private routes
	private := router.Group("/")
	private.Use(auth.AuthRequired()) // Apply the middleware to these routes
	{
		private.POST("/password/update", password.UpdatePassword)
		private.POST("/password/otp/validate", otp.ValidateUserOTP)
		// private.GET("/profile", profile.GetProfile)
        // private.PUT("/profile", profile.UpdateProfile)
	}

	return router
}
