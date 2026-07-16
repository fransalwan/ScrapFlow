// src/stores/scale.ts
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import api from '../services/api'
import type { ScaleDetailPayload, ScaleDetailUI } from '../types/scale'

export const useScaleStore = defineStore('scale', () => {
  const scaleDetails = ref<ScaleDetailUI[]>([])
  const isLoading = ref(false)
  const error = ref<string | null>(null)
  const activeFilter = ref<string | null>(null)

  // ✅ HELPER YANG SUDAH DIPERBAIKI & ANTI-GAGAL (Tetap sama)
  const transformToUI = (data: any): ScaleDetailUI => {
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
        id: data.invoice?.id || data.invoice?.invoice_id || 0,
        invoice_number: data.invoice?.invoice_number || '',
      },
      item: {
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

  const deleteScaleDetail = async (invoiceId: number, scaleId: number) => {
    try {
      // ✅ URL Nested yang benar
      await api.delete(`/invoices/${invoiceId}/scales/${scaleId}`)
      
      // Hapus dari state lokal
      scaleDetails.value = scaleDetails.value.filter(detail => detail.id !== scaleId)
    } catch (err: any) {
      console.error(`Failed to delete scale detail with id ${scaleId}:`, err)
      throw err
    }
  }

  // ✅ FIX: Terima 3 parameter (invoiceId, scaleId, & payload)
  const updateScaleDetail = async (invoiceId: number, scaleId: number, payload: ScaleDetailPayload) => {
    try {
      // ✅ URL Nested yang benar
      const res = await api.put(`/invoices/${invoiceId}/scales/${scaleId}`, payload)
      const updated = res.data.data || res.data
      const transformed = transformToUI(updated)
      
      // Update di state lokal
      const index = scaleDetails.value.findIndex(s => s.id === scaleId)
      if (index !== -1) {
        scaleDetails.value[index] = transformed
      }
      return transformed
    } catch (err: any) {
      console.error(`Failed to update scale detail with id ${scaleId}:`, err)
      throw err
    }
  }

  const cloneScaleDetailToFI = async (_invoiceId: number) => {
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