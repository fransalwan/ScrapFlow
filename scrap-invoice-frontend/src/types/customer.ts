export interface Customer {
  id: number
  name: string
  phone: string
  email: string
  address: string
  tier: 'silver' | 'gold' | 'platinum'
}