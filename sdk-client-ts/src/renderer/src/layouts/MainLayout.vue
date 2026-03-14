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
          <div v-if="userStore.user" class="user-info">
            <div class="user-avatar">
              <span>{{ userStore.user.nickname?.charAt(0) || userStore.user.username.charAt(0) }}</span>
            </div>
            <div class="user-details">
              <div class="welcome-text">
                WELCOME, {{ (userStore.user.nickname || userStore.user.username).toUpperCase() }}
              </div>
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

            <div v-else>
              <!-- 环境表格 -->
              <div class="environments-table-container">
                <table class="environments-table">
                  <thead>
                    <tr>
                      <th>环境名称</th>
                      <th>环境ID</th>
                      <th>状态</th>
                      <th>创建时间</th>
                      <th>操作</th>
                    </tr>
                  </thead>
                  <tbody>
                    <template v-for="(browser, index) in browserStore.browsers" :key="browser?.id || index">
                      <tr v-if="index < browserStore.pageSize" :class="getEnvironmentRowClass(browser?.status!)">
                        <td>{{ browser?.envName || '' }}</td>
                        <td>{{ browser?.envId || '' }}</td>
                        <td>
                          <span :class="getStatusBadgeClass(browser?.status)">
                            {{ getStatusText(browser?.status) }}
                          </span>
                        </td>
                        <td>{{ formatDateTime(browser?.createdAt) }}</td>
                        <td>
                          <div class="table-actions">
                            <button
                              class="action-btn start small"
                              :disabled="
                                browserStore.startedDict.has(browser?.envId) ||
                                browserStore.startingDict.has(browser?.envId) ||
                                browserStore.closingDict.has(browser?.envId)
                              "
                              @click="controlEnvironment(browser, 3)"
                            >
                              启动
                            </button>
                            <button
                              class="action-btn stop small"
                              @click="controlEnvironment(browser, 1)"
                              :disabled="
                                browserStore.startingDict.has(browser?.envId) ||
                                browserStore.closingDict.has(browser?.envId) ||
                                !browserStore.startedDict.has(browser?.envId)
                              "
                            >
                              停止
                            </button>
                            <button
                              class="action-btn edit small"
                              @click="editEnvironment(browser)"
                              :disabled="!browser"
                            >
                              编辑
                            </button>
                            <button
                              class="action-btn delete small"
                              @click="deleteEnvironment(browser?.id)"
                              :disabled="!browser"
                            >
                              删除
                            </button>
                          </div>
                        </td>
                      </tr>
                    </template>
                  </tbody>
                </table>

                <!-- 空状态 -->
                <div v-if="browserStore.browsers.length === 0" class="empty-state">
                  <div class="empty-icon">⚙️</div>
                  <p>暂无环境配置</p>
                  <button class="primary-btn" @click="showCreateModal = true">创建首个环境</button>
                </div>
              </div>

              <!-- 分页组件 -->
              <div v-if="browserStore.browsers.length > 0" class="pagination-container">
                <div class="pagination-info">
                  显示第 {{ (browserStore.currentPage - 1) * browserStore.pageSize + 1 }} -
                  {{ Math.min(browserStore.currentPage * browserStore.pageSize, browserStore.total) }}
                  条， 共 {{ browserStore.total }} 条记录
                </div>
                <div class="pagination-controls">
                  <button
                    class="pagination-btn"
                    :disabled="browserStore.currentPage <= 1"
                    @click="changePage(browserStore.currentPage - 1)"
                  >
                    上一页
                  </button>

                  <div class="pagination-pages">
                    <button
                      v-for="page in getPageNumbers()"
                      :key="page"
                      :class="['pagination-page', { active: page === browserStore.currentPage }]"
                      @click="changePage(page)"
                    >
                      {{ page }}
                    </button>
                  </div>

                  <button
                    class="pagination-btn"
                    :disabled="browserStore.currentPage >= Math.ceil(browserStore.total / browserStore.pageSize)"
                    @click="changePage(browserStore.currentPage + 1)"
                  >
                    下一页
                  </button>
                </div>
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
      <div class="modal-content large" @click.stop>
        <div class="modal-header">
          <h3>{{ editingEnvironment ? '编辑环境' : '创建新环境' }}</h3>
          <button class="close-btn" @click="closeModal">×</button>
        </div>

        <form @submit.prevent="saveEnvironment" class="modal-form">
          <!-- 基本信息 -->
          <div class="form-section">
            <h4>基本信息</h4>
            <div class="form-row">
              <div class="form-group">
                <label for="envName">环境名称 *</label>
                <input
                  id="envName"
                  v-model="environmentForm.envName"
                  type="text"
                  placeholder="请输入环境名称"
                  required
                  class="tech-input"
                />
              </div>
              <div class="form-group">
                <label for="envId">环境ID</label>
                <input
                  id="envId"
                  v-model.number="environmentForm.envId"
                  type="number"
                  placeholder="自动生成"
                  class="tech-input"
                  :disabled="!!editingEnvironment"
                />
                <small v-if="!editingEnvironment" class="help-text">留空则自动生成</small>
              </div>
            </div>
          </div>

          <!-- 浏览器配置 -->
          <div class="form-section">
            <h4>浏览器配置</h4>
            <div class="form-row">
              <div class="form-group">
                <label for="ua">User Agent</label>
                <input
                  id="ua"
                  v-model="environmentForm.ua"
                  type="text"
                  placeholder="Mozilla/5.0..."
                  class="tech-input"
                />
              </div>
              <div class="form-group">
                <label for="system">操作系统</label>
                <select id="system" v-model="environmentForm.system" class="tech-input" required>
                  <option value="">请选择系统</option>
                  <option value="Windows 10">Windows 10</option>
                  <option value="Windows 11">Windows 11</option>
                  <option value="macOS">macOS</option>
                  <option value="Linux">Linux</option>
                  <option value="Android">Android</option>
                  <option value="iOS">iOS</option>
                </select>
              </div>
            </div>
            <div class="form-row">
              <div class="form-group">
                <label for="kernel">内核</label>
                <input
                  id="kernel"
                  v-model="environmentForm.kernel"
                  type="text"
                  placeholder="Chrome"
                  class="tech-input"
                  required
                />
              </div>
              <div class="form-group">
                <label for="kernelVersion">内核版本</label>
                <input
                  id="kernelVersion"
                  v-model="environmentForm.kernelVersion"
                  type="text"
                  placeholder="98.0.4758.102"
                  class="tech-input"
                  required
                />
              </div>
            </div>
          </div>

          <!-- 硬件配置 -->
          <div class="form-section">
            <h4>硬件配置</h4>
            <div class="form-row">
              <div class="form-group">
                <label for="cpu">CPU核心数</label>
                <input id="cpu" v-model.number="environmentForm.cpu" type="number" placeholder="4" class="tech-input" />
              </div>
              <div class="form-group">
                <label for="mem">内存(GB)</label>
                <input id="mem" v-model.number="environmentForm.mem" type="number" placeholder="8" class="tech-input" />
              </div>
            </div>
            <div class="form-row">
              <div class="form-group">
                <label for="deviceName">设备名称</label>
                <input
                  id="deviceName"
                  v-model="environmentForm.deviceName"
                  type="text"
                  placeholder="My Computer"
                  class="tech-input"
                />
              </div>
              <div class="form-group">
                <label for="mac">MAC地址</label>
                <input
                  id="mac"
                  v-model="environmentForm.mac"
                  type="text"
                  placeholder="00:00:00:00:00:00"
                  class="tech-input"
                />
              </div>
            </div>
          </div>

          <!-- 网络配置 -->
          <div class="form-section">
            <h4>网络配置</h4>
            <div class="form-row">
              <div class="form-group">
                <label for="publicIp">公网IP</label>
                <input
                  id="publicIp"
                  v-model="environmentForm.publicIp"
                  type="text"
                  placeholder="192.168.1.1"
                  class="tech-input"
                />
              </div>
              <div class="form-group">
                <label for="proxy">代理设置</label>
                <input
                  id="proxy"
                  v-model="environmentForm.proxy"
                  type="text"
                  placeholder="http://proxy:8080"
                  class="tech-input"
                />
              </div>
            </div>
            <div class="form-row">
              <div class="form-group">
                <label for="zone">时区</label>
                <select id="zone" v-model="environmentForm.zone" class="tech-input">
                  <option value="">请选择时区</option>
                  <option value="UTC+8">UTC+8 (中国标准时间)</option>
                  <option value="UTC+0">UTC+0 (格林威治时间)</option>
                  <option value="UTC-5">UTC-5 (美国东部时间)</option>
                  <option value="UTC+1">UTC+1 (中欧时间)</option>
                </select>
              </div>
              <div class="form-group">
                <label for="ipChannel">IP通道</label>
                <input
                  id="ipChannel"
                  v-model="environmentForm.ipChannel"
                  type="text"
                  placeholder="default"
                  class="tech-input"
                />
              </div>
            </div>
          </div>

          <!-- 功能开关 -->
          <div class="form-section form-section2">
            <h4>功能开关</h4>
            <div class="form-row">
              <div class="form-group checkbox-group">
                <label>
                  <input type="checkbox" v-model="environmentForm.hardware" true-value="1" false-value="0" />
                  硬件加速
                </label>
              </div>
              <div class="form-group checkbox-group">
                <label>
                  <input type="checkbox" v-model="environmentForm.webGl" true-value="1" false-value="0" />
                  WebGL支持
                </label>
              </div>
              <div class="form-group checkbox-group">
                <label>
                  <input type="checkbox" v-model="environmentForm.canvas" true-value="1" false-value="0" />
                  Canvas支持
                </label>
              </div>
            </div>
            <div class="form-row">
              <div class="form-group checkbox-group">
                <label>
                  <input type="checkbox" v-model="environmentForm.audioContext" true-value="1" false-value="0" />
                  音频上下文
                </label>
              </div>
              <div class="form-group checkbox-group">
                <label>
                  <input v-model="environmentForm.mediaDevice" type="checkbox" true-value="1" false-value="0" />
                  媒体设备
                </label>
              </div>
              <div class="form-group checkbox-group">
                <label>
                  <input type="checkbox" v-model="environmentForm.bluetooth" true-value="1" false-value="0" />
                  蓝牙支持
                </label>
              </div>
            </div>
          </div>

          <!-- 其他配置 -->
          <div class="form-section">
            <h4>其他配置</h4>
            <div class="form-row">
              <div class="form-group">
                <label for="customerId">客户ID</label>
                <input
                  id="customerId"
                  v-model="environmentForm.customerId"
                  type="text"
                  placeholder="customer_123"
                  class="tech-input"
                  disabled
                />
              </div>
              <div class="form-group">
                <label for="serial">序列号</label>
                <input
                  id="serial"
                  v-model="environmentForm.serial"
                  type="text"
                  placeholder="SN123456789"
                  class="tech-input"
                />
              </div>
            </div>
            <div class="form-group">
              <label for="remark">备注</label>
              <textarea
                id="remark"
                v-model="environmentForm.remark"
                placeholder="请输入备注信息"
                rows="3"
                class="tech-input"
              ></textarea>
            </div>
          </div>

          <div class="modal-actions">
            <button type="button" class="secondary-btn" @click="closeModal">取消</button>
            <button type="submit" class="primary-btn" :disabled="isSaving">
              {{ isSaving ? '处理中...' : editingEnvironment ? '更新' : '创建' }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { useUserStore } from '@/stores/user'
import { useBrowserStore } from '@/stores/browser'
import { useRouter } from 'vue-router'
import { SdkService } from '@/services'
import type { Browser, BrowserDto } from '@/services'

type EnvironmentForm = Partial<Browser>

const userStore = useUserStore()
const browserStore = useBrowserStore()
const router = useRouter()

const activeTab = ref('environments')
const showCreateModal = ref(false)
const editingEnvironment = ref<BrowserDto | null>(null)
const isSaving = ref(false)
// const isActionLoading = ref(false);  // 添加动作加载状态

// 窗口高度相关
const windowHeight = ref(window.innerHeight)
const headerHeight = ref(80) // 头部高度估算

// 计算主内容区域可用高度
// const mainContentHeight = computed(() => {
//   return `${windowHeight.value - headerHeight.value}px`;
// });

// 监听窗口大小变化
const handleResize = (): void => {
  windowHeight.value = window.innerHeight
}

const tabs = [
  { key: 'environments', label: '环境管理' },
  { key: 'settings', label: '系统设置' }
]

const environmentForm = ref<EnvironmentForm>({
  envName: '',
  envId: undefined,
  // 浏览器配置
  ua: '',
  system: '',
  kernel: 'Chrome',
  kernelVersion: '134',
  // 硬件配置
  cpu: 4,
  mem: 8,
  deviceName: '',
  mac: '',
  // 网络配置
  publicIp: '',
  proxy: '',
  zone: '',
  ipChannel: '',
  // 功能开关
  hardware: 0,
  webGl: 0,
  canvas: 4,
  audioContext: 0,
  mediaDevice: 0,
  bluetooth: 0,
  // 其他配置
  customerId: '',
  serial: '',
  remark: ''
})

// 根据status值获取环境卡片的CSS类
// const getEnvironmentCardClass = (status: number): string => {
//   const baseClass = 'environment-card';
//   switch (status) {
//     case 1:  // 停止状态
//       return `${baseClass} environment-stopped`;
//     case 3:  // 启动状态
//       return `${baseClass} environment-active`;
//     default:
//       return `${baseClass} environment-unknown`;
//   }
// };

// 根据status值获取环境行的CSS类
const getEnvironmentRowClass = (status: number): string => {
  switch (status) {
    case 1:
      return 'environment-row-stopped'
    case 3:
      return 'environment-row-active'
    default:
      return 'environment-row-unknown'
  }
}

// const formatDate = (dateString: string) => {
//   return new Date(dateString).toLocaleString('zh-CN');
// };

// 安全获取浏览器属性的辅助函数
// const safeGetBrowserProp = <T>(browser: Browser | undefined, prop: keyof Browser): T | undefined => {
//   if (!browser) return undefined;
//   return browser[prop] as T;
// };

const changePage = async (page: number): Promise<void> => {
  try {
    await browserStore.loadBrowsers(page)
  } catch (error) {
    console.error('Failed to change page:', error)
  }
}

const getPageNumbers = (): number[] => {
  const totalPages = Math.ceil(browserStore.total / browserStore.pageSize)
  const currentPage = browserStore.currentPage
  const pages: number[] = []

  // 显示最多5个页码
  let start = Math.max(1, currentPage - 2)
  const end = Math.min(totalPages, start + 4)

  // 调整起始页码
  if (end - start < 4) {
    start = Math.max(1, end - 4)
  }

  for (let i = start; i <= end; i++) {
    pages.push(i)
  }

  return pages
}

const switchTab = (tabKey: string): void => {
  activeTab.value = tabKey
}

const handleLogout = (): void => {
  userStore.logout()
}

onMounted(async () => {
  router.push('/login')
})

// const showModal = () => {
//   showCreateModal.value = true;
//   editingEnvironment.value = null;
//   resetForm();
// };

const closeModal = (): void => {
  showCreateModal.value = false
  editingEnvironment.value = null
  resetForm()
  isSaving.value = false
}

const resetForm = (): void => {
  environmentForm.value = {
    envName: '',
    envId: undefined,
    ua: '',
    system: '',
    kernel: 'Chrome',
    kernelVersion: '134',
    cpu: 4,
    mem: 8,
    deviceName: '',
    mac: '',
    publicIp: '',
    proxy: '',
    zone: '',
    ipChannel: '',
    hardware: 0,
    webGl: 0,
    canvas: 4,
    audioContext: 0,
    mediaDevice: 0,
    bluetooth: 0,
    customerId: '',
    serial: '',
    remark: ''
  }
}

const controlEnvironment = async (item: BrowserDto, status: number): Promise<void> => {
  if (!item.id) {
    console.warn('Invalid browser id for control operation')
    return
  }

  try {
    if (status === 3) {
      const code = await SdkService.open({
        envs: [
          {
            envId: item.envId!,
            args: []
          }
        ]
      })
      if (code === 1) {
        browserStore.startingDict.set(item.envId, item)
      } else {
        // 启动失败
        alert(`启动：${item.envName} 环境失败`)
      }
    }

    if (status === 1) {
      const code = await SdkService.close({
        envs: [item.envId!]
      })
      if (code === 1) {
        browserStore.closingDict.set(item.envId, item)
      } else {
        // 关闭失败
        alert(`停止：${item.envName}环境失败`)
      }
    }
  } catch (error) {
    console.error('Failed to control environment:', error)
    alert('操作失败，请重试')
  }
}

const editEnvironment = (browser: BrowserDto | undefined): void => {
  if (!browser) {
    console.warn('Cannot edit undefined browser')
    return
  }

  console.log('editEnvironment', browser)
  editingEnvironment.value = browser
  environmentForm.value = {
    envName: browser.envName,
    envId: browser.envId,
    ua: browser.data!.finger.ua || '',
    system: browser.data!.finger.system || '',
    kernel: browser.data!.finger.kernel || '',
    kernelVersion: browser.data!.finger.kernelVersion || '',
    cpu: browser.data!.finger.cpu || 4,
    mem: browser.data!.finger.mem || 8,
    deviceName: browser.data!.finger.deviceName || '',
    mac: browser.data!.finger.mac || '',
    publicIp: browser.data!.publicIp || '',
    proxy: browser.data!.proxy || '',
    zone: browser.data!.finger.zone || '',
    ipChannel: browser.data!.ipChannel || '',
    hardware: browser.data!.finger.hardware || 0,
    webGl: browser.data!.finger.webGl || 0,
    canvas: browser.data!.finger.canvas || 4,
    audioContext: browser.data!.finger.audioContext || 0,
    mediaDevice: browser.data!.finger.mediaDevice || 0,
    bluetooth: browser.data!.finger.bluetooth || 0,
    customerId: browser.data!.customerId || '',
    serial: browser.data!.serial || '',
    remark: browser.data!.remark || ''
  }
  showCreateModal.value = true
}

const saveEnvironment = async (): Promise<void> => {
  if (!environmentForm.value.envName!.trim()) return

  isSaving.value = true

  try {
    const browserData = {
      ...editingEnvironment.value?.data,
      envId: environmentForm.value.envId || '',
      envName: environmentForm.value.envName!,
      // bridgeProxy: '',
      // cookie: '',
      customerId: environmentForm.value.customerId || '',
      ipChannel: environmentForm.value.ipChannel || '',
      publicIp: environmentForm.value.publicIp || '',
      proxy: environmentForm.value.proxy || '',
      // region: '',
      remark: environmentForm.value.remark || '',
      serial: environmentForm.value.serial || '',
      finger: {
        audioContext: (environmentForm.value.audioContext || 0) * 1,
        // battery: 0,
        bluetooth: (environmentForm.value.bluetooth || 0) * 1,
        canvas: (environmentForm.value.canvas || 4) * 1,
        // clientRects: 0,
        cpu: (environmentForm.value.cpu || 4) * 1,
        deviceName: environmentForm.value.deviceName || '',
        // doNotTrack: 0,
        // dpi: '1920x1080',
        // enableGc: 0,
        // enablenotice: 1,
        // enableopen: 1,
        // enableOpenNumber: 0,
        // enablepic: 1,
        // enableScanPort: 0,
        // enablesound: 1,
        // enablevideo: 1,
        // font: {
        //   enable: 0,
        //   list: ['Arial', 'Helvetica', 'Times New Roman'],
        // },
        // fontFinger: 0,
        // gcTime: 0,
        // geographic: {
        //   accuracy: 'high',
        //   enable: 1,
        //   latitude: '39.9042',
        //   longitude: '116.4074',
        //   useip: 1,
        // },
        hardware: (environmentForm.value.hardware || 0) * 1,
        kernel: environmentForm.value.kernel || '',
        kernelVersion: environmentForm.value.kernelVersion || '',
        language: ['zh-CN', 'en-US'],
        mac: environmentForm.value.mac || '',
        mediaDevice: (environmentForm.value.mediaDevice || 0) * 1,
        mem: (environmentForm.value.mem || 8) * 1,
        // picSize: '1920x1080',
        // scanPort: '',
        // speechVoices: 1,
        system: environmentForm.value.system || '',
        ua: environmentForm.value.ua || '',
        // uaVersion: '134',
        // uilanguage: [],
        webGl: (environmentForm.value.webGl || 0) * 1,
        // webGlInfo: 0,
        // webGlRenderer: '',
        // webGlVendor: '',
        // webRTC: 5,
        // webRTCIP: '127.0.0.1',
        // widowssize: '1920x1080',
        zone: environmentForm.value.zone || ''
      }
    }

    if (editingEnvironment.value) {
      browserData.finger = {
        ...editingEnvironment.value.data!.finger,
        audioContext: (environmentForm.value.audioContext || 0) * 1,
        bluetooth: (environmentForm.value.bluetooth || 0) * 1,
        canvas: (environmentForm.value.canvas || 4) * 1,
        cpu: (environmentForm.value.cpu || 4) * 1,
        deviceName: environmentForm.value.deviceName || '',
        hardware: (environmentForm.value.hardware || 0) * 1,
        kernel: environmentForm.value.kernel || '',
        kernelVersion: environmentForm.value.kernelVersion || '',
        language: editingEnvironment.value.data!.finger.language || [],
        mac: environmentForm.value.mac || '',
        mediaDevice: (environmentForm.value.mediaDevice || 0) * 1,
        mem: (environmentForm.value.mem || 8) * 1,
        system: environmentForm.value.system || '',
        ua: environmentForm.value.ua || '',
        webGl: (environmentForm.value.webGl || 0) * 1,
        zone: environmentForm.value.zone || ''
      }
      // 更新现有环境
      await browserStore.updateBrowser({
        id: editingEnvironment.value.id!,
        envName: environmentForm.value.envName!,
        envId: environmentForm.value.envId!,
        userId: userStore.user?.id || 0,
        // data: JSON.stringify(browserData),
        data: browserData
      })
    } else {
      // 创建新环境
      await browserStore.createBrowser({
        // envName: environmentForm.value.name!,
        envName: environmentForm.value.envName!,
        envId: environmentForm.value.envId,
        userId: userStore.user?.id || 0,
        // data: JSON.stringify(browserData),
        data: browserData
      })
    }

    closeModal()
  } catch (error) {
    console.error('Failed to save environment:', error)
    alert('操作失败，请重试')
  } finally {
    isSaving.value = false
  }
}

const deleteEnvironment = async (id: number | undefined): Promise<void> => {
  if (!id) {
    console.warn('Cannot delete environment with invalid id')
    return
  }

  if (confirm('确定要删除这个环境吗？此操作不可撤销。')) {
    try {
      await browserStore.deleteBrowser(id)
    } catch (error) {
      console.error('Failed to delete environment:', error)
      alert('删除失败，请重试')
    }
  }
}

const formatDateTime = (dateString: string | undefined): string => {
  if (!dateString) return ''
  return new Date(dateString).toLocaleString('zh-CN', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit'
  })
}

