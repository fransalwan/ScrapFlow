// src/types/invoice.ts
export interface Invoice {
  invoice_id: number
  invoice_number: string
  customer_id: number
  invoice_date?: string
  status?: 'draft' | 'paid' | 'unpaid' | 'finalized'
  total_weight: number
  total_price: number
  payment_method: string
  note: string
  created_at: string
  customer?: {
    id_customer: number
    name: string
    phone?: string
    email?: string
    address?: string
    tier?: string
  }
}

export interface InvoiceForm {
  customer_id: number
  invoice_date: string // ← WAJIB ADA
  status?: 'draft' | 'paid' | 'unpaid' | 'finalized' // ← TAMBAHIN
  payment_method?: string
  note?: string
}