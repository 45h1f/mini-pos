<script setup lang="ts">
import { ref } from 'vue'

// In production, fetch these metrics via Wails backend bindings
const todaySales = ref(12450) // $124.50 in cents
const todayCount = ref(18)
const lowStockProducts = ref([
  { id: '3', name: 'Cold Brew Bottle', sku: 'CF-200', stock: 2, threshold: 5 },
  { id: '4', name: 'Avocado Toast', sku: 'FD-100', stock: 1, threshold: 5 },
])
</script>

<template>
  <div class="dashboard-layout">
    <div class="metrics-grid">
      <div class="metric-card card">
        <div class="metric-icon blue-icon">💰</div>
        <div class="metric-info">
          <h3>Today's Sales</h3>
          <p class="metric-value">${{ (todaySales / 100).toFixed(2) }}</p>
        </div>
      </div>
      
      <div class="metric-card card">
        <div class="metric-icon green-icon">🧾</div>
        <div class="metric-info">
          <h3>Transactions</h3>
          <p class="metric-value">{{ todayCount }}</p>
        </div>
      </div>
    </div>
    
    <div class="dashboard-widgets">
      <div class="widget card warning-widget">
        <div class="widget-header">
          <h3>⚠️ Low Stock Alerts</h3>
          <span class="badge">{{ lowStockProducts.length }}</span>
        </div>
        
        <div v-if="lowStockProducts.length === 0" class="empty-state">
          All products are sufficiently stocked.
        </div>
        
        <div class="stock-list" v-else>
          <div v-for="product in lowStockProducts" :key="product.id" class="stock-item">
            <div class="stock-details">
              <h4>{{ product.name }}</h4>
              <span>{{ product.sku }}</span>
            </div>
            <div class="stock-level">
              <span class="qty critical">{{ product.stock }}</span>
              <span class="qty-label">left</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.dashboard-layout {
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

.metrics-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1.5rem;
}

.metric-card {
  display: flex;
  align-items: center;
  gap: 1.5rem;
  padding: 2rem;
}

.metric-icon {
  font-size: 3rem;
  padding: 1rem;
  border-radius: 16px;
  background-color: var(--hover-bg);
}

.metric-info h3 {
  margin: 0;
  font-size: 1.1rem;
  color: var(--sidebar-text);
  font-weight: 500;
}

.metric-value {
  margin: 0.5rem 0 0 0;
  font-size: 2.2rem;
  font-weight: 700;
  color: var(--text-color);
}

.dashboard-widgets {
  display: grid;
  grid-template-columns: 1fr;
}

.widget {
  padding: 1.5rem;
}

.widget-header {
  display: flex;
  align-items: center;
  gap: 1rem;
  margin-bottom: 1.5rem;
}

.widget-header h3 {
  margin: 0;
  color: #d97706; /* amber-600 */
}

.badge {
  background-color: #fef3c7; /* amber-100 */
  color: #92400e; /* amber-900 */
  padding: 0.2rem 0.6rem;
  border-radius: 999px;
  font-weight: 700;
  font-size: 0.9rem;
}

.stock-list {
  display: flex;
  flex-direction: column;
}

.stock-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem 0;
  border-bottom: 1px solid var(--border-color);
}

.stock-item:last-child {
  border-bottom: none;
  padding-bottom: 0;
}

.stock-details h4 {
  margin: 0 0 0.2rem 0;
  font-size: 1.05rem;
}

.stock-details span {
  font-size: 0.85rem;
  color: var(--sidebar-text);
}

.stock-level {
  display: flex;
  align-items: baseline;
  gap: 0.3rem;
}

.qty {
  font-size: 1.4rem;
  font-weight: 700;
}

.critical {
  color: #dc2626; /* red-600 */
}

.qty-label {
  color: var(--sidebar-text);
  font-size: 0.9rem;
}
</style>
