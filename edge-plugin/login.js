function usc_login(userValue, passValue) {
    let href = window.location.href;
    let loginAc = /.*success/
    if (loginAc.test(href)) {
        window.close()
        return
    }

    let username = document.getElementById('username')
    username.value = userValue
    let password = document.getElementById('password')
    password.value = passValue

    let login = document.getElementById('login')
    login.click()
}

  
  // 读取配置数据
  function loadConfig(callback) {
    chrome.storage.sync.get(['config'], function(result) {
      const config = result.config;
      console.log('Config loaded:', config);
      callback(config);
    });
  }
  


  
  // 读取配置
    loadConfig(function(config) {
        usc_login(config.username, config.password)
    });

// http://210.43.112.9/cgi-bin/rad_user_info