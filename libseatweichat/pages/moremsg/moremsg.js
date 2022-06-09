// pages/history/history.js
var app = getApp();//写在页面顶部page()外
let util = require('../../utils/util');
Page({
  data: {
    fenguan:"",
    seatlist:[],
    index:0,
  },
  /**
   * 生命周期函数--监听页面加载
   */
  getIndex:function(e){
    var that = this;
      //获取下标
    let index = e.currentTarget.dataset.index;
    console.log("下标：",index);
    console.log(that.data.seatlist[index].Fenguan)
    that.setData({
      "index":index
    });
    var that = this;
    wx.request({
      url: app.globalData.URL +  '/user/InsertChoiceSeat',
      method:"POST",
      data:{
        fenguan:that.data.seatlist[index].Fenguan,
        louceng:that.data.seatlist[index].Louceng,
        bianhao:that.data.seatlist[index].Bianhao,
        number:that.data.seatlist[index].Number,
        begintime:"08:00",
        endtime:"18:00",
        date:util.formatTime(new Date()),
        status:1,
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
          wx.showToast({
            title: "success",
            icon:"success",
            duration: 2000
          })
         }else{
          wx.showToast({
            title: "选座失败",
            icon:"error",
            duration: 2000
          })
         }
       },
       fail(res){
         console.log(res.errMsg);
       }
    })
  },
  onLoad: function (options) {
    var that = this;
    var fenguan =options.fenguan;
    console.log(fenguan)
    that.setData({
      fenguan: fenguan
    })
    that.requinfo();
  },
  requinfo:function(){
    var that = this;
    wx.request({
      url: app.globalData.URL +  '/user/showSeatInfo',
      method:"POST",
      data:{
        fenguan:that.data.fenguan
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
            seatlist:res.data.body,
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
    var that = this;

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