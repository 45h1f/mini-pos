<script setup lang="ts">
import { ref, onMounted } from 'vue'

const isDarkMode = ref(false)

const toggleTheme = () => {
  isDarkMode.value = !isDarkMode.value
  if (isDarkMode.value) {
    document.documentElement.classList.add('dark-theme')
  } else {
    document.documentElement.classList.remove('dark-theme')
  }
}

onMounted(() => {
  // Check system preference
  if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
    toggleTheme()
  }
})
</script>

<template>
  <div class="app-container">
    <aside class="sidebar">
      <div class="logo">
        <h2>Mini POS</h2>
      </div>
      <nav class="nav-links">
        <router-link to="/" class="nav-item">Dashboard</router-link>
        <router-link to="/sales" class="nav-item">Sales / POS</router-link>
        <router-link to="/inventory" class="nav-item">Inventory</router-link>
        <router-link to="/customers" class="nav-item">Customers</router-link>
        <router-link to="/settings" class="nav-item">Settings</router-link>
      </nav>
      <div class="sidebar-footer">
        <button class="theme-toggle" @click="toggleTheme">
          {{ isDarkMode ? '☀️ Light Mode' : '🌙 Dark Mode' }}
        </button>
      </div>
    </aside>

    <main class="main-content">
      <header class="header">
        <h1>{{ $route.name }}</h1>
        <div class="user-info">
          <span>Demo User</span>
        </div>
      </header>
      
      <div class="page-content">
        <router-view />
      </div>
    </main>
  </div>
</template>

<style scoped>
.app-container {
  display: flex;
  height: 100vh;
  width: 100vw;
  overflow: hidden;
}

.sidebar {
  width: 250px;
  background-color: var(--sidebar-bg);
  color: var(--sidebar-text);
  display: flex;
  flex-direction: column;
  border-right: 1px solid var(--border-color);
  transition: all 0.3s ease;
}

.logo {
  padding: 1.5rem;
  border-bottom: 1px solid var(--border-color);
}

.logo h2 {
  margin: 0;
  font-size: 1.5rem;
  font-weight: 700;
  color: var(--primary-color);
}

.nav-links {
  display: flex;
  flex-direction: column;
  padding: 1rem 0;
  flex: 1;
}

.nav-item {
  padding: 0.8rem 1.5rem;
  color: var(--sidebar-text);
  text-decoration: none;
  font-weight: 500;
  transition: background-color 0.2s, color 0.2s;
}

.nav-item:hover, .nav-item.router-link-active {
  background-color: var(--primary-color-light);
  color: var(--primary-color);
  border-right: 3px solid var(--primary-color);
}

.sidebar-footer {
  padding: 1rem;
  border-top: 1px solid var(--border-color);
}

.theme-toggle {
  width: 100%;
  padding: 0.8rem;
  background-color: transparent;
  color: var(--sidebar-text);
  border: 1px solid var(--border-color);
  border-radius: 8px;
  cursor: pointer;
  font-weight: 600;
  transition: all 0.2s;
}

.theme-toggle:hover {
  background-color: var(--hover-bg);
}

.main-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  background-color: var(--bg-color);
  color: var(--text-color);
  overflow: hidden;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem 2rem;
  background-color: var(--header-bg);
  border-bottom: 1px solid var(--border-color);
}

.header h1 {
  margin: 0;
  font-size: 1.25rem;
  font-weight: 600;
}

.user-info {
  font-weight: 500;
}

.page-content {
  flex: 1;
  padding: 2rem;
  overflow-y: auto;
}
</style>
