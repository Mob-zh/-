// pages/Stu_AttendanceRecord/Stu_AttendanceRecord.js
Page({

  /**
   * 页面的初始数据
   */
  data: {
    records: [
      { "date": "2024-12-01", "time":"16:07","result": "出勤" },
      { "date": "2024-12-02","time":"17:25", "result": "未出勤" },
      { "date": "2024-12-03", "time":"16:02","result": "出勤" },
      { "date": "2024-12-04","time":"17:30", "result": "未出勤" }
    ]
  },

  /**
   * 获取考勤记录
   */
  fetchAttendanceRecords() {
    const app = getApp(); // 假设JWT存储在全局变量中
    wx.request({
      url: "http://localhost:8080/api/student/attendanceRecords", // 替换为实际后端接口
      method: "GET",
      header: {
        "Content-Type": "application/json",
        "Authorization": app.globalData.jwt // 使用JWT认证
      },
      success: (res) => {
        if (res.statusCode === 200 && res.data) {
          this.setData({
            attendanceRecords: res.data.records // 假设返回数据格式为 { records: [...] }
          });
        } else {
          wx.showToast({
            title: "获取考勤数据失败",
            icon: "none"
          });
        }
      },
      fail: (err) => {
        console.error("请求失败:", err);
        wx.showToast({
          title: "网络请求失败",
          icon: "none"
        });
      }
    });
  },

  /**
   * 申请补签
   */
  applyReSign(e) {
    const date = e.currentTarget.dataset.date;
    console.log(date);
    wx.navigateTo({
      url: `/pages/Stu_ApplyReSign/Stu_ApplyReSign?date=${date}`
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