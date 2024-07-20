package controller

import (
	"net/http"
	"src/src/initializer"
	"src/src/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProductByID(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	var product models.Product
	initializer.DBClient.First(&product, id)

	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

func GetProducts(c *gin.Context) {
	var products []models.Product
	var count int64

	// Get filter parameters from URL
	category := c.Query("category")
	priceMin := c.Query("price_min")
	priceMax := c.Query("price_max")
	page := c.Query("page")
	pageSize := c.Query("page_size")
	sortOrder := c.Query("sort_order")

	// Default values for pagination
	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)
	if pageInt <= 0 {
		pageInt = 1
	}
	if pageSizeInt <= 0 {
		pageSizeInt = 10
	}

	// Default sort order
	if sortOrder != "asc" {
		sortOrder = "desc"
	}

	// Build query
	query := initializer.DBClient.Model(&models.Product{})
	if category != "" {
		query = query.Where("category = ?", category)
	}
	if priceMin != "" {
		priceMinFloat, _ := strconv.ParseFloat(priceMin, 64)
		query = query.Where("price >= ?", priceMinFloat)
	}
	if priceMax != "" {
		priceMaxFloat, _ := strconv.ParseFloat(priceMax, 64)
		query = query.Where("price <= ?", priceMaxFloat)
	}

	// Get total count for pagination
	query.Count(&count)

	// Apply pagination and sorting
	query.Order("created_at " + sortOrder).
		Offset((pageInt - 1) * pageSizeInt).
		Limit(pageSizeInt).
		Find(&products)

	// Send response
	c.JSON(http.StatusOK, gin.H{
		"total_count": count,
		"page":        pageInt,
		"page_size":   pageSizeInt,
		"products":    products,
	})
}
