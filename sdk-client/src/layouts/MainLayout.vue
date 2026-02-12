<template>
  <div class="main-layout">
    <!-- 头部 -->
    <header class="main-header">
      <div class="header-left">
        <div class="logo-glow">
          <h1 class="app-title">CONSOLE</h1>
          <div class="subtitle">控制台</div>
        </div>
      </div>
      
      <div class="header-right">
        <div class="user-panel">
          <div class="user-info" v-if="userStore.user">
            <div class="user-avatar">
              <span>{{ userStore.user.nickname?.charAt(0) || userStore.user.username.charAt(0) }}</span>
            </div>
            <div class="user-details">
              <div class="welcome-text">WELCOME, {{ (userStore.user.nickname || userStore.user.username).toUpperCase() }}</div>
              <div class="user-status">STATUS: <span class="online-indicator">ONLINE</span></div>
            </div>
          </div>
          <button class="logout-btn" @click="handleLogout">
            <span>LOGOUT</span>
          </button>
        </div>
      </div>
    </header>

    <div class="layout-container">
      <!-- 左侧导航栏 -->
      <nav class="sidebar-nav">
        <div class="nav-items">
          <button 
            v-for="tab in tabs" 
            :key="tab.key"
            :class="['nav-item', { active: activeTab === tab.key }]"
            @click="switchTab(tab.key)"
          >
            <span class="nav-text">{{ tab.label.toUpperCase() }}</span>
            <div class="nav-glow"></div>
          </button>
        </div>
      </nav>

      <!-- 主内容区域 -->
      <main class="main-content">
        <div class="content-wrapper">
          <!-- 浏览器环境管理 -->
          <div v-if="activeTab === 'environments'" class="tab-content">
            <div class="content-header">
              <div class="header-section">
                <h2>ENVIRONMENT MANAGEMENT</h2>
                <p>浏览器环境配置与监控</p>
              </div>
              <button class="primary-btn" @click="showCreateModal = true">
                <span>CREATE NEW ENVIRONMENT</span>
              </button>
            </div>
            
            <!-- 加载状态 -->
            <div v-if="browserStore.loading" class="loading-state">
              <div class="spinner"></div>
              <p>LOADING ENVIRONMENTS...</p>
            </div>
            
            <div v-else class="environments-grid">
              <div 
                v-for="browser in browserStore.browsers" 
                :key="browser.id"
                :class="getEnvironmentCardClass(browser.status)"
              >
                <div class="card-header">
                  <div class="header-content">
                    <h3>{{ browser.name.toUpperCase() }}</h3>
                  </div>
                  <div class="card-actions">
                    <button 
                      v-if="browser.status === 1" 
                      class="action-btn start" 
                      @click="startBrowser(browser.id)"
                      :disabled="isActionLoading"
                    >
                      START
                    </button>
                    <button 
                      v-else-if="browser.status === 3" 
                      class="action-btn stop" 
                      @click="stopBrowser(browser.id)"
                      :disabled="isActionLoading"
                    >
                      STOP
                    </button>
                    <button class="action-btn edit" @click="editEnvironment(browser)">
                      EDIT
                    </button>
                    <button class="action-btn delete" @click="deleteEnvironment(browser.id)">
                      DELETE
                    </button>
                  </div>
                </div>
                <div class="card-body">
                  <div class="info-grid">
                    <div class="info-item">
                      <span class="label">CONFIGURATION:</span>
                      <span class="value">{{ browser.data || 'No configuration' }}</span>
                    </div>
                    <div class="info-item">
                      <span class="label">ENV ID:</span>
                      <span class="value">{{ browser.envId }}</span>
                    </div>
                    <div class="info-item">
                      <span class="label">CREATED:</span>
                      <span class="value">{{ formatDate(browser.createdAt) }}</span>
                    </div>
                    <div class="info-item">
                      <span class="label">UPDATED:</span>
                      <span class="value">{{ formatDate(browser.updatedAt) }}</span>
                    </div>
                  </div>
                </div>
              </div>
              
              <!-- 空状态 -->
              <div v-if="browserStore.browsers.length === 0 && !browserStore.loading" class="empty-state">
                <div class="empty-icon">⚙️</div>
                <p>NO ENVIRONMENTS CONFIGURED</p>
                <button class="primary-btn" @click="showCreateModal = true">
                  CREATE FIRST ENVIRONMENT
                </button>
              </div>
            </div>
          </div>

          <!-- 系统设置 -->
          <div v-if="activeTab === 'settings'" class="tab-content">
            <div class="content-header">
              <h2>SYSTEM SETTINGS</h2>
              <p>系统配置与管理</p>
            </div>
            <div class="settings-panel">
              <div class="setting-card">
                <h3>SYSTEM STATUS</h3>
                <div class="status-grid">
                  <div class="status-item">
                    <span class="status-label">CPU USAGE:</span>
                    <div class="progress-bar">
                      <div class="progress-fill" style="width: 65%"></div>
                    </div>
                  </div>
                  <div class="status-item">
                    <span class="status-label">MEMORY:</span>
                    <div class="progress-bar">
                      <div class="progress-fill" style="width: 42%"></div>
                    </div>
                  </div>
                  <div class="status-item">
                    <span class="status-label">STORAGE:</span>
                    <div class="progress-bar">
                      <div class="progress-fill" style="width: 28%"></div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </main>
    </div>

    <!-- 创建/编辑环境模态框 -->
    <div v-if="showCreateModal" class="modal-overlay" @click="closeModal">
      <div class="modal-content" @click.stop>
        <div class="modal-header">
          <h3>{{ editingEnvironment ? 'EDIT ENVIRONMENT' : 'CREATE NEW ENVIRONMENT' }}</h3>
          <button class="close-btn" @click="closeModal">×</button>
        </div>
        
        <form @submit.prevent="saveEnvironment" class="modal-form">
          <div class="form-group">
            <label for="envName">ENVIRONMENT NAME</label>
            <input
              id="envName"
              v-model="environmentForm.name"
              type="text"
              placeholder="ENTER ENVIRONMENT NAME"
              required
              class="tech-input"
            />
          </div>
          
          <div class="form-group">
            <label for="envData">CONFIGURATION DATA</label>
            <textarea
              id="envData"
              v-model="environmentForm.data"
              placeholder="ENTER ENVIRONMENT CONFIGURATION DATA"
              rows="6"
              class="tech-input"
            ></textarea>
          </div>
          
          <div class="modal-actions">
            <button type="button" class="secondary-btn" @click="closeModal">
              CANCEL
            </button>
            <button type="submit" class="primary-btn" :disabled="isSaving">
              {{ isSaving ? 'PROCESSING...' : (editingEnvironment ? 'UPDATE' : 'CREATE') }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, computed } from 'vue';
