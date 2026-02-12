<template>
  <div class="login-container">
    <div class="login-box">
      <div class="login-header">
        <div class="header-glow">
          <h2>FINGERPRINT BROWSER DEMO</h2>
          <div class="subtitle">指纹浏览器</div>
        </div>
        <div class="divider"></div>
      </div>
      
      <form @submit.prevent="handleLogin" class="login-form">
        <div class="form-group">
          <label for="username">USERNAME</label>
          <div class="input-wrapper">
            <input
              id="username"
              v-model="loginForm.username"
              type="text"
              placeholder="ENTER YOUR USERNAME"
              required
              :disabled="isLoading"
              class="tech-input"
            />
            <div class="input-glow"></div>
          </div>
        </div>
        
        <div class="form-group">
          <label for="password">PASSWORD</label>
          <div class="input-wrapper">
            <input
              id="password"
              v-model="loginForm.password"
              type="password"
              placeholder="ENTER YOUR PASSWORD"
              required
              :disabled="isLoading"
              class="tech-input"
            />
            <div class="input-glow"></div>
          </div>
        </div>
        
        <button 
          type="submit" 
          class="login-button"
          :disabled="isLoading || !isFormValid"
        >
          <span class="button-text">{{ isLoading ? 'AUTHENTICATING...' : 'LOGIN' }}</span>
          <div class="button-glow"></div>
        </button>
      </form>
    </div>
    
    <!-- 背景装饰元素 -->
    <div class="background-elements">
      <div class="grid-lines"></div>
      <div class="floating-orbs">
        <div class="orb orb-1"></div>
        <div class="orb orb-2"></div>
        <div class="orb orb-3"></div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';
import { useUserStore } from '@/stores/user';

const userStore = useUserStore();

const loginForm = ref({
  username: '13800138000',
  password: 'Tt123456'
});

const isLoading = ref(false);

const isFormValid = computed(() => {
  return loginForm.value.username.trim() !== '' && 
         loginForm.value.password.trim() !== '';
});

const handleLogin = async () => {
  if (!isFormValid.value) return;
  
  isLoading.value = true;
  
  try {
    await userStore.login(loginForm.value.username, loginForm.value.password);
    // 登录成功后的跳转由store处理
  } catch (error) {
    console.error('Login error:', error);
    // 错误处理可以在这里添加更友好的提示
  } finally {
    isLoading.value = false;
  }
};
</script>

<style scoped>
.login-container {
  min-height: 100vh;
  background: #0a0a0a;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px;
  position: relative;
  overflow: hidden;
  font-family: 'Courier New', monospace;
}

/* 背景装饰 */
.background-elements {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  pointer-events: none;
}

.grid-lines {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-image: 
    linear-gradient(rgba(0, 255, 255, 0.05) 1px, transparent 1px),
    linear-gradient(90deg, rgba(0, 255, 255, 0.05) 1px, transparent 1px);
  background-size: 30px 30px;
  animation: grid-move 20s linear infinite;
}

@keyframes grid-move {
  0% { transform: translate(0, 0); }
  100% { transform: translate(30px, 30px); }
}

.floating-orbs {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
}

.orb {
  position: absolute;
  border-radius: 50%;
  background: radial-gradient(circle, rgba(0, 255, 255, 0.3) 0%, transparent 70%);
  filter: blur(20px);
  animation: float 15s infinite ease-in-out;
}

.orb-1 {
  width: 200px;
  height: 200px;
  top: 10%;
  left: 5%;
  animation-delay: 0s;
}

.orb-2 {
  width: 150px;
  height: 150px;
  top: 60%;
  right: 10%;
  animation-delay: -5s;
}

.orb-3 {
  width: 100px;
  height: 100px;
  bottom: 20%;
  left: 15%;
  animation-delay: -10s;
}

@keyframes float {
  0%, 100% { transform: translate(0, 0) scale(1); opacity: 0.3; }
  25% { transform: translate(20px, -20px) scale(1.1); opacity: 0.5; }
  50% { transform: translate(-10px, 20px) scale(0.9); opacity: 0.4; }
  75% { transform: translate(15px, 10px) scale(1.05); opacity: 0.6; }
}

