<script setup lang="ts">
// ✅ FIX: Hapus defineProps dan defineEmits dari import
import { ref, nextTick, computed } from 'vue'

const props = defineProps<{
  summary: Record<string, any[]>
  grandTotal: Record<string, number>
  visible: boolean
}>()

const emit = defineEmits(['close'])

const printRef = ref<HTMLElement | null>(null)

const fiSummary = computed(() => props.summary['FI'] || [])
const fiGrandTotal = computed(() => props.grandTotal['FI'] || 0)

function formatRupiah(value: number) {
  return new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
    minimumFractionDigits: 0,
  }).format(value)
}

function handlePrint() {
  if (!printRef.value) return

  const printContent = printRef.value.innerHTML
  const originalContent = document.body.innerHTML

  document.body.innerHTML = printContent
  window.print()
  document.body.innerHTML = originalContent

  nextTick(() => {
    window.location.reload()
  })
}
</script>

<template>
  <div
    class="fixed inset-0 z-50 flex items-center justify-center bg-black/50"
    v-if="visible"
  >
    <div
      ref="printRef"
      class="print-area bg-white p-2 text-xs w-[260px] font-mono text-black relative"
    >
      <button
        @click="$emit('close')"
        class="absolute top-1 right-1 text-red-500 text-base no-print"
      >
        ❌
      </button>

      <h3 class="font-bold text-center text-sm mb-2">🧾 Summary Timbangan FI</h3>

      <div class="mb-2">
        <div
          v-for="item in fiSummary"
          :key="item.item_id"
          class="mb-1"
        >
          <div class="font-semibold">{{ item.item_name }}</div>
          <div class="flex justify-between">
            <span>{{ item.total_weight }} kg × {{ formatRupiah(item.price_per_kg) }}</span>
            <span>{{ formatRupiah(item.sub_total_price) }}</span>
          </div>
        </div>

        <div class="border-t border-black border-dashed my-2"></div>

        <div class="flex justify-between font-bold text-sm">
          <span>Grand Total</span>
          <span>{{ formatRupiah(fiGrandTotal) }}</span>
        </div>
      </div>

      <button
        class="no-print mt-3 w-full bg-black text-white py-1 text-xs rounded hover:bg-gray-800"
        @click="handlePrint"
      >
        Cetak PDF
      </button>
    </div>
  </div>
</template>

<style scoped>
@media print {
  body * {
    visibility: hidden;
  }

  .print-area,
  .print-area * {
    visibility: visible;
  }

  .print-area {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
  }

  .no-print {
    display: none !important;
  }
}
</style>