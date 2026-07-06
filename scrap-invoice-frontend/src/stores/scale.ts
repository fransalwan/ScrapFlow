// src/stores/scale.ts
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import api from '../services/api'
import type { ScaleDetailResponse, ScaleDetailPayload, ScaleDetailUI } from '../types/scale'

export const useScaleStore = defineStore('scale', () => {
  // State
  const scaleDetails = ref<ScaleDetailUI[]>([])
  const isLoading = ref(false)
  const error = ref<string | null>(null)
  const activeFilter = ref<string | null>(null)

  // Helper: Transform backend response ke UI format
  const transformToUI = (item: ScaleDetailResponse): ScaleDetailUI => ({
    id: item.id,
    weight: item.weight,
    alas_weight: item.alas_weight,
    photo: item.photo,
    scale_type: item.scale_type,
    created_at: item.created_at,
    invoice: {
      id: item.invoice.invoice_id,
      invoice_number: item.invoice.invoice_number,
    },
    item: {
      id: item.item.item_id,
      name: item.item.item_name,
      category: item.item.category?.item_category_name || '',
    },
  })

  // Actions
  const fetchScaleDetails = async (invoiceId: number) => {
    isLoading.value = true
    error.value = null
    try {
      const res = await api.get(`/invoices/${invoiceId}/scales`)
      const rawData: ScaleDetailResponse[] = res.data.data || res.data || []
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
      const newDetail: ScaleDetailResponse = res.data.data || res.data
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
      console.log('Deleting scale detail with ID:', id)
      // Backend endpoint: DELETE /api/invoices/:invoiceId/scales
      // Tapi kita perlu invoice_id, jadi ambil dari data yang ada
      const detail = scaleDetails.value.find(d => d.id === id)
      if (!detail) {
        throw new Error(`Scale detail with id ${id} not found`)
      }
      
      // Note: Sesuaikan endpoint ini sama backend lu
      // Kalau backend expect DELETE /api/scales/:id, pakai ini:
      await api.delete(`/scales/${id}`)
      
      // Atau kalau backend expect DELETE /api/invoices/:invoiceId/scales/:scaleId:
      // await api.delete(`/invoices/${detail.invoice.id}/scales/${id}`)
      
      scaleDetails.value = scaleDetails.value.filter(detail => detail.id !== id)
    } catch (err: any) {
      console.error(`Failed to delete scale detail with id ${id}:`, err)
      throw err
    }
  }

  const updateScaleDetail = async (id: number, payload: ScaleDetailPayload) => {
    try {
      const detail = scaleDetails.value.find(d => d.id === id)
      if (!detail) {
        throw new Error(`Scale detail with id ${id} not found`)
      }

      // Note: Sesuaikan endpoint ini sama backend lu
      const res = await api.put(`/scales/${id}`, payload)
      const updated: ScaleDetailResponse = res.data.data || res.data
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

    // Buang semua item FI lama dulu
    scaleDetails.value = scaleDetails.value.filter(detail => detail.scale_type !== 'FI')

    // Clone TL ke FI dengan id unik baru (dummy clone)
    const cloned: ScaleDetailUI[] = tlItems.map(detail => ({
      ...detail,
      id: Date.now() + Math.random(), // id dummy
      scale_type: 'FI',
    }))

    scaleDetails.value.push(...cloned)
    
    // Note: Kalau mau sync ke backend, uncomment ini
    // for (const item of cloned) {
    //   const payload: ScaleDetailPayload = {
    //     item_id: item.item.id,
    //     weight: item.weight,
    //     alas_weight: item.alas_weight,
    //     scale_type: 'FI',
    //     photo: item.photo,
    //   }
    //   await createScaleDetail(invoiceId, payload)
    // }
  }

  const setFilter = (filter: string | null) => {
    activeFilter.value = filter
  }

  // Computed
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