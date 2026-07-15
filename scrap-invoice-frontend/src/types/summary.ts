// src/types/summary.ts

export interface SummaryItem {
  item_id: number
  item_name: string
  price_per_kg: number
  total_weight: number
  sub_total_price: number
}

export interface SummaryData {
  [scaleType: string]: SummaryItem[] // Contoh: { "FI": [...], "TL": [...] }
}

export interface SummaryResponse {
  data: SummaryData
  grandTotal: Record<string, number> // Contoh: { "FI": 300000, "TL": 50000 }
}