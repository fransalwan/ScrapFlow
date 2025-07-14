// stores/customer.ts
import type { Customer } from '../types/customer'
import { defineStore } from 'pinia'
import axios from '../lib/axios' // pakai instance custom


export const useCustomerStore = defineStore('customer', {
  state: () => ({
    customers: [] as Customer[],
  }),

  actions: {
    async fetchCustomers() {
      try {
        const res = await axios.get('/customers') // baseURL harus diset di axios
        this.customers = res.data.data // asumsi response bentuknya { data: [...] }
      } catch (err) {
        console.error('Failed to fetch customers:', err)
      }
    },

    async createCustomer(data: Omit<Customer, 'id'>) {
      try {
        const res = await axios.post('/customer', data)
        this.customers.push(res.data.data) // tambahkan item baru ke list
      } catch (err) {
        console.error('Failed to create customer:', err)
      }
    },

    async updateCustomer(id: number, data: Omit<Customer, 'id'>) {
      try {
        const res = await axios.put(`/customer/${id}`, data)
        const updated = res.data.data
        const index = this.customers.findIndex(c => c.id === id)
        if (index !== -1) this.customers[index] = updated
      } catch (err) {
        console.error('Failed to update customer:', err)
      }
    },

    async deleteCustomer(id: number) {
      try {
        await axios.delete(`/customer/${id}`)
        this.customers = this.customers.filter(c => c.id !== id)
      } catch (err) {
        console.error('Failed to delete customer:', err)
      }
    },
  },
})
