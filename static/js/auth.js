export async function auth() {
  var isOk = false;
  var mes = '';
  // 获取验证本地token
  var access_token = localStorage.getItem('access_token');
  console.log('验证access_token', access_token);
  // 验证access_token
  if (access_token != null) {
    await fetch('/user/auth', {
      method: 'POST', // 使用POST方法发送数据
      headers: {
        Authorization: access_token,
      },
    })
      .then((response) => response.json())
      .then((data) => {
        if (data == 'success') {
          isOk = true;
        }
      })
      .catch((error) => {
        console.log('认证失败：' + String(error));
      });
  }
  // 验证成功
  if (isOk) {
    return true;
  }
  // token验证失败
  // 使用refresh token
  var refresh_token = localStorage.getItem('refresh_token');
  if (refresh_token != null) {
    await fetch('/user/refresh', {
      method: 'POST', // 使用POST方法发送数据
      headers: {
        Authorization: refresh_token,
      },
    })
      .then((response) => {
        if (response.status == 200) {
          isOk = true;
        }
        return response.json();
      })
      .then((data) => {
        if (isOk) {
          console.log(data);
          localStorage.setItem('access_token', data.access_token);
        } else {
          throw new Error(data);
        }
      })
      .catch((error) => {
        console.log('认证失败：' + String(error));
      });
  }
  return isOk;
}

export async function login(username, password) {
  var mes = '';
  var ok = false;
  await fetch('/user/login', {
    method: 'POST', // 使用POST方法发送数据
    headers: {
      'Content-Type': 'application/json', // 设置内容类型为JSON
    },
    body: JSON.stringify({
      // 将登录信息转换为JSON字符串
      username,
      password,
    }),
  })
    .then((response) => response.json())
    .then((data) => {
      mes = data.message;
      if (data.status == 'ok') {
        ok = true;
        localStorage.setItem('access_token', data.access_token);
        localStorage.setItem('refresh_token', data.refresh_token);
      } else {
        ok = false;
      }
    })
    .catch((error) => {
      mes = '登录失败，请重试：' + String(error);
      console.log(mes);
      ok = false;
    });
  return { mes, ok };
}
