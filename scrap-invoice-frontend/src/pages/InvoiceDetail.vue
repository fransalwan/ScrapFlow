<template>
  <MainLayout>

 
    <div class="mb-6 space-y-2">
      <h2 class="text-2xl font-bold">🧾 Invoice #{{ invoiceId }}</h2>

      <div>
        <span class="text-gray-600 font-semibold">Customer: </span>
        <span
          v-if="editingField !== 'customer'"
          @dblclick="startEditing('customer')"
          class="text-blue-600 cursor-pointer hover:underline"
        >
          {{ invoice.customer }}
        </span>
        <input
          v-else
          v-model="invoice.customer"
          @blur="stopEditing"
          @keydown.enter.prevent="stopEditing"
          type="text"
          class="border rounded px-2 py-1 w-64"
          autofocus
        />
      </div>

      <div>
        <span class="text-gray-600 font-semibold">Tanggal: </span>
        <span
          v-if="editingField !== 'date'"
          @dblclick="startEditing('date')"
          class="text-blue-600 cursor-pointer hover:underline"
        >
          {{ formatDate(invoice.date) }}
        </span>
        <input
          v-else
          v-model="invoice.date"
          @blur="stopEditing"
          @keydown.enter.prevent="stopEditing"
          type="date"
          class="border rounded px-2 py-1"
          autofocus
        />
      </div>
    </div>

    <div class="bg-white rounded shadow overflow-hidden">
      <table class="min-w-full text-sm text-left">
        <thead class="bg-gray-100 border-b text-gray-700">
          <tr>
            <th class="px-4 py-3">Item</th>
            <th class="px-4 py-3 text-right">Qty</th>
            <th class="px-4 py-3 text-right">Harga</th>
            <th class="px-4 py-3 text-right">Total</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="summary in summaries"
            :key="summary.item"
            class="border-b hover:bg-gray-50"
          >
            <td class="px-4 py-3">{{ summary.item }}</td>
            <td class="px-4 py-3 text-right">{{ summary.qty }}</td>
            <td class="px-4 py-3 text-right">
  <span
    v-if="editingPriceItem !== summary.item"
    @dblclick="startEditingPrice(summary.item)"
    class="cursor-pointer hover:underline text-blue-600"
  >
    Rp {{ summary.price.toLocaleString() }}
  </span>
  <input
    v-else
    v-model.number="summary.price"
    @blur="stopEditingPrice"
    @keydown.enter.prevent="stopEditingPrice"
    type="number"
    class="border rounded px-1 py-0.5 w-24 text-right"
    autofocus
  />
</td>

            <td class="px-4 py-3 text-right font-semibold">
              Rp {{ (summary.qty * summary.price).toLocaleString() }}
            </td>
          </tr>
          <tr class="font-bold">
            <td class="px-4 py-3 text-right" colspan="3">Total</td>
            <td class="px-4 py-3 text-right">
              Rp {{ total.toLocaleString() }}
            </td>
          </tr>
        </tbody>
      </table>
      
    </div>

    <div class="flex gap-2 mb-4">
      <router-link
  :to="`/invoice/${invoiceId}/scales`"
  class="inline-block mt-6 bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
>
  🔍 Lihat Detail Scale
</router-link>
<button
        @click="showPreview = true"
        class="bg-green-600  hover:bg-green-700 inline-block mt-6  text-white px-4 py-2 rounded"
      >
        🖨 Cetak Invoice
      </button>
      

    </div>
    <!-- Tombol Back -->
  <router-link
    to="/invoice-list"
    class="inline-flex items-center text-sm px-3 py-1 bg-gray-100 rounded hover:bg-gray-200"

  >
    ← Kembali ke Invoice List
  </router-link>
    

    <!-- Modal Preview Cetak -->
    <div
      v-if="showPreview"
      class="fixed inset-0 z-50 bg-black bg-opacity-40 flex items-center justify-center"
    >
      <div class="bg-white p-6 rounded shadow w-full max-w-3xl print:shadow-none print:p-0">
     <div
  id="print-content"
  class="text-xs w-full max-w-xs mx-auto font-mono p-2"
>
  <h2 class="text-center font-bold mb-2">🧾 Invoice #{{ invoiceId }}</h2>
  <p>Customer: {{ invoice.customer }}</p>
  <p>Tanggal: {{ formatDate(invoice.date) }}</p>
  <hr class="my-2 border-t border-dashed" />

  <table class="w-full mb-2 text-xs">
    <thead>
      <tr>
        <th class="text-left">Item</th>
        <th class="text-right">Qty</th>
        <th class="text-right">Harga</th>
      </tr>
    </thead>
    <tbody>
      <tr
        v-for="summary in summaries"
        :key="summary.item"
        class="border-b"
      >
        <td>{{ summary.item }}</td>
        <td class="text-right">{{ summary.qty }}</td>
        <td class="text-right">Rp {{ summary.price.toLocaleString() }}</td>
      </tr>
    </tbody>
  </table>

  <hr class="my-2 border-t border-dashed" />

  <p class="text-right font-bold">
    Total: Rp {{ total.toLocaleString() }}
  </p>

  <p class="text-center mt-4">Terima kasih 🙏</p>
</div>



        <!-- Tombol Aksi -->
        <div class="flex justify-end gap-2 print:hidden">
          <button @click="showPreview = false" class="px-4 py-2 bg-gray-300 rounded">Tutup</button>
          <button @click="printInvoice" class="px-4 py-2 bg-green-600 text-white rounded">Print</button>
        </div>
      </div>
    </div>
  </MainLayout>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRoute } from 'vue-router'
import MainLayout from '../layouts/MainLayout.vue'

const route = useRoute()
const invoiceId = route.params.id

const invoice = ref({
  customer: 'PT. Baja Abadi',
  date: '2025-07-09',
})

const summaries = ref([
  { item: 'Besi', qty: 100, price: 10000 },
  { item: 'Aluminium', qty: 50, price: 20000 },
])

const total = computed(() =>
  summaries.value.reduce((acc, item) => acc + item.qty * item.price, 0)
)

const editingField = ref<null | 'customer' | 'date'>(null)

const showPreview = ref(false)

const editingPriceItem = ref<string | null>(null)

function startEditingPrice(itemName: string) {
  editingPriceItem.value = itemName
}

function stopEditingPrice() {
  editingPriceItem.value = null
}

function printInvoice() {
  const printContent = document.getElementById('print-content')?.innerHTML
  if (!printContent) return

  const original = document.body.innerHTML
  document.body.innerHTML = printContent
  window.print()
  document.body.innerHTML = original
  location.reload()
}

function startEditing(field: 'customer' | 'date') {
  editingField.value = field
}

function stopEditing() {
  editingField.value = null
}

function formatDate(dateStr: string): string {
  const options: Intl.DateTimeFormatOptions = {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  }
  return new Date(dateStr).toLocaleDateString('id-ID', options)
}
</script>