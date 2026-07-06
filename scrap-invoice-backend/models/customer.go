package models

import "time"

type Customer struct {
	ID        int       `gorm:"column:customer_id;primaryKey" json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Email     string    `gorm:"unique" json:"email"` // <-- Ini masalahnya
	Address   string    `json:"address"`
	Tier      string    `json:"tier"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
