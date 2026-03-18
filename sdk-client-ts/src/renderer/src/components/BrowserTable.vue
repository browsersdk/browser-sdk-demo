<template>
  <div>
    <!-- 加载状态 -->
    <div v-if="loading" class="loading-state">
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
            <template v-for="(browser, index) in browsers" :key="browser?.id || index">
              <tr v-if="index < pageSize" :class="getRowClass(browser?.status!)">
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
                        startedDict.has(browser?.envId ?? '') ||
                        startingDict.has(browser?.envId ?? '') ||
                        closingDict.has(browser?.envId ?? '')
                      "
                      @click="emit('start', browser)"
                    >启动</button>
                    <button
                      class="action-btn advanced small"
                      :disabled="
                        startedDict.has(browser?.envId ?? '') ||
                        startingDict.has(browser?.envId ?? '') ||
                        closingDict.has(browser?.envId ?? '')
                      "
                      @click="emit('advancedStart', browser)"
                    >高级启动</button>
                    <button
                      class="action-btn stop small"
                      :disabled="
                        startingDict.has(browser?.envId ?? '') ||
                        closingDict.has(browser?.envId ?? '') ||
                        !startedDict.has(browser?.envId ?? '')
                      "
                      @click="emit('stop', browser)"
                    >停止</button>
                    <button
                      class="action-btn edit small"
                      :disabled="!browser"
                      @click="emit('edit', browser)"
                    >编辑</button>
                    <button
                      class="action-btn delete small"
                      :disabled="!browser"
                      @click="emit('delete', browser?.id)"
                    >删除</button>
                  </div>
                </td>
              </tr>
            </template>
          </tbody>
        </table>

        <!-- 空状态 -->
        <div v-if="browsers.length === 0" class="empty-state">
          <div class="empty-icon">⚙️</div>
          <p>暂无环境配置</p>
          <button class="primary-btn" @click="emit('create')">创建首个环境</button>
        </div>
      </div>

      <!-- 分页 -->
      <div v-if="browsers.length > 0" class="pagination-container">
        <div class="pagination-info">
          显示第 {{ (currentPage - 1) * pageSize + 1 }} -
          {{ Math.min(currentPage * pageSize, total) }} 条，共 {{ total }} 条记录
        </div>
        <div class="pagination-controls">
          <button
            class="pagination-btn"
            :disabled="currentPage <= 1"
            @click="emit('pageChange', currentPage - 1)"
          >上一页</button>

          <div class="pagination-pages">
            <button
              v-for="page in pageNumbers"
              :key="page"
              :class="['pagination-page', { active: page === currentPage }]"
              @click="emit('pageChange', page)"
            >{{ page }}</button>
          </div>

          <button
            class="pagination-btn"
            :disabled="currentPage >= Math.ceil(total / pageSize)"
            @click="emit('pageChange', currentPage + 1)"
          >下一页</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import type { BrowserDto } from '@/services'

const props = defineProps<{
  browsers: BrowserDto[]
  loading: boolean
  currentPage: number
  pageSize: number
  total: number
  startedDict: Set<string>
  startingDict: Map<string, BrowserDto>
  closingDict: Map<string, BrowserDto>
}>()

const emit = defineEmits<{
  start: [browser: BrowserDto]
  stop: [browser: BrowserDto]
  advancedStart: [browser: BrowserDto]
  edit: [browser: BrowserDto]
  delete: [id: number | undefined]
  create: []
  pageChange: [page: number]
}>()

const pageNumbers = computed(() => {
  const totalPages = Math.ceil(props.total / props.pageSize)
  const current = props.currentPage
  let start = Math.max(1, current - 2)
  const end = Math.min(totalPages, start + 4)
  if (end - start < 4) start = Math.max(1, end - 4)
  const pages: number[] = []
  for (let i = start; i <= end; i++) pages.push(i)
  return pages
})

const getRowClass = (status: number): string => {
  switch (status) {
    case 3: return 'environment-row-active'
    case 1: return 'environment-row-stopped'
    default: return 'environment-row-unknown'
  }
}

const getStatusBadgeClass = (status: number | undefined): string => {
  switch (status) {
    case 3: return 'status-badge running'
    case 1: return 'status-badge stopped'
    default: return 'status-badge unknown'
  }
}

const getStatusText = (status: number | undefined): string => {
  switch (status) {
    case 3: return '运行'
    case 1: return '停止'
    default: return '未知'
  }
}

const formatDateTime = (dateString: string | undefined): string => {
  if (!dateString) return ''
  return new Date(dateString).toLocaleString('zh-CN', {
    year: 'numeric', month: '2-digit', day: '2-digit',
    hour: '2-digit', minute: '2-digit', second: '2-digit'
  })
}
</script>
