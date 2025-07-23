// dto/scale_summary.go
package dto

type ScaleSummary struct {
	ItemID      int     `json:"item_id"`
	ItemName    string  `json:"item_name"`
	Category    string  `json:"category"`
	TotalWeight float64 `json:"total_weight"`
	TotalAlas   float64 `json:"total_alas"`
	NetWeight   float64 `json:"net_weight"`
}
