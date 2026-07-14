export interface Customer {
  id: number
  name: string
  phone: string
  email: string
  address: string
  tier: 'silver' | 'gold' | 'platinum'
  created_at?: string
  updated_at?: string
}

export interface CustomerForm {
  name: string
  phone: string
  email: string
  address: string
  tier: 'silver' | 'gold' | 'platinum'
}