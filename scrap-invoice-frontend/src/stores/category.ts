// stores/category.ts
import { defineStore } from 'pinia'
import axios from 'axios'

import type { Category } from '../types/category'


export const useCategoryStore = defineStore('category', {
  state: () => ({
    categories: [] as Category[],
  }),
  actions: {
    async fetchCategories() {
      const res = await axios.get('http://localhost:8080/api/categories')
      this.categories = res.data.data
    },

    async createCategory(data: { item_category_name: string }) {
      const res = await axios.post('http://localhost:8080/api/category', data)
      this.categories.push(res.data.data)
    },

    async updateCategory(id: number, data: { item_category_name: string }) {
      const res = await axios.put(`http://localhost:8080/api/category/${id}`, data)
      const idx = this.categories.findIndex((c) => c.id === id)
      if (idx !== -1) this.categories[idx] = res.data.data
    },

    async deleteCategory(id: number) {
      await axios.delete(`http://localhost:8080/api/category/${id}`)
      this.categories = this.categories.filter((c) => c.id !== id)
    },
  },
})
