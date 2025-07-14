<script setup lang="ts">
import { ref, onMounted } from 'vue'
import MainLayout from '../layouts/MainLayout.vue'
import type { Customer } from '../types/customer.ts'
import { useCustomerStore } from '../stores/customer'

import { useToast } from 'vue-toastification'

const toast = useToast()


const store = useCustomerStore()

const modalOpen = ref(false)
const isEditing = ref(false)
const editingId = ref<number | null>(null)

const form = ref<Omit<Customer, 'id'>>({
  name: '',
  phone: '',
  email: '',
  address: '',
  tier: 'silver',
})

onMounted(() => {
  store.fetchCustomers()
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

async function saveCustomer() {
  try {
    if (isEditing.value && editingId.value !== null) {
      await store.updateCustomer(editingId.value, form.value)
      toast.success('Customer updated!')
    } else {
      await store.createCustomer(form.value)
      toast.success('Customer created!')
    }

    await store.fetchCustomers()
    resetModal()
  } catch (err) {
    toast.error('Failed to save customer')
    console.error('Failed to save customer:', err)
  }
}


async function deleteCustomer(id: number) {
  try {
    await store.deleteCustomer(id)
    toast.success('Customer deleted!')
    await store.fetchCustomers() // buat refresh list setelah delete
  } catch (err) {
    toast.error('Failed to delete customer')
    console.error('Failed to delete customer:', err)
  }
}


function resetForm() {
  form.value = {
    name: '',
    phone: '',
    email: '',
    address: '',
    tier: 'silver',
  }
}


function resetModal() {
  modalOpen.value = false
  isEditing.value = false
  resetForm()
}
</script>

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
            <th class="px-4 py-3">Email</th>
            <th class="px-4 py-3">Phone</th>
            <th class="px-4 py-3">Address</th>
            <th class="px-4 py-3">Tier</th>
            <th class="px-4 py-3">Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr
  v-for="customer in store.customers"
  :key="customer.id"
  class="border-b hover:bg-gray-50"
>
            <td class="px-4 py-3">{{ customer.name }}</td>
            <td class="px-4 py-3">{{ customer.email }}</td>
            <td class="px-4 py-3">{{ customer.phone }}</td>
            <td class="px-4 py-3">{{ customer.address }}</td>
            <td class="px-4 py-3">{{ customer.tier }}</td>
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
  <label class="block mb-1">Email</label>
  <input
    v-model="form.email"
  type="email"
  class="w-full border rounded px-3 py-2"
  required
  pattern="[^@\s]+@[^@\s]+\.[^@\s]+"
  title="Masukkan email yang valid"
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
          <div class="mb-4">
  <label class="block mb-1">Tier</label>
  <select
    v-model="form.tier"
    class="w-full border rounded px-3 py-2"
    required
  >
    <option value="silver">Silver</option>
    <option value="gold">Gold</option>
    <option value="platinum">Platinum</option>
  </select>
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

