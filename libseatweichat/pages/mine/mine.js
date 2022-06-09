// pages/mine/mine.js
var app = getApp();//写在页面顶部page()外
Page({

  /**
   * 页面的初始数据
   */
  data: {
    name:"获取信息",
    integral:0,
    user_image:"../../static/头像.jpg"
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
   var that = this;
    setInterval(function () {
      that.requinfo();
    },1000)    //代表1秒钟发送一次请求
  },
  requinfo:function(){
    var that = this;
    wx.request({
      url: app.globalData.URL +  '/user/showUserInfo',
      method:"POST",
      data:{
      },
      dataType:"json",
      header:{
        "content-type": "application/x-www-form-urlencoded",
        "content-type": "application/json",
        "Cookie":wx.getStorageSync('cookieKey')
      },
      success(res){
        console.log(res.data);
        if(res.data.msg == "ok"){
          console.log(res.data.body[0].Name);
          console.log(res.data.body[0].Integral);
          that.setData({
            name : res.data.body[0].Name,
            integral : res.data.body[0].Integral,
          });
        }
      },
      fail(res){
        console.log(res.errMsg);
      }
    })
  },

  tofuwu:function(){
    wx.navigateTo({
      url: '../problemback/problemback',
    })
  },
  tosetting:function(){
    wx.navigateTo({
      url: '../setting/setting',
    })
  },
  toadmin:function(){
    wx.navigateTo({
      url: '../toadmin/toadmin',
    })
  },
  toaboutme:function(){
    wx.navigateTo({
      url: '../aboutme/aboutme',
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