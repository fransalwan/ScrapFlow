<template>
  <MainLayout>
    <div class="mb-6">
      <h2 class="text-2xl font-bold mb-1">⚖️ Scale Detail - Invoice #{{ invoiceId }}</h2>
      <p class="text-gray-500">Rekap penimbangan per item</p>
    </div>

    <!-- CRUD Item + Penimbangan -->
    <div v-for="(item, iIndex) in scaleDetails" :key="item.item" class="mb-6">
      <div class="flex items-center justify-between mb-2">
        <div v-if="editingItem !== iIndex">
          <h3 class="text-lg font-bold cursor-pointer hover:underline" @dblclick="editItem(iIndex)">
            {{ item.item }}
          </h3>
        </div>
        <div v-else class="flex items-center gap-2">
          <input
            v-model="item.item"
            @blur="stopEditItem"
            @keydown.enter.prevent="stopEditItem"
            type="text"
            class="border px-2 py-1 rounded"
            autofocus
          />
        </div>
        <button
          @click="deleteItem(iIndex)"
          class="text-red-600 hover:underline text-sm"
        >
          🗑 Hapus Barang
        </button>
      </div>

      <div class="bg-white rounded shadow">
        <div
          v-for="(weight, wIndex) in item.weights"
          :key="wIndex"
          class="flex items-center justify-between border-b last:border-none px-4 py-2"
        >
          <div class="text-sm">Penimbangan ke-{{ wIndex + 1 }}</div>
          <div class="flex items-center gap-2">
            <input
              v-if="editingWeight === `${iIndex}-${wIndex}`"
              type="number"
              v-model.number="item.weights[wIndex]"
              @blur="editingWeight = null"
              @keydown.enter.prevent="editingWeight = null"
              class="w-20 border rounded px-2 py-1 text-right"
              autofocus
            />
            <span
              v-else
              @dblclick="editingWeight = `${iIndex}-${wIndex}`"
              class="text-right w-20 cursor-pointer hover:underline"
            >
              {{ weight }} kg
            </span>
            <button @click="deleteWeight(iIndex, wIndex)" class="text-red-600 hover:underline text-sm">
              🗑
            </button>
          </div>
        </div>
        <div class="px-4 py-2">
          <button
            @click="addWeight(iIndex)"
            class="text-blue-600 hover:underline text-sm"
          >
            ➕ Tambah Penimbangan
          </button>
        </div>
      </div>
    </div>

    <!-- Tambah Barang -->
    <div class="mt-6">
      <input
        v-model="newItemName"
        type="text"
        placeholder="Nama barang baru"
        class="border rounded px-3 py-2 mr-2"
      />
      <button
        @click="addItem"
        class="bg-blue-600 mt-2 text-white px-4 py-2 rounded hover:bg-blue-700"
      >
        ➕ Tambah Barang
      </button>
    </div>

    <!-- Tombol Back -->
  <router-link
    :to="`/invoice/${invoiceId}`"
    class="inline-flex items-center text-sm px-3 py-1 bg-gray-100 rounded hover:bg-gray-200 mt-6"

  >
    ← Kembali ke Invoice List
  </router-link>
  </MainLayout>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import MainLayout from '../layouts/MainLayout.vue'

const route = useRoute()
const invoiceId = route.params.id

const scaleDetails = ref([
  { item: 'Besi', weights: [40, 30] },
  { item: 'Aluminium', weights: [20, 30] },
])

const editingWeight = ref<string | null>(null)
const editingItem = ref<number | null>(null)
const newItemName = ref('')

function addWeight(index: number) {
  scaleDetails.value[index].weights.push(0)
}

function deleteWeight(itemIndex: number, weightIndex: number) {
  scaleDetails.value[itemIndex].weights.splice(weightIndex, 1)
}

function addItem() {
  if (!newItemName.value.trim()) return
  scaleDetails.value.push({ item: newItemName.value.trim(), weights: [] })
  newItemName.value = ''
}

function deleteItem(index: number) {
  scaleDetails.value.splice(index, 1)
}

function editItem(index: number) {
  editingItem.value = index
}

function stopEditItem() {
  editingItem.value = null
}
</script>
