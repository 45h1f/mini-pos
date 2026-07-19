import { createRouter, createWebHashHistory } from 'vue-router'
import MainLayout from '../layouts/MainLayout.vue'

const routes = [
  {
    path: '/',
    component: MainLayout,
    children: [
      {
        path: '',
        name: 'Dashboard',
        component: () => import('../pages/Dashboard.vue')
      },
      {
        path: 'sales',
        name: 'Sales',
        component: () => import('../pages/Sales.vue')
      },
      {
        path: 'inventory',
        name: 'Inventory',
        component: () => import('../pages/Inventory.vue')
      },
      {
        path: 'customers',
        name: 'Customers',
        component: () => import('../pages/Customers.vue')
      },
      {
        path: 'settings',
        name: 'Settings',
        component: () => import('../pages/Settings.vue')
      }
    ]
  }
]

const router = createRouter({
  // Wails recommends Hash History for its routing to work seamlessly
  history: createWebHashHistory(),
  routes
})

export default router
