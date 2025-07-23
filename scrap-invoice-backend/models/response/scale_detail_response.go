package response

import "time"

type ScaleDetailResponse struct {
	ID         int         `json:"id"`
	Invoice    InvoiceInfo `json:"invoice"`
	Item       ItemInfo    `json:"item"`
	Weight     float64     `json:"weight"`
	AlasWeight float64     `json:"alas_weight"`
	Photo      string      `json:"photo"`
	ScaleType  string      `json:"scale_type"`
	CreatedAt  time.Time   `json:"created_at"`
}

type InvoiceInfo struct {
	ID            int    `json:"id"`
	InvoiceNumber string `json:"invoice_number"`
}

type ItemInfo struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}
