<!-- InvoiceDetail.vue -->
<template>
  <MainLayout>
    <div class="mb-6">
      <h2 class="text-2xl font-bold mb-1">⚖️ Scale Detail - Invoice #{{ invoiceId }}</h2>
      <p class="text-gray-500">Rekap penimbangan berdasarkan item</p>
    </div>

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
        <tbody class="bg-white divide-y divide-gray-200" v-if="summaryItems?.data">
  <tr v-for="item in summaryItems.data" :key="item.item_id">
    <td class="px-6 py-4">{{ item.item_name }}</td>
    <td class="px-6 py-4">{{ formatRupiah(item.price_per_kg) }} </td>
    <td class="px-6 py-4">{{ item.total_weight }} Kg</td>
    <td class="px-6 py-4">{{ formatRupiah(item.sub_total_price) }}</td>
  </tr>
  <tr class="bg-gray-50 font-semibold">
    <td colspan="3" class="px-6 py-4 text-right">Grand Total</td>
    <td class="px-6 py-4">{{ formatRupiah(summaryItems.grandTotal) }}</td>
  </tr>
</tbody>
      </table>
    </div>

    <!-- Navigasi -->
    <div class="mt-6 flex gap-2">
      <router-link :to="`/invoice/${invoiceId}/scales`" class="text-sm px-3 py-1 bg-gray-100 rounded hover:bg-gray-200">
        ← Kembali ke Invoice List
      </router-link>
      <button  class="text-sm px-3 py-1 text-white bg-blue-600 rounded hover:bg-blue-500">
        Preview Invoice
      </button>
    </div>
  </MainLayout>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'
import { useSummaryStore } from '../stores/summaryStore'
import type { SummaryResponse } from '../types/summary'
import MainLayout from '../layouts/MainLayout.vue'
import { storeToRefs } from 'pinia'

const route = useRoute()
const invoiceId = Number(route.params.id)

const summaryData = ref<SummaryResponse | null>(null)
const summaryStore = useSummaryStore()
const { summaryItems } = storeToRefs(summaryStore)

const loading = ref(true)

onMounted(async () => {
  await summaryStore.fetchSummary(invoiceId)
  loading.value = false
})

watch(() => summaryStore.summaryItems, (val) => {
  console.log('Summary changed:', val)
})

function formatRupiah(value: number) {
  return new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
    minimumFractionDigits: 0,
  }).format(value)
}


</script>

