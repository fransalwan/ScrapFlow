package models

import "time"

type ScaleDetail struct {
	ID         int       `gorm:"column:scale_detail_id;primaryKey" json:"id"`
	SummaryID  *int      `gorm:"column:summary_id" json:"summary_id"` // pointer supaya bisa null
	InvoiceID  int       `gorm:"column:invoice_id" json:"invoice_id"` // foreign key baru
	ItemID     int       `gorm:"column:item_id" json:"item_id"`
	Weight     float64   `gorm:"column:weight" json:"weight"`
	AlasWeight float64   `gorm:"column:alas_weight" json:"alas_weight"`
	Photo      string    `gorm:"column:photo" json:"photo"`
	ScaleType  string    `gorm:"column:scale_type" json:"scale_type"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`

	Item    Item    `gorm:"foreignKey:ItemID" json:"item"`
	Invoice Invoice `gorm:"foreignKey:InvoiceID" json:"invoice"`
}
