import { defineStore } from 'pinia'
import { ref } from 'vue'
import { ApiService } from '@/services'
import type { BrowserDto } from '@/services'

export const useBrowserStore = defineStore('browser', () => {
  /** 已启动 */
  const startedDict = ref(new Set<string>())
  /** 启动中 */
  const startingDict = ref(new Map<string, BrowserDto>())
  /** 关闭中 */
  const closingDict = ref(new Map<string, BrowserDto>())
  const browsers = ref<BrowserDto[]>([])
  const loading = ref(false)
  const currentPage = ref(1)
  const pageSize = ref(10)
  const total = ref(0)
  const currentEnvName = ref<string>()

  const loadBrowsers = async (page: number = 1, name?: string): Promise<void> => {
    loading.value = true
    try {
      const response = await ApiService.getBrowserList({
        page,
        size: pageSize.value,
        envName: name
      })

      browsers.value = response.list
      total.value = response.total
      currentPage.value = page
      currentEnvName.value = name
    } catch (error) {
      console.error('Failed to load browsers:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  const getBrowser = async (id: number): Promise<BrowserDto> => {
    try {
      return await ApiService.getBrowser(id)
    } catch (error) {
      console.error('Failed to get browser:', error)
      throw error
    }
  }

  const createBrowser = async (browser: BrowserDto): Promise<BrowserDto> => {
    try {
      const newBrowser = await ApiService.createBrowser(browser)
      // 添加到列表开头
      browsers.value.unshift(newBrowser)
      total.value += 1
      return newBrowser
    } catch (error) {
      console.error('Failed to create browser:', error)
      throw error
    }
  }

  const updateBrowser = async (browser: BrowserDto): Promise<BrowserDto> => {
    try {
      const updatedBrowser = await ApiService.updateBrowser(browser)
      // 更新列表中的对应项
      const index = browsers.value.findIndex((b) => b.id === browser.id)
      if (index !== -1) {
        browsers.value[index] = updatedBrowser
      }
      return updatedBrowser
    } catch (error) {
      console.error('Failed to update browser:', error)
      throw error
    }
  }

  const updateBrowserStatus = async (browser: BrowserDto): Promise<BrowserDto> => {
    try {
      const updatedBrowser = await ApiService.updateBrowserStatus(browser)
      // 更新列表中的对应项
      const index = browsers.value.findIndex((b) => b.id === browser.id)
      if (index !== -1) {
        browsers.value[index] = updatedBrowser
      }
      return updatedBrowser
    } catch (error) {
      console.error('Failed to update browser:', error)
      throw error
    }
  }

  /**
   * 更新本地浏览器状态（直接修改本地列表，不调用 API）
   * @param envId 环境 ID
   * @param status 目标状态（1=停止，3=运行）
   * @returns 是否更新成功
   */
  const updateLocalBrowserStatus = (envId: string, status: number): boolean => {
    const browserIndex = browsers.value.findIndex(
      (b) => b.envId === envId || String(b.envId) === envId
    )

    if (browserIndex !== -1) {
      browsers.value[browserIndex].status = status
      console.log(`本地状态已更新: envId=${envId}, status=${status}`)
      return true
    } else {
      console.warn(`未找到对应的浏览器记录: envId=${envId}`)
      return false
    }
  }

  /**
   * 同步更新浏览器状态到服务端和本地列表（统一的状态同步方法）
   * @param data 浏览器数据（包含 id 和 envId）
   * @param status 目标状态（1=停止，3=运行）
   * @param syncToServer 是否同步到服务端（默认 true）
   * @returns Promise<void>
   */
  const syncBrowserStatus = async (
    data: { id?: number; envId?: string },
    status: number,
    syncToServer: boolean = true
  ): Promise<void> => {
    try {
      // 1. 同步到服务端（如果需要）
      if (syncToServer && data.id) {
        await ApiService.updateBrowserStatus({
          id: data.id,
          envId: data.envId?.toString(),
          status
        })
      }

      // 2. 更新本地列表
      if (data.envId) {
        updateLocalBrowserStatus(data.envId, status)
      }
    } catch (error) {
      console.error('同步浏览器状态失败:', error)
      // 状态同步失败不影响前端状态
      // 本地列表状态保持不变
    }
  }

  /**
   * 处理环境启动成功
   * @param envId 环境 ID
   */
  const handleOpenSuccess = (envId: string): void => {
    startingDict.value.delete(envId)
    startedDict.value.add(envId)
    console.log(`环境启动成功: envId=${envId}`)
  }

  /**
   * 处理环境关闭成功
   * @param envId 环境 ID
   */
  const handleCloseSuccess = (envId: string): void => {
    closingDict.value.delete(envId)
    startedDict.value.delete(envId)
    console.log(`环境关闭成功: envId=${envId}`)
  }

  /**
   * 处理环境启动失败
   * @param envId 环境 ID
   * @param errorMsg 错误信息
   */
  const handleOpenFailed = (envId: string, errorMsg?: string): void => {
    if (startingDict.value.has(envId)) {
      const item = startingDict.value.get(envId)
      startingDict.value.delete(envId)

      const message = errorMsg || `启动：${item?.envName || envId} 环境失败`
      alert(message)

      console.error(`环境启动失败: envId=${envId}`)
    }
  }

  /**
   * 处理环境关闭失败
   * @param envId 环境 ID
   * @param errorMsg 错误信息
   */
  const handleCloseFailed = (envId: string, errorMsg?: string): void => {
    if (closingDict.value.has(envId)) {
      const item = closingDict.value.get(envId)
      closingDict.value.delete(envId)

      const message = errorMsg || `停止：${item?.envName || envId} 环境失败`
      alert(message)

      console.error(`环境关闭失败: envId=${envId}`)
    }
  }

  const deleteBrowser = async (id: number): Promise<void> => {
    try {
      const code = await ApiService.deleteBrowser([id])

      if (code === 200) {
        loadBrowsers(currentPage.value, currentEnvName.value)
      }
    } catch (error) {
      console.error('Failed to delete browser:', error)
      throw error
    }
  }

  const deleteMultipleBrowsers = async (ids: number[]): Promise<void> => {
    try {
      await ApiService.deleteBrowser(ids)
      // 从列表中批量移除
      browsers.value = browsers.value.filter((b) => !ids.includes(b.id!))
      total.value -= ids.length
    } catch (error) {
      console.error('Failed to delete browsers:', error)
      throw error
    }
  }

  return {
    startedDict,
    startingDict,
    closingDict,
    browsers,
    loading,
    currentPage,
    pageSize,
    total,
    loadBrowsers,
    getBrowser,
    createBrowser,
    updateBrowser,
    updateBrowserStatus,
    syncBrowserStatus,
    updateLocalBrowserStatus,
    handleOpenSuccess,
    handleCloseSuccess,
    handleOpenFailed,
    handleCloseFailed,
    deleteBrowser,
    deleteMultipleBrowsers
  }
})
