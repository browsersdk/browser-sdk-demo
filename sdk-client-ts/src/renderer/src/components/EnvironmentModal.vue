<template>
  <div class="modal-overlay" @click="emit('close')">
    <div class="modal-content large" @click.stop>
      <div class="modal-header">
        <h3>{{ isEdit ? '编辑环境' : '创建新环境' }}</h3>
        <button class="close-btn" @click="emit('close')">×</button>
      </div>

      <form @submit.prevent="handleSubmit" class="modal-form">
        <!-- 必填项 -->
        <div class="form-section required-section">
          <h4 class="section-title">
            <span class="icon">⚙️</span>
            <span>必填配置</span>
          </h4>
          <div class="form-row">
            <div class="form-group">
              <label>环境名称 *</label>
              <input v-model="form.envName" type="text" placeholder="请输入环境名称" required class="tech-input" />
            </div>
            <div class="form-group">
              <label>内核 *</label>
              <input v-model="form.kernel" type="text" placeholder="Chrome" class="tech-input" required />
            </div>
          </div>
          <div class="form-row">
            <div class="form-group">
              <label>内核版本 *</label>
              <input v-model="form.kernelVersion" type="text" placeholder="98.0.4758.102" class="tech-input" required />
            </div>
            <div class="form-group">
              <label>操作系统 *</label>
              <select v-model="form.system" class="tech-input" required>
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
        </div>

        <!-- 基本信息扩展 -->
        <div class="collapsible-section">
          <div class="section-header" @click="toggleSection('basic')">
            <h4 class="section-title"><span class="icon">📋</span><span>基本信息扩展</span></h4>
            <span class="toggle-icon" :class="{ expanded: expanded.basic }">{{ expanded.basic ? '▼' : '▶' }}</span>
          </div>
          <div v-show="expanded.basic" class="section-content">
            <div class="form-row">
              <div class="form-group">
                <label>环境ID</label>
                <input v-model.number="form.envId" type="number" placeholder="自动生成" class="tech-input" :disabled="isEdit" />
                <small v-if="!isEdit" class="help-text">留空则自动生成</small>
              </div>
              <div class="form-group">
                <label>客户ID</label>
                <input v-model="form.customerId" type="text" placeholder="customer_123" class="tech-input" disabled />
              </div>
            </div>
            <div class="form-group">
              <label>序列号</label>
              <input v-model="form.serial" type="text" placeholder="SN123456789" class="tech-input" />
            </div>
          </div>
          <div class="form-group" style="padding: 0 20px 16px;">
            <label>备注</label>
            <textarea v-model="form.remark" placeholder="请输入备注信息" rows="3" class="tech-input"></textarea>
          </div>
        </div>

        <!-- 浏览器配置 -->
        <div class="collapsible-section">
          <div class="section-header" @click="toggleSection('browser')">
            <h4 class="section-title"><span class="icon">🌐</span><span>浏览器配置</span></h4>
            <span class="toggle-icon" :class="{ expanded: expanded.browser }">{{ expanded.browser ? '▼' : '▶' }}</span>
          </div>
          <div v-show="expanded.browser" class="section-content">
            <div class="form-group">
              <label>User Agent</label>
              <textarea v-model="form.ua" placeholder="Mozilla/5.0..." rows="3" class="tech-input"></textarea>
            </div>
          </div>
        </div>

        <!-- 硬件配置 -->
        <div class="collapsible-section">
          <div class="section-header" @click="toggleSection('hardware')">
            <h4 class="section-title"><span class="icon">💻</span><span>硬件配置</span></h4>
            <span class="toggle-icon" :class="{ expanded: expanded.hardware }">{{ expanded.hardware ? '▼' : '▶' }}</span>
          </div>
          <div v-show="expanded.hardware" class="section-content">
            <div class="form-row">
              <div class="form-group">
                <label>CPU核心数</label>
                <input v-model.number="form.cpu" type="number" placeholder="4" class="tech-input" />
              </div>
              <div class="form-group">
                <label>内存(GB)</label>
                <input v-model.number="form.mem" type="number" placeholder="8" class="tech-input" />
              </div>
            </div>
            <div class="form-row">
              <div class="form-group">
                <label>设备名称</label>
                <input v-model="form.deviceName" type="text" placeholder="My Computer" class="tech-input" />
              </div>
              <div class="form-group">
                <label>MAC地址</label>
                <input v-model="form.mac" type="text" placeholder="00:00:00:00:00:00" class="tech-input" />
              </div>
            </div>
          </div>
        </div>

        <!-- 网络配置 -->
        <div class="collapsible-section">
          <div class="section-header" @click="toggleSection('network')">
            <h4 class="section-title"><span class="icon">🌍</span><span>网络配置</span></h4>
            <span class="toggle-icon" :class="{ expanded: expanded.network }">{{ expanded.network ? '▼' : '▶' }}</span>
          </div>
          <div v-show="expanded.network" class="section-content">
            <div class="form-row">
              <div class="form-group">
                <label>公网IP</label>
                <input v-model="form.publicIp" type="text" placeholder="192.168.1.1" class="tech-input" />
              </div>
              <div class="form-group">
                <label>代理设置</label>
                <input v-model="form.proxy" type="text" placeholder="http://proxy:8080" class="tech-input" />
              </div>
            </div>
            <div class="form-row">
              <div class="form-group">
                <label>时区</label>
                <select v-model="form.zone" class="tech-input">
                  <option value="">请选择时区</option>
                  <option value="UTC+8">UTC+8 (中国标准时间)</option>
                  <option value="UTC+0">UTC+0 (格林威治时间)</option>
                  <option value="UTC-5">UTC-5 (美国东部时间)</option>
                  <option value="UTC+1">UTC+1 (中欧时间)</option>
                </select>
              </div>
              <div class="form-group">
                <label>IP通道</label>
                <input v-model="form.ipChannel" type="text" placeholder="default" class="tech-input" />
              </div>
            </div>
          </div>
        </div>

        <!-- 功能开关 -->
        <div class="collapsible-section">
          <div class="section-header" @click="toggleSection('features')">
            <h4 class="section-title"><span class="icon">⚡</span><span>功能开关</span></h4>
            <span class="toggle-icon" :class="{ expanded: expanded.features }">{{ expanded.features ? '▼' : '▶' }}</span>
          </div>
          <div v-show="expanded.features" class="section-content form-section2">
            <div class="form-row">
              <div class="form-group checkbox-group">
                <label><input type="checkbox" v-model="form.hardware" true-value="1" false-value="0" />硬件加速</label>
              </div>
              <div class="form-group checkbox-group">
                <label><input type="checkbox" v-model="form.webGl" true-value="1" false-value="0" />WebGL支持</label>
              </div>
              <div class="form-group checkbox-group">
                <label><input type="checkbox" v-model="form.canvas" true-value="1" false-value="0" />Canvas支持</label>
              </div>
            </div>
            <div class="form-row">
              <div class="form-group checkbox-group">
                <label><input type="checkbox" v-model="form.audioContext" true-value="1" false-value="0" />音频上下文</label>
              </div>
              <div class="form-group checkbox-group">
                <label><input v-model="form.mediaDevice" type="checkbox" true-value="1" false-value="0" />媒体设备</label>
              </div>
              <div class="form-group checkbox-group">
                <label><input type="checkbox" v-model="form.bluetooth" true-value="1" false-value="0" />蓝牙支持</label>
              </div>
            </div>
          </div>
        </div>

        <div class="modal-actions">
          <button type="button" class="secondary-btn" @click="emit('close')">取消</button>
          <button type="submit" class="primary-btn" :disabled="saving">
            {{ saving ? '处理中...' : isEdit ? '更新' : '创建' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, watch } from 'vue'
import type { Browser } from '@/services'

type EnvironmentForm = Partial<Browser>

const props = defineProps<{
  isEdit: boolean
  saving: boolean
  initialForm: EnvironmentForm
}>()

const emit = defineEmits<{
  close: []
  submit: [form: EnvironmentForm]
}>()

const form = reactive<EnvironmentForm>({ ...props.initialForm })

watch(() => props.initialForm, (val) => {
  Object.assign(form, val)
}, { deep: true })

const expanded = reactive({
  basic: false,
  browser: false,
  hardware: false,
  network: false,
  features: false
})

const toggleSection = (key: keyof typeof expanded) => {
  expanded[key] = !expanded[key]
}

const handleSubmit = () => {
  if (!form.envName?.trim()) return
  emit('submit', { ...form })
}
</script>
