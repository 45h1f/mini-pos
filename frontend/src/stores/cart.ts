import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export interface CartItem {
  product_id: string
  name: string
  price: number // in cents
  qty: number
  line_total: number
}

export const useCartStore = defineStore('cart', () => {
  const items = ref<CartItem[]>([])
  const taxRate = ref(0.0) // E.g., 0.10 for 10%

  // Getters
  const subtotal = computed(() => {
    return items.value.reduce((sum, item) => sum + item.line_total, 0)
  })
  
  const tax = computed(() => {
    return Math.round(subtotal.value * taxRate.value)
  })
  
  const discount = ref(0)
  
  const total = computed(() => {
    return subtotal.value + tax.value - discount.value
  })

  // Actions
  const addItem = (product: any, qty: number = 1) => {
    const existing = items.value.find(i => i.product_id === product.id)
    if (existing) {
      existing.qty += qty
      existing.line_total = existing.qty * existing.price
    } else {
      items.value.push({
        product_id: product.id,
        name: product.name,
        price: product.price,
        qty: qty,
        line_total: product.price * qty
      })
    }
  }

  const removeItem = (productId: string) => {
    items.value = items.value.filter(i => i.product_id !== productId)
  }

  const updateQty = (productId: string, qty: number) => {
    const item = items.value.find(i => i.product_id === productId)
    if (item) {
      if (qty <= 0) {
        removeItem(productId)
      } else {
        item.qty = qty
        item.line_total = item.qty * item.price
      }
    }
  }

  const clearCart = () => {
    items.value = []
    discount.value = 0
  }

  return {
    items,
    taxRate,
    subtotal,
    tax,
    discount,
    total,
    addItem,
    removeItem,
    updateQty,
    clearCart
  }
})
