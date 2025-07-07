// src/stores/invoiceStore.js
import { defineStore } from 'pinia'
import axios from 'axios'


export const useInvoiceStore = defineStore('invoice', {
  state: () => ({
    invoices: [],
    isLoading: false,
    error: null
  }),

  actions: {
    async fetchInvoices() {
      console.log("Fetching Invoices...")
      this.isLoading = true
      try {
        const res = await axios.get('http://localhost:8080/api/invoices')
        console.log("Fetched Invoices:", res.data.data)
        this.invoices = res.data.data
        this.error = null
      } catch (err) {
        this.error = err.message || 'Failed to fetch invoices'
      } finally {
        this.isLoading = false
      }
    }
  }
})