import { useUserStore } from '@/stores/user';
import { useBrowserStore } from '@/stores/browser';
import { useRouter } from 'vue-router';

interface Browser {
  id: number;
  name: string;
  envId: number;
  userId: number;
  data: string;
  createdAt: string;
  updatedAt: string;
  status: number;  // 1: 停止, 3: 启动
}

interface EnvironmentForm {
  name: string;
  data: string;
}

const userStore = useUserStore();
const browserStore = useBrowserStore();
const router = useRouter();

const activeTab = ref('environments');
const showCreateModal = ref(false);
const editingEnvironment = ref<Browser | null>(null);
const isSaving = ref(false);
const isActionLoading = ref(false);  // 添加动作加载状态

// 窗口高度相关
const windowHeight = ref(window.innerHeight);
const headerHeight = ref(80); // 头部高度估算

// 计算主内容区域可用高度
const mainContentHeight = computed(() => {
  return `${windowHeight.value - headerHeight.value}px`;
});

// 监听窗口大小变化
const handleResize = () => {
  windowHeight.value = window.innerHeight;
};

onMounted(async () => {
  window.addEventListener('resize', handleResize);
  handleResize(); // 初始化
  
  if (!userStore.isAuthenticated) {
    router.push('/login');
  } else {
    // 恢复localStorage中的用户信息
    userStore.restoreUserInfo();
    // 确保用户信息是最新的
    userStore.loadUserInfo();
    // 加载浏览器环境列表
    try {
      await browserStore.loadBrowsers();
    } catch (error) {
      console.error('Failed to load environments:', error);
    }
  }
});

onUnmounted(() => {
  window.removeEventListener('resize', handleResize);
});

const tabs = [
  { key: 'environments', label: '环境管理' },
  { key: 'settings', label: '系统设置' }
];

const environmentForm = ref<EnvironmentForm>({
  name: '',
  data: ''
});

// 根据status值获取环境卡片的CSS类
const getEnvironmentCardClass = (status: number): string => {
  const baseClass = 'environment-card';
  switch (status) {
    case 1:  // 停止状态
      return `${baseClass} environment-stopped`;
    case 3:  // 启动状态
      return `${baseClass} environment-active`;
    default:
      return `${baseClass} environment-unknown`;
  }
};

// 启动浏览器
const startBrowser = async (id: number) => {
  if (isActionLoading.value) return;
  
  isActionLoading.value = true;
  try {
    // 这里需要调用启动浏览器的API
    // 暂时模拟API调用
    await new Promise(resolve => setTimeout(resolve, 1000));
    
    // 更新浏览器状态
    const browser = browserStore.browsers.find(b => b.id === id);
    if (browser) {
      browser.status = 3;  // 设置为启动状态
    }
    
    console.log(`Browser ${id} started successfully`);
  } catch (error) {
    console.error('Failed to start browser:', error);
    alert('启动浏览器失败，请重试');
  } finally {
    isActionLoading.value = false;
  }
};

// 停止浏览器
const stopBrowser = async (id: number) => {
  if (isActionLoading.value) return;
  
  isActionLoading.value = true;
  try {
    // 这里需要调用停止浏览器的API
    // 暂时模拟API调用
    await new Promise(resolve => setTimeout(resolve, 1000));
    
    // 更新浏览器状态
    const browser = browserStore.browsers.find(b => b.id === id);
    if (browser) {
      browser.status = 1;  // 设置为停止状态
    }
    
    console.log(`Browser ${id} stopped successfully`);
  } catch (error) {
    console.error('Failed to stop browser:', error);
    alert('停止浏览器失败，请重试');
  } finally {
    isActionLoading.value = false;
  }
};

const switchTab = (tabKey: string) => {
  activeTab.value = tabKey;
};

const handleLogout = () => {
  userStore.logout();
  router.push('/login');
};

const showModal = () => {
  showCreateModal.value = true;
  editingEnvironment.value = null;
  environmentForm.value = { name: '', data: '' };
};

const closeModal = () => {
  showCreateModal.value = false;
  editingEnvironment.value = null;
  environmentForm.value = { name: '', data: '' };
  isSaving.value = false;
};

const editEnvironment = (browser: Browser) => {
  editingEnvironment.value = browser;
  environmentForm.value = {
    name: browser.name,
    data: browser.data
  };
  showCreateModal.value = true;
};

