// src/stores/customer.ts
import { defineStore } from 'pinia'
import { ref } from 'vue'
import api from '../services/api'

// Type Customer - di-export biar bisa dipake di komponen
export interface Customer {
  id_customer: number
  name: string
  phone: string
  email: string
  address: string
  tier: string
  created_at?: string
  updated_at?: string
}

export interface CustomerPayload {
  name: string
  phone?: string
  email?: string
  address?: string
  tier?: string
}

export const useCustomerStore = defineStore('customer', () => {
  // State
  const customers = ref<Customer[]>([])
  const loading = ref(false)
  const error = ref<string | null>(null)

  // Actions
  const fetchCustomers = async () => {
    loading.value = true
    error.value = null
    try {
      const res = await api.get('/customers')
      // Handle response format dari backend
      customers.value = res.data.data || res.data || []
    } catch (err: any) {
      console.error('Fetch customers failed:', err)
      error.value = err.response?.data?.error || 'Failed to fetch customers'
      throw err
    } finally {
      loading.value = false
    }
  }

  const createCustomer = async (payload: CustomerPayload) => {
    try {
      const res = await api.post('/customers', payload)
      return res.data
    } catch (err: any) {
      console.error('Create customer failed:', err)
      throw err
    }
  }

  const updateCustomer = async (id: number, payload: CustomerPayload) => {
    try {
      const res = await api.put(`/customers/${id}`, payload)
      return res.data
    } catch (err: any) {
      console.error('Update customer failed:', err)
      throw err
    }
  }

  const deleteCustomer = async (id: number) => {
    try {
      const res = await api.delete(`/customers/${id}`)
      return res.data
    } catch (err: any) {
      console.error('Delete customer failed:', err)
      throw err
    }
  }

  return {
    customers,
    loading,
    error,
    fetchCustomers,
    createCustomer,
    updateCustomer,
    deleteCustomer,
  }
})