import { createRouter, createWebHistory } from 'vue-router'
import Dashboard from '../pages/Dashboard.vue'
import CustomerPage from '../pages/CustomerPage.vue'
import InvoiceList from '../pages/InvoiceList.vue'
import InvoiceDetail from '../pages/InvoiceDetail.vue'
import ScaleDetail from '../pages/ScaleDetail.vue'
import ItemCategoryPage from '../pages/ItemCategoryPage.vue'
import ItemPage from '../pages/ItemPage.vue'
import Login from '../pages/Login.vue'
import { useAuthStore } from '../stores/auth'

const routes = [
  // 1. Route Public
  { 
    path: '/login', 
    name: 'Login', 
    component: Login 
  },
  
  // 2. Route Dashboard 
  { 
    path: '/', 
    name: 'Dashboard', 
    component: Dashboard, 
    meta: { 
      requiresAuth: true, 
      // allowedRoles: ['staff', 'operator'] 
    } 
  },
  
  // 3. Route Lain-lain -> SEMUA Role boleh (Admin, Staff, Operator)
  { 
    path: '/invoices', 
    name: 'InvoiceList', 
    component: InvoiceList, 
    meta: { 
      requiresAuth: true, 
      allowedRoles: ['admin', 'staff', 'operator'] 
    } 
  },
  { 
    path: '/invoice/:id/scales', 
    name: 'ScaleDetail', 
    component: ScaleDetail, 
    meta: { 
      requiresAuth: true, 
      allowedRoles: ['admin', 'staff', 'operator'] 
    } 
  },
  { 
    path: '/invoice/:id/summary', 
    name: 'InvoiceDetail', 
    component: InvoiceDetail, 
    meta: { 
      requiresAuth: true, 
      allowedRoles: ['admin', 'staff', 'operator'] 
    } 
  },
  { 
    path: '/customers', 
    name: 'Customers', 
    component: CustomerPage, 
    meta: { 
      requiresAuth: true, 
      allowedRoles: ['admin', 'staff', 'operator'] 
    } 
  },
  { 
    path: '/item-categories', 
    name: 'ItemCategoryPage', 
    component: ItemCategoryPage, 
    meta: { 
      requiresAuth: true, 
      allowedRoles: ['admin', 'staff', 'operator'] 
    } 
  },
  { 
    path: '/items', 
    name: 'Items', 
    component: ItemPage, 
    meta: { 
      requiresAuth: true, 
      allowedRoles: ['admin', 'staff', 'operator'] 
    } 
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach((to, _from, next) => {
  const auth = useAuthStore()

  // 1. Belum login tapi mau ke protected route -> lempar ke Login
  if (to.meta.requiresAuth && !auth.isLoggedIn) {
    next('/login')
    return
  }

  // 2. Udah login
  if (auth.isLoggedIn) {
    // Cek apakah role user diizinkan di route ini
    const allowedRoles = to.meta.allowedRoles as string[] | undefined
    
    if (allowedRoles && !allowedRoles.includes(auth.userRole)) {
      // Role tidak diizinkan (misal: Admin maksa ke Dashboard)
      // Lempar ke halaman aman default buat Admin (Invoices)
      next('/invoices') 
      return
    }

    // 3. Udah login tapi maksa ke halaman /login
    if (to.path === '/login') {
      // Redirect ke halaman default sesuai role
      if (auth.userRole === 'admin') {
        next('/invoices') // Admin gak boleh ke dashboard, jadi ke invoices
      } else {
        next('/') // Staff/Operator ke dashboard
      }
      return
    }
  }

  // 4. Aman, lanjut
  next()
})

export default router