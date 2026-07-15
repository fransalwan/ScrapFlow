// controllers/scale_detail_controller.go
package controllers

import (
	"fmt"
	"net/http"
	"scrap-invoice-backend/config"
	"scrap-invoice-backend/dto"
	"scrap-invoice-backend/models"
	"scrap-invoice-backend/models/response"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetScaleDetailsByInvoiceID(c *gin.Context) {
	invoiceID := c.Param("id") // pakai nama yang sama kayak di route
	var details []models.ScaleDetail

	if err := config.DB.
		Preload("Item.Category").
		Preload("Invoice").
		Where("invoice_id = ?", invoiceID).
		Find(&details).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Scale details not found"})
		return
	}

	var responseData []response.ScaleDetailResponse
	for _, d := range details {
		responseData = append(responseData, response.ScaleDetailResponse{
			ID:         d.ID,
			Weight:     d.Weight,
			AlasWeight: d.AlasWeight,
			Photo:      d.Photo,
			ScaleType:  d.ScaleType,
			CreatedAt:  d.CreatedAt,
			Invoice: response.InvoiceInfo{
				ID:            d.Invoice.ID,
				InvoiceNumber: d.Invoice.InvoiceNumber,
			},
			Item: response.ItemInfo{
				ItemID:   d.Item.ID,
				ItemName: d.Item.ItemName,
				Category: d.Item.Category.ItemCategoryName,
			},
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": responseData})
}

func CreateScaleDetail(c *gin.Context) {
	// Convert :id
	invoiceID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid invoice ID"})
		return
	}

	// Bind ke DTO
	var input dto.CreateScaleDetailInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Map ke model
	scaleDetail := models.ScaleDetail{
		InvoiceID:  invoiceID,
		ItemID:     input.ItemID,
		Weight:     input.Weight,
		AlasWeight: input.AlasWeight,
		Photo:      input.Photo,
		ScaleType:  input.ScaleType,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if input.ScaleType == "TL" {

		// clone to FI
		clone := scaleDetail
		clone.ScaleType = "FI"
		clone.ID = 0 // biar auto increment
		clone.CreatedAt = time.Now()
		clone.UpdatedAt = time.Now()

		if err := config.DB.Create(&clone).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal kloning scale detail ke FI"})
			return
		}
	}

	// Simpan
	if err := config.DB.Create(&scaleDetail).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal simpan scale detail"})
		return
	}

	// Load relasi
	if err := config.DB.
		Preload("Item.Category").
		Preload("Invoice").
		First(&scaleDetail, scaleDetail.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal load relasi"})
		return
	}

	// Mapping ke response DTO
	response := dto.ScaleDetailResponse{
		ID:         scaleDetail.ID,
		Weight:     scaleDetail.Weight,
		AlasWeight: scaleDetail.AlasWeight,
		Photo:      scaleDetail.Photo,
		ScaleType:  scaleDetail.ScaleType,
		CreatedAt:  scaleDetail.CreatedAt,
		Invoice: dto.InvoiceInfo{
			ID:            scaleDetail.Invoice.ID,
			InvoiceNumber: scaleDetail.Invoice.InvoiceNumber,
		},
		Item: dto.ItemInfo{
			ID:       scaleDetail.Item.ID,
			Name:     scaleDetail.Item.ItemName,
			Category: scaleDetail.Item.Category.ItemCategoryName,
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Scale detail created",
		"data":    response,
	})
}

func UpdateScaleDetail(c *gin.Context) {
	// 1. Ambil parameter dari URL
	invoiceIDParam := c.Param("invoice_id")
	scaleIDParam := c.Param("scale_id")

	// 2. Convert string ke uint (GORM defaultnya uint, kalau model lu int, ganti jadi strconv.Atoi)
	invoiceID, err := strconv.ParseUint(invoiceIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid invoice ID format"})
		return
	}

	scaleID, err := strconv.ParseUint(scaleIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid scale ID format"})
		return
	}

	// 3. Cari data existing berdasarkan scale_id
	var existing models.ScaleDetail
	if err := config.DB.First(&existing, scaleID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Scale detail not found"})
		return
	}

	// 4. 🔒 VALIDASI KEPEMILIKAN (Security)
	// Pastikan scale detail ini benar-benar milik invoice yang diminta di URL
	if existing.InvoiceID != int(invoiceID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Scale detail does not belong to this invoice"})
		return
	}

	// 5. Bind input ke struct lokal (LEBIH AMAN daripada bind ke models.ScaleDetail langsung)
	// Ini mencegah user jahat mengupdate field terlarang seperti InvoiceID atau CreatedAt
	var input struct {
		ItemID     int     `json:"item_id" binding:"required"`
		Weight     float64 `json:"weight" binding:"required,min=0"` // Ganti ke int jika model lu int
		AlasWeight float64 `json:"alas_weight"`
		Photo      string  `json:"photo"`
		ScaleType  string  `json:"scale_type" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// 6. Update hanya field yang diizinkan
	existing.ItemID = input.ItemID
	existing.Weight = input.Weight
	existing.AlasWeight = input.AlasWeight
	existing.Photo = input.Photo
	existing.ScaleType = input.ScaleType
	existing.UpdatedAt = time.Now()

	// 7. Simpan ke database
	if err := config.DB.Save(&existing).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update scale detail: " + err.Error()})
		return
	}

	// 8. Preload ulang untuk keperluan response
	var fullDetail models.ScaleDetail
	if err := config.DB.
		Preload("Item").
		Preload("Item.Category"). // Sesuaikan dengan nama relasi di model lu
		Preload("Invoice").
		First(&fullDetail, scaleID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated detail"})
		return
	}

	// 9. Map ke response (DISESUAIKAN dengan format JSON yang frontend harapkan: "id" & "name")
	// Kita pakai map[string]interface{} biar fleksibel dan nggak perlu ubah struct response yang lama
	responseData := map[string]interface{}{
		"id":          fullDetail.ID,
		"weight":      fullDetail.Weight,
		"alas_weight": fullDetail.AlasWeight,
		"photo":       fullDetail.Photo,
		"scale_type":  fullDetail.ScaleType,
		"created_at":  fullDetail.CreatedAt,
		"invoice": map[string]interface{}{
			"id":             fullDetail.Invoice.ID, // Atau fullDetail.InvoiceID tergantung model
			"invoice_number": fullDetail.Invoice.InvoiceNumber,
		},
		"item": map[string]interface{}{
			"id":       fullDetail.Item.ID,                        // ✅ Frontend expect "id"
			"name":     fullDetail.Item.ItemName,                  // ✅ Frontend expect "name" (sesuaikan dengan field model lu)
			"category": fullDetail.Item.Category.ItemCategoryName, // Sesuaikan dengan field model lu
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Scale detail updated successfully",
		"data":    responseData,
	})
}

func DeleteScaleDetail(c *gin.Context) {
	invoiceIDParam := c.Param("invoice_id")
	scaleIDParam := c.Param("scale_id")

	fmt.Printf("🔍 Menerima request: invoice_id=%s, scale_id=%s\n", invoiceIDParam, scaleIDParam)

	// Convert ke integer
	scaleID, err := strconv.Atoi(scaleIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid scale ID format"})
		return
	}

	invoiceID, _ := strconv.Atoi(invoiceIDParam)

	var detail models.ScaleDetail
	// GORM akan otomatis mencari berdasarkan primary key (scale_detail_id)
	if err := config.DB.First(&detail, scaleID).Error; err != nil {
		fmt.Printf("❌ Data tidak ditemukan di DB untuk ID: %d\n", scaleID)
		c.JSON(http.StatusNotFound, gin.H{"error": "Scale detail not found"})
		return
	}

	// Validasi kepemilikan
	if detail.InvoiceID != invoiceID {
		fmt.Printf("⚠️ Mismatch: Detail punya invoice_id %d, tapi request minta invoice_id %d\n", detail.InvoiceID, invoiceID)
		c.JSON(http.StatusForbidden, gin.H{"error": "Scale detail does not belong to this invoice"})
		return
	}

	// Hapus data
	if err := config.DB.Delete(&detail).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Scale detail deleted successfully"})
}
