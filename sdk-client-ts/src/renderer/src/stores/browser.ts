import { defineStore } from 'pinia'
import { ref } from 'vue'
import { ApiService } from '@/services'
import type { BrowserDto } from '@/services'

export const useBrowserStore = defineStore('browser', () => {
  /** 已启动 */
  const startedDict = ref(new Set())
  /** 启动中 */
  const startingDict = ref(new Map())
  /** 关闭中 */
  const closingDict = ref(new Map())
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
    deleteBrowser,
    deleteMultipleBrowsers
  }
})
