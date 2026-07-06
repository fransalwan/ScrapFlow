import axios from 'axios'
// import { useAuthStore } from '../stores/auth'

const instance = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api',
})

// instance.interceptors.request.use((config) => {
//   const auth = useAuthStore()
//   if (auth.token) {
//     config.headers.Authorization = `Bearer ${auth.token}`
//   }
//   return config
// })

export default instance
