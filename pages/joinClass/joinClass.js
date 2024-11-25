// pages/joinClass/joinClass.js
Page({

  /**
   * 页面的初始数据
   */
  data: {
    data: {
      identity: "", // 用户身份
      classCode: "", // 班级码
    },
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(options) {
    // 模拟从全局获取用户身份
    const app = getApp();
    const userRole = app.globalData.userRole;

    this.setData({
      identity: userRole === "teacher" ? "老师" : "学生", // 回显身份
    });
  },
  onClassCodeInput(e) {
    this.setData({ classCode: e.detail.value });
  },

  joinClass() {
    const { classCode } = this.data;
    if (!classCode) {
      wx.showToast({
        title: "班级码不能为空",
        icon: "none",
      });
      return;
    }

    // 模拟提交加入班级请求
    wx.showLoading({ title: "加入中..." });
    setTimeout(() => {
      wx.hideLoading();
      wx.showToast({
        title: "加入成功！",
        icon: "success",
      });
      // 跳转到班级详情页面
      wx.redirectTo({
        url: `/pages/classDetail/classDetail?classCode=${classCode}`,
      });
    }, 1000);
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