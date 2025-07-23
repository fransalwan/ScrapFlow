<template>
  <MainLayout>
    <!-- Header -->
    <div class="mb-6">
      <h2 class="text-2xl font-bold mb-1">⚖️ Scale Detail - Invoice #{{ invoiceId }}</h2>
      <p class="text-gray-500">Rekap penimbangan per item</p>
    </div>

    <!-- Daftar Penimbangan -->
    <div v-for="(item, iIndex) in scaleDetails" :key="item.id" class="mb-6">
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
  <!-- Bagian kiri: Berat dan alas -->
  <div class="text-right">
    <div class="flex items-center gap-2">
      <span class="w-24">{{ item.weight }} kg</span>
    </div>
    <span class="text-xs text-gray-400">- {{ item.alas_weight || 0 }} kg (alas)</span>
  </div>

  <!-- Bagian kanan: Icon edit dan hapus -->
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

          <div class="mb-4">
            <label class="block mb-1">Tipe Penimbangan</label>
            <select v-model="form.scale_type" class="w-full border rounded px-3 py-2" required>
              <option disabled value="">Pilih tipe</option>
              <option value="Netto">Netto</option>
              <option value="Bruto">Bruto</option>
              <option value="Kg">Kg</option>
            </select>
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
      <router-link :to="`/invoice/${invoiceId}/summary`" class="text-sm px-3 py-1 bg-gray-100 rounded hover:bg-gray-200">
        Summary
      </router-link>
    </div>
  </MainLayout>
</template>

<script setup lang="ts">
import { ref, watch, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import MainLayout from '../layouts/MainLayout.vue'
import { useScaleStore } from '../stores/scale'
import { useItemStore } from '../stores/items'
import { useToast } from 'vue-toastification'
import { Pencil, Trash2 } from 'lucide-vue-next'
import type { ScaleDetailPayload, ScaleDetailResponse } from '../types/scale'

const selectedItem = ref(null)
const isEdit = ref(false)

const handleEdit = (item: ScaleDetailResponse) => {
  selectedItem.value = item // simpan semua data asli (termasuk id)

  form.value = {
    item_id: item.item.id, // ambil dari relasi nested
    weight: item.weight,
    alas_weight: item.alas_weight,
    photo: item.photo, // optional tergantung apakah kamu pakai upload atau url
    scale_type: item.scale_type,
  }

  isEdit.value = true
}

async function handleSubmitScaleDetail() {
  if (!form.value.item_id || form.value.weight == null || form.value.weight <= 0 || !form.value.scale_type) {
    toast.error('Barang, berat (>= 0), dan tipe wajib diisi.')
    return
  }

  isSubmitting.value = true

  try {
    if (isEdit.value && selectedItem.value?.id) {
      await scaleStore.updateScaleDetail(selectedItem.value.id, {
        item_id: form.value.item_id,
        weight: form.value.weight,
        alas_weight: form.value.alas_weight ?? 0,
        photo: form.value.photo,
        scale_type: form.value.scale_type,
      })

      toast.success('Data timbangan berhasil diperbarui!')
    } else {
      await scaleStore.createScaleDetail(invoiceId, {
        item_id: form.value.item_id,
        weight: form.value.weight,
        alas_weight: form.value.alas_weight ?? 0,
        photo: form.value.photo,
        scale_type: form.value.scale_type,
      })

      toast.success('Data timbangan berhasil ditambahkan!')
    }

    closeModal()
    await scaleStore.fetchScaleDetails(invoiceId)
  } catch (error) {
    console.error('Gagal menyimpan data scale detail:', error)
    toast.error('Terjadi kesalahan saat menyimpan data.')
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

const toast = useToast()

// Store & Route
const route = useRoute()
const invoiceId = Number(route.params.id)
const scaleStore = useScaleStore()
const itemStore = useItemStore()

// State
const showModal = ref(false)
const isSubmitting = ref(false)

const form = ref({
  item_id: null,
  weight: null,
  alas_weight: null,
  photo: '',
  scale_type: '',
})

const scaleDetails = ref<any[]>([])

// Lifecycle
onMounted(() => {
  scaleStore.fetchScaleDetails(invoiceId)
  itemStore.fetchItems()
})

watch(
  () => scaleStore.scaleDetails,
  (details) => {
    scaleDetails.value = details.map((d) => ({
      id: d.id, // tambahin ini
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

// Actions
function openModal() {
  resetForm()
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

async function handleDeleteScaleDetail(scaleDetailId: number) {
  const confirmDelete = window.confirm('Yakin mau hapus data timbangan ini?')
  if (!confirmDelete) return

  try {
    await scaleStore.deleteScaleDetail(scaleDetailId)

    toast.success('Data timbangan berhasil dihapus.')
    await scaleStore.fetchScaleDetails(invoiceId) // refetch tanpa reload
  } catch (error) {
    console.error('Gagal menghapus scale detail:', error)
    toast.error('Gagal menghapus data.')
  }
}
</script>

