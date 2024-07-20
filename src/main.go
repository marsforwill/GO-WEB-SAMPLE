package main

import (
	"src/src/controller"
	"src/src/initializer"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnv()
	initializer.ConnectToDB()
}

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/products/:id", controller.GetProductByID)
	r.Run() // listen and serve on 0.0.0.0:8080
}
