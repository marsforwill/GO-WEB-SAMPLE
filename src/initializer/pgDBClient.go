package initializer

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBClient *gorm.DB

// connect cloud(azure) db server so that we can use DBClient later
func ConnectToDB() {
	var err error
	dbConnString := os.Getenv("DB_Connect")
	DBClient, err = gorm.Open(postgres.Open(dbConnString), &gorm.Config{})
	if DBClient != nil && err == nil {
		log.Println("postgre DB client init success")
	} else {
		log.Fatalf("postgre DB client init error: %v", err)
	}
}
