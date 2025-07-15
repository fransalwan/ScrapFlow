package models

import "time"

type Invoice struct {
	ID            int     `gorm:"column:invoice_id;primaryKey" json:"id"`
	InvoiceNumber string  `gorm:"column:invoice_number" json:"invoice_number"`
	CustomerID    int     `gorm:"column:customer_id" json:"customer_id"`
	TotalWeight   float64 `gorm:"column:total_weight" json:"total_weight"`
	TotalPrice    float64 `gorm:"column:total_price" json:"total_price"`
	PaymentMethod string  `gorm:"column:payment_method" json:"payment_method"`
	Note          string  `gorm:"column:note" json:"note"`
	Status        string  `gorm:"column:status" json:"status"`
	CreatedBy     string  `gorm:"column:created_by" json:"created_by"`

	// Field baru
	InvoiceDate time.Time `gorm:"column:invoice_date" json:"invoice_date"`

	CreatedAt time.Time `gorm:"column:created_at" json:"-"` // DB only
	UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`

	Customer  Customer  `gorm:"foreignKey:CustomerID" json:"customer"`
	Summaries []Summary `gorm:"foreignKey:InvoiceID" json:"summaries"`
}
