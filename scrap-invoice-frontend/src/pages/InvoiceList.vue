<template>
  <MainLayout>  
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-2xl font-bold">👥 Invoice List</h2>
      <button
        class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded"
        @click="openModal('create')"
      >
        + New Invoice List
      </button>
    </div>


    <div class="bg-white shadow rounded overflow-x-auto">
      <table class="min-w-full text-sm text-left">
        <thead class="bg-gray-100 border-b text-gray-700">
          <tr>
            <th class="px-4 py-3">Invoice</th>
            <th class="px-4 py-3">Customer</th>
            <th class="px-4 py-3">Tanggal</th>
            <th class="px-4 py-3">Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="invoice in invoices"
            :key="invoice.id"
            class="hover:bg-gray-50 cursor-pointer"
            @click="goToDetail(invoice.id)"
          >
            <td class="px-4 py-3 font-mono">{{ invoice.number }}</td>
            <td class="px-4 py-3">{{ invoice.customer }}</td>
            <td class="px-4 py-3">{{ invoice.date }}</td>
            <td class="px-4 py-3 space-x-2">
              <button
                @click.stop="deleteInvoice(invoice.id)"
                class="text-red-600 hover:underline"
              >
                Delete
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

      


    <!-- Modal -->
    <div
      v-if="modalOpen"
      class="fixed inset-0 bg-black bg-opacity-30 flex items-center justify-center z-50 px-2"
    >
      <div class="bg-white p-4 sm:p-6 rounded w-[90%] max-w-md shadow">
        <h3 class="text-xl font-bold mb-4">
          {{ isEditing ? 'Edit Invoice' : 'Create Invoice' }}
        </h3>
        <form @submit.prevent="saveInvoice">
          <div class="mb-4">
            <label class="block mb-1">Invoice #</label>
            <input
              v-model="form.number"
              type="text"
              class="w-full border rounded px-3 py-2"
              required
            />
          </div>
          <div class="mb-4">
            <label class="block mb-1">Customer</label>
            <input
              v-model="form.customer"
              type="text"
              class="w-full border rounded px-3 py-2"
              required
            />
          </div>
          <div class="mb-4">
            <label class="block mb-1">Tanggal</label>
            <input
              v-model="form.date"
              type="date"
              class="w-full border rounded px-3 py-2"
              required
            />
          </div>
          <div class="mb-4">
            <label class="block mb-1">Total (Rp)</label>
            <input
              v-model.number="form.total"
              type="number"
              class="w-full border rounded px-3 py-2"
              required
            />
          </div>
          <div class="flex justify-end gap-2">
            <button type="button" @click="resetModal" class="px-4 py-2 bg-gray-200 rounded">
              Cancel
            </button>
            <button type="submit" class="px-4 py-2 bg-blue-600 text-white rounded">
              {{ isEditing ? 'Update' : 'Create' }}
            </button>
          </div>
        </form>
      </div>
    </div>
  
</MainLayout>

</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import MainLayout from '../layouts/MainLayout.vue'

interface Invoice {
  id: number
  number: string
  customer: string
  date: string
  total: number
}

const router = useRouter()

const invoices = ref<Invoice[]>([
  {
    id: 1,
    number: 'INV-001',
    customer: 'PT. Baja Jaya',
    date: '2025-07-01',
    total: 13500000,
  },
  {
    id: 2,
    number: 'INV-002',
    customer: 'CV. Efrata',
    date: '2025-07-05',
    total: 7400000,
  },
])

const modalOpen = ref(false)
const isEditing = ref(false)
const editingId = ref<number | null>(null)

const form = ref({
  number: '',
  customer: '',
  date: '',
  total: 0,
})

function goToDetail(id: number) {
  router.push(`/invoice/${id}`)
}

function openModal(mode: 'create' | 'edit', invoice?: Invoice) {
  modalOpen.value = true
  isEditing.value = mode === 'edit'
  if (mode === 'edit' && invoice) {
    editingId.value = invoice.id
    form.value = { ...invoice }
  } else {
    resetForm()
  }
}

function saveInvoice() {
  if (isEditing.value && editingId.value !== null) {
    const idx = invoices.value.findIndex((i) => i.id === editingId.value)
    if (idx !== -1) {
      invoices.value[idx] = { id: editingId.value, ...form.value }
    }
  } else {
    const newId = Date.now()
    invoices.value.push({ id: newId, ...form.value })
  }
  resetModal()
}

function deleteInvoice(id: number) {
  invoices.value = invoices.value.filter((i) => i.id !== id)
}

function resetForm() {
  form.value = { number: '', customer: '', date: '', total: 0 }
  editingId.value = null
}

function resetModal() {
  modalOpen.value = false
  isEditing.value = false
  resetForm()
}
</script>
