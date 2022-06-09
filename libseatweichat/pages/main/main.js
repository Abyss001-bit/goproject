// pages/main/main.js
var app = getApp();//写在页面顶部page()外
Page({

  /**
   * 页面的初始数据
   */
  data: {
    num:1,
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
  gotoreload:function(e){
    this.onLoad()
  },
  gotore:function(e){
    wx.navigateTo({
      url: '../recom/recom',
    })
  },
  gotohis:function(){
    wx.navigateTo({
      url: '../history/history',
    })
  },
  gotomine:function(){
    wx.navigateTo({
      url: '../mine/mine',
    })
  },
  qiandao:function(){
    var that =this;
    wx.scanCode({
      // onlyFromCamera: true,
      success(res) {
        console.log(res.result);
        var n = that.data.num + 1
        that.setData({
          num:n
        });
        wx.request({
          url: app.globalData.URL +  '/user/signInSeat',
          method:"POST",
          data:{
            seatstrinngs:res.result,
            status:that.data.num,
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
              wx.showToast({
                title: '扫码成功',
                icon:'success',
                duration:2000
              });
            }
          },
          fail(res){
            console.log(res.errMsg);
          }
        })
        if (n == 2) {
          that.setData({
            num:0
          });
        }
      },
      fail(res){
        console.log("----------------" + res.result);
        wx.showToast({
          title: '扫码失败',
          icon:'error',
          duration:2000
        });
      }
    })
  },
  tuijian:function(){
// 判断用户受否满足推荐要求：选座位超过20次以上
// 推荐方法：空位最多的教室选第一个
wx.navigateTo({
  url: '../tuijian/tuijian',
})
  },
  quxiao:function(){
    // 0 已完成
    // 1 待完成
    // 2 完成中 扫码一次 
    // 3 爽约未完成 
    // 即选中数据库中状态为 1 的进行数据删除
    wx.navigateTo({
      url: '../quxiao/quxiao',
    })
  },
  tomoremsg:function(e){
    // 查看该馆信息
    var that = this;
    let index = e.currentTarget.dataset.index;
    //将对象转为string
    console.log(that.data.seatlist[index].Fenname)
    var fenguan = that.data.seatlist[index].Fenname;
    // console.log("-----------" + fenguan)
    wx.navigateTo({
      url: '../moremsg/moremsg?fenguan=' + fenguan,
    })
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    var that = this;
    setInterval(function () {
      that.requinfo();
    }, 3000)    //代表1秒钟发送一次请求
  },
  requinfo:function(){
    var that = this;
    wx.request({
      url: app.globalData.URL +  '/user/ShowFenguanInfo',
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