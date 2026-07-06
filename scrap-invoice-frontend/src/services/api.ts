import axios from 'axios'

const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL || 'http://localhost:8080/api',
  headers: {
    'Content-Type': 'application/json',
  },
})

// Request interceptor - TEMPELKAN TOKEN
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    
    // DEBUG: Log request data
    if (config.method === 'post' || config.method === 'put') {
      console.log('🌐 API Request:', {
        url: config.url,
        method: config.method,
        data: config.data,
      })
    }
    
    return config
  },
  (error) => Promise.reject(error)
)

// Response interceptor - HANDLE 401
api.interceptors.response.use(
  (response) => response,
  (error) => {
    console.error('❌ API Error:', error.response?.status, error.response?.data)
    
    if (error.response?.status === 401) {
      console.log('🚫 Token invalid/expired, clearing storage...')
      localStorage.removeItem('token')
      localStorage.removeItem('user')
      window.location.href = '/login'
    }
    
    return Promise.reject(error)
  }
)

export default api