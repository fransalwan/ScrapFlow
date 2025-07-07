<script setup>
import { onMounted } from 'vue'
import { useInvoiceStore } from '@/stores/invoiceStore'

const invoiceStore = useInvoiceStore()

onMounted(() => {
  invoiceStore.fetchInvoices()
})
</script>

<template>
  <div class="p-4">
    <h1 class="text-2xl font-bold mb-4">Daftar Invoice</h1>

    <div v-if="invoiceStore.isLoading">Loading...</div>
    <div v-else-if="invoiceStore.error">{{ invoiceStore.error }}</div>
    <div v-else>
      <div
        v-for="invoice in invoiceStore.invoices"
        :key="invoice.id"
        class="border p-3 rounded mb-2"
      >
        <div class="flex justify-between">
          <div>
            <p class="font-semibold">No: {{ invoice.invoice_number }}</p>
            <p>Customer: {{ invoice.customer.name }}</p>
          </div>
          <!-- <router-link :to="`/invoice/${invoice.id}`" class="text-blue-500">Detail</router-link> -->
        </div>
      </div>
    </div>
  </div>
</template>
