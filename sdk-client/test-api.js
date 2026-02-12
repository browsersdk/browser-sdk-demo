// API测试脚本
const BASE_URL = 'http://localhost:7888';

async function testLogin() {
  try {
    console.log('测试登录API...');
    
    const response = await fetch(`${BASE_URL}/api/app/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        code: '666666',
        devId: '',
        inviteCode: '',
        name: '',
        os: '',
        osVer: '',
        password: 'Tt123456',
        src: '',
        tid: '',
        username: '13800138000',
        uuid: '',
        ver: ''
      })
    });

    const data = await response.json();
    console.log('登录响应:', JSON.stringify(data, null, 2));
    
    if (data.code === 200) {
      console.log('✅ 登录成功');
      console.log('访问令牌:', data.data.accessToken.substring(0, 20) + '...');
      return data.data.accessToken;
    } else {
      console.log('❌ 登录失败:', data.msg);
      return null;
    }
  } catch (error) {
    console.error('❌ 登录请求失败:', error.message);
    return null;
  }
}

async function testUserInfo(token) {
  if (!token) {
    console.log('没有有效的访问令牌');
    return;
  }

  try {
    console.log('测试获取用户信息API...');
    
    const response = await fetch(`${BASE_URL}/api/app/user/info`, {
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      }
    });

    const data = await response.json();
    console.log('用户信息响应:', JSON.stringify(data, null, 2));
    
    if (data.code === 200) {
      console.log('✅ 获取用户信息成功');
      console.log('用户昵称:', data.data.nickname);
      console.log('用户名:', data.data.username);
    } else {
      console.log('❌ 获取用户信息失败:', data.msg);
    }
  } catch (error) {
    console.error('❌ 获取用户信息请求失败:', error.message);
  }
}

async function main() {
  console.log('开始API测试...\n');
  
  const token = await testLogin();
  console.log('');
  await testUserInfo(token);
  
  console.log('\nAPI测试完成');
}

// 运行测试
main();