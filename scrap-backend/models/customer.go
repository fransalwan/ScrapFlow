// customer.go
package models

import "time"

type Customer struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	Notes     string    `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
}
