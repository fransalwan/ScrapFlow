<template>
  <MainLayout>
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-2xl font-bold">👥 Dashboard</h2>
    </div>

    <!-- Flex untuk mobile, grid mulai lg ke atas -->
    <div class="flex flex-col gap-4 lg:grid lg:grid-cols-3">
      <div class="bg-white p-4 rounded shadow">
        <p class="text-gray-500 text-sm mb-1">Total Invoices</p>
        <p class="text-2xl font-bold">
          {{ dashboard.invoiceCount !== null ? dashboard.invoiceCount : '...' }}
        </p>
      </div>
      <div class="bg-white p-4 rounded shadow">
        <p class="text-gray-500 text-sm mb-1">Customers</p>
        <p class="text-2xl font-bold">
          {{ dashboard.customerCount !== null ? dashboard.customerCount : '...' }}
        </p>
      </div>
      <div class="bg-white p-4 rounded shadow">
        <p class="text-gray-500 text-sm mb-1">Total Revenue</p>
        <p class="text-2xl font-bold text-green-600">Coming Soon 💸</p>
      </div>
    </div>
  </MainLayout>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import MainLayout from '../layouts/MainLayout.vue'
import { useDashboardStore } from '../stores/dashboard'

const dashboard = useDashboardStore()

onMounted(() => {
  dashboard.fetchDashboardCounts()
})
</script>
