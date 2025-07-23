import { defineStore } from 'pinia'
import type { SummaryItem } from '../types/summary'
import { getSummaryByInvoiceId } from '../services/summaryService'

export const useSummaryStore = defineStore('summary', {
  state: () => ({
    summaryItems: [] as SummaryItem[],
    isLoading: false,
    error: null as string | null,
  }),

  actions: {
    async fetchSummary(invoiceId: number) {
      this.isLoading = true
      this.error = null
      try {
        this.summaryItems = await getSummaryByInvoiceId(invoiceId)
      } catch (err: any) {
        this.error = err.message || 'Unknown error'
      } finally {
        this.isLoading = false
      }
    },

    clearSummary() {
      this.summaryItems = []
    }
  }
})
