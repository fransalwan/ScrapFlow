import axios from 'axios'
import type { SummaryItem } from '../types/summary'
const BASE_URL = 'http://localhost:8080/api'

export const getSummaryByInvoiceId = async (id: number): Promise<SummaryItem[]> => {
  try {
    const response = await axios.get(`${BASE_URL}/invoice/${id}/summary`)
    console.log('Summary response:', response.data)
    return response.data 
  } catch (error) {
    console.error('Failed to fetch summary:', error)
    return []
  }
}