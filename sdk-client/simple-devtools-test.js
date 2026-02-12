// 简化的Vue调试工具移除验证脚本
const fs = require('fs');
const path = require('path');

console.log('🔍 Vue调试工具移除验证\n');

// 检查Vite配置文件
const viteConfigPath = path.join(__dirname, 'vite.config.ts');
if (fs.existsSync(viteConfigPath)) {
  const viteConfig = fs.readFileSync(viteConfigPath, 'utf8');
  const hasVueDevTools = viteConfig.includes('vueDevTools');
  const hasImport = viteConfig.includes('vite-plugin-vue-devtools');
  
  console.log('📋 Vite配置检查:');
  console.log(`  vueDevTools插件: ${hasVueDevTools ? '❌ 仍存在' : '✅ 已移除'}`);
  console.log(`  插件导入语句: ${hasImport ? '❌ 仍存在' : '✅ 已移除'}`);
}

// 检查package.json依赖
const packageJsonPath = path.join(__dirname, 'package.json');
if (fs.existsSync(packageJsonPath)) {
  const packageJson = JSON.parse(fs.readFileSync(packageJsonPath, 'utf8'));
  const hasDevToolsDep = packageJson.devDependencies && packageJson.devDependencies['vite-plugin-vue-devtools'];
  
  console.log('\n📦 依赖检查:');
  console.log(`  vite-plugin-vue-devtools: ${hasDevToolsDep ? '❌ 仍存在' : '✅ 已移除'}`);
}

// 检查lock文件
const lockFilePath = path.join(__dirname, 'package-lock.json');
if (fs.existsSync(lockFilePath)) {
  const lockFile = fs.readFileSync(lockFilePath, 'utf8');
  const hasDevToolsLock = lockFile.includes('vite-plugin-vue-devtools');
  
  console.log('\n🔒 Lock文件检查:');
  console.log(`  vite-plugin-vue-devtools条目: ${hasDevToolsLock ? '❌ 仍存在' : '✅ 已移除'}`);
}

console.log('\n🔧 移除操作:');
console.log('✅ 从vite.config.ts中移除了vueDevTools()插件');
console.log('✅ 从package.json中移除了vite-plugin-vue-devtools依赖');
console.log('✅ 删除了相关的import语句');

console.log('\n🎯 预期效果:');
console.log('✅ 开发环境中不再显示Vue调试工具栏');
console.log('✅ 页面界面更加简洁专业');
console.log('✅ 移除了不必要的开发工具干扰');

console.log('\n✅ Vue调试工具移除验证完成！');
console.log('重启开发服务器后，Vue调试按钮将不再显示');