const getStatusBadgeClass = (status: number | undefined): string => {
  if (status === undefined) return 'status-badge unknown'

  switch (status) {
    case 1:
      return 'status-badge stopped'
    case 3:
      return 'status-badge running'
    default:
      return 'status-badge unknown'
  }
}

const getStatusText = (status: number | undefined): string => {
  if (status === undefined) return '未知'

  switch (status) {
    case 1:
      return '停止'
    case 3:
      return '运行'
    default:
      return '未知'
  }
}

onMounted(async () => {
  window.addEventListener('resize', handleResize)
  handleResize() // 初始化

  if (!userStore.isAuthenticated) {
    router.push('/login')
  } else {
    // 恢复localStorage中的用户信息
    userStore.restoreUserInfo()
    // 确保用户信息是最新的
    userStore.loadUserInfo()

    console.log('userStore.user', userStore.user)

    // 加载浏览器环境列表
    try {
      await browserStore.loadBrowsers()
    } catch (error) {
      console.error('Failed to load environments:', error)
    }
  }
})

onUnmounted(() => {
  window.removeEventListener('resize', handleResize)
})
</script>

<style lang="less" scoped>
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
  0% {
    transform: translate(0, 0);
  }
  100% {
    transform: translate(40px, 40px);
  }
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
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
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

  .label,
  .value {
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

  .primary-btn,
  .secondary-btn {
    width: 100%;
    padding: 12px;
  }
}

