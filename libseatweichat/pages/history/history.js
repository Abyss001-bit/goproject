// pages/history/history.js
var app = getApp();//写在页面顶部page()外
Page({

  /**
   * 页面的初始数据
   */
  data: {
    historylist:[],
    i:0,
    viewcolor:'',
  },
  // 0 已完成
  // 1 待完成
  // 2 完成中 扫码一次 
  // 3 爽约未完成 
  tip0:function(e){
    var that = this;
    that.setData({
      i:0,
      viewcolor:"rgb(125, 128, 126)",
    });
  },
  tip1:function(e){
    var that = this;
    that.setData({
      i:1,
      viewcolor:"rgb(50, 213, 241)",
    });
  },
  tip2:function(e){
    var that = this;
    that.setData({
      i:2,
      viewcolor:"rgb(96, 243, 152)",
    });
  },
  tip3:function(e){
    var that = this;
    that.setData({
      i:3,
      viewcolor:"rgb(190, 81, 81)",
    });
  },
  /**
   * 生命周期函数--监听页面加载
   */

  onLoad: function (options) {
    var that = this;
    setInterval(function () {
      that.requinfo();
    }, 1000)    //代表1秒钟发送一次请求
  },
  requinfo:function(){
    var that = this;
    wx.request({
      url: app.globalData.URL +  '/user/ShowHistory',
      method:"POST",
      data:{

      },
      dataType:"json",
      header:{
        "content-type": "application/x-www-form-urlencoded",
        "content-type": "application/json",
        "Cookie":wx.getStorageSync('cookieKey')
      },
      success:function(res){
        console.log(res.data);
        if(res.data.msg == "ok"){
          that.setData({
            historylist:res.data.body,
          });
        }
      },
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