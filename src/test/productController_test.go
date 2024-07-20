package main

import (
	"net/http"
	"net/http/httptest"
	"src/src/controller"
	"src/src/initializer"
	"src/src/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// here we use the same database test for convenience, because we wont build another db instance. ideally we should isolate the db between prod and test env 
var testDBConnString = "host=devpostgre.postgres.database.azure.com user=azureuser password=mars12345678. dbname=postgres port=5432 sslmode=require"

func TestGetProductByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.GET("/products/:id", controller.GetProductByID)

	db, err := gorm.Open(postgres.Open(testDBConnString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	initializer.DBClient = db
	initializer.DBClient.AutoMigrate(&models.Product{})

	product := models.Product{
		ID:            101,
		Name:          "Test Product",
		Category:      "Test Category",
		Price:         10.99,
		Description:   "This is a test product",
		StockQuantity: 100,
	}
	initializer.DBClient.Create(&product)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products/101", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Product")
}

func TestGetProducts(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.GET("/products", controller.GetProducts)

	db, err := gorm.Open(postgres.Open(testDBConnString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	initializer.DBClient = db
	initializer.DBClient.AutoMigrate(&models.Product{})

	products := []models.Product{
		{ID: 103, Name: "Product 1", Category: "Category 1", Price: 20.0, Description: "Description 1", StockQuantity: 10},
		{ID: 104, Name: "Product 2", Category: "Category 2", Price: 30.0, Description: "Description 2", StockQuantity: 20},
	}
	for _, product := range products {
		initializer.DBClient.Create(&product)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products?category=Category 1&price_min=10&price_max=25&page=1&page_size=1&sort_order=asc", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Product 1")
	assert.NotContains(t, w.Body.String(), "Product 2")
}
