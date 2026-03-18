<template>
  <div class="modal-overlay" @click="emit('close')">
    <div class="modal-content large" @click.stop>
      <div class="modal-header">
        <h3>高级启动 - {{ browser?.envName }}</h3>
        <button class="close-btn" @click="emit('close')">×</button>
      </div>

      <div class="modal-form">
        <!-- 启动参数 -->
        <div class="form-section">
          <h4>启动参数 (Args)</h4>
          <div class="form-group">
            <label>启动命令行参数（每行一个）</label>
            <textarea
              v-model="form.argsText"
              class="tech-input"
              rows="4"
              placeholder="例如：&#10;--disable-blink-features=AutomationControlled&#10;--disable-infobars&#10;--no-first-run"
            ></textarea>
            <small class="help-text">Chrome启动参数，每行一个</small>
          </div>
        </div>

        <!-- URLs -->
        <div class="form-section">
          <h4>启动URL (URLs)</h4>
          <div class="form-group">
            <label>要打开的URL列表（每行一个）</label>
            <textarea
              v-model="form.urlsText"
              class="tech-input"
              rows="3"
              placeholder="例如：&#10;https://www.example.com&#10;https://www.google.com"
            ></textarea>
            <small class="help-text">启动时自动打开的网页URL，每行一个</small>
          </div>
        </div>

        <!-- Cookies -->
        <div class="form-section">
          <h4>Cookies 配置</h4>
          <div class="form-group">
            <label>Cookie JSON 配置</label>
            <textarea
              v-model="form.cookiesJson"
              class="tech-input"
              rows="8"
              placeholder='[&#10;  {&#10;    "name": "session_id",&#10;    "value": "abc123",&#10;    "domain": ".example.com",&#10;    "path": "/",&#10;    "httpOnly": true,&#10;    "secure": true&#10;  }&#10;]'
            ></textarea>
            <small class="help-text">JSON格式的Cookie配置</small>
          </div>
          <div class="form-group">
            <button type="button" class="secondary-btn" style="width:auto;padding:8px 16px" @click="validateCookies">
              验证JSON格式
            </button>
            <span v-if="validationMsg" :class="['validation-msg', validationOk ? 'valid' : 'invalid']">
              {{ validationMsg }}
            </span>
          </div>
        </div>

        <div class="modal-actions">
          <button type="button" class="secondary-btn" @click="emit('close')">取消</button>
          <button type="button" class="primary-btn" :disabled="starting" @click="handleStart">
            {{ starting ? '启动中...' : '启动' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import type { BrowserDto } from '@/services'
import type { IOpenCookie } from '@type/sdk'

interface AdvancedStartForm {
  argsText: string
  urlsText: string
  cookiesJson: string
}

const props = defineProps<{
  browser: BrowserDto | null
  starting: boolean
}>()

const emit = defineEmits<{
  close: []
  start: [params: { args: string[]; urls: string[]; cookies?: IOpenCookie[] }]
}>()

const form = reactive<AdvancedStartForm>({
  argsText: '',
  urlsText: '',
  cookiesJson: ''
})

const validationMsg = ref('')
const validationOk = ref(false)

const validateCookies = (): boolean => {
  try {
    const jsonStr = form.cookiesJson.trim()
    if (!jsonStr) {
      validationMsg.value = 'JSON为空，将不设置Cookies'
      validationOk.value = true
      return true
    }
    const cookies: IOpenCookie[] = JSON.parse(jsonStr)
    if (!Array.isArray(cookies)) throw new Error('必须是数组')
    for (let i = 0; i < cookies.length; i++) {
      if (!cookies[i].name || !cookies[i].value) throw new Error(`第${i + 1}个cookie缺少name或value`)
      if (!cookies[i].domain) throw new Error(`第${i + 1}个cookie缺少domain`)
    }
    validationMsg.value = `✓ JSON格式正确，包含${cookies.length}个Cookie`
    validationOk.value = true
    return true
  } catch (error) {
    validationMsg.value = `✗ ${error instanceof Error ? error.message : 'JSON格式错误'}`
    validationOk.value = false
    return false
  }
}

const handleStart = () => {
  const args = form.argsText.split('\n').map(l => l.trim()).filter(Boolean)
  const urls = form.urlsText.split('\n').map(l => l.trim()).filter(Boolean)

  let cookies: IOpenCookie[] | undefined
  const jsonStr = form.cookiesJson.trim()
  if (jsonStr) {
    if (!validateCookies()) {
      alert('Cookies JSON格式错误，请先验证格式')
      return
    }
    try {
      cookies = JSON.parse(jsonStr)
    } catch {
      return
    }
  }

  emit('start', { args, urls, cookies: cookies?.length ? cookies : undefined })
}
</script>
