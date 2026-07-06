// src/stores/item.ts
import { defineStore } from 'pinia'
import { ref } from 'vue'
import api from '../services/api'
import type { Item, ItemFormPayload } from '../types/item'

export const useItemStore = defineStore('item', () => {
  // State
  const items = ref<Item[]>([])
  const loading = ref(false)
  const error = ref<string | null>(null)

  // Actions
  const fetchItems = async () => {
    loading.value = true
    error.value = null
    try {
      const res = await api.get('/items')
      items.value = res.data.data || res.data || []
    } catch (err: any) {
      console.error('Failed to fetch items:', err)
      error.value = err.response?.data?.error || 'Failed to fetch items'
      throw err
    } finally {
      loading.value = false
    }
  }

  const createItem = async (data: ItemFormPayload) => {
    try {
      const res = await api.post('/items', data)  // <-- Fix: /item → /items
      await fetchItems()
      return res.data
    } catch (err: any) {
      console.error('Failed to create item:', err)
      throw err
    }
  }

  const updateItem = async (id: number, data: ItemFormPayload) => {
    try {
      const res = await api.put(`/items/${id}`, data)  // <-- Fix: /item → /items
      await fetchItems()
      return res.data
    } catch (err: any) {
      console.error(`Failed to update item with id ${id}:`, err)
      throw err
    }
  }

  const deleteItem = async (id: number) => {
    try {
      const res = await api.delete(`/items/${id}`)  // <-- Fix: /item → /items
      await fetchItems()
      return res.data
    } catch (err: any) {
      console.error(`Failed to delete item with id ${id}:`, err)
      throw err
    }
  }

  return {
    items,
    loading,
    error,
    fetchItems,
    createItem,
    updateItem,
    deleteItem,
  }
})