const saveEnvironment = async () => {
  if (!environmentForm.value.name.trim()) return;
  
  isSaving.value = true;
  
  try {
    if (editingEnvironment.value) {
      // 更新现有环境
      await browserStore.updateBrowser({
        id: editingEnvironment.value.id,
        name: environmentForm.value.name,
        userId: userStore.user?.id || 0,
        data: environmentForm.value.data
      });
    } else {
      // 创建新环境
      await browserStore.createBrowser({
        name: environmentForm.value.name,
        userId: userStore.user?.id || 0,
        data: environmentForm.value.data
      });
    }
    
    closeModal();
  } catch (error) {
    console.error('Failed to save environment:', error);
    alert('操作失败，请重试');
  } finally {
    isSaving.value = false;
  }
};

const deleteEnvironment = async (id: number) => {
  if (confirm('确定要删除这个环境吗？此操作不可撤销。')) {
    try {
      await browserStore.deleteBrowser(id);
    } catch (error) {
      console.error('Failed to delete environment:', error);
      alert('删除失败，请重试');
    }
  }
};

const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleString('zh-CN');
};

onMounted(async () => {
  if (!userStore.isAuthenticated) {
    router.push('/login');
  } else {
    // 恢复localStorage中的用户信息
    userStore.restoreUserInfo();
    // 确保用户信息是最新的
    userStore.loadUserInfo();
    // 加载浏览器环境列表
    try {
      await browserStore.loadBrowsers();
    } catch (error) {
      console.error('Failed to load environments:', error);
    }
  }
});
</script>

<style scoped>
/* 隐藏滚动条的全局样式 */
::-webkit-scrollbar {
  width: 0px;
  height: 0px;
  background: transparent;
}

::-webkit-scrollbar-track {
  background: transparent;
}

::-webkit-scrollbar-thumb {
  background: transparent;
}

/* Firefox 隐藏滚动条 */
html {
  scrollbar-width: none;
}

/* IE/Edge 隐藏滚动条 */
body {
  -ms-overflow-style: none;
}

.main-layout {
  min-height: 100vh;
  height: 100vh; /* 固定高度以适应窗口 */
  background: #0a0a0a;
  font-family: 'Courier New', monospace;
  position: relative;
  overflow: hidden; /* 防止整体滚动 */
}


/* 背景网格效果 */
.main-layout::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-image: 
    linear-gradient(rgba(0, 255, 255, 0.03) 1px, transparent 1px),
    linear-gradient(90deg, rgba(0, 255, 255, 0.03) 1px, transparent 1px);
  background-size: 40px 40px;
  animation: gridMove 30s linear infinite;
  pointer-events: none;
  z-index: 0;
}

@keyframes gridMove {
  0% { transform: translate(0, 0); }
  100% { transform: translate(40px, 40px); }
}

.main-header {
  background: rgba(15, 15, 25, 0.95);
  backdrop-filter: blur(15px);
  padding: 20px 40px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: 0 0 20px rgba(0, 255, 255, 0.1);
  border-bottom: 1px solid rgba(0, 255, 255, 0.2);
  position: relative;
  z-index: 10;
  flex-shrink: 0; /* 防止头部被压缩 */
}

.header-left .logo-glow {
  position: relative;
}

.app-title {
  color: #00ffff;
  font-size: 20px;
  font-weight: 600;
  margin: 0 0 8px 0;
  letter-spacing: 2px;
  text-shadow: 0 0 10px rgba(0, 255, 255, 0.7);
}

.subtitle {
  color: rgba(255, 255, 255, 0.6);
  font-size: 12px;
  letter-spacing: 1px;
}

.user-panel {
  display: flex;
  align-items: center;
  gap: 25px;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 15px;
  padding: 12px 20px;
  background: rgba(5, 5, 15, 0.6);
  border-radius: 8px;
  border: 1px solid rgba(0, 255, 255, 0.2);
}

