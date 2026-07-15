import axios from 'axios'
import type { ScaleDetailPayload, ScaleDetailResponse } from '../types/scale.ts'

const BASE_URL = 'http://localhost:8080/api'

export async function fetchScaleDetailsAPI(invoiceId: number): Promise<ScaleDetailResponse[]> {
  const res = await axios.get(`${BASE_URL}/invoices/${invoiceId}/scales`)
  return res.data.data
}

export async function createScaleDetailAPI(invoiceId: number, payload: ScaleDetailPayload) {
  const res = await axios.post(`${BASE_URL}/invoices/${invoiceId}/scales`, payload)
  return res.data.data
}

export async function updateScaleDetailAPI(id: number, payload: any) {
  const res = await axios.put(`${BASE_URL}/invoices/${id}/scales`, payload)
  console.log('Updated scale detail:', res.data.data)
  return res.data.data
}

export async function deleteScaleDetailAPI(invoiceId: number): Promise<void> {
  await axios.delete(`${BASE_URL}/invoices/${invoiceId}/scales`)
}