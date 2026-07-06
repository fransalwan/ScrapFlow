// src/stores/invoice.ts
import { defineStore } from 'pinia'
import { ref } from 'vue'
import api from '../services/api'
import type { Invoice, InvoiceForm } from '../types/invoice'

export const useInvoiceStore = defineStore('invoice', () => {
  // State - mulai kosong, hanya diisi dari database
  const invoices = ref<Invoice[]>([])
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  // Actions
  const fetchInvoices = async () => {
    isLoading.value = true
    error.value = null
    try {
      const res = await api.get('/invoices')
      // Ambil data dari response backend, tanpa dummy data
      invoices.value = res.data.data || res.data || []
    } catch (err: any) {
      console.error('Failed to fetch invoices:', err)
      error.value = err.response?.data?.error || 'Failed to fetch invoices'
      throw err
    } finally {
      isLoading.value = false
    }
  }

 const createInvoice = async (payload: InvoiceForm) => {
  try {
    console.log('📡 Store sending to API:', payload)
    console.log('📡 customer_id type:', typeof payload.customer_id)
    
    const res = await api.post('/invoices', payload)
    const newInvoice = res.data.data || res.data
    
    invoices.value.unshift(newInvoice)
    return newInvoice
  } catch (err: any) {
    console.error('Failed to create invoice:', err)
    console.error('Error response:', err.response?.data)
    throw err
  }
}

  const updateInvoice = async (id: number, payload: InvoiceForm) => {
    try {
      const res = await api.put(`/invoices/${id}`, payload)
      const updated = res.data.data || res.data
      
      // Update di array
      const idx = invoices.value.findIndex((inv) => inv.invoice_id === id)
      if (idx !== -1) {
        invoices.value[idx] = updated
      }
      return updated
    } catch (err: any) {
      console.error(`Failed to update invoice with id ${id}:`, err)
      throw err
    }
  }

  const deleteInvoice = async (id: number) => {
    try {
      await api.delete(`/invoices/${id}`)
      // Hapus dari array
      invoices.value = invoices.value.filter((inv) => inv.invoice_id !== id)
    } catch (err: any) {
      console.error(`Failed to delete invoice with id ${id}:`, err)
      throw err
    }
  }

  const getTodayInvoiceCount = (): number => {
    const today = new Date().toISOString().slice(0, 10)
    return invoices.value.filter((inv) =>
      inv.created_at?.startsWith(today)
    ).length
  }

  return {
    invoices,
    isLoading,
    error,
    fetchInvoices,
    createInvoice,
    updateInvoice,
    deleteInvoice,
    getTodayInvoiceCount,
  }
})