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
  const [minutes, seconds] = e.detail.value.split(':').map(Number); // 分割并转换为数字
  const res = minutes * 60 + seconds; // 计算总秒数
  this.setData({
    signInDuration: e.detail.value,
    remainingTime:res
  });
},

 // 启动签到
 startSignIn() {
  const { classId } = this.data; // 假设班级 ID 存储在 data 中

  // 向后端请求签到码
  wx.request({
    url: `http://localhost:8080/api/startSignIn`, // 替换为实际接口
    method: "POST",
    header: {
      'Authorization': app.globalData.jwt, // 添加用户凭证
      'Content-Type': 'application/json'
    },
    data: {
      classId: classId
    },
    success: (res) => {
      if (res.statusCode === 200 && res.data.signInCode) {
        // 将返回的签到码存储到全局变量和页面数据
        const newSignInCode = res.data.signInCode;
        app.globalData.currentSignInCode = newSignInCode;

        // 设置页面数据
        this.setData({
          isSignInStarted: true,
          signInCode: newSignInCode,
          signInCount: 0, // 重置签到人数
          remainingTime: 60 // 重置倒计时
        });

        // 启动倒计时
        this.startCountdown();
      } else {
        wx.showToast({
          title: res.data.message || "无法获取签到码",
          icon: "none"
        });
      }
    },
    fail: (err) => {
      wx.showToast({
        title: "网络错误，无法发起签到",
        icon: "none"
      });
      console.error(err);
    }
  });
},

// 启动倒计时
startCountdown() {
  const timer = setInterval(() => {
    if (this.data.remainingTime <= 0) {
      clearInterval(timer); // 停止计时
      this.setData({ timer: null }); // 清空定时器

      // 清空全局签到码
      app.globalData.currentSignInCode = null;

      // 提示签到已结束
      wx.showToast({
        title: "签到已结束",
        icon: "none"
      });

      // 重置页面状态
      this.resetPageState();
    } else {
      this.setData({
        remainingTime: this.data.remainingTime - 1
      });
    }
  }, 1000);

  this.setData({ timer });
},

// 重置页面状态
resetPageState() {
  this.setData({
    isSignInStarted: false,
    signInCode: "",
    signInCount: 0,
    remainingTime: 60
  });
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