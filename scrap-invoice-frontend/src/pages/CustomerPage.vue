<template>
  <MainLayout>
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-2xl font-bold">👥 Customers</h2>
      <button
        class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded"
        @click="openModal('create')"
      >
        + New Customer
      </button>
    </div>

    <div class="bg-white shadow rounded overflow-x-auto">
      <table class="min-w-full text-sm text-left">
        <thead class="bg-gray-100 border-b text-gray-700">
          <tr>
            <th class="px-4 py-3">Name</th>
            <th class="px-4 py-3">Phone</th>
            <th class="px-4 py-3">Address</th>
            <th class="px-4 py-3">Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="customer in customers"
            :key="customer.id"
            class="border-b hover:bg-gray-50"
          >
            <td class="px-4 py-3">{{ customer.name }}</td>
            <td class="px-4 py-3">{{ customer.phone }}</td>
            <td class="px-4 py-3">{{ customer.address }}</td>
            <td class="px-4 py-3 space-x-2">
              <button
                class="text-blue-600 hover:underline"
                @click="openModal('edit', customer)"
              >
                Edit
              </button>
              <button
                class="text-red-600 hover:underline"
                @click="deleteCustomer(customer.id)"
              >
                Delete
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Modal -->
    <div v-if="modalOpen" class="fixed inset-0 bg-black bg-opacity-30 flex items-center justify-center z-50">
      <div class="bg-white p-6 rounded w-full max-w-md shadow">
        <h3 class="text-xl font-bold mb-4">{{ isEditing ? 'Edit' : 'Add' }} Customer</h3>
        <form @submit.prevent="saveCustomer">
          <div class="mb-4">
            <label class="block mb-1">Name</label>
            <input
              v-model="form.name"
              type="text"
              class="w-full border rounded px-3 py-2"
              required
            />
          </div>
          <div class="mb-4">
            <label class="block mb-1">Phone</label>
            <input
              v-model="form.phone"
              type="text"
              class="w-full border rounded px-3 py-2"
              required
            />
          </div>
          <div class="mb-4">
            <label class="block mb-1">Address</label>
            <textarea
              v-model="form.address"
              class="w-full border rounded px-3 py-2"
              required
            ></textarea>
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
import MainLayout from '../layouts/MainLayout.vue'

interface Customer {
  id: number
  name: string
  phone: string
  address: string
}

const customers = ref<Customer[]>([
  { id: 1, name: 'PT. Baja Jaya', phone: '081234567890', address: 'Jl. Besi No. 1' },
  { id: 2, name: 'CV. Efrata', phone: '085678123456', address: 'Jl. Baja Mulia No. 12' },
])

const modalOpen = ref(false)
const isEditing = ref(false)
const editingId = ref<number | null>(null)

const form = ref({
  name: '',
  phone: '',
  address: '',
})

function openModal(mode: 'create' | 'edit', customer?: Customer) {
  modalOpen.value = true
  isEditing.value = mode === 'edit'
  if (mode === 'edit' && customer) {
    editingId.value = customer.id
    form.value = { ...customer }
  } else {
    resetForm()
  }
}

function saveCustomer() {
  if (isEditing.value && editingId.value !== null) {
    const index = customers.value.findIndex(c => c.id === editingId.value)
    if (index !== -1) {
      customers.value[index] = { id: editingId.value, ...form.value }
    }
  } else {
    const newId = Date.now()
    customers.value.push({ id: newId, ...form.value })
  }
  resetModal()
}

function deleteCustomer(id: number) {
  customers.value = customers.value.filter(c => c.id !== id)
}

function resetForm() {
  form.value = { name: '', phone: '', address: '' }
  editingId.value = null
}

function resetModal() {
  modalOpen.value = false
  isEditing.value = false
  resetForm()
}
</script>
