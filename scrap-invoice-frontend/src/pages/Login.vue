<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-100">
    <div class="bg-white p-6 rounded-lg shadow w-full max-w-sm">
      <h2 class="text-2xl font-bold mb-4 text-center">Login</h2>

      <form @submit.prevent="handleLogin">
        <input
          v-model="form.username"
          type="text"
          placeholder="Username"
          class="w-full mb-3 px-4 py-2 border rounded"
        />
        <input
          v-model="form.password"
          type="password"
          placeholder="Password"
          class="w-full mb-3 px-4 py-2 border rounded"
        />

        <button
          type="submit"
          class="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700"
        >
          {{ loading ? 'Logging in...' : 'Login' }}
        </button>

        <p v-if="error" class="text-red-500 mt-3 text-sm text-center">
          {{ error }}
        </p>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useAuthStore } from '../stores/auth'
import { useRouter } from 'vue-router'

const auth = useAuthStore()
const router = useRouter()

const form = ref({
  username: '',
  password: '',
})

const loading = ref(false)
const error = ref('')

const handleLogin = async () => {
  error.value = ''
  loading.value = true
  const success = await auth.login(form.value)

  if (success) {
    router.push('/') // redirect ke dashboard
  } else {
    error.value = 'Username atau password salah'
  }

  loading.value = false
}
</script>
