<template>
  <!-- Overlay (mobile only) -->
  <div
    v-if="isOpen"
    class="fixed inset-0 bg-black bg-opacity-50 z-40 md:hidden"
    @click="$emit('close')"
  ></div>

  <!-- Sidebar -->
<aside
  :class="[
    'w-48 md:w-52 bg-gray-900 text-white fixed md:sticky top-0 left-0 z-50 transition-transform transform',
    isOpen ? 'translate-x-0' : '-translate-x-full',
    'md:translate-x-0 md:z-auto',
  ]"
  style="min-height: 100vh"
>
    <div class="p-6 text-2xl font-bold border-b border-gray-700">
      🧾 InvoiceApp
    </div>
    <nav class="flex-1 p-4 space-y-2 pb-24">
      <RouterLink
        to="/"
        class="block py-2 px-4 rounded hover:bg-gray-700"
        active-class="bg-gray-800"
      >
        Dashboard
      </RouterLink>
      <RouterLink
        to="/invoices"
        class="block py-2 px-4 rounded hover:bg-gray-700 border-gray-700 border-b"
        active-class="bg-gray-800"
      >
        Invoice List
      </RouterLink>
      <RouterLink
        to="/customers"
        class="block py-2 px-4 rounded hover:bg-gray-700"
        active-class="bg-gray-800"
      >
        Customers
      </RouterLink>
      <RouterLink
        to="/item-categories"
        class="block py-2 px-4 rounded hover:bg-gray-700"
        active-class="bg-gray-800"
      >
        Category Item
      </RouterLink>
      <RouterLink
        to="/items"
        class="block py-2 px-4 rounded hover:bg-gray-700"
        active-class="bg-gray-800"
      >
        Item
      </RouterLink>
    </nav>
  </aside>
</template>

<script setup lang="ts">
defineProps<{
  isOpen: boolean
}>()

defineEmits(['close'])
</script>
