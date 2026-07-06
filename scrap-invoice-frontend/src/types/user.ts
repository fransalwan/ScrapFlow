export interface User {
  user_id: number
  name: string
  email: string
  password?: string
  role: 'admin' | 'staff' | 'operator'
  created_at: string
  updated_at: string
}

export interface LoginRequest {
  email: string
  password: string
}

export interface LoginResponse {
  message: string
  token: string
  user: User
}