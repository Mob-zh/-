// pages/Tea_StartSignIn/Tea_StartSignIn.js
Page({

  /**
   * 页面的初始数据
   */
  data: {
    classInfo:{},
    signInDuration: "00:30",  // 默认签到有效时间30分钟
    signInCode: "------",  // 默认签到码
    isSignInStarted: false,  // 是否已开始签到
    signInCount: 0,  // 当前签到人数
    remainingTime: 30,  // 剩余签到时间，单位秒
    timer: null,  // 用于存储计时器引用
    signInStatus: "notStarted"  // 签到状态（未开始，进行中，已结束）
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(options) {
    const id = options.classId; // 获取传递的课程ID

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
    console.log(this.data.classInfo);
  },
 // 签到有效时间变更时更新
 onTimeChange(e) {
  this.setData({
    signInDuration: e.detail.value
  });
},

// 启动签到
startSignIn() {
  // 隐藏开始签到按钮
  this.setData({
    isSignInStarted: true
  });

  // 启动倒计时
  this.startCountdown();

  // 模拟签到人数，通常此部分数据来自服务器
  this.setData({
    signInCount: 5  // 假设已经有5人签到
  });

  // 模拟签到成功，设置签到码
  this.setData({
    signInCode: this.generateSignInCode()
  });
},

// 启动倒计时
startCountdown() {
  const timer = setInterval(() => {
    if (this.data.remainingTime <= 0) {
      clearInterval(timer); // 停止计时
      wx.showToast({
        title: "签到已结束",
        icon: "none"
      });
    } else {
      this.setData({
        remainingTime: this.data.remainingTime - 1
      });
    }
  }, 1000);

  this.setData({ timer });
},

// 生成六位签到码
generateSignInCode() {
  return Math.floor(100000 + Math.random() * 900000).toString();
},

// 页面卸载时清理定时器
onUnload() {
  if (this.data.timer) {
    clearInterval(this.data.timer);
  }
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