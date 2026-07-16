import api from './api' 

export const getSummaryByInvoiceId = async (id: number) => {
  try {
    console.log(`Fetching summary for invoice ID: ${id}`)
    
    const response = await api.get(`/invoices/${id}/summary`)
    
    return response.data 
  } catch (error) {
    console.error('Failed to fetch summary:', error)
    
    throw error 
  }
}