.login-box {
  background: rgba(15, 15, 25, 0.85);
  backdrop-filter: blur(20px);
  border-radius: 15px;
  padding: 45px 35px;
  width: 100%;
  max-width: 420px;
  box-shadow: 
    0 0 30px rgba(0, 255, 255, 0.2),
    inset 0 0 20px rgba(128, 0, 255, 0.1);
  border: 1px solid rgba(0, 255, 255, 0.3);
  position: relative;
  z-index: 10;
}

.login-header {
  text-align: center;
  margin-bottom: 28px;
}

.header-glow {
  position: relative;
  padding: 15px;
}

.header-glow h2 {
  color: #00ffff;
  font-size: 22px;
  font-weight: 600;
  margin: 0 0 12px 0;
  letter-spacing: 2px;
  text-shadow: 0 0 10px rgba(0, 255, 255, 0.7);
  animation: text-glow 2s ease-in-out infinite alternate;
}

@keyframes text-glow {
  0% { text-shadow: 0 0 10px rgba(0, 255, 255, 0.7); }
  100% { text-shadow: 0 0 15px rgba(0, 255, 255, 0.9), 0 0 20px rgba(128, 0, 255, 0.5); }
}

.subtitle {
  color: rgba(255, 255, 255, 0.7);
  font-size: 14px;
  letter-spacing: 1px;
}

.divider {
  height: 1px;
  background: linear-gradient(90deg, transparent, rgba(0, 255, 255, 0.5), transparent);
  margin: 20px auto;
  width: 60%;
}

.login-form {
  margin-bottom: 28px;
}

.form-group {
  margin-bottom: 28px;
}

.form-group label {
  display: block;
  color: rgba(0, 255, 255, 0.8);
  margin-bottom: 12px;
  font-size: 12px;
  letter-spacing: 1px;
  font-weight: 500;
}

.input-wrapper {
  position: relative;
}

.tech-input {
  width: 100%;
  padding: 15px 20px;
  border: 1px solid rgba(0, 255, 255, 0.3);
  border-radius: 8px;
  background: rgba(5, 5, 15, 0.6);
  color: #ffffff;
  font-size: 14px;
  font-family: 'Courier New', monospace;
  transition: all 0.3s ease;
  position: relative;
  z-index: 2;
}

.tech-input::placeholder {
  color: rgba(255, 255, 255, 0.4);
  font-style: italic;
}

.tech-input:focus {
  outline: none;
  border-color: #00ffff;
  box-shadow: 0 0 15px rgba(0, 255, 255, 0.4);
}

.tech-input:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.input-glow {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  border-radius: 8px;
  background: linear-gradient(90deg, transparent, rgba(0, 255, 255, 0.1), transparent);
  opacity: 0;
  transition: opacity 0.3s ease;
  pointer-events: none;
  z-index: 1;
}

.tech-input:focus + .input-glow {
  opacity: 1;
}

.login-button {
  width: 100%;
  padding: 16px;
  border: 1px solid rgba(0, 255, 255, 0.4);
  border-radius: 8px;
  background: linear-gradient(45deg, rgba(0, 255, 255, 0.1), rgba(128, 0, 255, 0.1));
  color: #00ffff;
  font-size: 14px;
  font-weight: 600;
  font-family: 'Courier New', monospace;
  letter-spacing: 1px;
  cursor: pointer;
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
}

.login-button:hover:not(:disabled) {
  border-color: #00ffff;
  box-shadow: 0 0 20px rgba(0, 255, 255, 0.6);
  transform: translateY(-2px);
}

.login-button:disabled {
  opacity: 0.4;
  cursor: not-allowed;
  transform: none;
}

.button-text {
  position: relative;
  z-index: 2;
}

.button-glow {
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(0, 255, 255, 0.3), transparent);
  transition: left 0.5s ease;
}

.login-button:hover:not(:disabled) .button-glow {
  left: 100%;
}

</style>