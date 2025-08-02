<template>
  <MainLayout>
    <!-- Tabel untuk Setiap Scale Type -->
    <div v-for="(items, type) in summaryItems.data" :key="type" class="mb-10">
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
            <tr class="bg-gray-50 font-semibold">
              <td colspan="3" class="px-6 py-4 text-right">Grand Total</td>
              <td class="px-6 py-4">
                {{ formatRupiah(summaryItems.grandTotal?.[type] ?? 0) }}
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Navigasi -->
    <div class="mt-6 flex gap-2">
      <router-link
        :to="`/invoice/${invoiceId}/scales`"
        class="text-sm px-3 py-1 bg-gray-100 rounded hover:bg-gray-200"
      >
        ← Kembali ke Invoice List
      </router-link>

      <button
        @click="showModal = true"
        class="text-sm px-3 py-1 text-white bg-blue-600 rounded hover:bg-blue-500"
      >
        🖨 Preview Ringkasan FI
      </button>
    </div>

    <!-- Modal Preview Invoice -->
    <InvoicePreviewModal
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
import InvoicePreviewModal from '../components/InvoicePreviewModal.vue'

const route = useRoute()
const invoiceId = Number(route.params.id)
const showModal = ref(false)

const summaryStore = useSummaryStore()
const { summaryItems } = storeToRefs(summaryStore)

const loading = ref(true)

onMounted(async () => {
  await summaryStore.fetchSummary(invoiceId)
  loading.value = false
})

function formatRupiah(value: number) {
  return new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
    minimumFractionDigits: 0,
  }).format(value)
}
</script>
