<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useCartStore } from '../stores/cart'

import Receipt from '../components/Receipt.vue'

// Note: In production, these will be replaced with Wails bindings (e.g. GetAllProducts())
const cart = useCartStore()
const search = ref('')
const barcodeBuffer = ref('')
const wedgeTimeout = ref<number | null>(null)
const completedSaleData = ref<any>(null)

// Mock products
const products = ref([
  { id: '1', name: 'Artisan Espresso', sku: 'CF-100', price: 450, stock: 100 },
  { id: '2', name: 'Almond Croissant', sku: 'BK-001', price: 320, stock: 40 },
  { id: '3', name: 'Cold Brew Bottle', sku: 'CF-200', price: 500, stock: 20 },
  { id: '4', name: 'Avocado Toast', sku: 'FD-100', price: 850, stock: 15 },
  { id: '5', name: 'Oat Milk Latte', sku: 'CF-101', price: 550, stock: 80 },
  { id: '6', name: 'Matcha Frappe', sku: 'CF-300', price: 600, stock: 40 },
])

const filteredProducts = computed(() => {
  const query = search.value.toLowerCase()
  return products.value.filter(p => 
    p.name.toLowerCase().includes(query) || 
    p.sku.toLowerCase().includes(query)
  )
})

const addToCart = (product: any) => {
  cart.addItem(product)
}

// Wedge Scanner Support
const handleKeydown = (e: KeyboardEvent) => {
  // Ignore if typing in an input field (user is manually searching)
  if (e.target instanceof HTMLInputElement || e.target instanceof HTMLTextAreaElement) {
    return
  }
  
  if (e.key === 'Enter') {
    if (barcodeBuffer.value.length > 2) {
      // Find product by SKU or Barcode
      const p = products.value.find(prod => prod.sku === barcodeBuffer.value || prod.id === barcodeBuffer.value)
      if (p) {
        addToCart(p)
      }
    }
    barcodeBuffer.value = ''
  } else if (e.key.length === 1) {
    barcodeBuffer.value += e.key
    // Clear buffer rapidly if it's slow human typing instead of a fast wedge scanner
    if (wedgeTimeout.value) clearTimeout(wedgeTimeout.value)
    wedgeTimeout.value = window.setTimeout(() => { barcodeBuffer.value = '' }, 40) // 40ms threshold
  }
}

const checkout = () => {
  if (cart.items.length === 0) return
  
  // Snap shot the cart state for the receipt
  completedSaleData.value = {
    invoiceNo: 'INV-' + Math.floor(Math.random() * 1000000),
    items: [...cart.items],
    subtotal: cart.subtotal,
    tax: cart.tax,
    discount: cart.discount,
    total: cart.total
  }

  // Trigger print dialog (Vue nextTick wait ensures component renders)
  setTimeout(() => {
    window.print()
    cart.clearCart()
  }, 100)
}

onMounted(() => {
  window.addEventListener('keydown', handleKeydown)
})

onUnmounted(() => {
  window.removeEventListener('keydown', handleKeydown)
})
</script>

