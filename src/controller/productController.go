package controller

import (
	"net/http"
	"src/src/initializer"
	"src/src/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	defaultPageSize  = 10
	defaultSortOrder = "desc"
)

// Retrieves a product by its ID.
func GetProductByID(c *gin.Context) {
	// get id from url
	id := c.Param("id")
	var product models.Product

	if err := initializer.DBClient.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

// Retrieves a list of products with optional filters for category and price range, along with pagination and sorting options.
func GetProducts(c *gin.Context) {
	var products []models.Product
	var count int64

	query, pageInt, pageSizeInt, sortOrder := buildGetProductsQuery(c)

	// Get total count for pagination
	if err := query.Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count products"})
		return
	}

	if err := query.Order("created_at " + sortOrder).
		Offset((pageInt - 1) * pageSizeInt).
		Limit(pageSizeInt).
		Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	// Send response
	c.JSON(http.StatusOK, gin.H{
		"total_count": count,
		"page":        pageInt,
		"page_size":   pageSizeInt,
		"products":    products,
	})
}

func buildGetProductsQuery(c *gin.Context) (*gorm.DB, int, int, string) {
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
		pageSizeInt = defaultPageSize
	}

	// Default sort order
	if sortOrder != "asc" {
		sortOrder = defaultSortOrder
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

	return query, pageInt, pageSizeInt, sortOrder
}
