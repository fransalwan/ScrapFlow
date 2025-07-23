package dto

import "time"

type CreateScaleDetailInput struct {
	ItemID     int     `json:"item_id" binding:"required"`
	Weight     float64 `json:"weight" binding:"required"`
	AlasWeight float64 `json:"alas_weight"`
	Photo      string  `json:"photo"`
	ScaleType  string  `json:"scale_type" binding:"required"`
}

type ScaleDetailResponse struct {
	ID         int         `json:"id"`
	Weight     float64     `json:"weight"`
	AlasWeight float64     `json:"alas_weight"`
	Photo      string      `json:"photo"`
	ScaleType  string      `json:"scale_type"`
	CreatedAt  time.Time   `json:"created_at"`
	Invoice    InvoiceInfo `json:"invoice"`
	Item       ItemInfo    `json:"item"`
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
