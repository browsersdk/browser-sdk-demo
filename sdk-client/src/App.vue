<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { RouterView } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { WebSocketService, SdkService } from '@/services'

const userStore = useUserStore()
const skService = WebSocketService.getInstance()
const port = ref(65535)

watch(
  () => userStore.isAuthenticated,
  async (nVal) => {
    console.log('nVal发生变化', nVal)
    if (nVal) {
      // 初始化SDK
      const code = await SdkService.appInit(port.value)
      console.log('初始化', code)
      if (code === 0) {
        skService.connect(port.value)
      }
      if (code !== 0) alert('Sdk初始化失败!')
    }
  },
  {
    immediate: true,
  },
)

onMounted(() => {
  // 初始化用户认证状态
  userStore.initializeAuth()
})
</script>

<template>
  <RouterView />
</template>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  background: #f5f7fa;
}

#app {
  min-height: 100vh;
}
</style>
