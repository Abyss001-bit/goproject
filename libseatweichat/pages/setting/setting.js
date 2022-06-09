// pages/setting/setting.js
var app = getApp();//写在页面顶部page()外
Page({

  /**
   * 页面的初始数据
   */
  data: {

  },


  tochangeinfo:function(e){
    wx.navigateTo({
      url: '../changeinfo/changeinfo',
    })
  },
  tologout:function(e){
    wx.request({
      url: app.globalData.URL + '/user/loginOut',
      method:"POST",
      data:{
      },
      header:{
        "content-type": "application/x-www-form-urlencoded",
        "content-type": "application/json", 
        "Cookie":wx.getStorageSync('cookieKey')
      },
      success(res){
        console.log(res.data);
        if(res.data.msg == "ok"){
          wx.removeStorage({
            key: 'cookieKey',
            success (res) {
              console.log(res)
            }
          });
          wx.showToast({
            title: '退出登录成功',
            icon:'success',
            duration:2000
          });
          var pages = getCurrentPages(); // 当前页面
          var beforePage = pages[0]; // 前一个页面
          setTimeout(function(){
             beforePage.onLoad(); // 执行前一个页面的onLoad方法
          },1000)
          // console.log(wx.getStorageSync('cookieKey'))
          wx.redirectTo({
            url: '../login/login',
          })
        }else{
          wx.showToast({
            title: '退出登录失败',
            icon:'error',
            duration:2000
          });
        }
      },
      fail(res){
        console.log(res.data);
      }
    })
  },
  onLoad: function (options) {

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