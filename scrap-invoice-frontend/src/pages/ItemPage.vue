<script setup lang="ts">
import { ref, onMounted } from 'vue'
import MainLayout from '../layouts/MainLayout.vue'
import { useItemStore } from '../stores/items'
import { useCategoryStore } from '../stores/category'
import type { Item, ItemFormInput, ItemFormPayload } from '../types/item'
import { useToast } from 'vue-toastification'

const toast = useToast()
const itemStore = useItemStore()
const categoryStore = useCategoryStore()

const modalOpen = ref(false)
const isEditing = ref(false)
const editingId = ref<number | null>(null)

const form = ref<ItemFormInput>({
  name: '',
  category_id: null,
  price_per_kg: null,
})

onMounted(() => {
  itemStore.fetchItems()
  categoryStore.fetchCategories()
})

function openModal(mode: 'create' | 'edit', item?: Item) {
  modalOpen.value = true
  isEditing.value = mode === 'edit'

  if (mode === 'edit' && item) {
    editingId.value = item.id
    form.value = {
      name: item.item_name,
      category_id: item.item_category_id,
      price_per_kg: item.price_per_kg,
    }
  } else {
    resetForm()
  }
}

async function saveItem() {
  try {
    const payload: ItemFormPayload = {
      item_name: form.value.name,
      item_category: form.value.category_id!,
      price_per_kg: form.value.price_per_kg!,
    }

    if (isEditing.value && editingId.value !== null) {
      await itemStore.updateItem(editingId.value, payload)
      toast.success('Item updated!')
    } else {
      await itemStore.createItem(payload)
      toast.success('Item created!')
    }

    await itemStore.fetchItems()
    resetModal()
  } catch (err) {
    console.error('Failed to save item:', err)
    toast.error('Failed to save item')
  }
}

async function deleteItem(id: number) {
  try {
    await itemStore.deleteItem(id)
    await itemStore.fetchItems()
    toast.success('Item deleted!')
  } catch (err) {
    console.error('Failed to delete item:', err)
    toast.error('Failed to delete item')
  }
}

function resetForm() {
  form.value = {
    name: '',
    category_id: null,
    price_per_kg: null,
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
      <h2 class="text-2xl font-bold">🧾 Items</h2>
      <button
        class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded"
        @click="openModal('create')"
      >
        + New Item
      </button>
    </div>

    <div class="bg-white shadow rounded overflow-x-auto">
      <table class="min-w-full text-sm text-left">
        <thead class="bg-gray-100 border-b text-gray-700">
          <tr>
            <th class="px-4 py-3">Name</th>
            <th class="px-4 py-3">Category</th>
            <th class="px-4 py-3">Price / Kg</th>
            <th class="px-4 py-3">Actions</th>

          </tr>
        </thead>
        <tbody>
          <tr
            v-for="item in itemStore.items"
            :key="item.id"
            class="border-b hover:bg-gray-50"
          >
            <td class="px-4 py-3">{{ item.item_name }}</td>
            <td class="px-4 py-3">{{ item.category.item_category_name }}</td>
            <td class="px-4 py-3">Rp {{ item.price_per_kg.toLocaleString('id-ID') }}</td>
            <td class="px-4 py-3 space-x-2">
              <button class="text-blue-600 hover:underline" @click="openModal('edit', item)">Edit</button>
              <button class="text-red-600 hover:underline" @click="deleteItem(item.id)">Delete</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Modal -->
    <div v-if="modalOpen" class="fixed inset-0 bg-black bg-opacity-30 flex items-center justify-center z-50">
      <div class="bg-white p-6 rounded w-full max-w-md shadow">
        <h3 class="text-xl font-bold mb-4">{{ isEditing ? 'Edit' : 'Add' }} Item</h3>
        <form @submit.prevent="saveItem">
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
            <label class="block mb-1">Category</label>
            <select
              v-model="form.category_id"
              class="w-full border rounded px-3 py-2"
              required
            >
              <option disabled value="">-- Select Category --</option>
              <option
  v-for="cat in categoryStore.categories"
  :key="cat.id"
  :value="cat.id"
>
  {{ cat.item_category_name }}
</option>
            </select>
          </div>
          <div class="mb-4">
  <label class="block mb-1">Price Per Kg</label>
  <input
    v-model="form.price_per_kg"
    type="number"
    step="0.01"
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