/* 超小屏幕优化 */

/* 表格样式 */
.environments-table-container {
  background: rgba(15, 15, 25, 0.7);
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 5px 20px rgba(0, 0, 0, 0.3);
  border: 1px solid rgba(0, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  margin-bottom: 25px;
}

.environments-table {
  width: 100%;
  border-collapse: collapse;
  background: transparent;
}

.environments-table th {
  background: rgba(5, 5, 15, 0.8);
  color: #00ffff;
  font-weight: 600;
  text-align: left;
  padding: 15px 20px;
  border-bottom: 1px solid rgba(0, 255, 255, 0.2);
  font-size: 14px;
  letter-spacing: 1px;
}

.environments-table td {
  padding: 15px 20px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.05);
  color: rgba(255, 255, 255, 0.9);
  font-size: 13px;
}

.environments-table tr:last-child td {
  border-bottom: none;
}

.environments-table tr:hover {
  background: rgba(0, 255, 255, 0.05);
}

/* 环境行状态样式 */
.environment-row-active {
  background: rgba(0, 255, 0, 0.05);
}

.environment-row-active:hover {
  background: rgba(0, 255, 0, 0.1);
}

.environment-row-stopped {
  background: rgba(255, 0, 0, 0.05);
}

.environment-row-stopped:hover {
  background: rgba(255, 0, 0, 0.1);
}

