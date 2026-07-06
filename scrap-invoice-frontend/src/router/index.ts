import { createRouter, createWebHistory} from 'vue-router'
import Dashboard from '../pages/Dashboard.vue'
import CustomerPage from '../pages/CustomerPage.vue'
import InvoiceList from '../pages/InvoiceList.vue'
import InvoiceDetail from '../pages/InvoiceDetail.vue'
import ScaleDetail from '../pages/ScaleDetail.vue'
import ItemCategoryPage from '../pages/ItemCategoryPage.vue'
import ItemPage from '../pages/ItemPage.vue'
// import Login from '../pages/Login.vue'
import { useAuthStore } from '../stores/auth'

const routes = [
  // { path: '/login', component: Login },
  { path: '/', name: 'Dashboard', component: Dashboard },
  { path: '/invoices', name: 'InvoiceList', component: InvoiceList },
  { path: '/invoice/:id/scales', name: 'ScaleDetail', component: ScaleDetail },
  { path: '/invoice/:id/summary', name: 'InvoiceDetail', component: InvoiceDetail },
  { path: '/customers', name: 'Customers', component: CustomerPage },
  { path: '/item-categories', name: 'ItemCategoryPage', component: ItemCategoryPage },
  { path: '/items', name: 'Items', component: ItemPage },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach((to, from, next) => {
  const auth = useAuthStore()
  if (to.meta.requiresAuth && !auth.isLoggedIn) {
    next('/login')
  } else {
    next()
  }
})

export default router
