package models

import "time"

type Customer struct {
	ID        int       `gorm:"column:id_customer;primaryKey" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Phone     string    `gorm:"column:phone" json:"phone"`
	Email     string    `gorm:"column:email" json:"email"`
	Address   string    `gorm:"column:address" json:"address"`
	Tier      string    `gorm:"column:tier" json:"tier"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
