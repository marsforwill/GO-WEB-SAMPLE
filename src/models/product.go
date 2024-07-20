package models

import (
	"time"

	"gorm.io/gorm"
)

// Product represents the product model for the products table
type Product struct {
	ID            uint           `gorm:"primaryKey"`                              // Auto-increment primary key
	Name          string         `gorm:"type:varchar(255);not null"`              // Product name, not null
	Category      string         `gorm:"type:varchar(255);not null"`              // Product category, not null
	Price         float64        `gorm:"type:numeric(10,2);not null"`             // Product price with 2 decimal precision, not null
	Description   string         `gorm:"type:text"`                               // Product description, nullable
	StockQuantity int            `gorm:"type:int;not null;default:0"`             // Stock quantity, default value is 0, not null
	CreatedAt     *time.Time     `gorm:"type:timestamptz;not null;default:now()"` // Creation timestamp, default value is current time, not null
	UpdatedAt     *time.Time     `gorm:"type:timestamptz;not null;default:now()"` // Update timestamp, default value is current time, not null
	DeletedAt     gorm.DeletedAt `gorm:"type:timestamptz"`                        // Deletion timestamp, nullable, for soft delete
}
