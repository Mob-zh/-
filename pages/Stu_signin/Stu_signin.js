// pages/Stu_signin/Stu_signin.js
Page({

  /**
   * 页面的初始数据
   */
  data: {
    jwt:"",
    classId:'',
    hideTabbar: false, // 控制 tabbar 显示隐藏
    singInCode:""
  },

  handleInput(e) {
    const value = e.detail.value; // 获取输入框的值
    this.setData({
      singInCode: value // 动态更新对应字段的值
    });
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(options) {
      const id = options.classId;
      const app = getApp();
      this.setData({
        jwt: app.globalData.userjwt,
        classId:id
      })

      // 如果页面参数中存在 hideTabbar 参数且为 true，则隐藏 tabbar
      if (options.hideTabbar === "true") {
        this.setData({
          hideTabbar: true
        });
      }
  },
  signIn(){
    const that = this; // 保存上下文
    const sign_code = this.data.singInCode;
    wx.request({
      url: "http://localhost:8080/student/"+this.data.classId+"/sign", // 替换为你的 API 地址
      method: "POST",
      header: {
        "Content-Type": "application/json",
        "Authorization": this.data.jwt
      },
      data:{
        "sign_code":sign_code
      },
      success(res) {
        if (res.statusCode === 200 ) {
          console.log("签到成功:", res.data);
          wx.showToast({
            title: "签到成功",
            icon: "none"
          });
        } else {
          console.error("签到失败:", res);
          wx.showToast({
            title: "签到失败",
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