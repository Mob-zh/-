// pages/classDetail/classDetail.js
Page({

  /**
   * 页面的初始数据
   */
  data: {
    jwt:'',
    classId: null, // 课程ID
    classInfo: {}, // 课程详细信息
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(options) {
    const id = options.id; // 获取传递的课程ID
    const app = getApp();
    this.setData({
      jwt: app.globalData.userjwt
    })
    this.setData({ classId: id });
    this.fetchClassDetail();

  },

  fetchClassDetail(){
    const that = this; // 保存上下文
    wx.request({
      url: "http://localhost:8080/student/"+this.data.classId+"/info", // 替换为你的 API 地址
      method: "GET",
      header: {
        "Content-Type": "application/json",
        "Authorization": this.data.jwt
      },
      success(res) {
        if (res.statusCode === 200 && res.data) {
          console.log("获取班级详情成功:", res.data);
          that.setData({
            classInfo: res.data // 假设 API 返回的是一个班级数组
          });
          //console.log(this.data.classList);
        } else {
          console.error("获取班级详情失败:", res);
          wx.showToast({
            title: "获取班级详情失败",
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


  SignIn(){
    const classId = this.data.classId; 
    console.log(classId);
    wx.navigateTo({
      url: `/pages/Stu_signin/Stu_signin?classId=${classId}&hideTabbar=true`
    });
  },

  AttendanceRecord(){
    const classId = this.data.classId; 
    console.log(classId);
    wx.navigateTo({
      url: `/pages/Stu_AttendanceRecord/Stu_AttendanceRecord?classId=${classId}`
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