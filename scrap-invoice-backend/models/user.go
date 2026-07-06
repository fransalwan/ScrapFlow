package models

import (
	"time"
)

type User struct {
	UserID    int       `gorm:"primaryKey;autoIncrement;column:user_id" json:"user_id"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	Role      string    `gorm:"default:'staff'" json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func (User) TableName() string {
	return "users"
}
