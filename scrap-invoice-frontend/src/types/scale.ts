export interface ScaleDetailResponse {
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
  id?: number
  item_id: number
  weight: number
  alas_weight?: number
  scale_type?: string
  photo?: string
}
