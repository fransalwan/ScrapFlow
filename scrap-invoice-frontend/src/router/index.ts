import { createRouter, createWebHistory} from 'vue-router'
import Dashboard from '../pages/Dashboard.vue'
import CustomerPage from '../pages/CustomerPage.vue'
import InvoiceList from '../pages/InvoiceList.vue'
import InvoiceDetail from '../pages/InvoiceDetail.vue'
import ScaleDetail from '../pages/ScaleDetail.vue'

const routes = [
  { path: '/', name: 'Dashboard', component: Dashboard },
  { path: '/customers', name: 'Customers', component: CustomerPage },
  { path: '/invoice/:id/scales', name: 'ScaleDetail', component: ScaleDetail },
  { path: '/invoice-list', name: 'InvoiceList', component: InvoiceList },
  { path: '/invoice/:id', name: 'InvoiceDetail', component: InvoiceDetail },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
