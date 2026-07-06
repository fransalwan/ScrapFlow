// src/stores/category.ts
import { defineStore } from 'pinia'
import { ref } from 'vue'
import api from '../services/api'

// Type Category - di-export biar bisa dipake di komponen
export interface Category {
  id: number
  item_category_name: string
  created_at?: string
  updated_at?: string
}

export const useCategoryStore = defineStore('category', () => {
  // State
  const categories = ref<Category[]>([])
  const loading = ref(false)
  const error = ref<string | null>(null)

  // Actions
  const fetchCategories = async () => {
    loading.value = true
    error.value = null
    try {
      const res = await api.get('/categories')
      // Handle response dari backend (biasanya format { data: [...] })
      categories.value = res.data.data || res.data || []
    } catch (err: any) {
      console.error('Fetch categories failed:', err)
      error.value = err.response?.data?.error || 'Failed to fetch categories'
      throw err
    } finally {
      loading.value = false
    }
  }

  const createCategory = async (payload: { item_category_name: string }) => {
    try {
      const res = await api.post('/categories', payload)
      return res.data
    } catch (err: any) {
      console.error('Create category failed:', err)
      throw err
    }
  }

  const updateCategory = async (id: number, payload: { item_category_name: string }) => {
    try {
      const res = await api.put(`/categories/${id}`, payload)
      return res.data
    } catch (err: any) {
      console.error('Update category failed:', err)
      throw err
    }
  }

  const deleteCategory = async (id: number) => {
    try {
      const res = await api.delete(`/categories/${id}`)
      return res.data
    } catch (err: any) {
      console.error('Delete category failed:', err)
      throw err
    }
  }

  return {
    categories,
    loading,
    error,
    fetchCategories,
    createCategory,
    updateCategory,
    deleteCategory,
  }
})