.user-avatar {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  background: linear-gradient(45deg, #00ffff, #8000ff);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-weight: bold;
  font-size: 16px;
}

.user-details {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.welcome-text {
  color: #00ffff;
  font-size: 12px;
  letter-spacing: 1px;
}

.user-status {
  color: rgba(255, 255, 255, 0.7);
  font-size: 11px;
}

.online-indicator {
  color: #00ff00;
  font-weight: bold;
}

.logout-btn {
  background: linear-gradient(45deg, rgba(255, 107, 107, 0.2), rgba(238, 90, 82, 0.2));
  color: #ff6b6b;
  border: 1px solid rgba(255, 107, 107, 0.4);
  padding: 10px 20px;
  border-radius: 6px;
  cursor: pointer;
  font-size: 12px;
  font-weight: 500;
  letter-spacing: 1px;
  transition: all 0.3s ease;
  font-family: 'Courier New', monospace;
}

.logout-btn:hover {
  background: linear-gradient(45deg, rgba(255, 107, 107, 0.3), rgba(238, 90, 82, 0.3));
  box-shadow: 0 0 15px rgba(255, 107, 107, 0.4);
  transform: translateY(-1px);
}

.main-nav {
  background: rgba(10, 10, 20, 0.9);
  padding: 0 40px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.3);
  position: relative;
  z-index: 10;
}

.nav-tabs {
  display: flex;
  gap: 0;
  position: relative;
}

.nav-tab {
  padding: 18px 30px;
  border: none;
  background: transparent;
  color: rgba(255, 255, 255, 0.6);
  font-size: 13px;
  cursor: pointer;
  position: relative;
  transition: all 0.3s ease;
  font-family: 'Courier New', monospace;
  letter-spacing: 1px;
  overflow: hidden;
}

.nav-tab:hover {
  color: #00ffff;
  background: rgba(0, 255, 255, 0.05);
}

.nav-tab.active {
  color: #00ffff;
}

.tab-glow {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(90deg, transparent, rgba(0, 255, 255, 0.2), transparent);
  opacity: 0;
  transition: opacity 0.3s ease;
}

.nav-tab:hover .tab-glow {
  opacity: 1;
}

.nav-tab.active::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  height: 2px;
  background: linear-gradient(90deg, #00ffff, #8000ff);
  box-shadow: 0 0 10px rgba(0, 255, 255, 0.5);
}

.layout-container {
  display: flex;
  height: calc(100vh - v-bind(headerHeight) * 1px); /* 动态计算剩余高度 */
  flex-shrink: 0;
}

.sidebar-nav {
  width: 220px;
  background: rgba(10, 10, 20, 0.95);
  backdrop-filter: blur(15px);
  border-right: 1px solid rgba(0, 255, 255, 0.2);
  box-shadow: 5px 0 15px rgba(0, 0, 0, 0.3);
  padding: 25px 0;
  position: relative;
  z-index: 10;
  flex-shrink: 0;
  overflow-y: auto; /* 侧边栏内部可滚动 */
}

.nav-items {
  display: flex;
  flex-direction: column;
  gap: 5px;
}

.nav-item {
  padding: 16px 25px;
  border: none;
  background: transparent;
  color: rgba(255, 255, 255, 0.6);
  font-size: 13px;
  cursor: pointer;
  position: relative;
  transition: all 0.3s ease;
  font-family: 'Courier New', monospace;
  letter-spacing: 1px;
  text-align: left;
  width: 100%;
  overflow: hidden;
}

.nav-item:hover {
  color: #00ffff;
  background: rgba(0, 255, 255, 0.08);
  transform: translateX(5px);
}

.nav-item.active {
  color: #00ffff;
  background: rgba(0, 255, 255, 0.12);
  border-left: 3px solid #00ffff;
}

.nav-glow {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(90deg, transparent, rgba(0, 255, 255, 0.15), transparent);
  opacity: 0;
  transition: opacity 0.3s ease;
}

.nav-item:hover .nav-glow {
  opacity: 1;
}

.nav-item.active::after {
  content: '';
  position: absolute;
  top: 0;
  right: 0;
  bottom: 0;
  width: 2px;
  background: linear-gradient(180deg, #00ffff, #8000ff);
  box-shadow: 0 0 10px rgba(0, 255, 255, 0.5);
}

.main-content {
  flex: 1;
  padding: 30px 40px;
  position: relative;
  z-index: 5;
  background: rgba(5, 5, 15, 0.3);
  overflow-y: auto; /* 主内容区可滚动 */
  height: 100%; /* 占满剩余空间 */
}

.content-wrapper {
  max-width: 1200px;
  margin: 0 auto;
  min-height: 100%; /* 确保内容区域占满高度 */
}

.content-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 30px;
  padding-bottom: 20px;
  border-bottom: 1px solid rgba(0, 255, 255, 0.1);
}

.header-section h2 {
  color: #00ffff;
  font-size: 24px;
  margin: 0 0 10px 0;
  letter-spacing: 2px;
  text-shadow: 0 0 10px rgba(0, 255, 255, 0.5);
}

.header-section p {
  color: rgba(255, 255, 255, 0.6);
  font-size: 14px;
  margin: 0;
}

.primary-btn {
  background: linear-gradient(45deg, rgba(0, 255, 255, 0.15), rgba(128, 0, 255, 0.15));
  color: #00ffff;
  border: 1px solid rgba(0, 255, 255, 0.4);
  padding: 12px 24px;
  border-radius: 6px;
  cursor: pointer;
  font-size: 12px;
  font-weight: 500;
  letter-spacing: 1px;
  transition: all 0.3s ease;
  font-family: 'Courier New', monospace;
  position: relative;
  overflow: hidden;
}

.primary-btn:hover {
  background: linear-gradient(45deg, rgba(0, 255, 255, 0.25), rgba(128, 0, 255, 0.25));
  box-shadow: 0 0 20px rgba(0, 255, 255, 0.6);
  transform: translateY(-2px);
  border-color: #00ffff;
}

.environments-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(380px, 1fr));
  gap: 25px;
}

.environment-card {
  background: rgba(15, 15, 25, 0.7);
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 5px 20px rgba(0, 0, 0, 0.3);
  border: 1px solid rgba(0, 255, 255, 0.1);
  transition: all 0.3s ease;
  backdrop-filter: blur(10px);
}

/* 停止状态的环境卡片 */
.environment-stopped {
  background: rgba(240, 240, 240, 0.15);
  border: 1px solid rgba(200, 200, 200, 0.3);
  box-shadow: 0 5px 20px rgba(200, 200, 200, 0.2);
}

.environment-stopped:hover {
  background: rgba(240, 240, 240, 0.25);
  border: 1px solid rgba(200, 200, 200, 0.5);
  box-shadow: 0 10px 30px rgba(200, 200, 200, 0.3);
}

/* 启动状态的环境卡片 */
.environment-active {
  background: rgba(0, 255, 0, 0.1);
  border: 1px solid rgba(0, 255, 0, 0.3);
  box-shadow: 0 5px 20px rgba(0, 255, 0, 0.2);
}

.environment-active:hover {
  background: rgba(0, 255, 0, 0.15);
  border: 1px solid rgba(0, 255, 0, 0.5);
  box-shadow: 0 10px 30px rgba(0, 255, 0, 0.3);
}

