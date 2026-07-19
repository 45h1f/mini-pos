<script setup lang="ts">
import { ref } from 'vue'

// In production, these import from Wails bindings:
// import { BackupDatabase, RestoreDatabase } from '../../wailsjs/go/app/App'

const backupStatus = ref('')
const restoreStatus = ref('')

const handleBackup = async () => {
  backupStatus.value = 'Preparing backup dialog...'
  try {
    // Simulating Wails native dialog for UI development
    // const path = await BackupDatabase()
    setTimeout(() => {
      backupStatus.value = `Successfully backed up database to selected location!`
      setTimeout(() => { backupStatus.value = '' }, 4000)
    }, 1200)
  } catch (err: any) {
    backupStatus.value = 'Error: ' + err.message
  }
}

const handleRestore = async () => {
  restoreStatus.value = 'Awaiting backup file selection...'
  try {
    // Simulating Wails native dialog for UI development
    // const success = await RestoreDatabase()
    setTimeout(() => {
      restoreStatus.value = `Restore staging complete. Please restart the app.`
    }, 1500)
  } catch (err: any) {
    restoreStatus.value = 'Error: ' + err.message
  }
}
</script>

<template>
  <div class="settings-layout">
    <div class="card settings-section">
      <h2 class="section-title">Data Safety & Backup</h2>
      <p class="section-desc">Create secure offline backups of your entire SQLite POS database or restore from an existing backup file.</p>
      
      <div class="action-grid">
        <div class="action-card">
          <div class="icon">💾</div>
          <div class="action-info">
            <h3>Backup Database</h3>
            <p>Save a secure snapshot of all products, sales, and settings to your computer.</p>
            <button class="btn btn-primary" @click="handleBackup" :disabled="backupStatus.length > 0 && backupStatus !== 'Successfully backed up database to selected location!'">
              Export Backup
            </button>
            <span v-if="backupStatus" class="status-msg success-msg">{{ backupStatus }}</span>
          </div>
        </div>

        <div class="action-card danger-zone">
          <div class="icon">🔄</div>
          <div class="action-info">
            <h3>Restore Database</h3>
            <p>Overwrite current data with a previous backup. This action cannot be undone.</p>
            <button class="btn btn-danger" @click="handleRestore" :disabled="restoreStatus.length > 0">
              Import Restore
            </button>
            <span v-if="restoreStatus" class="status-msg warning-msg">{{ restoreStatus }}</span>
          </div>
        </div>
      </div>
    </div>
    
    <div class="card settings-section">
      <h2 class="section-title">Store Configuration</h2>
      <p class="section-desc">Manage receipt details and tax rates.</p>
      
      <form class="settings-form" @submit.prevent>
        <div class="form-group">
          <label>Store Name</label>
          <input type="text" value="Mini POS Demo Store" class="form-input" />
        </div>
        
        <div class="form-group">
          <label>Tax Rate (%)</label>
          <input type="number" step="0.1" value="0" class="form-input" />
        </div>
        
        <div class="form-group">
          <label>Receipt Footer Message</label>
          <textarea class="form-input" rows="3">Thank you for your business! Please come again.</textarea>
        </div>
        
        <button class="btn btn-primary">Save Settings</button>
      </form>
    </div>
  </div>
</template>

<style scoped>
.settings-layout {
  display: flex;
  flex-direction: column;
  gap: 2rem;
  max-width: 900px;
  margin: 0 auto;
}

.section-title {
  font-size: 1.4rem;
  margin-bottom: 0.5rem;
  color: var(--primary-color);
}

.section-desc {
  color: var(--sidebar-text);
  margin-bottom: 2rem;
  font-size: 0.95rem;
}

.action-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1.5rem;
}

.action-card {
  display: flex;
  gap: 1rem;
  padding: 1.5rem;
  background-color: var(--bg-color);
  border: 1px solid var(--border-color);
  border-radius: 12px;
}

.action-card .icon {
  font-size: 2.5rem;
}

.action-info h3 {
  margin-top: 0;
  margin-bottom: 0.5rem;
  font-size: 1.1rem;
}

.action-info p {
  font-size: 0.9rem;
  color: var(--sidebar-text);
  margin-bottom: 1rem;
  line-height: 1.4;
}

.btn-danger {
  background-color: #ef4444;
  color: white;
}

.btn-danger:hover {
  background-color: #dc2626;
}

.danger-zone {
  border-color: #fca5a5;
  background-color: rgba(254, 226, 226, 0.05);
}

:root.dark-theme .danger-zone {
  border-color: #7f1d1d;
  background-color: rgba(127, 29, 29, 0.1);
}

.status-msg {
  display: block;
  margin-top: 0.8rem;
  font-size: 0.85rem;
  font-weight: 500;
}

.success-msg { color: #10b981; }
.warning-msg { color: #f59e0b; }

.settings-form {
  max-width: 500px;
  display: flex;
  flex-direction: column;
  gap: 1.2rem;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.form-group label {
  font-weight: 600;
  font-size: 0.9rem;
}

.form-input {
  padding: 0.8rem;
  border-radius: 8px;
  border: 1px solid var(--border-color);
  background-color: var(--bg-color);
  color: var(--text-color);
  font-family: inherit;
  transition: all 0.2s;
}

.form-input:focus {
  outline: none;
  border-color: var(--primary-color);
  box-shadow: 0 0 0 3px var(--primary-color-light);
}
</style>
