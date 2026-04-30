<template>
  <div class="main-layout" :class="{ 'mobile-view': isMobile }">
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
          <button class="logout-btn" @click="userStore.logout()">LOGOUT</button>
        </div>
      </div>
    </header>

    <div class="layout-container">
      <!-- 侧边栏 -->
      <nav class="sidebar-nav">
        <div class="nav-items">
          <button
            v-for="tab in tabs"
            :key="tab.key"
            :class="['nav-item', { active: activeTab === tab.key }]"
            @click="activeTab = tab.key"
          >{{ tab.label.toUpperCase() }}</button>
        </div>
      </nav>

      <!-- 主内容 -->
      <main class="main-content">
        <div class="content-wrapper">

          <!-- 环境管理 -->
          <div v-if="activeTab === 'environments'" class="tab-content">
            <div class="content-header">
              <div>
                <h2>ENVIRONMENT MANAGEMENT</h2>
                <p>浏览器环境配置与监控</p>
              </div>
              <button class="primary-btn" @click="openCreateModal">CREATE NEW ENVIRONMENT</button>
            </div>

            <BrowserTable
              :browsers="browserStore.browsers"
              :loading="browserStore.loading"
              :current-page="browserStore.currentPage"
              :page-size="browserStore.pageSize"
              :total="browserStore.total"
              :started-dict="browserStore.startedDict"
              :starting-dict="browserStore.startingDict"
              :closing-dict="browserStore.closingDict"
              @start="(b) => controlEnvironment(b, 3)"
              @stop="(b) => controlEnvironment(b, 1)"
              @advanced-start="openAdvancedStart"
              @edit="editEnvironment"
              @delete="deleteEnvironment"
              @create="openCreateModal"
              @page-change="browserStore.loadBrowsers"
            />
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
                    <div class="progress-bar"><div class="progress-fill" style="width: 65%"></div></div>
                  </div>
                  <div class="status-item">
                    <span class="status-label">MEMORY:</span>
                    <div class="progress-bar"><div class="progress-fill" style="width: 42%"></div></div>
                  </div>
                  <div class="status-item">
                    <span class="status-label">STORAGE:</span>
                    <div class="progress-bar"><div class="progress-fill" style="width: 28%"></div></div>
                  </div>
                </div>
              </div>
            </div>
          </div>

        </div>
      </main>
    </div>

    <!-- 创建/编辑环境模态框 -->
    <EnvironmentModal
      v-if="showCreateModal"
      :is-edit="!!editingBrowser"
      :saving="isSaving"
      :initial-form="environmentForm"
      @close="closeCreateModal"
      @submit="saveEnvironment"
    />

    <!-- 高级启动模态框 -->
    <AdvancedStartModal
      v-if="showAdvancedStartModal"
      :browser="advancedStartBrowser"
      :starting="isStarting"
      @close="showAdvancedStartModal = false"
      @start="performAdvancedStart"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { useBrowserStore } from '@/stores/browser'
import { SdkService } from '@/services'
import type { Browser, BrowserDto } from '@/services'
import type { IOpenCookie } from '@type/sdk'
import BrowserTable from '@/components/BrowserTable.vue'
import EnvironmentModal from '@/components/EnvironmentModal.vue'
import AdvancedStartModal from '@/components/AdvancedStartModal.vue'
import '@/styles/layout.css'
import '@/styles/minimal-theme.css'

type EnvironmentForm = Partial<Browser>

const DEFAULT_FORM: EnvironmentForm = {
  envName: '', envId: undefined, ua: '', system: '',
  kernel: 'Chrome', kernelVersion: '134',
  cpu: 4, mem: 8, deviceName: '', mac: '',
  publicIp: '', proxy: '', zone: '', ipChannel: '',
  hardware: 0, webGl: 0, canvas: 4,
  audioContext: 0, mediaDevice: 0, bluetooth: 0,
  customerId: '', serial: '', remark: ''
}

const userStore = useUserStore()
const browserStore = useBrowserStore()
const router = useRouter()

