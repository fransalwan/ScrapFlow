// stores/dashboard.ts
import { defineStore } from 'pinia'
import axios from '../lib/axios'

export const useDashboardStore = defineStore('dashboard', {
  state: () => ({
    invoiceCount: null as number | null,
    customerCount: null as number | null,
  }),

  actions: {
    async fetchDashboardCounts() {
      try {
        const [invoiceRes, customerRes] = await Promise.all([
          axios.get('http://localhost:8080/api/dashboard/invoices/count'),
          axios.get('http://localhost:8080/api/dashboard/customers/count'),
        ])
        this.invoiceCount = invoiceRes.data.count
        this.customerCount = customerRes.data.count
      } catch (err) {
        console.error('Failed to fetch dashboard counts:', err)
      }
    },
  },
})
