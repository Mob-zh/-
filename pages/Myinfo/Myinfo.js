// pages/Stu_message/Stu_message.js
Page({

  /**
   * 页面的初始数据
   */
  data: {
    userInfo: {
      userRole: '' ,// 默认空值
      avatarUrl: "/image/person.png", // 默认头像
      nickName: "用户昵称",
      phone: "" // 手机号
    }
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(options) {
    // 模拟从全局获取用户数据
    const app = getApp();
    const userInfo = app.globalData.userInfo || {};
    this.setData({ userInfo });
    this.setData({
      userRole: app.globalData.userRole // 赋值全局变量
    });
  },
  editProfile() {
    wx.navigateTo({
      url: '/pages/editProfile/editProfile' // 跳转到修改资料页面
    });
  },

  changepasswd() {
    wx.navigateTo({
      url: '/pages/changepasswd/changepasswd' // 跳转到修改资料页面
    });
  },

  logout() {
    wx.showModal({
      title: "提示",
      content: "确定要退出登录吗？",
      success: (res) => {
        if (res.confirm) {
          wx.clearStorageSync(); // 清除本地存储
          wx.redirectTo({
            url: '/pages/login/login' // 跳转到登录页面
          });
        }
      }
    });
  },

  navigateToHome() {
    // 根据 userRole 跳转到不同页面
    if (this.data.userRole === 'student') {
      wx.navigateTo({
        url: '/pages/Stu_choose/Stu_choose'
      });
    } else {
      wx.navigateTo({
        url: '/pages/Tea_choose/Tea_choose'
      });
    }
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