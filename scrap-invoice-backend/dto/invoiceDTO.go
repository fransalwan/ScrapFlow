package dto

import (
	"scrap-invoice-backend/models"
	"time"
)

type InvoiceInput struct {
	CustomerID    int     `json:"customer_id" binding:"required"`
	InvoiceDate   string  `json:"invoice_date" binding:"required"`
	Status        *string `json:"status"`
	PaymentMethod *string `json:"payment_method"`
	Note          *string `json:"note"`
}

func (i *InvoiceInput) ToModel(createdAt time.Time, invoiceDate time.Time, invoiceNumber string) models.Invoice {
	invoice := models.Invoice{
		CustomerID:    i.CustomerID,
		InvoiceNumber: invoiceNumber,
		CreatedAt:     createdAt, // ← from time.Now()
		InvoiceDate:   invoiceDate,
		Status:        "draft",
	}
	if i.Status != nil {
		invoice.Status = *i.Status
	}
	if i.PaymentMethod != nil {
		invoice.PaymentMethod = *i.PaymentMethod
	}
	if i.Note != nil {
		invoice.Note = *i.Note
	}
	return invoice
}
