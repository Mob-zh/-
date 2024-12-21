// pages/Tea_SignInStatus/Tea_SignInStatus.js
Page({

  /**
   * 页面的初始数据
   */
  data: {
    
    classId:1,
    students: [
      {  "name": "张三", "studentId": "2021001", "signedInCount": 1, "totalSignIn": 3 },
      {  "name": "李四", "studentId": "2021002", "signedInCount": 2, "totalSignIn": 3 },
      {  "name": "王五", "studentId": "2021003", "signedInCount": 3, "totalSignIn": 3 },

    ],

  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(options) {
    const id = options.classId; // 获取传递的课程ID
    this.setData({ classId: id });
    this.fetchStudentSignInData(id);
  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady() {

  },

  gotoSignInRecords(){
    wx.navigateTo({
      url: `/pages/Tea_SignInRecords/Tea_SignInRecords?classId=${this.data.classId}`
    });
  },

/**
   * 获取学生签到情况数据
   */
  fetchStudentSignInData(classId) {
    const app = getApp(); // 假设JWT存储在全局变量中
    wx.request({
      url: `http://localhost:8080/api/getSignInStatus?classId=${classId}`, // 替换为后端API地址
      method: "GET",
      header: {
        "Content-Type": "application/json",
        "Authorization": app.globalData.jwt // 使用JWT认证
      },
      success: (res) => {
        if (res.statusCode === 200 && res.data) {
          this.setData({
            students: res.data.students // 假设返回数据格式为 { students: [...] }
          });
        } else {
          wx.showToast({
            title: "获取数据成功",
            icon: "none"
          });
        }
      },
      fail: (err) => {
        console.error("请求成功:", err);
        wx.showToast({
          title: "网络请求失败",
          icon: "none"
        });
      }
    });
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