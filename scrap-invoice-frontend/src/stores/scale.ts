import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { ScaleDetailResponse } from '../types/scale'
import {
  fetchScaleDetailsAPI,
  createScaleDetailAPI,
  deleteScaleDetailAPI,
  updateScaleDetailAPI
} from '../services/scaleService'

export const useScaleStore = defineStore('scale', () => {
  const scaleDetails = ref<ScaleDetailResponse[]>([])
  const isLoading = ref(false)
  const activeFilter = ref<string | null>(null)

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
    console.log('Deleting scale detail with ID:', id)
    await deleteScaleDetailAPI(id)
    scaleDetails.value = scaleDetails.value.filter(detail => detail.id !== id)
  }

  async function updateScaleDetail(id: number, payload: any) {
    const updated = await updateScaleDetailAPI(id, payload)
    const index = scaleDetails.value.findIndex(s => s.id === id)
    if (index !== -1) {
      scaleDetails.value[index] = updated
    }
    return updated
  }

  async function cloneScaleDetailToFI(invoiceId: number) {
    const tlItems = scaleDetails.value.filter(detail => detail.scale_type === 'TL')

    // Buang semua item FI lama dulu
    scaleDetails.value = scaleDetails.value.filter(detail => detail.scale_type !== 'FI')

    // Clone TL ke FI dengan id unik baru (dummy clone)
    const cloned = tlItems.map(detail => ({
      ...detail,
      id: Date.now() + Math.random(), // id dummy, disesuaikan dengan backend kalau ada support
      scale_type: 'FI'
    }))

    scaleDetails.value.push(...cloned)
  }

  function setFilter(filter: string | null) {
    activeFilter.value = filter
  }

  const filteredScaleDetails = computed(() => {
    if (!activeFilter.value) return scaleDetails.value
    return scaleDetails.value.filter(item => item.scale_type === activeFilter.value)
  })

  return {
    scaleDetails,
    isLoading,
    fetchScaleDetails,
    createScaleDetail,
    deleteScaleDetail,
    updateScaleDetail,
    cloneScaleDetailToFI,
    activeFilter,
    setFilter,
    filteredScaleDetails
  }
})
