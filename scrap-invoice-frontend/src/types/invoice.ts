// src/types/invoice.ts
export interface Invoice {
  id: number
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
  invoice_date: string
  status?: 'draft' | 'paid' | 'unpaid' | 'finalized'
  payment_method?: string
  note?: string
}