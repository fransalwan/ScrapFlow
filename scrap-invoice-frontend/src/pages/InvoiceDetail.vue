<template>
  <MainLayout>
    <!-- Loading State -->
    <div v-if="loading" class="flex justify-center items-center h-64">
      <p class="text-gray-500">Memuat data ringkasan...</p>
    </div>

   <!-- Cari bagian ini di template InvoiceDetail.vue lu -->

<!-- Tabel untuk Setiap Scale Type -->
<div v-if="!loading && summaryItems?.data" v-for="(items, type) in summaryItems.data" :key="type" class="mb-10">
  <h3 class="text-lg font-bold text-gray-700 mb-2">
    Rekap Timbangan {{ type }}
  </h3>

  <div class="overflow-x-auto bg-white rounded-xl shadow">
    <table class="min-w-full divide-y divide-gray-200">
      <thead class="bg-gray-100">
        <tr>
          <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Item</th>
          <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Harga/kg</th>
          <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Berat Total (kg)</th>
          <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Subtotal</th>
        </tr>
      </thead>

      <tbody>
        <tr v-for="item in items" :key="item.item_id">
          <td class="px-6 py-4">{{ item.item_name }}</td>
          <td class="px-6 py-4">{{ formatRupiah(item.price_per_kg) }}</td>
          <td class="px-6 py-4">{{ item.total_weight }} Kg</td>
          <td class="px-6 py-4">{{ formatRupiah(item.sub_total_price) }}</td>
        </tr>
        
        <!-- ✅ FIX GRAND TOTAL: Gunakan optional chaining dan fallback -->
        <tr class="bg-gray-50 font-semibold border-t-2 border-gray-200">
          <td colspan="3" class="px-6 py-4 text-right">Grand Total {{ type }}</td>
          <td class="px-6 py-4 text-blue-700">
            {{ formatRupiah(summaryItems?.grandTotal?.[type]) }}
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</div>

<!-- Pesan Jika Data Kosong -->
<div v-else-if="!loading" class="text-center py-10 text-gray-500">
  Belum ada data ringkasan untuk invoice ini.
</div>

    <!-- Navigasi -->
    <div class="mt-6 flex gap-2">
      <!-- ✅ FIX: Teks disesuaikan dengan tujuan route -->
      <router-link
        :to="`/invoice/${invoiceId}/scales`"
        class="text-sm px-3 py-1 bg-gray-100 rounded hover:bg-gray-200 flex items-center gap-1"
      >
        ← Kembali ke Detail Timbangan
      </router-link>

      <!-- Tombol ini hanya muncul jika ada data FI -->
      <button
        v-if="summaryItems?.data?.['FI']"
        @click="showModal = true"
        class="text-sm px-3 py-1 text-white bg-blue-600 rounded hover:bg-blue-500 flex items-center gap-1"
      >
        🖨 Preview Ringkasan FI
      </button>
    </div>

    <!-- Modal Preview Invoice -->
    <InvoicePreviewModal
      v-if="showModal && summaryItems?.data"
      :summary="summaryItems.data"
      :grandTotal="summaryItems.grandTotal"
      :visible="showModal"
      @close="showModal = false"
    />
  </MainLayout>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useSummaryStore } from '../stores/summaryStore'
import { storeToRefs } from 'pinia'
import MainLayout from '../layouts/MainLayout.vue'
import InvoicePreviewModal from '../components/InvoicePreviewModal.vue' // Jika dipakai

const route = useRoute()
const invoiceId = Number(route.params.id)
const showModal = ref(false)
const loading = ref(true)

const summaryStore = useSummaryStore()
// ✅ TypeScript sekarang tahu persis bentuk data summaryItems
const { summaryItems } = storeToRefs(summaryStore)

onMounted(async () => {
  try {
    await summaryStore.fetchSummary(invoiceId)
  } catch (error) {
    console.error('Gagal memuat summary di Invoice Detail:', error)
  } finally {
    loading.value = false
  }
})

// ✅ FIX: Terima number | undefined, dan gunakan ?? 0 untuk fallback
function formatRupiah(value: number | undefined) {
  return new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
    minimumFractionDigits: 0,
  }).format(value ?? 0) 
}
</script>