.environment-row-unknown {
  background: rgba(255, 255, 0, 0.05);
}

.environment-row-unknown:hover {
  background: rgba(255, 255, 0, 0.1);
}

/* 表格操作按钮 */
.table-actions {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
}

.action-btn.small {
  padding: 4px 10px;
  font-size: 10px;
  min-width: 50px;
  height: 24px;
}

/* 分页样式 */
.pagination-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px;
  background: rgba(15, 15, 25, 0.5);
  border-radius: 12px;
  border: 1px solid rgba(0, 255, 255, 0.1);
  margin-top: 20px;
}

.pagination-info {
  color: rgba(255, 255, 255, 0.7);
  font-size: 13px;
}

.pagination-controls {
  display: flex;
  align-items: center;
  gap: 15px;
}

.pagination-btn {
  background: rgba(0, 255, 255, 0.1);
  color: #00ffff;
  border: 1px solid rgba(0, 255, 255, 0.3);
  padding: 8px 15px;
  border-radius: 6px;
  cursor: pointer;
  font-size: 12px;
  font-family: 'Courier New', monospace;
  transition: all 0.2s ease;
}

.pagination-btn:hover:not(:disabled) {
  background: rgba(0, 255, 255, 0.2);
  box-shadow: 0 0 10px rgba(0, 255, 255, 0.4);
}

