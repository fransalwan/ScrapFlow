// stores/item.ts
import { defineStore } from 'pinia'
import axios from '../lib/axios'
import type { Item, ItemFormInput, ItemFormPayload } from '../types/item'

export const useItemStore = defineStore('item', {
  state: () => ({
    items: [] as Item[],
  }),

  actions: {
    async fetchItems() {
      try {
        const res = await axios.get('/items')
        this.items = res.data.data
      } catch (err) {
        console.error('Failed to fetch items:', err)
        throw err
      }
    },

    async createItem(data: ItemFormPayload) {
      try {
        await axios.post('/item', data)
        await this.fetchItems()
      } catch (err) {
        console.error('Failed to create item:', err)
        throw err
      }
    },

    async updateItem(id: number, data: ItemFormPayload) {
      try {
        await axios.put(`/item/${id}`, data)
        await this.fetchItems()
      } catch (err) {
        console.error(`Failed to update item with id ${id}:`, err)
        throw err
      }
    },

    async deleteItem(id: number) {
      try {
        await axios.delete(`/item/${id}`)
        await this.fetchItems()
      } catch (err) {
        console.error(`Failed to delete item with id ${id}:`, err)
        throw err
      }
    },
  },
})
