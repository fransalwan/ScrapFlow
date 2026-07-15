// src/stores/summaryStore.ts
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { getSummaryByInvoiceId } from '../services/summaryService'
import type { SummaryResponse } from '../types/summary'

export const useSummaryStore = defineStore('summary', () => {
  // ✅ Berikan tipe data eksplisit di sini
  const summaryItems = ref<SummaryResponse | null>(null)
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  const fetchSummary = async (invoiceId: number) => {
    isLoading.value = true
    error.value = null
    try {
      const response = await getSummaryByInvoiceId(invoiceId)
      summaryItems.value = response // Response sudah sesuai tipe SummaryResponse
    } catch (err: any) {
      console.error('Failed to fetch summary:', err)
      error.value = err.response?.data?.error || 'Failed to fetch summary'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  return {
    summaryItems,
    isLoading,
    error,
    fetchSummary,
  }
})