// 响应式状态
const isMobile = ref(false)
const screenWidth = ref(window.innerWidth)

const tabs = [
  { key: 'environments', label: '环境管理' },
  { key: 'settings', label: '系统设置' }
]
const activeTab = ref('environments')

// 模态框状态
const showCreateModal = ref(false)
const editingBrowser = ref<BrowserDto | null>(null)
const isSaving = ref(false)
const environmentForm = ref<EnvironmentForm>({ ...DEFAULT_FORM })

// 高级启动状态
const showAdvancedStartModal = ref(false)
const advancedStartBrowser = ref<BrowserDto | null>(null)
const isStarting = ref(false)

// ── 响应式处理 ──────────────────────────────────────
const handleResize = () => {
  screenWidth.value = window.innerWidth
  isMobile.value = window.innerWidth < 768
}

// ── 创建/编辑环境 ──────────────────────────────────
const openCreateModal = () => {
  editingBrowser.value = null
  environmentForm.value = { ...DEFAULT_FORM }
  showCreateModal.value = true
}

const editEnvironment = (browser: BrowserDto | undefined) => {
  if (!browser) return
  editingBrowser.value = browser
  environmentForm.value = {
    envName: browser.envName,
    envId: browser.envId,
    ua: browser.data?.finger.ua || '',
    system: browser.data?.finger.system || '',
    kernel: browser.data?.finger.kernel || '',
    kernelVersion: browser.data?.finger.kernelVersion || '',
    cpu: browser.data?.finger.cpu || 4,
    mem: browser.data?.finger.mem || 8,
    deviceName: browser.data?.finger.deviceName || '',
    mac: browser.data?.finger.mac || '',
    publicIp: browser.data?.publicIp || '',
    proxy: browser.data?.proxy || '',
    zone: browser.data?.finger.zone || '',
    ipChannel: browser.data?.ipChannel || '',
    hardware: browser.data?.finger.hardware || 0,
    webGl: browser.data?.finger.webGl || 0,
    canvas: browser.data?.finger.canvas || 4,
    audioContext: browser.data?.finger.audioContext || 0,
    mediaDevice: browser.data?.finger.mediaDevice || 0,
    bluetooth: browser.data?.finger.bluetooth || 0,
    customerId: browser.data?.customerId || '',
    serial: browser.data?.serial || '',
    remark: browser.data?.remark || ''
  }
  showCreateModal.value = true
}

const closeCreateModal = () => {
  showCreateModal.value = false
  editingBrowser.value = null
  isSaving.value = false
}

const saveEnvironment = async (form: EnvironmentForm) => {
  if (!form.envName?.trim()) return
  isSaving.value = true
  try {
    const fingerBase = {
      audioContext: (form.audioContext || 0) * 1,
      bluetooth:    (form.bluetooth || 0) * 1,
      canvas:       (form.canvas || 4) * 1,
      cpu:          (form.cpu || 4) * 1,
      deviceName:   form.deviceName || '',
      hardware:     (form.hardware || 0) * 1,
      kernel:       form.kernel || '',
      kernelVersion: form.kernelVersion || '',
      language:     ['zh-CN', 'en-US'],
      mac:          form.mac || '',
      mediaDevice:  (form.mediaDevice || 0) * 1,
      mem:          (form.mem || 8) * 1,
      system:       form.system || '',
      ua:           form.ua || '',
      webGl:        (form.webGl || 0) * 1,
      zone:         form.zone || ''
    }

    const browserData = {
      ...editingBrowser.value?.data,
      envId:      form.envId?.toString() || '',
      envName:    form.envName!,
      customerId: form.customerId || '',
      ipChannel:  form.ipChannel || '',
      publicIp:   form.publicIp || '',
      proxy:      form.proxy || '',
      remark:     form.remark || '',
      serial:     form.serial || '',
      finger: editingBrowser.value
        ? { ...editingBrowser.value.data?.finger, ...fingerBase, language: editingBrowser.value.data?.finger.language || [] }
        : fingerBase
    }

    if (editingBrowser.value) {
      await browserStore.updateBrowser({
        id: editingBrowser.value.id!,
        envName: form.envName!,
        envId: form.envId?.toString(),
        userId: userStore.user?.id || 0,
        data: browserData
      })
    } else {
      await browserStore.createBrowser({
        envName: form.envName!,
        envId: form.envId?.toString(),
        userId: userStore.user?.id || 0,
        data: browserData
      })
    }
    closeCreateModal()
  } catch (error) {
    console.error('Failed to save environment:', error)
    alert('操作失败，请重试')
  } finally {
    isSaving.value = false
  }
}

