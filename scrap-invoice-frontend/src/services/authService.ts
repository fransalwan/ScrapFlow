// src/services/authService.ts
import api from './api'  // <-- Pastikan path-nya bener
import type { LoginRequest, LoginResponse } from '../types/user'

export const authService = {
  login: async (credentials: LoginRequest): Promise<LoginResponse> => {
    const response = await api.post<LoginResponse>('/login', credentials)
    return response.data
  },
}