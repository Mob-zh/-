// pages/classDetail/classDetail.js
Page({

  /**
   * 页面的初始数据
   */
  data: {

  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(options) {
    const classId = options.id;
    this.setData({ classId });

    // 模拟获取班级详情
    const allClasses = [
      { id: 1, name: "软件工程", description: "2024年秋", teacher: "周老师" },
      { id: 2, name: "编译原理", description: "2024年秋", teacher: "袁老师" },
      { id: 3, name: "数据结构", description: "2023年秋", teacher: "李老师" }
    ];
    const classDetail = allClasses.find((cls) => cls.id == classId);
    this.setData({ classDetail });
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