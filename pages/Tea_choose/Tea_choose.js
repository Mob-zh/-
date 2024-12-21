// pages/Tea_choose.js
Page({

  /**
   * 页面的初始数据
   */
  data: {
    jwt:'',
    classList: [
      { id: 1, name: "软件工程", description: "1班" },
      { id: 2, name: "软件工程", description: "2班" },
    ]
  },
  goToClassDetail(event) {
    const classId = event.currentTarget.dataset.id;
    
    console.log(classId);
    wx.navigateTo({
      url: `/pages/Tea_classdetail/Tea_classdetail?classId=${classId}`
    });
  },
  CreateClass() {
    wx.navigateTo({
      url: '/pages/createClass/createClass'
    });
  },

  /**
   * 请求服务器获取班级列表
   */
  fetchClassList() {
    const that = this; // 保存上下文
    wx.request({
      url: "http://localhost:8080/teacher/home", // 替换为你的 API 地址
      method: "GET",
      header: {
        "Content-Type": "application/json",
        "Authorization": this.data.jwt
      },
      success(res) {
        if (res.statusCode === 200 && res.data) {
          console.log("获取班级列表成功:", res.data);
          that.setData({
            classList: res.data // 假设 API 返回的是一个班级数组
          });
          console.log(this.data.jwt)
        } else {
          console.error("获取班级列表失败:", res);
          wx.showToast({
            title: "获取班级列表失败",
            icon: "none"
          });
        }
      },
      fail(err) {
        console.error("请求失败:", err);
        wx.showToast({
          title: "请求失败",
          icon: "none"
        });
      }
    });
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(options) {
    const app = getApp();
    this.setData({
      jwt: app.globalData.userjwt
    })
    this.fetchClassList();
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