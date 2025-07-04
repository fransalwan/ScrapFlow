package models

import "time"

type Invoice struct {
	ID            int       `gorm:"primaryKey" json:"id"`
	InvoiceNumber string    `gorm:"column:invoice_number" json:"invoice_number"`
	CustomerID    int       `gorm:"column:customer_id" json:"customer_id"`
	CreatedBy     string    `gorm:"column:created_by" json:"created_by"`
	TotalWeight   float64   `gorm:"column:total_weight" json:"total_weight"`
	TotalPrice    float64   `gorm:"column:total_price" json:"total_price"`
	PaymentMethod string    `gorm:"column:payment_method" json:"payment_method"`
	Note          string    `gorm:"column:note" json:"note"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	Summaries     []Summary `gorm:"foreignKey:InvoiceID" json:"summaries"`
}

type Summary struct {
	ID           int `gorm:"primaryKey" json:"id"`
	InvoiceID    int `gorm:"column:invoice_id" json:"invoice_id"`
	ItemID       int `gorm:"column:item_id" json:"item_id"`
	TotalWeight  float64
	Subtotal     float64
	ScaleDetails []ScaleDetail `gorm:"foreignKey:SummaryID"`
}

type ScaleDetail struct {
	ID         int
	SummaryID  int
	ItemID     int
	Weight     float64
	AlasWeight string
	ScaleType  string
	Photo      string
}
