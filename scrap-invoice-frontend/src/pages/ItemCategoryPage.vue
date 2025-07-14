<script setup lang="ts">
import { ref, onMounted } from 'vue'
import MainLayout from '../layouts/MainLayout.vue'
import { useCategoryStore } from '../stores/category'
import type { Category } from '../stores/category.ts'
import { useToast } from 'vue-toastification'

const toast = useToast()
const store = useCategoryStore()

const modalOpen = ref(false)
const isEditing = ref(false)
const editingId = ref<number | null>(null)

// Pakai 'name' biar tetap simple di form
const form = ref({
  name: '',
})

onMounted(() => {
  store.fetchCategories()
})

function openModal(mode: 'create' | 'edit', category?: Category) {
  modalOpen.value = true
  isEditing.value = mode === 'edit'
  if (mode === 'edit' && category) {
    editingId.value = category.id
    form.value.name = category.item_category_name // 👈 disesuaikan saat edit
  } else {
    resetForm()
  }
}

async function saveCategory() {
  const payload = { item_category_name: form.value.name }

  try {
    if (isEditing.value && editingId.value !== null) {
      await store.updateCategory(editingId.value, payload)
      toast.success('Category updated!')
    } else {
      await store.createCategory(payload)
      toast.success('Category created!')
    }

    await store.fetchCategories()
    resetModal()
  } catch (err) {
    console.error('Failed to save category:', err)
    toast.error('Failed to save category')
  }
}

async function deleteCategory(id: number) {
  try {
    await store.deleteCategory(id)
    await store.fetchCategories()
    toast.success('Category deleted!')
  } catch (err) {
    console.error('Failed to delete category:', err)
    toast.error('Failed to delete category')
  }
}

function resetForm() {
  form.value = { name: '' }
  editingId.value = null
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
      <h2 class="text-2xl font-bold">📦 Item Categories</h2>
      <button
        class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded"
        @click="openModal('create')"
      >
        + New Category
      </button>
    </div>

    <div class="bg-white shadow rounded overflow-x-auto">
      <table class="min-w-full text-sm text-left">
        <thead class="bg-gray-100 border-b text-gray-700">
          <tr>
            <th class="px-4 py-3">Name</th>
            <th class="px-4 py-3">Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="category in store.categories"
            :key="category.id"
            class="border-b hover:bg-gray-50"
          >
            <td class="px-4 py-3">{{ category.item_category_name }}</td>
            <td class="px-4 py-3 space-x-2">
              <button
                class="text-blue-600 hover:underline"
                @click="openModal('edit', category)"
              >
                Edit
              </button>
              <button
                class="text-red-600 hover:underline"
                @click="deleteCategory(category.id)"
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
      class="fixed inset-0 bg-black bg-opacity-30 flex items-center justify-center z-50"
    >
      <div class="bg-white p-6 rounded w-full max-w-md shadow">
        <h3 class="text-xl font-bold mb-4">
          {{ isEditing ? 'Edit' : 'Add' }} Category
        </h3>
        <form @submit.prevent="saveCategory">
          <div class="mb-4">
            <label class="block mb-1">Name</label>
            <input
              v-model="form.name"
              type="text"
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
