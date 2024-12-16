// pages/Stu_signin/Stu_signin.js
Page({

  /**
   * 页面的初始数据
   */
  data: {
    hideTabbar: false // 控制 tabbar 显示隐藏
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(options) {
      // 如果页面参数中存在 hideTabbar 参数且为 true，则隐藏 tabbar
      if (options.hideTabbar === "true") {
        this.setData({
          hideTabbar: true
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