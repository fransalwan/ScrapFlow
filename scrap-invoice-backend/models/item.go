package models

import "time"

type Item struct {
	ID             int       `gorm:"column:item_id;primaryKey" json:"id"`
	ItemName       string    `gorm:"column:item_name" json:"item_name"`
	ItemCategoryID int       `gorm:"column:item_category_id" json:"item_category"`
	PricePerKg     float64   `gorm:"column:price_per_kg" json:"price_per_kg"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`

	Category ItemCategory `gorm:"foreignKey:ItemCategoryID;references:ID" json:"category"`
}
