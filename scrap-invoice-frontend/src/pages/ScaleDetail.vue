<template>
  <MainLayout>
    <!-- Header -->
    <div class="mb-6">
      <h2 class="text-2xl font-bold mb-1">⚖️ Scale Detail - Invoice #{{ invoiceId }}</h2>
      <p class="text-gray-500">Rekap penimbangan per item</p>
    </div>

    <!-- Filter Scale Type -->
    <div class="flex gap-2 mb-4">
      <button
        v-for="type in ['FI', 'TL', 'TG', 'TS']"
        :key="type"
        @click="setFilter(type)"
        :class="[
          'px-3 py-1 rounded text-sm',
          activeFilter === type ? 'bg-blue-600 text-white' : 'bg-gray-200 hover:bg-gray-300'
        ]"
      >
        {{ type }}
      </button>
    </div>

    <!-- Daftar Penimbangan -->
    <div v-for="(item, iIndex) in filteredScaleDetails" :key="item.id" class="mb-6">
      <div class="bg-white rounded shadow px-4 py-3 flex justify-between items-center gap-4 border-b">
        <!-- KIRI: Foto + Info -->
        <div class="flex items-center gap-4">
          <div v-if="item.photo" class="w-16 h-16 overflow-hidden rounded border">
            <img :src="item.photo" alt="Foto Timbangan" class="object-cover w-full h-full" />
          </div>
          <div class="text-sm flex flex-col gap-1">
            <span class="font-medium">Penimbangan ke-{{ iIndex + 1 }}</span>
            <span class="text-gray-500">({{ item.item_name }})</span>
            <div class="text-gray-500">
              Tipe Timbangan: <strong>{{ item.scale_type || '-' }}</strong>
            </div>
          </div>
        </div>

        <!-- KANAN: Berat + Alas -->
        <div class="flex justify-between items-center">
          <div class="text-right">
            <div class="flex items-center gap-2">
              <span class="w-24">{{ item.weight }} kg</span>
            </div>
            <span class="text-xs text-gray-400">- {{ item.alas_weight || 0 }} kg (alas)</span>
          </div>

          <div class="flex items-center gap-2 ml-4">
            <Pencil
              @click="openEditModal(item)"
              class="w-5 h-5 text-blue-500 hover:text-blue-700 cursor-pointer"
            />
            <Trash2
              @click="handleDeleteScaleDetail(item.id)"
              class="w-5 h-5 text-red-500 hover:text-red-700 cursor-pointer"
            />
          </div>
        </div>
      </div>
    </div>

    <!-- Tombol Tambah Timbangan -->
    <div class="mt-6">
      <button @click="openModal()" class="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700">
        ➕ Tambah Timbangan
      </button>
    </div>

    <!-- Modal Tambah Timbangan -->
    <div v-if="showModal" class="fixed inset-0 bg-black bg-opacity-30 flex items-center justify-center z-50">
      <div class="bg-white p-6 rounded w-full max-w-md shadow">
        <h3 class="text-xl font-bold mb-4">
          {{ isEdit ? 'Edit Timbangan' : 'Tambah Timbangan' }}
        </h3>

        <form @submit.prevent="handleSubmitScaleDetail">
          <div class="mb-4">
            <label class="block mb-1">Nama Barang</label>
            <select v-model="form.item_id" class="w-full border rounded px-3 py-2" required>
              <option disabled value="">Pilih barang</option>
              <option v-for="item in itemStore.items" :key="item.id" :value="item.id">
                {{ item.item_name }}
              </option>
            </select>
          </div>

          <div class="mb-4">
            <label class="block mb-1">Berat (kg)</label>
            <input v-model.number="form.weight" type="number" min="0" class="w-full border rounded px-3 py-2" required />
          </div>

          <div class="mb-4">
            <label class="block mb-1">Berat Alas (kg)</label>
            <input v-model.number="form.alas_weight" type="number" min="0" class="w-full border rounded px-3 py-2" />
          </div>

          <div class="mb-4">
            <label class="block mb-1">Photo</label>
            <input v-model="form.photo" type="text" class="w-full border rounded px-3 py-2" />
          </div>


          <div class="flex justify-end gap-2">
            <button type="button" @click="closeModal" class="px-4 py-2 bg-gray-200 rounded">Cancel</button>
            <button :disabled="isSubmitting" type="submit" class="px-4 py-2 bg-blue-600 text-white rounded">
              {{ isSubmitting ? 'Menyimpan...' : (isEdit ? 'Update' : 'Create') }}
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- Navigasi -->
    <div class="mt-6 flex gap-2">
      <router-link :to="`/invoices`" class="text-sm px-3 py-1 bg-gray-100 rounded hover:bg-gray-200">
        ← Kembali ke Invoice List
      </router-link>
      <router-link v-if="activeFilter === 'FI'" :to="`/invoice/${invoiceId}/summary`" class="text-sm px-3 py-1 bg-gray-100 rounded hover:bg-gray-200">
        Summary
      </router-link>
    </div>
  </MainLayout>
