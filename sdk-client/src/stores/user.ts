import { defineStore } from 'pinia';
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { ApiService } from '@/services/api';
import { TokenManager } from '@/utils/tokenManager';

interface User {
  id: number;
  username: string;
  phone: string;
  email: string;
  nickname: string;
  name: string;
  status: number;
  createdAt: string;
  updatedAt: string;
}

export const useUserStore = defineStore('user', () => {
  const router = useRouter();
  const user = ref<User | null>(null);
  const isAuthenticated = ref(false);
  const isLoading = ref(false);

  const initializeAuth = () => {
    isAuthenticated.value = TokenManager.isAuthenticated();
    if (isAuthenticated.value) {
      loadUserInfo();
    }
  };

  const login = async (username: string, password: string) => {
    isLoading.value = true;
    try {
      const tokens = await ApiService.login(username, password);
      TokenManager.setTokens(tokens);
      isAuthenticated.value = true;
      
      // 直接跳转到主控制台
      router.push('/dashboard');
      
      // 在控制台中异步加载用户信息
      loadUserInfo();
      
      return true;
    } catch (error) {
      console.error('Login failed:', error);
      throw error;
    } finally {
      isLoading.value = false;
    }
  };

  const loadUserInfo = async () => {
    if (!TokenManager.isAuthenticated()) {
      return;
    }

    try {
      const userInfo = await ApiService.getUserInfo();
      user.value = userInfo;
      // 同时保存到localStorage（TokenManager已经处理了tokens的存储）
      localStorage.setItem('user_info', JSON.stringify(userInfo));
    } catch (error) {
      console.error('Failed to load user info:', error);
      // 如果获取用户信息失败，清除认证状态
      logout();
    }
  };

  const logout = () => {
    ApiService.logout();
    user.value = null;
    isAuthenticated.value = false;
    // 清除localStorage中的用户信息
    localStorage.removeItem('user_info');
    router.push('/login');
  };

  // 从localStorage恢复用户信息
  const restoreUserInfo = () => {
    const storedUserInfo = localStorage.getItem('user_info');
    if (storedUserInfo && isAuthenticated.value) {
      try {
        user.value = JSON.parse(storedUserInfo);
      } catch (error) {
        console.error('Failed to parse stored user info:', error);
        localStorage.removeItem('user_info');
      }
    }
  };

  return {
    user,
    isAuthenticated,
    isLoading,
    initializeAuth,
    login,
    loadUserInfo,
    logout,
    restoreUserInfo
  };
});