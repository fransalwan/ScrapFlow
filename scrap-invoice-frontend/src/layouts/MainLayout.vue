<!-- MainLayout.vue -->
<template>
  <header class="flex justify-between items-center px-6 py-4 bg-gray-900 text-white border-b border-gray-800">
  <h1 class="text-xl font-bold">🧾 InvoiceApp</h1>

  <div v-if="auth.isLoggedIn" class="relative">
    <button @click="dropdown = !dropdown" class="text-sm font-medium hover:underline">
      {{ auth.user?.name || 'User' }} ⬇️
    </button>

    <div
      v-if="dropdown"
      class="absolute right-0 mt-2 w-40 bg-white text-gray-900 border rounded shadow z-50"
    >
      <button
        @click="logout"
        class="block w-full text-left px-4 py-2 text-sm hover:bg-gray-100"
      >
        Logout
      </button>
    </div>
  </div>
</header>


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
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const sidebarOpen = ref(false)


const auth = useAuthStore()
const router = useRouter()

const dropdown = ref(false)

const logout = () => {
  auth.logout()
  router.push('/login')
}
</script>