import { defineStore } from 'pinia';
import { ref } from 'vue';
import { ApiService } from '@/services/api';

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

interface BrowserDto {
  id?: number;
  name: string;
  envId?: number;  // 变为可选
  userId: number;
  data: string;
}

export const useBrowserStore = defineStore('browser', () => {
  const browsers = ref<Browser[]>([]);
  const loading = ref(false);
  const currentPage = ref(1);
  const pageSize = ref(10);
  const total = ref(0);

  const loadBrowsers = async (page: number = 1, name?: string) => {
    loading.value = true;
    try {
      const response = await ApiService.getBrowserList({
        page,
        size: pageSize.value,
        name
      });
      
      browsers.value = response.list;
      total.value = response.total;
      currentPage.value = page;
    } catch (error) {
      console.error('Failed to load browsers:', error);
      throw error;
    } finally {
      loading.value = false;
    }
  };

  const getBrowser = async (id: number): Promise<Browser> => {
    try {
      return await ApiService.getBrowser(id);
    } catch (error) {
      console.error('Failed to get browser:', error);
      throw error;
    }
  };

  const createBrowser = async (browser: BrowserDto): Promise<Browser> => {
    try {
      const newBrowser = await ApiService.createBrowser(browser);
      // 添加到列表开头
      browsers.value.unshift(newBrowser);
      total.value += 1;
      return newBrowser;
    } catch (error) {
      console.error('Failed to create browser:', error);
      throw error;
    }
  };

  const updateBrowser = async (browser: BrowserDto): Promise<Browser> => {
    try {
      const updatedBrowser = await ApiService.updateBrowser(browser);
      // 更新列表中的对应项
      const index = browsers.value.findIndex(b => b.id === browser.id);
      if (index !== -1) {
        browsers.value[index] = updatedBrowser;
      }
      return updatedBrowser;
    } catch (error) {
      console.error('Failed to update browser:', error);
      throw error;
    }
  };

  const deleteBrowser = async (id: number): Promise<void> => {
    try {
      await ApiService.deleteBrowser([id]);
      // 从列表中移除
      browsers.value = browsers.value.filter(b => b.id !== id);
      total.value -= 1;
    } catch (error) {
      console.error('Failed to delete browser:', error);
      throw error;
    }
  };

  const deleteMultipleBrowsers = async (ids: number[]): Promise<void> => {
    try {
      await ApiService.deleteBrowser(ids);
      // 从列表中批量移除
      browsers.value = browsers.value.filter(b => !ids.includes(b.id));
      total.value -= ids.length;
    } catch (error) {
      console.error('Failed to delete browsers:', error);
      throw error;
    }
  };

  return {
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
  };
});