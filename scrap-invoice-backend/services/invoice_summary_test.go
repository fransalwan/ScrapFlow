package services

import (
	"math"
	"testing"

	"scrap-invoice-backend/models"
)

func detail(invoiceID, itemID int, scaleType string, weight, alas, price float64) models.ScaleDetail {
	return models.ScaleDetail{
		InvoiceID:  invoiceID,
		ItemID:     itemID,
		ScaleType:  scaleType,
		Weight:     weight,
		AlasWeight: alas,
		Item:       models.Item{ID: itemID, ItemName: "item", PricePerKg: price},
	}
}

func near(a, b float64) bool { return math.Abs(a-b) < 1e-6 }

func TestSummarizeScaleDetails(t *testing.T) {
	tests := []struct {
		name      string
		details   []models.ScaleDetail
		scaleType string
		wantRows  int
		wantFirst ItemSummary
		wantGrand float64
	}{
		{
			// Test ini MENDOKUMENTASIKAN perilaku saat ini (known-issues #1).
			// Kalau aturan diputuskan "harga dari berat bersih", ubah test ini
			// bersamaan dengan kode dan PriceBasis.
			name:      "harga dihitung dari berat timbang, alas tidak mengurangi",
			details:   []models.ScaleDetail{detail(1, 10, ScaleTypeFI, 100, 5, 8800)},
			scaleType: ScaleTypeFI,
			wantRows:  1,
			wantFirst: ItemSummary{ItemID: 10, PricePerKg: 8800, TotalWeight: 100, TotalAlas: 5, NetWeight: 95, SubTotalPrice: 880000},
			wantGrand: 880000,
		},
		{
			name: "beberapa timbangan item yang sama dijumlahkan",
			details: []models.ScaleDetail{
				detail(1, 10, ScaleTypeTG, 40.5, 1, 1000),
				detail(1, 10, ScaleTypeTG, 59.5, 2, 1000),
			},
			scaleType: ScaleTypeTG,
			wantRows:  1,
			wantFirst: ItemSummary{ItemID: 10, PricePerKg: 1000, TotalWeight: 100, TotalAlas: 3, NetWeight: 97, SubTotalPrice: 100000},
			wantGrand: 100000,
		},
		{
			name: "tipe timbangan dipisah, tidak saling menjumlah",
			details: []models.ScaleDetail{
				detail(1, 10, ScaleTypeTL, 100, 0, 1000),
				detail(1, 10, ScaleTypeFI, 100, 0, 1000),
				detail(1, 10, ScaleTypeTS, 90, 0, 1000),
			},
			scaleType: ScaleTypeTS,
			wantRows:  1,
			wantFirst: ItemSummary{ItemID: 10, PricePerKg: 1000, TotalWeight: 90, NetWeight: 90, SubTotalPrice: 90000},
			wantGrand: 90000,
		},
		{
			name: "baris diurutkan berdasarkan item_id",
			details: []models.ScaleDetail{
				detail(1, 30, ScaleTypeFI, 1, 0, 3),
				detail(1, 20, ScaleTypeFI, 1, 0, 2),
				detail(1, 10, ScaleTypeFI, 1, 0, 1),
			},
			scaleType: ScaleTypeFI,
			wantRows:  3,
			wantFirst: ItemSummary{ItemID: 10, PricePerKg: 1, TotalWeight: 1, NetWeight: 1, SubTotalPrice: 1},
			wantGrand: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SummarizeScaleDetails(tt.details)
			rows := got.Data[tt.scaleType]
			if len(rows) != tt.wantRows {
				t.Fatalf("rows = %d, want %d", len(rows), tt.wantRows)
			}
			r, w := rows[0], tt.wantFirst
			if r.ItemID != w.ItemID || !near(r.PricePerKg, w.PricePerKg) || !near(r.TotalWeight, w.TotalWeight) ||
				!near(r.TotalAlas, w.TotalAlas) || !near(r.NetWeight, w.NetWeight) || !near(r.SubTotalPrice, w.SubTotalPrice) {
				t.Errorf("row = %+v, want %+v", r, w)
			}
			if !near(got.GrandTotal[tt.scaleType], tt.wantGrand) {
				t.Errorf("grandTotal[%s] = %v, want %v", tt.scaleType, got.GrandTotal[tt.scaleType], tt.wantGrand)
			}
			for i := 1; i < len(rows); i++ {
				if rows[i-1].ItemID > rows[i].ItemID {
					t.Errorf("rows tidak terurut: %+v", rows)
				}
			}
		})
	}
}

func TestSummarizeScaleDetailsEmpty(t *testing.T) {
	got := SummarizeScaleDetails(nil)
	if got.Data == nil || got.GrandTotal == nil || len(got.Data) != 0 {
		t.Fatalf("invoice tanpa timbangan harus menghasilkan map kosong, bukan nil: %+v", got)
	}
}

func TestCountInvoicesPerScaleType(t *testing.T) {
	got := CountInvoicesPerScaleType([]models.ScaleDetail{
		detail(1, 10, ScaleTypeFI, 1, 0, 1),
		detail(1, 11, ScaleTypeFI, 1, 0, 1),
		detail(2, 10, ScaleTypeFI, 1, 0, 1),
		detail(2, 10, ScaleTypeTG, 1, 0, 1),
	})
	if got[ScaleTypeFI] != 2 || got[ScaleTypeTG] != 1 || len(got) != 2 {
		t.Fatalf("got %v", got)
	}
}
