// pages/main/main.js
Page({

  /**
   * 页面的初始数据
   */
  data: {
    seatlist:[],
    swiperCurrent: 0,
    indicatorDots: true,
    autoplay: true,
    interval: 3000,
    duration: 800,
    circular:true,
    imgUrls: [
      // 'https://gitee.com/itanrong/picstore/raw/master/libseatimages/main1.jpg',
      // 'https://gitee.com/itanrong/picstore/raw/master/libseatimages/main2.jpg',
      // 'https://gitee.com/itanrong/picstore/raw/master/libseatimages/main3.jpg'
      'https://github.com/Abyss001-bit/picture/blob/master/static/main1.jpg?raw=true',
      'https://github.com/Abyss001-bit/picture/blob/master/static/main2.jpg?raw=true',
      'https://github.com/Abyss001-bit/picture/blob/master/static/main3.jpg?raw=true'
    ],
    links:[
      '../user/user',
      '../user/user',
      '../user/user'
    ]
  },
  //轮播图的切换事件
  swiperChange: function (e) {
    this.setData({
      swiperCurrent: e.detail.current
    })
  },
  //点击指示点切换
  chuangEvent: function (e) {
    this.setData({
      swiperCurrent: e.currentTarget.id
    })
  },
  //点击图片触发事件
  swipclick: function (e) {
    console.log(this.data.swiperCurrent);
    wx.switchTab({
      url: this.data.links[this.data.swiperCurrent]
    })
  },

  
  
  
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    var that = this;
    setInterval(function () {
      // that.requinfo();
    }, 1000)    //代表1秒钟发送一次请求
  },
  daoru:function(){
    wx.navigateTo({
      url: '../daoru/daoru',
    })
  },
  xiugai:function(){
    wx.navigateTo({
      url: '../xiugai/xiugai',
    })
  },
  shanchuseat:function(){
    wx.navigateTo({
      url: '../shanchuseat/shanchuseat',
    })
  },
  code:function(){
    wx.navigateTo({
      url: '../code/code',
    })
  },
  huifu:function(){
    wx.navigateTo({
      url: '../huifu/huifu',
    })
  },
  heimingdan:function(){
    wx.navigateTo({
      url: '../heimingdan/heimingdan',
    })
  },
  xiugaiuser:function(){
    wx.navigateTo({
      url: '../huifu/huifu',
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