<template>
  <div class="pos-layout">
    <!-- Left Side: Product Grid -->
    <div class="product-section">
      <div class="search-bar">
        <input 
          type="text" 
          v-model="search" 
          placeholder="Search products by name or SKU..." 
          class="search-input"
        />
      </div>
      
      <div class="product-grid">
        <div 
          class="product-card" 
          v-for="product in filteredProducts" 
          :key="product.id" 
          @click="addToCart(product)"
        >
          <div class="product-content">
            <h4 class="product-title">{{ product.name }}</h4>
            <span class="product-sku">{{ product.sku }}</span>
            <div class="product-footer">
              <span class="product-price">${{ (product.price / 100).toFixed(2) }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <!-- Right Side: Cart -->
    <div class="cart-section card">
      <div class="cart-header">
        <h3>Current Order</h3>
        <button class="btn-clear" @click="cart.clearCart()">Clear</button>
      </div>
      
      <div class="cart-items">
        <div v-if="cart.items.length === 0" class="empty-cart">
          <p>Cart is empty. Scan or select items.</p>
        </div>
        <div v-for="item in cart.items" :key="item.product_id" class="cart-item">
          <div class="item-details">
            <h5>{{ item.name }}</h5>
            <span class="item-price">${{ (item.price / 100).toFixed(2) }}</span>
          </div>
          <div class="qty-controls">
            <button class="qty-btn" @click="cart.updateQty(item.product_id, item.qty - 1)">-</button>
            <span class="qty">{{ item.qty }}</span>
            <button class="qty-btn" @click="cart.updateQty(item.product_id, item.qty + 1)">+</button>
          </div>
          <div class="item-total">
            ${{ (item.line_total / 100).toFixed(2) }}
          </div>
        </div>
      </div>
      
      <div class="cart-summary">
        <div class="summary-row">
          <span>Subtotal</span>
          <span>${{ (cart.subtotal / 100).toFixed(2) }}</span>
        </div>
        <div class="summary-row">
          <span>Tax</span>
          <span>${{ (cart.tax / 100).toFixed(2) }}</span>
        </div>
        <div class="summary-row grand-total">
          <span>Total</span>
          <span>${{ (cart.total / 100).toFixed(2) }}</span>
        </div>
        
        <button 
          class="btn btn-primary checkout-btn" 
          :disabled="cart.items.length === 0"
          @click="checkout"
        >
          Charge ${{ (cart.total / 100).toFixed(2) }}
        </button>
      </div>
    </div>
    
    <!-- Hidden Print Area -->
    <Receipt v-if="completedSaleData" :saleData="completedSaleData" />
  </div>
</template>

<style scoped>
.pos-layout {
  display: grid;
  grid-template-columns: 2fr 1.2fr;
  gap: 2rem;
  height: 100%;
}

/* Product Section */
.product-section {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.search-input {
  width: 100%;
  padding: 1rem 1.5rem;
  border-radius: 12px;
  border: 1px solid var(--border-color);
  background-color: var(--card-bg);
  color: var(--text-color);
  font-size: 1rem;
  box-shadow: var(--shadow-sm);
  transition: all 0.2s ease;
}

.search-input:focus {
  outline: none;
  border-color: var(--primary-color);
  box-shadow: 0 0 0 3px var(--primary-color-light);
}

.product-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
  gap: 1rem;
  overflow-y: auto;
  padding-bottom: 2rem;
}

.product-card {
  background-color: var(--card-bg);
  border: 1px solid var(--border-color);
  border-radius: 12px;
  padding: 1.2rem;
  cursor: pointer;
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  min-height: 140px;
}

.product-card:hover {
  transform: translateY(-4px) scale(1.02);
  border-color: var(--primary-color);
  box-shadow: var(--shadow-md);
  background-color: var(--hover-bg);
}

.product-card:active {
  transform: scale(0.98);
}

.product-title {
  font-size: 1.05rem;
  font-weight: 600;
  margin-bottom: 0.2rem;
}

.product-sku {
  font-size: 0.8rem;
  color: var(--sidebar-text);
}

.product-footer {
  margin-top: 1rem;
  display: flex;
  justify-content: flex-end;
}

.product-price {
  font-weight: 700;
  color: var(--primary-color);
  font-size: 1.1rem;
}

/* Cart Section */
.cart-section {
  display: flex;
  flex-direction: column;
  height: 100%;
  padding: 0;
  overflow: hidden;
}

.cart-header {
  padding: 1.5rem;
  border-bottom: 1px solid var(--border-color);
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.cart-header h3 {
  margin: 0;
  font-size: 1.2rem;
}

.btn-clear {
  background: none;
  border: none;
  color: #ef4444; /* red-500 */
  cursor: pointer;
  font-weight: 500;
  font-size: 0.9rem;
}

.btn-clear:hover {
  text-decoration: underline;
}

.cart-items {
  flex: 1;
  overflow-y: auto;
  padding: 1.5rem;
}

.empty-cart {
  text-align: center;
  color: var(--sidebar-text);
  margin-top: 2rem;
  font-style: italic;
}

.cart-item {
  display: grid;
  grid-template-columns: 2fr 1fr 1fr;
  align-items: center;
  padding: 1rem 0;
  border-bottom: 1px solid var(--border-color);
}

.cart-item:last-child {
  border-bottom: none;
}

.item-details h5 {
  margin: 0 0 0.2rem 0;
  font-size: 0.95rem;
}

.qty-controls {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
}

.qty-btn {
  width: 28px;
  height: 28px;
  border-radius: 6px;
  border: 1px solid var(--border-color);
  background-color: var(--bg-color);
  color: var(--text-color);
  cursor: pointer;
  font-weight: bold;
}

.qty-btn:hover {
  background-color: var(--hover-bg);
}

.qty {
  font-weight: 600;
  min-width: 20px;
  text-align: center;
}

.item-total {
  text-align: right;
  font-weight: 600;
}

/* Cart Summary */
.cart-summary {
  background-color: var(--hover-bg);
  padding: 1.5rem;
  border-top: 1px solid var(--border-color);
}

.summary-row {
  display: flex;
  justify-content: space-between;
  margin-bottom: 0.8rem;
  color: var(--sidebar-text);
}

.grand-total {
  font-size: 1.3rem;
  font-weight: 700;
  color: var(--text-color);
  margin-top: 1rem;
  padding-top: 1rem;
  border-top: 1px dashed var(--border-color);
  margin-bottom: 1.5rem;
}

.checkout-btn {
  width: 100%;
  padding: 1rem;
  font-size: 1.1rem;
  font-weight: 700;
  letter-spacing: 0.5px;
}

.checkout-btn:disabled {
  background-color: var(--border-color);
  color: var(--sidebar-text);
  cursor: not-allowed;
  box-shadow: none;
}
</style>
