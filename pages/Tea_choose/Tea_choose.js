// pages/Tea_choose.js
Page({

  /**
   * 页面的初始数据
   */
  data: {
    classList: [
      { id: 1, name: "软件工程", description: "1班" },
      { id: 2, name: "软件工程", description: "2班" },
    ]
  },
  goToClassDetail(event) {
    const classId = event.currentTarget.dataset.id;
    const className = event.currentTarget.dataset.name;
    const classDes = event.currentTarget.dataset.description;
    
    console.log(classId);
    wx.navigateTo({
      url: `/pages/Tea_StartSignIn/Tea_StartSignIn?classId=${classId}&name=${className}&description=${classDes}`
    });
  },
  CreateClass() {
    wx.navigateTo({
      url: '/pages/createClass/createClass'
    });
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(options) {

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