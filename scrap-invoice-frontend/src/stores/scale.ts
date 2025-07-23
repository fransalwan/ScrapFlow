import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { ScaleDetailResponse } from '../types/scale'
import { fetchScaleDetailsAPI, createScaleDetailAPI, deleteScaleDetailAPI, updateScaleDetailAPI } from '../services/scaleService'

export const useScaleStore = defineStore('scale', () => {
  const scaleDetails = ref<ScaleDetailResponse[]>([])
  const isLoading = ref(false)

  async function fetchScaleDetails(invoiceId: number) {
    isLoading.value = true
    try {
      const data = await fetchScaleDetailsAPI(invoiceId)
      scaleDetails.value = data
    } finally {
      isLoading.value = false
    }
  }

  async function createScaleDetail(invoiceId: number, payload: any) {
    const newDetail = await createScaleDetailAPI(invoiceId, payload)
    scaleDetails.value.push(newDetail)
  }

    async function deleteScaleDetail(id: number) {
    await deleteScaleDetailAPI(id)
    scaleDetails.value = scaleDetails.value.filter(detail => detail.id !== id)
  }

  async function updateScaleDetail(id: number, payload: any) {
  const updated = await updateScaleDetailAPI(id, payload)

  // Optional: Update data lokal biar gak perlu refetch
  const index = scaleDetails.value.findIndex(s => s.id === id)
  if (index !== -1) {
    scaleDetails.value[index] = updated
  }

  return updated
}

  return {
    scaleDetails,
    isLoading,
    fetchScaleDetails,
    createScaleDetail,
    deleteScaleDetail,
    updateScaleDetail
  }
})
