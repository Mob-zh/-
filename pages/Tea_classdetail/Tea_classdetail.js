// pages/Tea_classdetail/Tea_classdetail.js
Page({

  /**
   * 页面的初始数据
   */
  data: {
    classId: null, // 课程ID
    classInfo: {}, // 课程详细信息
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(options) {
    const id = options.id; // 获取传递的课程ID
    this.setData({ classId: id });
    // 模拟获取课程信息（可替换为接口请求）
    const allClasses = [
      { id: 1, name: "软件工程", description: "1班", teacher: "周老师" ,schedule: "周一到周五，8:00-10:00",location: "A301教室"},
      { id: 2, name: "软件工程", description: "2班", teacher: "周老师" ,schedule: "周一到周五，10:20-12:00",location: "A301教室"},
    ];
    const classDetail = allClasses.find((cls) => cls.id == id);
    if (classDetail) {
      this.setData({ classInfo: classDetail });
    } else {
      console.error("未找到对应课程信息");
    }
  },
  // 发起签到
  initiateSignIn() {
    wx.navigateTo({
      url: `/pages/Tea_StartSignIn/Tea_StartSignIn?classId=${classId}&name=${encodeURIComponent(name)}&description=${encodeURIComponent(description)}`
    });
  },

  // 查看签到情况
  viewSignInStatus() {
    wx.navigateTo({
      url: `/pages/Tea_SignInStatus/Tea_SignInStatus?classId=${this.data.classId}`
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