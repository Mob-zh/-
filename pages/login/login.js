//index.js
//获取应用实例
const app = getApp()
 let username=''
 let password=''
Page({
  data: {
    username: '',
    password: '',
    clientHeight:''
  },
  onLoad(){
    var that=this
    wx.getSystemInfo({ 
      success: function (res) { 
        console.log(res.windowHeight)
          that.setData({ 
              clientHeight:res.windowHeight
        }); 
      } 
    }) 
  },
  //获取输入款内容
  bed(e){
    username=e.detail.value
  },
  password(e){
    password=e.detail.value
  },
  //登录事件
  goadmin(){
    let flag = false  //表示账户是否存在,false为初始值
    if(username=='')
    {
      wx.showToast({
        icon:'none',
        title: '账号不能为空',
      })
    }else if(password==''){
      wx.showToast({
        icon:'none',
        title: '密码不能为空',
      })
    }else{
          
          if (username === '1') { //账户已存在
            flag=true;
            if (password !== '111111') {  //判断密码正确与否
              wx.showToast({  //显示密码错误信息
                title: '密码错误！！',
                icon: 'error',
                duration: 2500
              });
            } else {
              const app = getApp();
              app.globalData.userRole = 'student';
              wx.showToast({  //显示登录成功信息
                title: '登陆成功！！',
                icon: 'success',
                duration: 2500
              })
              flag=true;
              //跳转到非tartab页面（实际用）
              wx.navigateTo({
                url: '/pages/Stu_choose/Stu_choose',
              })
              
            }
          }
        
        
      
      
    }
  },
})
 
