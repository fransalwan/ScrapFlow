// src/types/scale.ts
export interface ScaleDetailResponse {
  id: number
  summary_id?: number
  invoice_id: number
  item_id: number
  weight: number
  alas_weight: number
  photo: string
  scale_type: string
  created_at: string
  updated_at: string
  item: {
    item_id: number
    item_name: string
    item_category_id: number
    price_per_kg: number
    category?: {
      item_category_id: number
      item_category_name: string
    }
  }
  invoice: {
    invoice_id: number
    invoice_number: string
    customer_id: number
    total_weight: number
    total_price: number
    payment_method: string
    note: string
    created_at: string
  }
}

export interface ScaleDetailPayload {
  item_id: number
  weight: number
  alas_weight?: number
  scale_type?: string
  photo?: string
}

// Helper type buat UI yang lebih simple
export interface ScaleDetailUI {
  id: number
  weight: number
  alas_weight: number
  photo: string
  scale_type: string
  created_at: string
  invoice: {
    id: number
    invoice_number: string
  }
  item: {
    id: number
    name: string
    category: string
  }
}