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
    'w-48 md:w-52 h-screen bg-gray-900 text-white fixed top-0 left-0 z-50 transition-transform transform',
    isOpen ? 'translate-x-0' : '-translate-x-full',
    'md:translate-x-0 md:relative md:z-auto'
  ]"
>
    <div class="p-6 text-2xl font-bold border-b border-gray-700">
      🧾 InvoiceApp
    </div>
    <nav class="flex-1 p-4 space-y-2">
      <RouterLink
        to="/"
        class="block py-2 px-4 rounded hover:bg-gray-700"
        active-class="bg-gray-800"
      >
        Dashboard
      </RouterLink>
      <RouterLink
        to="/customers"
        class="block py-2 px-4 rounded hover:bg-gray-700"
        active-class="bg-gray-800"
      >
        Customers
      </RouterLink>
      <RouterLink
        to="/invoice-list"
        class="block py-2 px-4 rounded hover:bg-gray-700"
        active-class="bg-gray-800"
      >
        Invoice List
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
