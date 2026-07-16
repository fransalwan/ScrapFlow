import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { authService } from '../services/authService'
import type { User, LoginRequest } from '../types/user'

export const useAuthStore = defineStore('auth', () => {
  const token = ref<string | null>(localStorage.getItem('token'))
  const user = ref<User | null>(
    localStorage.getItem('user') ? JSON.parse(localStorage.getItem('user')!) : null
  )

  const isLoggedIn = computed(() => !!token.value)
  const currentUser = computed(() => user.value)
  const userRole = computed(() => user.value?.role || '')
  const isAdmin = computed(() => user.value?.role === 'admin')

  const login = async (email: string, password: string) => {
    try {
      const credentials: LoginRequest = { email, password }
      const response = await authService.login(credentials)
      
      // Simpan ke localStorage
      localStorage.setItem('token', response.token)
      localStorage.setItem('user', JSON.stringify(response.user))
      
      // Update state
      token.value = response.token
      user.value = response.user
      
      return { success: true }
    } catch (error: any) {
      console.error('❌ Login failed:', error.response?.data)
      return { 
        success: false, 
        error: error.response?.data?.error || 'Login failed' 
      }
    }
  }

  const logout = () => {
    token.value = null
    user.value = null
    localStorage.removeItem('token')
    localStorage.removeItem('user')
  }

  return {
    token,
    user,
    isLoggedIn,
    currentUser,
    userRole,
    isAdmin,
    login,
    logout,
  }
})