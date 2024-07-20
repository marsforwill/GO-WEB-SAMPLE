package controller

import (
	"src/src/initializer"
	"src/src/models"

	"github.com/gin-gonic/gin"
)

func GetProductByID(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	var product models.Product
	initializer.DBClient.First(&product, id)

	c.JSON(200, gin.H{
		"product": product,
	})
}
