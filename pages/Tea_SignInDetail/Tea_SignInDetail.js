// pages/Tea_SignInDetail/Tea_SignInDetail.js
Page({

  /**
   * 页面的初始数据
   */
  data: {
    studentRecords: [
      { "id": 1, "name": "张三", "studentId": "20240101", "status": "出勤" },
      { "id": 2, "name": "李四", "studentId": "20240102", "status": "未出勤" },
      { "id": 3, "name": "王五", "studentId": "20240103", "status": "出勤" },
      { "id": 4, "name": "赵六", "studentId": "20240104", "status": "未出勤" },
      { "id": 5, "name": "孙七", "studentId": "20240105", "status": "出勤" },
      { "id": 6, "name": "周八", "studentId": "20240106", "status": "未出勤" },
      { "id": 7, "name": "吴九", "studentId": "20240107", "status": "出勤" },
      { "id": 8, "name": "郑十", "studentId": "20240108", "status": "出勤" },
    ],
    classId: null, // 当前班级ID
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(options) {
    const classId = options.classId;
    this.setData({ classId });
  },
  
  // 补签操作
  makeupSignIn(event) {
    const studentId = event.currentTarget.dataset.studentId; // 获取学生ID
    const updatedRecords = this.data.studentRecords.map((record) => {
      if (record.studentId === studentId && record.status === "未出勤") {
        return { ...record, status: "出勤" }; // 修改为出勤状态
      }
      return record;
    });

    this.setData({ studentRecords: updatedRecords }); // 更新数据
    wx.showToast({
      title: "补签成功",
      icon: "success",
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