.pagination-btn:disabled {
  opacity: 0.4;
  cursor: not-allowed;
}

.pagination-pages {
  display: flex;
  gap: 5px;
}

.pagination-page {
  background: rgba(255, 255, 255, 0.05);
  color: rgba(255, 255, 255, 0.7);
  border: 1px solid rgba(255, 255, 255, 0.1);
  padding: 8px 12px;
  border-radius: 4px;
  cursor: pointer;
  font-size: 12px;
  font-family: 'Courier New', monospace;
  transition: all 0.2s ease;
}

.pagination-page:hover {
  background: rgba(0, 255, 255, 0.1);
  color: #00ffff;
  border-color: rgba(0, 255, 255, 0.3);
}

.pagination-page.active {
  background: rgba(0, 255, 255, 0.2);
  color: #00ffff;
  border-color: #00ffff;
  box-shadow: 0 0 10px rgba(0, 255, 255, 0.3);
}

/* 模态框样式扩展 */
.modal-content.large {
  width: 800px;
  max-width: 95vw;
  max-height: 90vh;
  overflow-y: auto;
}

.form-section {
  margin-bottom: 25px;
  padding: 20px;
  background: rgba(5, 5, 15, 0.3);
  border-radius: 8px;
  border: 1px solid rgba(0, 255, 255, 0.1);
  &.form-section2 {
    .form-group {
      label {
        padding-left: 16px;
      }
    }
  }
}

