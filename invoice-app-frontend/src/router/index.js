// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router'
import InvoiceList from '../pages/InvoiceList.vue'
import InvoiceDetail from '../pages/InvoiceDetail.vue'

const routes = [
  { path: '/', name: 'InvoiceList', component: InvoiceList },
  { path: '/invoice/:id', name: 'InvoiceDetail', component: InvoiceDetail },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
