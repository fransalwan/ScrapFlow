package models

import "time"

type Transaction struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	CustomerID uint      `json:"customer_id"`
	Customer   Customer  `gorm:"foreignKey:CustomerID" json:"customer"`
	Material   string    `json:"material"`
	Weight     float64   `json:"weight"`
	PricePerKg float64   `json:"price_per_kg"`
	Total      float64   `json:"total"`
	CreatedAt  time.Time `json:"created_at"`
}
