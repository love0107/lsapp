package main

import (
	"lsapp/persistance"

	"github.com/gin-gonic/gin"
)

func main() {
	persistance.Init()
	// register the model as well like so  orm.RegisterModel(new(models.User))
	r := gin.Default()
	r.Run(":8080")
	//
}
