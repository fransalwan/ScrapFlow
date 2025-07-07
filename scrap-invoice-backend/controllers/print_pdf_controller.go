package controllers

import (
	"bytes"
	"fmt"
	"net/http"
	"scrap-invoice-backend/config"
	"scrap-invoice-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/phpdave11/gofpdf"
)

func GetInvoicePDF(c *gin.Context) {
	id := c.Param("id")
	var invoice models.Invoice

	err := config.DB.
		Preload("Customer").
		Preload("Summaries.Item.Category").
		First(&invoice, id).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
		return
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	// Header
	pdf.Cell(40, 10, "INVOICE")
	pdf.Ln(12)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("Invoice #: %s", invoice.InvoiceNumber))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Customer: %s", invoice.Customer.Name))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Date: %s", invoice.CreatedAt.Format("2006-01-02")))
	pdf.Ln(12)

	// Table Headers
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(60, 10, "Item", "1", 0, "", false, 0, "")
	pdf.CellFormat(40, 10, "Category", "1", 0, "", false, 0, "")
	pdf.CellFormat(40, 10, "Weight (kg)", "1", 0, "", false, 0, "")
	pdf.CellFormat(40, 10, "Price", "1", 1, "", false, 0, "")

	// Table Data
	pdf.SetFont("Arial", "", 12)
	for _, s := range invoice.Summaries {
		pdf.CellFormat(60, 10, s.Item.ItemName, "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 10, s.Item.Category.ItemCategoryName, "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 10, fmt.Sprintf("%.2f", s.TotalWeight), "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 10, fmt.Sprintf("Rp %.2f", s.SubTotalPrice), "1", 1, "", false, 0, "")
	}

	// Footer
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Total Weight: %.2f kg", invoice.TotalWeight))
	pdf.Ln(6)
	pdf.Cell(40, 10, fmt.Sprintf("Total Price: Rp %.2f", invoice.TotalPrice))
	pdf.Ln(6)
	pdf.Cell(40, 10, fmt.Sprintf("Payment: %s", invoice.PaymentMethod))
	pdf.Ln(6)
	pdf.Cell(40, 10, fmt.Sprintf("Note: %s", invoice.Note))

	// Save to buffer
	var buf bytes.Buffer
	err = pdf.Output(&buf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate PDF"})
		return
	}

	// Send PDF as response
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "inline; filename=invoice.pdf")
	c.Data(http.StatusOK, "application/pdf", buf.Bytes())
}
