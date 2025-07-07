package models

import "time"

type ItemCategory struct {
	ID               int       `gorm:"column:item_category_id;primaryKey" json:"id"`
	ItemCategoryName string    `gorm:"column:item_category_name" json:"item_category_name"`
	CreatedAt        time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at" json:"updated_at"`
}
