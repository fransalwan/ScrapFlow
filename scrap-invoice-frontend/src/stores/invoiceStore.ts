import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { Invoice, InvoiceForm } from '../types/invoice'
import {
  fetchInvoicesAPI,
  createInvoiceAPI,
  updateInvoiceAPI,
  deleteInvoiceAPI,
} from '../services/invoiceService'
import { formatInvoiceDate } from '../utils/format'

export const useInvoiceStore = defineStore('invoice', () => {
  const invoices = ref<Invoice[]>([])
  const isLoading = ref(false)

  async function fetchInvoices() {
    isLoading.value = true
    try {
      const rawInvoices = await fetchInvoicesAPI()
      invoices.value = rawInvoices.map((inv) => formatInvoiceDate(inv))
    } finally {
      isLoading.value = false
    }
  }

  async function createInvoice(payload: InvoiceForm) {
    // payload hanya berisi invoice_date, bukan created_at
    const newInvoice = await createInvoiceAPI(payload)
    invoices.value.unshift(formatInvoiceDate(newInvoice))
  }

  async function updateInvoice(id: number, payload: InvoiceForm) {
    const updated = await updateInvoiceAPI(id, payload)
    console.log(updated, "<<<<<")
    const idx = invoices.value.findIndex((inv) => inv.id === id)
    if (idx !== -1) invoices.value[idx] = formatInvoiceDate(updated)
  }

  async function deleteInvoice(id: number) {
    await deleteInvoiceAPI(id)
    invoices.value = invoices.value.filter((inv) => inv.id !== id)
  }

  function getTodayInvoiceCount(): number {
    const today = new Date().toISOString().slice(0, 10)
    return invoices.value.filter((inv) =>
      inv.created_at?.startsWith(today)
    ).length
  }

  return {
    invoices,
    isLoading,
    fetchInvoices,
    createInvoice,
    updateInvoice,
    deleteInvoice,
    getTodayInvoiceCount,
  }
})
