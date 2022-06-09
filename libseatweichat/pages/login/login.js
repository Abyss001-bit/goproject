// pages/login/login.js
var app = getApp();//写在页面顶部page()外
Page({

  /**
   * 页面的初始数据
   */
  data: {
    name:"",
    password:"",
  },
  nameInput: function (e) {
    var that = this;
    that.setData({
      name: e.detail.value
    })
  },
passwordInput: function (e) {
    var that = this;
    that.setData({
      password: e.detail.value
    })
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
  },
  onTapRegister:function(){  //按钮页面跳转
    wx.navigateTo({ //微信开发的内置跳转方法
      url: '../register/register', //这里是要跳转的路径
    })
  },
  onTapFindbackpassword:function(){
    wx.navigateTo({
      url: '../findbackpassword/findbackpassword',
    })
  },
  submit:function(e){
    wx.request({
      // 请求地址
      url: app.globalData.URL +  '/user/loginBypassword',
      //请求方法
      method:"POST",
      //请求参数
      data:{
        name:e.detail.value.name,
        password:e.detail.value.password
      },
      //请求头
      header:{
        "content-type": "application/x-www-form-urlencoded",
        "content-type": "application/json" // 默认
      },
      //请求成功回调
      success(res){
        console.log(res.data);
        if(res.data.msg == "ok"){
          // 保存cookie
          wx.setStorageSync('cookieKey', res.header["Set-Cookie"])
          var cookie = wx.getStorageSync('cookieKey')
          console.log(cookie);
          wx.showToast({
            title: '登陆成功',
            icon:'success',
            duration:2000,
          });
          console.log(res.data.body[0].Admin)
          if (res.data.body[0].Admin == true){
            // 管理员登录
            setTimeout(function(){
              wx.reLaunch({
                url: '/Admins/admin/admin',
              })
            },1000)
          }else{
            // 普通用户登录 
            setTimeout(function(){
              wx.reLaunch({
                url: '/pages/main/main',
              })
            },1000)
          }
        }else if (res.data.msg == "用户信誉积分不足,无权查看"){
          wx.showToast({
            title: '信誉积分不足',
            icon:'error',
            duration:2000,
          })
        }else{
          wx.showToast({
            title: '用户名或密码错误',
            icon:'error',
            duration:2000,
          })
        }
      },
      //请求失败回调
      fail(res){
        console.log(res.errMsg);
      }
    })
  },
  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady: function () {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow: function () {

  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide: function () {

  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload: function () {

  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh: function () {

  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom: function () {

  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage: function () {

  }
})