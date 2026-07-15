// ✅ 1. Import instance 'api' yang sudah punya interceptor JWT, BUKAN axios biasa
import api from './api' 
// import type { SummaryItem } from '../types/summary' 

export const getSummaryByInvoiceId = async (id: number) => {
  try {
    console.log(`Fetching summary for invoice ID: ${id}`)
    
    // ✅ 2. Ganti axios.get menjadi api.get
    const response = await api.get(`/invoices/${id}/summary`)
    
    console.log('Summary fetched successfully:', response.data)
    
    // Backend return format: { data: {...}, grandTotal: {...} }
    return response.data 
  } catch (error) {
    console.error('Failed to fetch summary:', error)
    
    // ✅ 3. Lebih baik throw error agar store/component bisa handle (misal: tampilkan toast)
    // Daripada return [] yang bikin error jadi "silent" dan susah di-debug
    throw error 
  }
}