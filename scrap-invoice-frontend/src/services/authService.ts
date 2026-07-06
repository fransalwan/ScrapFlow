// src/services/authService.ts
import api from './api'  // <-- Pastikan path-nya bener
import type { LoginRequest, LoginResponse } from '../types/user'

export const authService = {
  login: async (credentials: LoginRequest): Promise<LoginResponse> => {
    console.log('📡 Sending login request to /login')
    const response = await api.post<LoginResponse>('/login', credentials)
    console.log('📥 Login response received:', response.data)
    return response.data
  },
}