</template>

<script setup lang="ts">
import MainLayout from '../layouts/MainLayout.vue'
import { ref, watch, onMounted, computed } from 'vue'
import { useRoute } from 'vue-router'
import { useScaleStore } from '../stores/scale'
import { useItemStore } from '../stores/items'
import { useToast } from 'vue-toastification'
import { Pencil, Trash2 } from 'lucide-vue-next'
import Swal from 'sweetalert2'
import type { ScaleDetailResponse } from '../types/scale'

const route = useRoute()
const invoiceId = Number(route.params.id)
const scaleStore = useScaleStore()
const itemStore = useItemStore()
const toast = useToast()

const showModal = ref(false)
const isSubmitting = ref(false)
const isEdit = ref(false)
const selectedItem = ref<ScaleDetailResponse | null>(null)
const activeFilter = ref<string>('FI')

const form = ref({
  item_id: null,
  weight: null,
  alas_weight: null,
  photo: '',
  scale_type: '',
})

const scaleDetails = ref<any[]>([])

const filteredScaleDetails = computed(() => {
  return scaleDetails.value.filter(item => item.scale_type === activeFilter.value)
})

onMounted(() => {
  scaleStore.fetchScaleDetails(invoiceId)
  itemStore.fetchItems()
})

watch(
  () => scaleStore.scaleDetails,
  (details) => {
    scaleDetails.value = details.map(d => ({
      id: d.id,
      item_id: d.item.id,
      item_name: d.item.name,
      weight: d.weight,
      alas_weight: d.alas_weight,
      photo: d.photo,
      scale_type: d.scale_type,
    }))
  },
  { immediate: true }
)

function openModal() {
  resetForm()
  form.value.scale_type = activeFilter.value // ✅ force isi tipe berdasarkan filter yang aktif
  showModal.value = true
}

function closeModal() {
  resetForm()
  showModal.value = false
}

function resetForm() {
  form.value = {
    item_id: null,
    weight: null,
    alas_weight: null,
    photo: '',
    scale_type: '',
  }
  isEdit.value = false
  selectedItem.value = null
}

function setFilter(type: string) {
  activeFilter.value = type
}

async function handleSubmitScaleDetail() {
  const { item_id, weight, alas_weight = 0, photo, scale_type } = form.value

  if (!item_id || weight == null || weight <= 0 || !scale_type) {
    toast.error('Barang, berat (>= 0), dan tipe wajib diisi.')
    return
  }

  isSubmitting.value = true

  try {
    const payload = { item_id, weight, alas_weight, photo, scale_type }

    if (isEdit.value) {
      const scaleId = selectedItem.value?.id
      if (!scaleId) throw new Error('ID timbangan tidak ditemukan.')

      await scaleStore.updateScaleDetail(scaleId, payload)
      toast.success('Data timbangan berhasil diperbarui!')
    } else {
      await scaleStore.createScaleDetail(invoiceId, payload)
      toast.success('Data timbangan berhasil ditambahkan!')
    }

    closeModal()
    await scaleStore.fetchScaleDetails(invoiceId)
  } catch (error: any) {
    console.error('Gagal menyimpan data scale detail:', error)
    toast.error(error?.message || 'Terjadi kesalahan saat menyimpan data.')
  } finally {
    isSubmitting.value = false
  }
}

function openEditModal(item: any) {
  selectedItem.value = { ...item }
  form.value = {
    item_id: item.item_id,
    weight: item.weight,
    alas_weight: item.alas_weight,
    photo: item.photo,
    scale_type: item.scale_type,
  }
  isEdit.value = true
  showModal.value = true
}

async function handleDeleteScaleDetail(scaleDetailId: number) {
  const result = await Swal.fire({
    title: 'Yakin mau hapus data timbangan ini?',
    text: 'Data yang dihapus tidak bisa dikembalikan.',
    icon: 'warning',
    showCancelButton: true,
    confirmButtonColor: '#d33',
    cancelButtonColor: '#aaa',
    confirmButtonText: 'Ya, hapus',
    cancelButtonText: 'Batal',
  })

  if (!result.isConfirmed) return

  try {
    await scaleStore.deleteScaleDetail(scaleDetailId)
    toast.success('Data timbangan berhasil dihapus.')
    await scaleStore.fetchScaleDetails(invoiceId)
  } catch (error) {
    console.error('Gagal menghapus scale detail:', error)
    toast.error('Gagal menghapus data.')
  }
}
</script>
