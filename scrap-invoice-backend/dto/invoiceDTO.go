package dto

import (
	"scrap-invoice-backend/models"
	"time"
)

type InvoiceInput struct {
	CustomerID    int    `json:"customer_id" binding:"required"`
	InvoiceDate   string `json:"invoice_date"`
	Status        string `json:"status"`
	PaymentMethod string `json:"payment_method"`
	Note          string `json:"note"`
}

func (input InvoiceInput) ToModel(now time.Time, invoiceDate time.Time, invoiceNumber string) models.Invoice {
	return models.Invoice{
		CustomerID:    input.CustomerID,
		InvoiceNumber: invoiceNumber,
		InvoiceDate:   invoiceDate,
		Status:        input.Status,
		PaymentMethod: input.PaymentMethod,
		Note:          input.Note,
		TotalWeight:   0, // Default
		TotalPrice:    0, // Default
		CreatedAt:     now,
	}
}
