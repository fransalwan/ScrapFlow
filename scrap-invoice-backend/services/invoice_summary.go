// Package services berisi logika bisnis murni (tanpa HTTP, tanpa LLM).
// Controller, PDF, dan AI agent wajib memakai fungsi di sini supaya angka
// yang dilihat user selalu berasal dari satu sumber perhitungan.
package services

import (
	"sort"

	"scrap-invoice-backend/models"
)

// Tipe timbangan. Lihat spec/requirements.md bagian "Tipe timbangan".
const (
	ScaleTypeFI = "FI" // Full Invoice. Input TL otomatis disalin menjadi FI.
	ScaleTypeTL = "TL" // Timbang Lapangan
	ScaleTypeTG = "TG" // Timbang Gudang
	ScaleTypeTS = "TS" // Timbang Sortir
)

// ScaleTypeOrder adalah urutan tampil yang konsisten di UI dan agent.
var ScaleTypeOrder = []string{ScaleTypeFI, ScaleTypeTL, ScaleTypeTG, ScaleTypeTS}

// ScaleTypeNames memetakan kode ke nama lengkap.
var ScaleTypeNames = map[string]string{
	ScaleTypeFI: "Full Invoice",
	ScaleTypeTL: "Timbang Lapangan",
	ScaleTypeTG: "Timbang Gudang",
	ScaleTypeTS: "Timbang Sortir",
}

// PriceBasis menjelaskan cara subtotal dihitung SAAT INI, dalam bahasa manusia.
// Wajib diubah bersamaan dengan SummarizeScaleDetails. Agent menampilkan teks
// ini supaya user tahu dasar angkanya selama aturan harga belum diverifikasi
// (spec/known-issues.md #1 dan #2).
const PriceBasis = "Subtotal = berat timbang × harga master item saat ini. " +
	"Berat alas belum dikurangi, dan harga saat transaksi tidak disimpan."

// ItemSummary adalah ringkasan satu item dalam satu tipe timbangan.
// Tag JSON item_id, item_name, price_per_kg, total_weight, sub_total_price
// sama persis dengan response lama supaya frontend tidak perlu diubah.
type ItemSummary struct {
	ItemID        int     `json:"item_id"`
	ItemName      string  `json:"item_name"`
	PricePerKg    float64 `json:"price_per_kg"`
	TotalWeight   float64 `json:"total_weight"`
	TotalAlas     float64 `json:"total_alas"`
	NetWeight     float64 `json:"net_weight"`
	SubTotalPrice float64 `json:"sub_total_price"`
}

// InvoiceSummary adalah bentuk response GET /api/invoices/:id/summary.
type InvoiceSummary struct {
	Data       map[string][]ItemSummary `json:"data"`
	GrandTotal map[string]float64       `json:"grandTotal"`
}

// SummarizeScaleDetails mengelompokkan scale detail per tipe timbangan lalu per item.
// details harus sudah Preload("Item").
//
// PERILAKU SAAT INI, dipindah apa adanya dari summary_controller.go.
// Jangan diubah tanpa memperbarui spec/requirements.md dan test-nya:
//   - subtotal dihitung dari Weight; AlasWeight TIDAK mengurangi berat yang dihargai
//   - harga diambil dari harga master item saat ini (sd.Item.PricePerKg)
//
// TotalAlas dan NetWeight hanya informasi, belum dipakai untuk harga.
func SummarizeScaleDetails(details []models.ScaleDetail) InvoiceSummary {
	byType := map[string]map[int]*ItemSummary{}
	grand := map[string]float64{}

	for _, sd := range details {
		items, ok := byType[sd.ScaleType]
		if !ok {
			items = map[int]*ItemSummary{}
			byType[sd.ScaleType] = items
		}
		s, ok := items[sd.ItemID]
		if !ok {
			s = &ItemSummary{ItemID: sd.ItemID, ItemName: sd.Item.ItemName, PricePerKg: sd.Item.PricePerKg}
			items[sd.ItemID] = s
		}
		s.TotalWeight += sd.Weight
		s.TotalAlas += sd.AlasWeight
		s.NetWeight = s.TotalWeight - s.TotalAlas
		s.SubTotalPrice = s.TotalWeight * s.PricePerKg
		grand[sd.ScaleType] += sd.Weight * s.PricePerKg
	}

	out := InvoiceSummary{Data: make(map[string][]ItemSummary, len(byType)), GrandTotal: grand}
	for scaleType, items := range byType {
		rows := make([]ItemSummary, 0, len(items))
		for _, s := range items {
			rows = append(rows, *s)
		}
		// Dulu urutan acak karena iterasi map; sekarang stabil.
		sort.Slice(rows, func(i, j int) bool { return rows[i].ItemID < rows[j].ItemID })
		out.Data[scaleType] = rows
	}
	return out
}

// CountInvoicesPerScaleType menghitung jumlah invoice berbeda per tipe timbangan.
func CountInvoicesPerScaleType(details []models.ScaleDetail) map[string]int {
	seen := map[string]map[int]struct{}{}
	for _, sd := range details {
		if seen[sd.ScaleType] == nil {
			seen[sd.ScaleType] = map[int]struct{}{}
		}
		seen[sd.ScaleType][sd.InvoiceID] = struct{}{}
	}
	counts := make(map[string]int, len(seen))
	for t, ids := range seen {
		counts[t] = len(ids)
	}
	return counts
}
