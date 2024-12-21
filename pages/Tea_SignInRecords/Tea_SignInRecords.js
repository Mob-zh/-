// pages/Tea_SignInRecords/Tea_SignInRecords.js
Page({

  /**
   * 页面的初始数据
   */
  data: {
    classId:1,
    records: [
      { id: 1, date: "2024-12-01", time: "08:00"},
      { id: 2, date: "2024-12-02", time: "08:10"},
      { id: 3, date: "2024-12-03", time: "08:05"},
    ],
  },
  // 跳转到签到详情页面
  goToSignInDetail(event) {
    const recordId = event.currentTarget.dataset.id;
    const classId = this.data.classId;
    wx.navigateTo({
      url: `/pages/Tea_SignInDetail/Tea_SignInDetail?recordId=${recordId}&classId=${classId}`,
    });
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(options) {

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