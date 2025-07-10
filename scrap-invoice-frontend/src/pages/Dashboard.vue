<template>
  <MainLayout>
  <div class="flex justify-between items-center mb-6">
    <h2 class="text-2xl font-bold">👥 Dashboard</h2>
  </div>

  <!-- Flex untuk mobile, grid mulai lg ke atas -->
  <div class="flex flex-col gap-4 lg:grid lg:grid-cols-3">
    <div class="bg-white p-4 rounded shadow">Total Invoices: 123</div>
    <div class="bg-white p-4 rounded shadow">Customers: 45</div>
    <div class="bg-white p-4 rounded shadow">Total Revenue: $12,345</div>
  </div>
</MainLayout>

</template>

<script setup lang="ts">
import MainLayout from '../layouts/MainLayout.vue'
</script>
