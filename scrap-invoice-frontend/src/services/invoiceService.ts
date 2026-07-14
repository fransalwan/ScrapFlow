import type { Invoice, InvoiceForm } from '../types/invoice'
import axios from 'axios'

const BASE_URL = 'http://localhost:8080/api'

export async function fetchInvoicesAPI(): Promise<Invoice[]> {
  const res = await axios.get(`${BASE_URL}/invoices`)
  console.log("Fetched invoices:", res.data.data) 
  return res.data.data 
}

export async function createInvoiceAPI(payload: InvoiceForm): Promise<Invoice> {
  const res = await axios.post(`${BASE_URL}/invoices`, payload)
  return res.data.data 
}

export async function updateInvoiceAPI(id: number, payload: InvoiceForm): Promise<Invoice> {
  const res = await axios.put(`${BASE_URL}/invoices/${id}`, payload)
  return res.data.data
}

export async function deleteInvoiceAPI(id: number): Promise<void> {
  console.log("Deleting invoice with ID:", id)
  await axios.delete(`${BASE_URL}/invoices/${id}`)
}

