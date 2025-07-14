// Tipe yang datang dari backend (misalnya saat fetch semua item)
export interface Item {
  id: number
  item_name: string
  item_category_id: number // ini biasanya ikut dikirim juga
  price_per_kg: number
  created_at: string
  updated_at: string
  category: {
    item_category_id: number
    item_category_name: string
  }
}

// Tipe untuk input form dari UI
export interface ItemFormInput {
  name: string
  category_id: number | null
  price_per_kg: number | null
}

// Tipe payload untuk API create/update (sesuai dengan struktur backend expect-nya)
export interface ItemFormPayload {
  item_name: string
  item_category: number
  price_per_kg: number
}
