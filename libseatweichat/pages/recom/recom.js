// pages/recom/recom.js
var app = getApp();//写在页面顶部page()外
let util = require('../../utils/util');
Page({
  /**
   * 页面的初始数据
   */
  data: {
    indexfenguan:0,
    indexlouceng:0,
    indexbianhao:0,
    indexnumber:0,
    datetime:util.formatTime(new Date()),  //util工具获取当前时间-当前只用到日期
    begintime:'08:00',
    endtime:'18:00',
    funguan:[],
    louceng:[],
    bianhao:[],
    number:[]
  },

  changefenguan: function(e) {
    var that = this;
    var selectIndex = e.detail.value;
    wx.setStorageSync('Fenguan',that.data.fenguan[e.detail.value])
    that.setData({
      "indexfenguan":selectIndex
    })
  },
  changelouceng: function(e) {
    var that = this;
    var selectIndex = e.detail.value;
    wx.setStorageSync('Louceng', that.data.louceng[e.detail.value])
    that.setData({
      "indexlouceng":selectIndex
    })
    console.log(e);
  },
  changebianhao: function(e) {
    var that=this;
    var selectIndex = e.detail.value;
    wx.setStorageSync('Bianhao', that.data.bianhao[e.detail.value])
    this.setData({
      "indexbianhao":selectIndex
    })
    console.log(e);
  },
  changenumber: function(e) {
    var that = this;
    var selectIndex = e.detail.value;
    wx.setStorageSync('Number', that.data.number[e.detail.value])
    this.setData({
      "indexnumber":selectIndex
    })
    console.log(e);
  },
  datetime:function(e){
    var that = this;
    wx.setStorageSync('Datetime', e.detail.value)
    that.setData({
     datetime:e.detail.value
    })
   },
   begintime:function(e){
    var that = this;
    wx.setStorageSync('Begintime', e.detail.value)
    that.setData({
     begintime:e.detail.value
    })
   },
   endtime:function(e){
    var that = this;
    wx.setStorageSync('Endtime', e.detail.value)
    that.setData({
      endtime:e.detail.value
    })
   },
   fenguan:function(){
    var that = this;
    wx.request({
      url: app.globalData.URL +'/user/ShowFenguanName',
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
           that.setData({
             fenguan:res.data.body,
           });
         }
       },
       fail(res){
         console.log(res.errMsg);
       }
    })
  },

  louceng:function(e){
    var that = this;
    wx.request({
      url: app.globalData.URL +'/user/ShowFenguanLouceng',
      method:"POST",
      data:{
        fenguan:wx.getStorageSync('Fenguan').Fenguan,
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
           that.setData({
             louceng:res.data.body,
           });
         }
       },
       fail(res){
         console.log(res.errMsg);
       }
    })
  },
  bianhao:function(e){
    var that = this;
    wx.request({
      url: app.globalData.URL +'/user/ShowFenguanBianhao',
      method:"POST",
      data:{
        fenguan:wx.getStorageSync('Fenguan').Fenguan,
        louceng:wx.getStorageSync('Louceng').Louceng,
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
           that.setData({
             bianhao:res.data.body,
           });
         }
       },
       fail(res){
         console.log(res.errMsg);
       }
    })
  },
  number:function(e){
    var that = this;
    var f = wx.getStorageSync('Fenguan').Fenguan;
    var l = wx.getStorageSync('Louceng').Louceng;
    var b = wx.getStorageSync('Bianhao').Bianhao;
    console.log(f)
    console.log(l)
    console.log(b)
    wx.request({
      url: app.globalData.URL +'/user/ShowFenguanNumber',
      method:"POST",
      data:{
        fenguan:wx.getStorageSync('Fenguan').Fenguan,
        louceng:wx.getStorageSync('Louceng').Louceng,
        bianhao:wx.getStorageSync('Bianhao').Bianhao,
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
           that.setData({
             number:res.data.body,
           });
         }
       },
       fail(res){
         console.log(res.errMsg);
       }
    })
  },
  submit:function(e){
    var that = this;
    var f = wx.getStorageSync('Fenguan').Fenguan;
    var l = wx.getStorageSync('Louceng').Louceng;
    var b = wx.getStorageSync('Bianhao').Bianhao;
    var n= wx.getStorageSync('Number').Number;
    var be=wx.getStorageSync('Begintime');
    var en=wx.getStorageSync('Endtime');
    var da=wx.getStorageSync('Datetime');
    console.log(f);
    console.log(l);
    console.log(b);
    console.log(n);
    console.log(be);
    console.log(en);
    console.log(da);
    wx.request({
      url: app.globalData.URL +  '/user/InsertChoiceSeat',
      method:"POST",
      data:{
        fenguan:wx.getStorageSync('Fenguan').Fenguan,
        louceng:wx.getStorageSync('Louceng').Louceng,
        bianhao:wx.getStorageSync('Bianhao').Bianhao,
        number:wx.getStorageSync('Number').Number,
        begintime:wx.getStorageSync('Begintime'),
        endtime:wx.getStorageSync('Endtime'),
        date:wx.getStorageSync('Datetime'),
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

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    var that = this;
    setInterval(function () {
      // that.fenguan();
    }, 3000)
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