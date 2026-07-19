<script setup lang="ts">
import { defineProps } from 'vue'

const props = defineProps({
  saleData: {
    type: Object,
    required: true
  },
  storeName: {
    type: String,
    default: "Mini POS Demo Store"
  },
  footerText: {
    type: String,
    default: "Thank you for your business!"
  }
})

// Format a date safely
const formatDate = (date: Date) => {
  return new Intl.DateTimeFormat('default', {
    dateStyle: 'short',
    timeStyle: 'short'
  }).format(date)
}
</script>

<template>
  <div id="print-section" class="receipt-container" v-if="saleData">
    <div class="receipt-header">
      <h2>{{ storeName }}</h2>
      <p>Invoice: #{{ saleData.invoiceNo || 'PENDING' }}</p>
      <p>Date: {{ formatDate(new Date()) }}</p>
    </div>
    
    <div class="divider">--------------------------------</div>
    
    <div class="receipt-items">
      <div v-for="(item, index) in saleData.items" :key="index" class="receipt-item-row">
        <div class="item-name">{{ item.name }}</div>
        <div class="item-calc">
          {{ item.qty }} x ${{ (item.price / 100).toFixed(2) }}
          <span class="item-line-total">${{ (item.line_total / 100).toFixed(2) }}</span>
        </div>
      </div>
    </div>
    
    <div class="divider">--------------------------------</div>
    
    <div class="receipt-totals">
      <div class="total-line"><span>Subtotal:</span> <span>${{ (saleData.subtotal / 100).toFixed(2) }}</span></div>
      <div class="total-line"><span>Tax:</span> <span>${{ (saleData.tax / 100).toFixed(2) }}</span></div>
      <div class="total-line"><span>Discount:</span> <span>-${{ (saleData.discount / 100).toFixed(2) }}</span></div>
      <div class="total-line grand-total"><span>Total:</span> <span>${{ (saleData.total / 100).toFixed(2) }}</span></div>
    </div>
    
    <div class="divider">================================</div>
    
    <div class="receipt-footer">
      <p>{{ footerText }}</p>
    </div>
  </div>
</template>

<style scoped>
/* These styles apply to the print preview */
.receipt-container {
  display: flex;
  flex-direction: column;
  align-items: stretch;
}

.receipt-header, .receipt-footer {
  text-align: center;
}

.receipt-header h2 {
  margin: 0 0 5px 0;
  font-size: 16px;
}

.receipt-header p, .receipt-footer p {
  margin: 2px 0;
}

.divider {
  text-align: center;
  margin: 8px 0;
  overflow: hidden;
  white-space: nowrap;
}

.receipt-item-row {
  margin-bottom: 6px;
}

.item-name {
  font-weight: bold;
}

.item-calc {
  display: flex;
  justify-content: space-between;
  padding-left: 10px;
}

.item-line-total {
  text-align: right;
}

.receipt-totals {
  display: flex;
  flex-direction: column;
}

.total-line {
  display: flex;
  justify-content: space-between;
  margin-bottom: 3px;
}

.grand-total {
  font-weight: bold;
  font-size: 14px;
  margin-top: 5px;
}
</style>
