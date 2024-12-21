// pages/Tea_StartSignIn/Tea_StartSignIn.js
Page({

  /**
   * 页面的初始数据
   */
  data: {
    classInfo:{},
    classId:"",
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
      url: "http://localhost:8080/teacher/"+this.data.classId+"/info", // 替换为你的 API 地址
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
  const remainingTime = this.data.remainingTime;
  console.log(remainingTime);
  // 向后端请求签到码
  wx.request({
    url: `http://localhost:8080/teacher/`+this.data.classId+`/sign/start`, // 替换为实际接口
    method: "POST",
    header: {
      'Authorization': this.data.jwt, // 添加用户凭证
      'Content-Type': 'application/json'
    },
    data: {
       'checking_seconds':remainingTime
    },
    success: (res) => {
      if (res.statusCode === 200 && res.data.sign_in_code) {
        // 将返回的签到码存储到全局变量和页面数据
        const newSignInCode = res.data.sign_in_code;

        // 设置页面数据
        this.setData({
          isSignInStarted: true,
          signInCode: newSignInCode,
          signInCount: 0, // 重置签到人数
          remainingTime: remainingTime // 重置倒计时
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
// 启动签到
checkSignInNum() {

  // 向后端请求签到码
  wx.request({
    url: `http://localhost:8080/teacher/`+this.data.classId+`/sign/count`, // 替换为实际接口
    method: "GET",
    header: {
      'Authorization': this.data.jwt, // 添加用户凭证
      'Content-Type': 'application/json'
    },
    success: (res) => {
      if (res.statusCode === 200) {
        // 将返回的签到码存储到全局变量和页面数据
        const newSignInCount = res.data.signed_in_count;

        // 设置页面数据
        this.setData({
          signInCount: newSignInCount, // 重置签到人数
        });

      } else {
        wx.showToast({
          title: res.data.message,
          icon: "none"
        });
      }
    },
    fail: (err) => {
      wx.showToast({
        title: "网络错误,获取签到人数失败",
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
      this.checkSignInNum();
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