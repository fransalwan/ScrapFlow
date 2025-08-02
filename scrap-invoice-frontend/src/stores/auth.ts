import { defineStore } from 'pinia'
import axios from '../lib/axios'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: localStorage.getItem('token') || '',
    user: null as null | { username: string },
  }),

  getters: {
    isLoggedIn: (state) => !!state.token,
  },

  actions: {
    async login({ username, password }: { username: string; password: string }) {
      try {
        const res = await axios.post('http://localhost:8080/api/login', { username, password })
        this.token = res.data.token
        this.user = { username }
        localStorage.setItem('token', this.token)
        return true
      } catch (err) {
        console.error('Login failed:', err)
        return false
      }
    },

    logout() {
      this.token = ''
      this.user = null
      localStorage.removeItem('token')
    },
  },
})
