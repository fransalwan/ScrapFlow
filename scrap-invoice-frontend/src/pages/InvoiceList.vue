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

    <!-- Loading State -->
    <div v-if="invoiceStore.isLoading" class="text-center py-8">
      <p class="text-gray-500">Loading invoices...</p>
    </div>

    <!-- Error State -->
    <div v-else-if="invoiceStore.error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
      {{ invoiceStore.error }}
    </div>

    <!-- Empty State -->
    <div v-else-if="invoices.length === 0" class="text-center py-8">
      <p class="text-gray-500">Belum ada invoice. Klik "+ New Invoice" untuk membuat invoice pertama.</p>
    </div>

    <!-- Data Table -->
    <div v-else class="bg-white shadow rounded overflow-x-auto">
      <table class="min-w-full text-sm text-left">
        <thead class="bg-gray-100 border-b text-gray-700">
          <tr>
            <th class="px-4 py-3">Invoice Number</th>
            <th class="px-4 py-3">Customer</th>
            <th class="px-4 py-3">Tanggal</th>
            <th class="px-4 py-3">Total Weight</th>
            <th class="px-4 py-3">Total Price</th>
            <th class="px-4 py-3">Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="invoice in invoices"
            :key="invoice.id"
            class="hover:bg-gray-50 cursor-pointer"
            @click="goToScaleDetail(invoice.id)"
          >
            <td class="px-4 py-3 font-mono">{{ invoice.invoice_number }}</td>
            <td class="px-4 py-3">{{ invoice.customer?.name || '-' }}</td>
            <td class="px-4 py-3">{{ formatDateShort(invoice.invoice_date || invoice.created_at) }}</td>
            <td class="px-4 py-3">{{ invoice.total_weight }} kg</td>
            <td class="px-4 py-3">{{ formatCurrency(invoice.total_price) }}</td>
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
              <option disabled :value="0">Pilih customer</option>
  
  <!-- ✅ FLEXIBLE: Coba customer_id, kalau nggak ada pakai id -->
  <option
    v-for="customer in customers"
    :key="customer.id"
    :value="customer.id"
  >
    {{ customer.name }}
  </option>
</select>
          </div>

          <!-- TANGGAL INVOICE - PINDAH KE ATAS -->
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

          <!-- STATUS - GANTI KE SELECT DROPDOWN -->
          <div class="mb-4">
            <label class="block mb-1">Status</label>
            <select
              v-model="form.status"
              class="w-full border rounded px-3 py-2"
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
import { useInvoiceStore } from '../stores/invoice'
import { useCustomerStore } from '../stores/customer'
import type { InvoiceForm, Invoice } from '../types/invoice'
import MainLayout from '../layouts/MainLayout.vue'
import { useToast } from 'vue-toastification'

const toast = useToast()
const router = useRouter()
const invoiceStore = useInvoiceStore()
const customerStore = useCustomerStore()

const { invoices } = storeToRefs(invoiceStore)
const { customers } = storeToRefs(customerStore)

const modalOpen = ref(false)
const isEditing = ref(false)
const editingId = ref<number | null>(null)
const isSubmitting = ref(false)

const form = ref<InvoiceForm>({
  customer_id: 0,
  invoice_date: new Date().toISOString().slice(0, 10), // ← Default hari ini
  status: 'draft', // ← Default status
  payment_method: '',
  note: '',
})

onMounted(() => {
   invoiceStore.fetchInvoices()
   customerStore.fetchCustomers()
})

function goToScaleDetail(id: number) {
  router.push(`/invoice/${id}/scales`)
}

function formatDateShort(dateStr: string | undefined): string {
  if (!dateStr) return '-' // ✅ Fallback jika tanggal kosong
  
  return new Date(dateStr).toLocaleDateString('id-ID', {
    day: '2-digit',
    month: 'short',
    year: 'numeric',
  })
}

function formatCurrency(amount: number): string {
  return new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
  }).format(amount)
}

function openModal(mode: 'create' | 'edit', invoice?: Invoice) {
  modalOpen.value = true
  isEditing.value = mode === 'edit'

  if (mode === 'edit' && invoice) {
    editingId.value = invoice.id
    console.log("Editing invoice:", invoice)
    form.value = {
      customer_id: invoice.customer_id,
      invoice_date: invoice.invoice_date?.slice(0, 10) || invoice.created_at.slice(0, 10),
      status: invoice.status || 'draft',
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
    const payload = {
      customer_id: Number(form.value.customer_id) || 0,
      invoice_date: form.value.invoice_date,
      status: form.value.status || 'draft',
      payment_method: form.value.payment_method || '',
      note: form.value.note || '',
    }

    if (isEditing.value && editingId.value !== null) {
      await invoiceStore.updateInvoice(editingId.value, payload)
      toast.success('Invoice updated!')
    } else {
      await invoiceStore.createInvoice(payload)
      toast.success('Invoice created!')
    }

    resetModal()
  } catch (err) {
    console.error("❌ Gagal menyimpan invoice:", err)
  } finally {
    isSubmitting.value = false
  }
}

async function deleteInvoice(id: number) {
  try {
    toast.success('Deleting invoice...')
    await invoiceStore.deleteInvoice(id)
    await invoiceStore.fetchInvoices() 
  } catch (err) {
    toast.error('Failed to delete invoice')
  }
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