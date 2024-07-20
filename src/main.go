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

	r.GET("/product/:id", controller.GetProductByID)
	r.GET("/products", controller.GetProducts)
	r.Run() // listen and serve on 0.0.0.0:9000
}
