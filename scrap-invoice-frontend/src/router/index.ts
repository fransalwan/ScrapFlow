import { createRouter, createWebHistory} from 'vue-router'
import Dashboard from '../pages/Dashboard.vue'
import CustomerPage from '../pages/CustomerPage.vue'
import InvoiceList from '../pages/InvoiceList.vue'
import InvoiceDetail from '../pages/InvoiceDetail.vue'
import ScaleDetail from '../pages/ScaleDetail.vue'
import ItemCategoryPage from '../pages/ItemCategoryPage.vue'
import ItemPage from '../pages/ItemPage.vue'

const routes = [
  { path: '/', name: 'Dashboard', component: Dashboard },
  { path: '/invoice-list', name: 'InvoiceList', component: InvoiceList },
  { path: '/invoice/:id/scales', name: 'ScaleDetail', component: ScaleDetail },
  { path: '/invoice/:id', name: 'InvoiceDetail', component: InvoiceDetail },
  { path: '/customers', name: 'Customers', component: CustomerPage },
  { path: '/item-categories', name: 'ItemCategoryPage', component: ItemCategoryPage },
  { path: '/items', name: 'Items', component: ItemPage },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
