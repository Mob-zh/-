// pages/changepasswd/changepasswd.js
Page({

  /**
   * 页面的初始数据
   */
  data: {
    oldPassword: '', // 原密码
    newPassword: '', // 新密码
    confirmPassword: '', // 确认新密码

    identity:"",
    jwt:""
  },
  handleInput(e) {
    const field = e.currentTarget.dataset.field; // 获取字段名称
    const value = e.detail.value; // 获取输入框的值
    this.setData({
      [field]: value // 动态更新对应字段的值
    });
  },

  changePassword() {
    const  oldPassword = this.data.oldPassword;
    const  newPassword = this.data.newPassword;
    const  confirmPassword  = this.data.confirmPassword; // 解构 data 中的变量
    console.log(oldPassword,newPassword,confirmPassword);

    // 验证输入是否有效
    if (!oldPassword || !newPassword || !confirmPassword) {
      wx.showToast({
        title: '请填写完整信息',
        icon: 'none'
      });
      return;
    }
    if (newPassword !== confirmPassword) {
      wx.showToast({
        title: '新密码和确认密码不一致',
        icon: 'none'
      });
      return;
    }
    // 准备请求体
    const requestBody = {
      oldPassword: oldPassword,
      newPassword: newPassword,
      repeated_new_pwd: newPassword
    };
    

    // 发送 POST 请求
    wx.request({
      url: 'http://localhost:8080/'+this.data.identity+'/changePwd',
      method: 'POST',
      header: {
        'Content-Type': 'application/json',
        'Authorition':this.data.jwt
      },
      data: requestBody,
      success(res) {
        const statusCode = res.statusCode;
        if (statusCode === 200) {
          wx.showToast({
            title: '密码修改成功',
            icon: 'success',
          });
          // 在这里处理后端返回的数据

        } else {
          wx.showToast({
            title: res.data.error || '修改失败',
            icon: 'none',
          });
        }
      },
      fail(error) {
        wx.showToast({
          title: '请求失败，请稍后再试',
          icon: 'none',
        });
        console.error('请求失败:', error);
      }
    });
  },
  

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(options) {
    // 模拟从全局获取用户身份
    const app = getApp();
    const userRole = app.globalData.userRole;
    const userjwt = app.globalData.userjwt;
    this.setData({
      identity: userRole, // 回显身份
      jwt: userjwt
    });
  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady() {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow() {

  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide() {

  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload() {

  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh() {

  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom() {

  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage() {

  }
})