// pages/Stu_ApplyReSign/Stu_ApplyReSign.js
Page({

  /**
   * 页面的初始数据
   */
  data: {
    attendanceDate: '',
    attendanceResult: '',
    reason: '', // 补签理由
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(options) {
    // 获取传递的考勤记录信息
    const date = options.date;
    this.setData({
      attendanceDate: date,
      attendanceResult: "缺勤",
    });
  },
  onReasonInput(e) {
    this.setData({
      reason: e.detail.value,
    });
  },

  submitReSign() {
    const { attendanceDate, attendanceResult, reason } = this.data;
    if (!reason) {
      wx.showToast({
        title: '补签理由不能为空',
        icon: 'none',
      });
      return;
    }

    // 发送补签请求
    wx.request({
      url: 'http://localhost:8080/student/applyReSign', // 后端接口地址
      method: 'POST',
      data: {
        date: attendanceDate,
        result: attendanceResult,
        reason: reason,
      },
      success(res) {
        if (res.data.success) {
          wx.showToast({
            title: '补签申请已提交',
            icon: 'success',
            duration: 2000,
          });
          // 返回上一页
          wx.navigateBack();
        } else {
          wx.showToast({
            title: '提交失败，请重试',
            icon: 'none',
          });
        }
      },
      fail() {
        wx.showToast({
          title: '请求失败，请检查网络',
          icon: 'none',
        });
      },
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