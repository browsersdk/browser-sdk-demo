import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from '@/stores/user'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      redirect: '/dashboard'
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('../components/LoginForm.vue'),
      meta: { requiresGuest: true }
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: () => import('../layouts/MainLayout.vue'),
      meta: { requiresAuth: true }
    }
  ]
})

// 路由守卫
router.beforeEach((to, from, next) => {
  const userStore = useUserStore()
  
  // 初始化认证状态
  if (!userStore.isAuthenticated) {
    userStore.initializeAuth()
  }

  // 检查是否需要认证
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (!userStore.isAuthenticated) {
      next({
        path: '/login',
        query: { redirect: to.fullPath }
      })
    } else {
      next()
    }
  }
  // 检查是否需要访客权限（如登录页面）
  else if (to.matched.some(record => record.meta.requiresGuest)) {
    if (userStore.isAuthenticated) {
      next({ path: '/dashboard' })
    } else {
      next()
    }
  } else {
    next()
  }
})

export default router