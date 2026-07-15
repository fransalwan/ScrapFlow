// src/stores/scale.ts
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import api from '../services/api'
import type { ScaleDetailResponse, ScaleDetailPayload, ScaleDetailUI } from '../types/scale'

export const useScaleStore = defineStore('scale', () => {
  const scaleDetails = ref<ScaleDetailUI[]>([])
  const isLoading = ref(false)
  const error = ref<string | null>(null)
  const activeFilter = ref<string | null>(null)

  // ✅ HELPER YANG SUDAH DIPERBAIKI & ANTI-GAGAL
  const transformToUI = (data: any): ScaleDetailUI => {
    // Ambil category, handle kalau backend kirim string "Besi" atau object { item_category_name: "Besi" }
    const rawCategory = data.item?.category
    const categoryName = typeof rawCategory === 'string' 
      ? rawCategory 
      : (rawCategory?.item_category_name || '')

    return {
      id: data.id,
      weight: data.weight,
      alas_weight: data.alas_weight,
      photo: data.photo,
      scale_type: data.scale_type,
      created_at: data.created_at,
      invoice: {
        // Fallback: coba invoice.id, kalau tidak ada coba invoice.invoice_id
        id: data.invoice?.id || data.invoice?.invoice_id || 0,
        invoice_number: data.invoice?.invoice_number || '',
      },
      item: {
        // ✅ PRIORITAS: Baca format Postman (id & name), fallback ke format lama (item_id & item_name)
        id: data.item?.id || data.item?.item_id || 0,
        name: data.item?.name || data.item?.item_name || 'Unknown Item',
        category: categoryName,
      },
    }
  }

  const fetchScaleDetails = async (invoiceId: number) => {
    isLoading.value = true
    error.value = null
    try {
      const res = await api.get(`/invoices/${invoiceId}/scales`)
      const rawData = res.data.data || res.data || []
      scaleDetails.value = rawData.map(transformToUI)
    } catch (err: any) {
      console.error('Failed to fetch scale details:', err)
      error.value = err.response?.data?.error || 'Failed to fetch scale details'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const createScaleDetail = async (invoiceId: number, payload: ScaleDetailPayload) => {
    try {
      const res = await api.post(`/invoices/${invoiceId}/scales`, payload)
      const newDetail = res.data.data || res.data
      const transformed = transformToUI(newDetail)
      scaleDetails.value.push(transformed)
      return transformed
    } catch (err: any) {
      console.error('Failed to create scale detail:', err)
      throw err
    }
  }

  const deleteScaleDetail = async (id: number) => {
    try {
      await api.delete(`/scales/${id}`) // Sesuaikan endpoint jika perlu
      scaleDetails.value = scaleDetails.value.filter(detail => detail.id !== id)
    } catch (err: any) {
      console.error(`Failed to delete scale detail with id ${id}:`, err)
      throw err
    }
  }

  const updateScaleDetail = async (id: number, payload: ScaleDetailPayload) => {
    try {
      const res = await api.put(`/scales/${id}`, payload) // Sesuaikan endpoint jika perlu
      const updated = res.data.data || res.data
      const transformed = transformToUI(updated)
      
      const index = scaleDetails.value.findIndex(s => s.id === id)
      if (index !== -1) {
        scaleDetails.value[index] = transformed
      }
      return transformed
    } catch (err: any) {
      console.error(`Failed to update scale detail with id ${id}:`, err)
      throw err
    }
  }

  const cloneScaleDetailToFI = async (invoiceId: number) => {
    const tlItems = scaleDetails.value.filter(detail => detail.scale_type === 'TL')
    scaleDetails.value = scaleDetails.value.filter(detail => detail.scale_type !== 'FI')

    const cloned: ScaleDetailUI[] = tlItems.map(detail => ({
      ...detail,
      id: Date.now() + Math.random(), 
      scale_type: 'FI',
    }))

    scaleDetails.value.push(...cloned)
  }

  const setFilter = (filter: string | null) => {
    activeFilter.value = filter
  }

  const filteredScaleDetails = computed(() => {
    if (!activeFilter.value) return scaleDetails.value
    return scaleDetails.value.filter(item => item.scale_type === activeFilter.value)
  })

  return {
    scaleDetails,
    isLoading,
    error,
    fetchScaleDetails,
    createScaleDetail,
    deleteScaleDetail,
    updateScaleDetail,
    cloneScaleDetailToFI,
    activeFilter,
    setFilter,
    filteredScaleDetails,
  }
})