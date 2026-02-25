#!/usr/bin/env node

const { spawn } = require('child_process');
const { promisify } = require('util');
const setTimeoutPromise = promisify(setTimeout);

console.log('🚀 Starting Browser SDK Client development environment...\n');

// 启动Vite开发服务器
const viteProcess = spawn('npm', ['run', 'dev:web'], {
  cwd: process.cwd(),
  stdio: 'pipe',
  shell: true
});

let viteOutput = '';
viteProcess.stdout.on('data', (data) => {
  const output = data.toString();
  process.stdout.write(output);
  viteOutput += output;
  
  // 检测Vite服务器是否准备就绪
  if (output.includes('VITE v') && output.includes('ready in')) {
    console.log('\n✅ Vite development server is ready!\n');
    startElectron();
  }
});

viteProcess.stderr.on('data', (data) => {
  process.stderr.write(data.toString());
});

console.log('📦 Starting Vite development server...\n');

async function startElectron() {
  // 额外等待1秒确保服务器完全稳定
  await setTimeoutPromise(1000);
  
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
}

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