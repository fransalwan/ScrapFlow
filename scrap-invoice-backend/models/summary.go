package models

import "time"

type Summary struct {
	ID            int       `gorm:"column:summary_id;primaryKey" json:"id"`
	InvoiceID     int       `gorm:"column:invoice_id" json:"invoice_id"`
	ItemID        int       `gorm:"column:item_id" json:"item_id"`
	ItemName      string    `gorm:"column:item_name" json:"item_name"`
	TotalWeight   float64   `gorm:"column:total_weight" json:"total_weight"`
	SubTotalPrice float64   `gorm:"column:sub_total_price" json:"sub_total_price"`
	ScaleType     string    `gorm:"column:scale_type" json:"scale_type"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`

	// Relasi
	Item         Item          `gorm:"foreignKey:ItemID" json:"item"`
	ScaleDetails []ScaleDetail `gorm:"foreignKey:SummaryID" json:"scale_details"`
}
