<template>
  <MainLayout>  
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-2xl font-bold">👥 Invoice List</h2>
      <button
        class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded"
        @click="openModal('create')"
      >
        + New Invoice
      </button>
    </div>


    <div class="bg-white shadow rounded overflow-x-auto">
      <table class="min-w-full text-sm text-left">
        <thead class="bg-gray-100 border-b text-gray-700">
          <tr>
            <th class="px-4 py-3">Invoice Number</th>
            <th class="px-4 py-3">Customer</th>
            <th class="px-4 py-3">Tanggal</th>
            <th class="px-4 py-3">Status</th>
            <th class="px-4 py-3">Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="invoice in invoices"
            :key="invoice.id"
            class="hover:bg-gray-50 cursor-pointer"
            @click="goToDetail(invoice.id)"
          >
            <td class="px-4 py-3 font-mono">{{ invoice.invoice_number }}</td>
            <td class="px-4 py-3">{{ invoice.customer.name }}</td>
            <td class="px-4 py-3">{{ formatDateShort(invoice.invoice_date) }}</td>
            <td class="px-4 py-3">
  <span
    :class="[
      'px-2 py-1 text-xs rounded-full',
      invoice.status === 'paid' ? 'bg-green-100 text-green-700' :
      invoice.status === 'unpaid' ? 'bg-yellow-100 text-yellow-700' :
      'bg-gray-100 text-gray-700'
    ]"
  >
    {{ invoice.status }}
  </span>
</td>

            <td class="px-4 py-3 space-x-2">
  <button
    @click.stop="openModal('edit', invoice)"
    class="text-blue-600 hover:underline"
  >
    Edit
  </button>
  <button
    @click.stop="deleteInvoice(invoice.id)"
    class="text-red-600 hover:underline"
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
      class="fixed inset-0 bg-black bg-opacity-30 flex items-center justify-center z-50 px-2"
    >
      <div class="bg-white p-4 sm:p-6 rounded w-[90%] max-w-md shadow">
        <h3 class="text-xl font-bold mb-4">
          {{ isEditing ? 'Edit Invoice' : 'Create Invoice' }}
        </h3>
        <form @submit.prevent="saveInvoice">

  <div class="mb-4">
    <label class="block mb-1">Customer</label>
    <select
      v-model="form.customer_id"
      class="w-full border rounded px-3 py-2"
      required
    >
      <option disabled value="">Pilih customer</option>
      <option
        v-for="customer in customers"
        :key="customer.id"
        :value="customer.id"
      >
        {{ customer.name }}
      </option>
    </select>
  </div>

  <div class="mb-4">
  <label class="block mb-1">Tanggal Invoice</label>
  <input
    v-model="form.invoice_date"
    type="date"
    class="w-full border rounded px-3 py-2"
    required
  />
</div>

  <div class="mb-4">
  <label class="block mb-1">Payment Method</label>
  <input
    v-model="form.payment_method"
    type="text"
    class="w-full border rounded px-3 py-2"
    placeholder="e.g. Transfer, Cash"
  />
</div>
<div class="mb-4">
  <label class="block mb-1">Status</label>
  <select
    v-model="form.status"
    class="w-full border rounded px-3 py-2"
    required
  >
    <option value="draft">Draft</option>
    <option value="unpaid">Unpaid</option>
    <option value="paid">Paid</option>
    <option value="finalized">Finalized</option>
  </select>
</div>


<div class="mb-4">
  <label class="block mb-1">Note</label>
  <textarea
    v-model="form.note"
    class="w-full border rounded px-3 py-2"
    rows="2"
    placeholder="Tambahkan catatan..."
  ></textarea>
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

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { storeToRefs } from 'pinia'
import { useInvoiceStore } from '../stores/invoiceStore'
import { useCustomerStore } from '../stores/customer'
import type { InvoiceForm } from '../types/invoice'
import type { Invoice } from '../types/invoice'
import MainLayout from '../layouts/MainLayout.vue'

const router = useRouter()
const invoiceStore = useInvoiceStore()
const customerStore = useCustomerStore()

const { invoices } = storeToRefs(invoiceStore)
const { customers } = storeToRefs(customerStore)

const modalOpen = ref(false)
const isEditing = ref(false)
const editingId = ref<number | null>(null)

const form = ref<InvoiceForm>({
  customer_id: 0,
  invoice_date: '',     
  status: 'draft',
  payment_method: '',
  note: '',
})


const isSubmitting = ref(false)

onMounted(async () => {
  await invoiceStore.fetchInvoices()
  await customerStore.fetchCustomers()
})

function goToDetail(id: number) {
  router.push(`/invoice/${id}`)
}

function formatDateShort(dateStr: string): string {
  return new Date(dateStr).toLocaleDateString('id-ID', {
    day: '2-digit',
    month: 'short',
    year: 'numeric',
  })
}

function openModal(mode: 'create' | 'edit', invoice?: Invoice) {
  modalOpen.value = true
  isEditing.value = mode === 'edit'

  if (mode === 'edit' && invoice) {
    editingId.value = invoice.id
    form.value = {
  customer_id: invoice.customer.id,
  invoice_date: invoice.invoice_date.slice(0, 10),
  status: invoice.status,
  payment_method: invoice.payment_method ?? '',
  note: invoice.note ?? '',
}
  } else {
    resetForm()
  }
}

async function saveInvoice() {
  isSubmitting.value = true
  try {
    const payload: any = {
      customer_id: form.value.customer_id,
      invoice_date: form.value.invoice_date,
      status: form.value.status,
      payment_method: form.value.payment_method,
      note: form.value.note,
    }

    if (isEditing.value && editingId.value !== null) {
      await invoiceStore.updateInvoice(editingId.value, payload)
    } else {
      await invoiceStore.createInvoice(payload)
    }

    // Hanya reset modal kalau tidak error
    resetModal()
  } catch (err) {
    console.error("Gagal menyimpan invoice:", err)
  } finally {
    isSubmitting.value = false
  }
}


async function deleteInvoice(id: number) {
  await invoiceStore.deleteInvoice(id)
}

function resetForm() {
  form.value = {
  customer_id: 0,
  invoice_date: new Date().toISOString().slice(0, 10),
  status: 'draft',
  payment_method: '',
  note: '',
}
  editingId.value = null
}

function resetModal() {
  modalOpen.value = false
  isEditing.value = false
  resetForm()
}
</script>



