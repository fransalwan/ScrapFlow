import type { Customer } from './customer'

export interface Invoice {
  id: number
  invoice_number: string
  customer: Pick<Customer, 'id' | 'name'>
  invoice_date: string                      // tanggal invoice (visible di UI)
  created_at: string                        // tanggal buat invoice (opsional di UI, penting di logic)
  status: 'draft' | 'paid' | 'unpaid' | 'finalized'
  payment_method?: string
  note?: string
}

export interface InvoiceForm {
  customer_id: number
  invoice_date: string         // ← dikirim ke backend sebagai invoice_date
  status?: 'draft' | 'paid' | 'unpaid' | 'finalized'
  payment_method?: string
  note?: string
}