/* 未知状态的环境卡片 */
.environment-unknown {
  background: rgba(255, 255, 0, 0.1);
  border: 1px solid rgba(255, 255, 0, 0.3);
  box-shadow: 0 5px 20px rgba(255, 255, 0, 0.2);
}

.environment-unknown:hover {
  background: rgba(255, 255, 0, 0.15);
  border: 1px solid rgba(255, 255, 0, 0.5);
  box-shadow: 0 10px 30px rgba(255, 255, 0, 0.3);
}

.environment-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 10px 30px rgba(0, 255, 255, 0.2);
  border-color: rgba(0, 255, 255, 0.3);
}

.card-header {
  padding: 20px;
  background: rgba(5, 5, 15, 0.6);
  border-bottom: 1px solid rgba(0, 255, 255, 0.1);
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-content {
  display: flex;
  align-items: center;
  gap: 15px;
}

.card-header h3 {
  margin: 0;
  color: #00ffff;
  font-size: 16px;
  letter-spacing: 1px;
}

.card-actions {
  display: flex;
  gap: 10px;
  align-items: center;
  min-width: 200px; /* 确保按钮容器有最小宽度 */
  justify-content: flex-end;
}

.action-btn {
  padding: 6px 15px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 11px;
  font-weight: 500;
  letter-spacing: 1px;
  transition: all 0.2s ease;
  font-family: 'Courier New', monospace;
  min-width: 60px;
  text-align: center;
  height: 28px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.action-btn.start {
  background: rgba(0, 255, 0, 0.2);
  color: #00ff00;
  border: 1px solid rgba(0, 255, 0, 0.4);
}

.action-btn.start:hover:not(:disabled) {
  background: rgba(0, 255, 0, 0.3);
  box-shadow: 0 0 10px rgba(0, 255, 0, 0.4);
  transform: translateY(-1px);
}

.action-btn.stop {
  background: rgba(255, 0, 0, 0.2);
  color: #ff0000;
  border: 1px solid rgba(255, 0, 0, 0.4);
}

.action-btn.stop:hover:not(:disabled) {
  background: rgba(255, 0, 0, 0.3);
  box-shadow: 0 0 10px rgba(255, 0, 0, 0.4);
  transform: translateY(-1px);
}

.action-btn.edit {
  background: rgba(23, 162, 184, 0.2);
  color: #17a2b8;
  border: 1px solid rgba(23, 162, 184, 0.4);
}

.action-btn.edit:hover {
  background: rgba(23, 162, 184, 0.3);
  box-shadow: 0 0 10px rgba(23, 162, 184, 0.4);
}

.action-btn.delete {
  background: rgba(220, 53, 69, 0.2);
  color: #dc3545;
  border: 1px solid rgba(220, 53, 69, 0.4);
}

.action-btn.delete:hover {
  background: rgba(220, 53, 69, 0.3);
  box-shadow: 0 0 10px rgba(220, 53, 69, 0.4);
}

.action-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
  transform: none !important;
}

.card-body {
  padding: 25px;
}

.info-grid {
  display: grid;
  gap: 15px;
}

.info-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.label {
  color: rgba(0, 255, 255, 0.7);
  font-size: 12px;
  letter-spacing: 1px;
}

.value {
  color: rgba(255, 255, 255, 0.9);
  font-size: 12px;
  font-weight: 500;
}

.empty-state {
  grid-column: 1 / -1;
  text-align: center;
  padding: 80px 20px;
  color: rgba(255, 255, 255, 0.6);
}

.empty-icon {
  font-size: 48px;
  margin-bottom: 20px;
  filter: grayscale(100%) brightness(0.5);
}

.empty-state p {
  font-size: 16px;
  margin-bottom: 25px;
  letter-spacing: 1px;
}

.settings-panel {
  display: grid;
  gap: 25px;
}

.setting-card {
  background: rgba(15, 15, 25, 0.7);
  border-radius: 12px;
  padding: 25px;
  border: 1px solid rgba(0, 255, 255, 0.1);
  backdrop-filter: blur(10px);
}

.setting-card h3 {
  color: #00ffff;
  font-size: 18px;
  margin: 0 0 20px 0;
  letter-spacing: 1px;
}

.status-grid {
  display: grid;
  gap: 20px;
}

.status-item {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.status-label {
  color: rgba(255, 255, 255, 0.7);
  font-size: 12px;
  letter-spacing: 1px;
}

.progress-bar {
  width: 100%;
  height: 8px;
  background: rgba(255, 255, 255, 0.1);
  border-radius: 4px;
  overflow: hidden;
}

.progress-fill {
  height: 100%;
  background: linear-gradient(90deg, #00ffff, #8000ff);
  border-radius: 4px;
  transition: width 0.5s ease;
  box-shadow: 0 0 10px rgba(0, 255, 255, 0.5);
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.8);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  backdrop-filter: blur(5px);
}

.modal-content {
  background: rgba(15, 15, 25, 0.95);
  border-radius: 15px;
  padding: 30px;
  width: 90%;
  max-width: 500px;
  border: 1px solid rgba(0, 255, 255, 0.3);
  box-shadow: 0 0 30px rgba(0, 255, 255, 0.2);
  backdrop-filter: blur(20px);
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 25px;
  padding-bottom: 15px;
  border-bottom: 1px solid rgba(0, 255, 255, 0.2);
}

.modal-header h3 {
  color: #00ffff;
  font-size: 20px;
  margin: 0;
  letter-spacing: 1px;
}

.close-btn {
  background: transparent;
  border: none;
  color: rgba(255, 255, 255, 0.7);
  font-size: 24px;
  cursor: pointer;
  width: 30px;
  height: 30px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  transition: all 0.2s ease;
}

.close-btn:hover {
  background: rgba(255, 255, 255, 0.1);
  color: #ffffff;
}

.modal-form {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.form-group label {
  color: rgba(0, 255, 255, 0.8);
  font-size: 12px;
  letter-spacing: 1px;
  font-weight: 500;
}

.tech-input {
  padding: 12px 15px;
  border: 1px solid rgba(0, 255, 255, 0.3);
  border-radius: 6px;
  background: rgba(5, 5, 15, 0.6);
  color: #ffffff;
  font-size: 14px;
  font-family: 'Courier New', monospace;
  transition: all 0.3s ease;
}

.tech-input::placeholder {
  color: rgba(255, 255, 255, 0.4);
  font-style: italic;
}

.tech-input:focus {
  outline: none;
  border-color: #00ffff;
  box-shadow: 0 0 15px rgba(0, 255, 255, 0.4);
}

.modal-actions {
  display: flex;
  gap: 15px;
  justify-content: flex-end;
  margin-top: 10px;
}

.secondary-btn {
  background: rgba(255, 255, 255, 0.1);
  color: rgba(255, 255, 255, 0.8);
  border: 1px solid rgba(255, 255, 255, 0.3);
  padding: 10px 20px;
  border-radius: 6px;
  cursor: pointer;
  font-size: 12px;
  font-weight: 500;
  letter-spacing: 1px;
  transition: all 0.3s ease;
  font-family: 'Courier New', monospace;
}

.secondary-btn:hover {
  background: rgba(255, 255, 255, 0.2);
  border-color: rgba(255, 255, 255, 0.5);
}

.loading-state {
  text-align: center;
  padding: 60px 20px;
  color: rgba(255, 255, 255, 0.7);
}

.spinner {
  width: 40px;
  height: 40px;
  border: 3px solid rgba(0, 255, 255, 0.3);
  border-top: 3px solid #00ffff;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin: 0 auto 20px;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

.loading-state p {
  font-size: 16px;
  letter-spacing: 1px;
  margin: 0;
}

/* 响应式设计 - 自适应大小 */
@media screen and (max-width: 1400px) {
  .content-wrapper {
    max-width: 1000px;
    padding: 0 20px;
  }
  
  .environments-grid {
    grid-template-columns: repeat(auto-fill, minmax(350px, 1fr));
    gap: 20px;
  }
  
  .main-content {
    padding: 25px 30px;
  }
}

@media screen and (max-width: 1200px) {
  .content-wrapper {
    max-width: 900px;
  }
  
  .environments-grid {
    grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
    gap: 15px;
  }
  
  .main-content {
    padding: 20px 25px;
  }
  
  .header-section h2 {
    font-size: 22px;
  }
  
  .app-title {
    font-size: 18px;
  }
}

@media screen and (max-width: 992px) {
  .layout-container {
    flex-direction: column;
  }
  
  .sidebar-nav {
    width: 100%;
    height: auto;
    border-right: none;
    border-bottom: 1px solid rgba(0, 255, 255, 0.2);
    padding: 15px 0;
  }
  
  .nav-items {
    flex-direction: row;
    justify-content: center;
    flex-wrap: wrap;
    gap: 10px;
  }
  
  .nav-item {
    padding: 12px 20px;
    flex: 1;
    text-align: center;
    min-width: 120px;
  }
  
  .content-wrapper {
    max-width: 100%;
    padding: 0 15px;
  }
  
  .environments-grid {
    grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
    gap: 15px;
  }
  
  .main-content {
    padding: 20px 15px;
  }
  
  .content-header {
    flex-direction: column;
    gap: 15px;
    align-items: stretch;
  }
  
  .header-section {
    text-align: center;
  }
  
  .primary-btn {
    align-self: center;
  }
}

@media screen and (max-width: 768px) {
  .main-header {
    padding: 15px 20px;
    flex-direction: column;
    gap: 15px;
  }
  
  .header-left {
    text-align: center;
  }
  
  .user-panel {
    justify-content: center;
    width: 100%;
  }
  
  .environments-grid {
    grid-template-columns: 1fr;
    gap: 15px;
  }
  
  .environment-card {
    margin: 0 10px;
  }
  
  .card-header {
    flex-direction: column;
    gap: 15px;
    align-items: stretch;
  }
  
  .header-content {
    justify-content: center;
  }
  
  .card-actions {
    justify-content: center;
    min-width: unset;
    width: 100%;
  }
  
  .action-btn {
    flex: 1;
    min-width: unset;
  }
  
  .info-grid {
    grid-template-columns: 1fr;
    gap: 10px;
  }
  
  .modal-content {
    width: 95%;
    padding: 20px;
    margin: 10px;
  }
  
  .modal-header h3 {
    font-size: 18px;
  }
  
  .settings-panel {
    grid-template-columns: 1fr;
  }
  
  .status-grid {
    grid-template-columns: 1fr;
  }
}

@media screen and (max-width: 576px) {
  .main-header {
    padding: 12px 15px;
  }
  
  .app-title {
    font-size: 16px;
    letter-spacing: 1px;
  }
  
  .subtitle {
    font-size: 10px;
  }
  
  .user-info {
    padding: 8px 15px;
    gap: 10px;
  }
  
  .user-avatar {
    width: 30px;
    height: 30px;
    font-size: 14px;
  }
  
  .welcome-text {
    font-size: 11px;
  }
  
  .user-status {
    font-size: 10px;
  }
  
  .logout-btn {
    padding: 8px 15px;
    font-size: 11px;
  }
  
  .environments-grid {
    gap: 12px;
  }
  
  .environment-card {
    margin: 0 5px;
  }
  
  .card-header {
    padding: 15px;
  }
  
  .card-header h3 {
    font-size: 14px;
  }
  
  .card-body {
    padding: 15px;
  }
  
  .action-btn {
    padding: 5px 10px;
    font-size: 10px;
    height: 24px;
  }
  
  .info-item {
    flex-direction: column;
    align-items: flex-start;
    gap: 3px;
  }
  
  .label, .value {
    font-size: 11px;
  }
  
  .modal-content {
    padding: 15px;
  }
  
  .modal-header {
    margin-bottom: 20px;
    padding-bottom: 12px;
  }
  
  .modal-header h3 {
    font-size: 16px;
  }
  
  .form-group label {
    font-size: 11px;
  }
  
  .tech-input {
    padding: 10px 12px;
    font-size: 13px;
  }
  
  .modal-actions {
    flex-direction: column;
    gap: 10px;
  }
  
  .primary-btn, .secondary-btn {
    width: 100%;
    padding: 12px;
  }
}

/* 超小屏幕优化 */
@media screen and (max-width: 400px) {
  .main-header {
    padding: 10px 12px;
  }
  
  .app-title {
    font-size: 14px;
  }
  
  .environments-grid {
    gap: 10px;
  }
  
  .environment-card {
    margin: 0;
  }
  
  .card-header {
    padding: 12px;
  }
  
  .card-body {
    padding: 12px;
  }
  
  .action-btn {
    padding: 4px 8px;
    font-size: 9px;
    height: 22px;
  }
  
  .modal-content {
    padding: 12px;
    margin: 5px;
  }
  
  .close-btn {
    width: 25px;
    height: 25px;
    font-size: 20px;
  }
}

/* 高度自适应基础样式 */
.main-layout {
  height: 100vh;
  overflow: hidden;
}

/* 横屏手机优化 */
@media screen and (max-width: 900px) and (orientation: landscape) {
  .layout-container {
    flex-direction: row;
  }
  
  .sidebar-nav {
    width: 180px;
    height: 100vh;
    border-right: 1px solid rgba(0, 255, 255, 0.2);
    border-bottom: none;
    padding: 20px 0;
  }
  
  .nav-items {
    flex-direction: column;
    justify-content: flex-start;
  }
  
  .nav-item {
    padding: 14px 15px;
    text-align: left;
    min-width: unset;
  }
  
  .main-content {
    padding: 20px;
  }
}

/* 高分辨率屏幕优化 */
@media screen and (min-width: 2000px) {
  .content-wrapper {
    max-width: 1600px;
  }
  
  .environments-grid {
    grid-template-columns: repeat(auto-fill, minmax(420px, 1fr));
  }
  
  .environment-card {
    padding: 25px;
  }
  
  .card-header {
    padding: 25px;
  }
  
  .card-body {
    padding: 25px;
  }
}

/* 响应式设计 - 高度自适应优化 */
@media screen and (max-height: 700px) {
  .main-header {
    padding: 15px 30px;
  }
  
  .layout-container {
    height: calc(100vh - 60px); /* 调整头部高度计算 */
  }
  
  .sidebar-nav {
    padding: 15px 0;
  }
  
  .nav-item {
    padding: 12px 20px;
  }
  
  .main-content {
    padding: 20px 30px;
  }
  
  .content-header {
    margin-bottom: 20px;
    padding-bottom: 15px;
  }
}

@media screen and (max-height: 600px) {
  .main-header {
    padding: 12px 25px;
  }
  
  .layout-container {
    height: calc(100vh - 50px);
  }
  
  .sidebar-nav {
    padding: 10px 0;
  }
  
  .nav-item {
    padding: 10px 18px;
    font-size: 12px;
  }
  
  .main-content {
    padding: 15px 25px;
  }
  
  .environments-grid {
    gap: 15px;
  }
  
  .environment-card {
    padding: 15px;
  }
}

/* 超高屏幕优化 */
@media screen and (min-height: 1000px) {
  .layout-container {
    height: calc(100vh - 100px);
  }
  
  .main-content {
    padding: 40px 50px;
  }
  
  .content-wrapper {
    max-width: 1400px;
  }
  
  .environments-grid {
    gap: 35px;
  }
  
  .environment-card {
    padding: 30px;
  }
}

/* 宽度响应式保持不变 */
@media screen and (max-width: 1400px) {
  .content-wrapper {
    max-width: 1000px;
    padding: 0 20px;
  }
  
  .environments-grid {
    grid-template-columns: repeat(auto-fill, minmax(350px, 1fr));
    gap: 20px;
  }
  
  .main-content {
    padding: 25px 30px;
  }
}

@media screen and (max-width: 1200px) {
  .content-wrapper {
    max-width: 900px;
  }
  
  .environments-grid {
    grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
    gap: 15px;
  }
  
  .main-content {
    padding: 20px 25px;
  }
  
  .header-section h2 {
    font-size: 22px;
  }
  
  .app-title {
    font-size: 18px;
  }
}

@media screen and (max-width: 992px) {
  .layout-container {
    flex-direction: column;
    height: calc(100vh - v-bind(headerHeight) * 1px);
  }
  
  .sidebar-nav {
    width: 100%;
    height: auto;
    border-right: none;
    border-bottom: 1px solid rgba(0, 255, 255, 0.2);
    padding: 15px 0;
    max-height: 200px; /* 限制侧边栏最大高度 */
    overflow-y: auto;
  }
  
  .nav-items {
    flex-direction: row;
    justify-content: center;
    flex-wrap: wrap;
    gap: 10px;
  }
  
  .nav-item {
    padding: 12px 20px;
    flex: 1;
    text-align: center;
    min-width: 120px;
  }
  
  .content-wrapper {
    max-width: 100%;
    padding: 0 15px;
  }
  
  .environments-grid {
    grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
    gap: 15px;
  }
  
  .main-content {
    padding: 20px 15px;
    height: calc(100% - 200px); /* 减去侧边栏高度 */
  }
  
  .content-header {
    flex-direction: column;
    gap: 15px;
    align-items: stretch;
  }
  
  .header-section {
    text-align: center;
  }
  
  .primary-btn {
    align-self: center;
  }
}

@media screen and (max-width: 768px) {
  .main-header {
    padding: 15px 20px;
    flex-direction: column;
    gap: 15px;
  }
  
  .header-left {
    text-align: center;
  }
  
  .user-panel {
    justify-content: center;
    width: 100%;
  }
  
  .environments-grid {
    grid-template-columns: 1fr;
    gap: 15px;
  }
  
  .environment-card {
    margin: 0 10px;
  }
  
  .card-header {
    flex-direction: column;
    gap: 15px;
    align-items: stretch;
  }
  
  .header-content {
    justify-content: center;
  }
  
  .card-actions {
    justify-content: center;
    min-width: unset;
    width: 100%;
  }
  
  .action-btn {
    flex: 1;
    min-width: unset;
  }
  
  .info-grid {
    grid-template-columns: 1fr;
    gap: 10px;
  }
  
  .modal-content {
    width: 95%;
    padding: 20px;
    margin: 10px;
  }
  
  .modal-header h3 {
    font-size: 18px;
  }
  
  .settings-panel {
    grid-template-columns: 1fr;
  }
  
  .status-grid {
    grid-template-columns: 1fr;
  }
}

@media screen and (max-width: 576px) {
  .main-header {
    padding: 12px 15px;
  }
  
  .app-title {
    font-size: 16px;
    letter-spacing: 1px;
  }
  
  .subtitle {
    font-size: 10px;
  }
  
  .user-info {
    padding: 8px 15px;
    gap: 10px;
  }
  
  .user-avatar {
    width: 30px;
    height: 30px;
    font-size: 14px;
  }
  
  .welcome-text {
    font-size: 11px;
  }
  
  .user-status {
    font-size: 10px;
  }
  
  .logout-btn {
    padding: 8px 15px;
    font-size: 11px;
  }
  
  .environments-grid {
    gap: 12px;
  }
  
  .environment-card {
    margin: 0 5px;
  }
  
  .card-header {
    padding: 15px;
  }
  
  .card-header h3 {
    font-size: 14px;
  }
  
  .card-body {
    padding: 15px;
  }
  
  .action-btn {
    padding: 5px 10px;
    font-size: 10px;
    height: 24px;
  }
  
  .info-item {
    flex-direction: column;
    align-items: flex-start;
    gap: 3px;
  }
  
  .label, .value {
    font-size: 11px;
  }
  
  .modal-content {
    padding: 15px;
  }
  
  .modal-header {
    margin-bottom: 20px;
    padding-bottom: 12px;
  }
  
  .modal-header h3 {
    font-size: 16px;
  }
  
  .form-group label {
    font-size: 11px;
  }
  
  .tech-input {
    padding: 10px 12px;
    font-size: 13px;
  }
  
  .modal-actions {
    flex-direction: column;
    gap: 10px;
  }
  
  .primary-btn, .secondary-btn {
    width: 100%;
    padding: 12px;
  }
}

/* 超小屏幕优化 */
@media screen and (max-width: 400px) {
  .main-header {
    padding: 10px 12px;
  }
  
  .app-title {
    font-size: 14px;
  }
  
  .environments-grid {
    gap: 10px;
  }
  
  .environment-card {
    margin: 0;
  }
  
  .card-header {
    padding: 12px;
  }
  
  .card-body {
    padding: 12px;
  }
  
  .action-btn {
    padding: 4px 8px;
    font-size: 9px;
    height: 22px;
  }
  
  .modal-content {
    padding: 12px;
    margin: 5px;
  }
  
  .close-btn {
    width: 25px;
    height: 25px;
    font-size: 20px;
  }
}

/* 高度自适应基础样式 */
.main-layout {
  height: 100vh;
  overflow: hidden;
}

/* 横屏手机优化 */
@media screen and (max-width: 900px) and (orientation: landscape) {
  .layout-container {
    flex-direction: row;
  }
  
  .sidebar-nav {
    width: 180px;
    height: 100vh;
    border-right: 1px solid rgba(0, 255, 255, 0.2);
    border-bottom: none;
    padding: 20px 0;
  }
  
  .nav-items {
    flex-direction: column;
    justify-content: flex-start;
  }
  
  .nav-item {
    padding: 14px 15px;
    text-align: left;
    min-width: unset;
  }
  
  .main-content {
    padding: 20px;
  }
}

/* 高分辨率屏幕优化 */
@media screen and (min-width: 2000px) {
  .content-wrapper {
    max-width: 1600px;
  }
  
  .environments-grid {
    grid-template-columns: repeat(auto-fill, minmax(420px, 1fr));
  }
  
  .environment-card {
    padding: 25px;
  }
  
  .card-header {
    padding: 25px;
  }
  
  .card-body {
    padding: 25px;
  }
}
</style>
