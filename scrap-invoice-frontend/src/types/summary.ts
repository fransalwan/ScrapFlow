export interface SummaryItem {
  item_id: number
  item_name: string
  price_per_kg: number
  total_weight: number
  sub_total_price: number
}

export interface SummaryResponse {
  data: SummaryItem[]
  grandTotal: number
}