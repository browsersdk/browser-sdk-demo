#!/usr/bin/env node

const { spawn } = require('child_process');
const path = require('path');

console.log('🚀 Starting Browser SDK Client development environment...\n');

// 启动Vite开发服务器
const viteProcess = spawn('npm', ['run', 'dev:web'], {
  cwd: process.cwd(),
  stdio: 'inherit',
  shell: true
});

console.log('📦 Vite development server starting on http://localhost:5173\n');

// 等待Vite服务器启动后再启动Electron
setTimeout(() => {
  console.log('🖥️  Starting Electron application...\n');
  
  const electronProcess = spawn('npm', ['run', 'dev:electron'], {
    cwd: process.cwd(),
    stdio: 'inherit',
    shell: true,
    env: {
      ...process.env,
      NODE_ENV: 'development'
    }
  });

  electronProcess.on('close', (code) => {
    console.log(`\nElectron process exited with code ${code}`);
    viteProcess.kill();
    process.exit(code);
  });

  electronProcess.on('error', (error) => {
    console.error('Failed to start Electron:', error);
    viteProcess.kill();
    process.exit(1);
  });
}, 5000); // 增加到5秒等待时间

// 处理退出信号
process.on('SIGINT', () => {
  console.log('\n🛑 Shutting down development environment...');
  viteProcess.kill();
  process.exit(0);
});

process.on('SIGTERM', () => {
  console.log('\n🛑 Shutting down development environment...');
  viteProcess.kill();
  process.exit(0);
});