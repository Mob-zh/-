//index.js
//获取应用实例
const app = getApp()

Page({
  data: {
    username: '',
    password: '',
    clientHeight:''
  },
  onLoad(){
    var that=this
    wx.getSystemInfo({ 
      success: function (res) { 
        console.log(res.windowHeight)
          that.setData({ 
              clientHeight:res.windowHeight
        }); 
      } 
    }) 
  },

  goadmin() {
    const username = this.data.username; // 获取用户名
    const password = this.data.password; // 获取密码
  
    // 检查输入是否为空
    if (username === '') {
      wx.showToast({
        icon: 'none',
        title: '账号不能为空',
      });
      return;
    }
  
    if (password === '') {
      wx.showToast({
        icon: 'none',
        title: '密码不能为空',
      });
      return;
    }
  
    // 发送请求到后端
    wx.request({
      url: 'http://localhost:8080/student/Login', // 后端接口地址
      method: 'POST',
      header: {
        'Content-Type': 'application/json',
      },
      data: {
        user_id: username,
        pwd: password,
      },
      success(res) {
        // 根据后端返回的结果处理
        if (res.data.success) { 
          wx.showToast({
            title: '登录成功',
            icon: 'success',
          });
        const app = getApp(); // 获取全局 App 实例
        app.globalData.userjwt = res.data.token; 
          
          // 跳转到 Stu_choose 页面
          wx.navigateTo({
            url: '/pages/Stu_choose/Stu_choose',
          });
        } else {
          wx.showToast({
            title: res.data.message || '登录失败', // 假设后端返回 `message` 字段表示错误信息
            icon: 'none',
          });
        }
      },
      fail(error) {
        // 请求失败处理
        wx.showToast({
          title: '请求失败，请稍后再试',
          icon: 'none',
        });
        console.error('请求失败:', error);
      },
    });
  },
  

  //获取输入款内容
  bed(e){
    this.data.username=e.detail.value
  },
  password(e){
    this.data.password=e.detail.value
  },

  
})
 