const deleteEnvironment = async (id: number | undefined) => {
  if (!id) return
  if (confirm('确定要删除这个环境吗？此操作不可撤销。')) {
    try {
      await browserStore.deleteBrowser(id)
    } catch (error) {
      console.error('Failed to delete environment:', error)
      alert('删除失败，请重试')
    }
  }
}

// ── 环境控制 ──────────────────────────────────────
const controlEnvironment = async (item: BrowserDto, status: number) => {
  if (!item.id) return

  try {
    if (status === 3) {
      // 启动环境
      const code = await SdkService.open({ envs: [{ envId: item.envId!, args: [] }] })
      if (code === 1) {
        // 添加到启动中
        browserStore.startingDict.set(item.envId!, item)

        // 预同步服务端状态为运行中
        await browserStore.syncBrowserStatus({
          id: item.id,
          envId: item.envId?.toString()
        }, 3)
      } else {
        alert(`启动：${item.envName} 环境失败`)
      }
    } else if (status === 1) {
      // 停止环境
      const code = await SdkService.close({ envs: [item.envId!] })
      if (code === 1) {
        // 添加到关闭中
        browserStore.closingDict.set(item.envId!, item)

        // 预同步服务端状态为停止
        await browserStore.syncBrowserStatus({
          id: item.id,
          envId: item.envId?.toString()
        }, 1)
      } else {
        alert(`停止：${item.envName} 环境失败`)
      }
    }
  } catch (error) {
    console.error('Failed to control environment:', error)
    alert('操作失败，请重试')
  }
}

// ── 高级启动 ──────────────────────────────────────
const openAdvancedStart = (browser: BrowserDto | undefined) => {
  if (!browser) return
  advancedStartBrowser.value = browser
  showAdvancedStartModal.value = true
}

const performAdvancedStart = async (params: { args: string[]; urls: string[]; cookies?: IOpenCookie[] }) => {
  if (!advancedStartBrowser.value) return
  isStarting.value = true
  try {
    const code = await SdkService.open({
      envs: [{
        envId: advancedStartBrowser.value.envId!,
        args: params.args.length ? params.args : undefined,
        urls: params.urls.length ? params.urls : undefined,
        cookies: params.cookies
      }]
    })
    if (code === 1) {
      browserStore.startingDict.set(advancedStartBrowser.value.envId!, advancedStartBrowser.value)
      showAdvancedStartModal.value = false
    } else {
      alert(`高级启动：${advancedStartBrowser.value.envName} 环境失败`)
    }
  } catch (error) {
    console.error('Failed to perform advanced start:', error)
    alert('高级启动失败，请重试')
  } finally {
    isStarting.value = false
  }
}

// ── 生命周期 ──────────────────────────────────────
onMounted(async () => {
  // 监听窗口大小变化
  handleResize()
  window.addEventListener('resize', handleResize)

  if (!userStore.isAuthenticated) {
    router.push('/login')
  } else {
    userStore.restoreUserInfo()
    userStore.loadUserInfo()
    try {
      await browserStore.loadBrowsers()
      // 查询 SDK 中正在运行的环境，初始化 startedDict
      const runningEnvIds = await SdkService.info()
      browserStore.startedDict.clear()
      runningEnvIds.forEach((id) => browserStore.startedDict.add(id))
    } catch (error) {
      console.error('Failed to load environments:', error)
    }
  }
})

onUnmounted(() => {
  window.removeEventListener('resize', handleResize)
})
</script>
