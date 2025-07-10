<!-- MainLayout.vue -->
<template>
  <div class="flex min-h-screen flex-col md:flex-row">
    <!-- Sidebar -->
    <Sidebar :isOpen="sidebarOpen" @close="sidebarOpen = false" />

    <!-- Content -->
    <div class="flex-1 bg-gray-50 p-4 pt-6 z-0 max-w-screen-xl mx-auto">
      <!-- Toggle button for mobile -->
      <button
        class="md:hidden mb-4 text-gray-700"
        @click="sidebarOpen = true"
      >
        ☰ Menu
      </button>

      <slot />
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref } from 'vue'
import Sidebar from '../components/Sidebar.vue'

const sidebarOpen = ref(false)
</script>