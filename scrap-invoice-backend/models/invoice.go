package models

import "time"

type Invoice struct {
	ID            int       `gorm:"column:invoice_id;primaryKey" json:"id"`
	InvoiceNumber string    `gorm:"column:invoice_number" json:"invoice_number"`
	CustomerID    int       `gorm:"column:customer_id" json:"customer_id"`
	TotalWeight   float64   `gorm:"column:total_weight" json:"total_weight"`
	TotalPrice    float64   `gorm:"column:total_price" json:"total_price"`
	PaymentMethod string    `gorm:"column:payment_method" json:"payment_method"`
	Note          string    `gorm:"column:note" json:"note"`
	Status        string    `json:"status"`
	CreatedBy     string    `json:"created_by"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Customer      Customer  `gorm:"foreignKey:CustomerID" json:"customer"`
	Summaries     []Summary `gorm:"foreignKey:InvoiceID" json:"summaries"`
}
