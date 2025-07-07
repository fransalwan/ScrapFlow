<template>
  <div class="p-4">
    <h1 class="text-2xl font-bold mb-4">Detail Invoice #{{ invoice.value.invoice_number }}</h1>

    <div class="mb-4">
      <p class="font-semibold">Customer: {{ invoice.value.customer.name }}</p>
      <p>Total Berat: {{ invoice.value.total_weight }} kg</p>
      <p>Total Harga: Rp{{ invoice.value.total_price.toLocaleString() }}</p>
    </div>

    <h2 class="text-lg font-semibold mt-4">Ringkasan (Summaries)</h2>
    <ul class="list-disc ml-5">
      <li v-for="s in invoice.value.summaries" :key="s.id">
        {{ s.item.item_name }} - {{ s.total_weight }} kg × Rp{{ s.item.price_per_kg.toLocaleString() }} = Rp{{ s.sub_total_price.toLocaleString() }}
      </li>
    </ul>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import axios from 'axios'

const route = useRoute()
const invoice = ref({})

onMounted(async () => {
  const res = await axios.get(`http://localhost:8080/api/invoices/${route.params.id}`)
  invoice.value = res.data.data
})
</script>
