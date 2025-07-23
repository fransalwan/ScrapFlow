package models

import "time"

type ScaleDetail struct {
	ID         int       `gorm:"column:scale_detail_id;primaryKey" json:"id"`
	SummaryID  *int      `gorm:"column:summary_id" json:"summary_id"`          // nullable, OK
	InvoiceID  int       `gorm:"column:invoice_id;not null" json:"invoice_id"` // FK, required
	ItemID     int       `gorm:"column:item_id;not null" json:"item_id"`       // FK, required
	Weight     float64   `gorm:"column:weight" json:"weight"`
	AlasWeight float64   `gorm:"column:alas_weight" json:"alas_weight"`
	Photo      string    `gorm:"column:photo" json:"photo"`
	ScaleType  string    `gorm:"column:scale_type" json:"scale_type"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	// --- Relationships ---
	Item    Item    `gorm:"foreignKey:ItemID;references:item_id" json:"item"`
	Invoice Invoice `gorm:"foreignKey:InvoiceID;references:invoice_id" json:"invoice"`

	// Optional: Uncomment jika lo pakai summary table
	Summary Summary `gorm:"foreignKey:SummaryID;references:summary_id" json:"summary"`
}