.form-section h4 {
  color: #00ffff;
  margin: 0 0 20px 0;
  font-size: 16px;
  letter-spacing: 1px;
  padding-bottom: 10px;
  border-bottom: 1px solid rgba(0, 255, 255, 0.2);
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 20px;
  margin-bottom: 15px;
}

.form-row:last-child {
  margin-bottom: 0;
}

.checkbox-group {
  display: flex;
  align-items: flex-start;
}

.checkbox-group label {
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
  color: rgba(255, 255, 255, 0.9);
  font-size: 13px;
}

.checkbox-group input[type='checkbox'] {
  width: 16px;
  height: 16px;
  accent-color: #00ffff;
}

.help-text {
  color: rgba(255, 255, 255, 0.5);
  font-size: 11px;
  margin-top: 5px;
  display: block;
}

.tech-input:disabled {
  opacity: 0.6;
  cursor: not-allowed;
  background: rgba(255, 255, 255, 0.05);
}

/* 响应式设计 */
@media (max-width: 768px) {
  .environments-table-container {
    overflow-x: auto;
  }

  .environments-table {
    min-width: 800px;
  }

  .form-row {
    grid-template-columns: 1fr;
  }

  .modal-content.large {
    width: 95vw;
    padding: 20px;
  }

  .pagination-container {
    flex-direction: column;
    gap: 15px;
    text-align: center;
  }

  .table-actions {
    justify-content: flex-start;
  }

  .action-btn.small {
    min-width: 45px;
    padding: 3px 8px;
  }
}

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

  .label,
  .value {
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

  .primary-btn,
  .secondary-btn {
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
