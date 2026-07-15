// src/types/scale.ts
export interface ScaleDetailResponse {
  id: number
  summary_id?: number
  invoice_id?: number // Opsional, karena di Postman nested di dalam "invoice"
  item_id?: number    // Opsional, karena di Postman nested di dalam "item"
  weight: number
  alas_weight: number
  photo: string
  scale_type: string
  created_at: string
  updated_at?: string
  
  // ✅ SESUAIKAN DENGAN OUTPUT POSTMAN
  item: {
    id: number             // Bukan item_id
    name: string           // Bukan item_name
    category?: string | { item_category_name?: string } // Bisa string langsung atau object
  }
  invoice: {
    id: number             // Bukan invoice_id (berdasarkan Postman: "id": 2)
    invoice_number: string
    customer_id?: number
    total_weight?: number
    total_price?: number
    payment_method?: string
    note?: string
    created_at?: string
  }
}


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

export interface ScaleDetailPayload {
  item_id: number
  weight: number
  alas_weight: number
  photo: string
  scale_type: string
}