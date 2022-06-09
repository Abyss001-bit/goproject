// pages/findbackpassword/findbackpassword.js
var app = getApp();//写在页面顶部page()外
var timeInterval = null //倒计时函数

Page({

  /**
   * 页面的初始数据
   */
  data: {
    phonenumber:"",
    btnText: '获取验证码', //倒计时 
    currentTime: 60,//限制60s
    isClick: false,//获取验证码按钮，默认允许点击
    codeColor:"rgb(233,214,199)",
  },
phoneInput: function (e) {
  this.setData({
    phonenumber: e.detail.value
  })
},
sentcode:function(e){
  var that = this;
  wx.request({
    // 请求地址
    url: app.globalData.URL +  '/user/sendCode',
    //请求方法
    method:"POST",
    //请求参数
    data:{
      phonenumber:this.data.phonenumber
    },
    //请求头
    header:{
      "content-type": "application/x-www-form-urlencoded",
      "content-type": "application/json" // 默认
    },
    //请求成功回调
    success(res){
      that.sendMsg();
      console.log(res.data);
    },
    //请求失败回调
    fail(res){
      console.log(res.errMsg);
    }
  })
},
sendMsg() {
  //第二步获取验证码
    let that = this;
  //设置button是否可点击，倒计时期间不可点击
    that.setData({
      isClick: true,
      codeColor:"#0271c1",
    })
    var currentTime = that.data.currentTime;
  //开始倒计时
    timeInterval = setInterval(function () {
      currentTime--;//倒计时
      that.setData({
        btnText: currentTime + '秒后获取'
      })
      if (currentTime <= 0) {
        clearInterval(timeInterval )//重置倒计时
        that.setData({
          btnText: '获取验证码',
          currentTime: 60,
          isClick: false
        })
      }
    }, 1000);
  },
  submit:function(e){
    wx.request({
      // 请求地址
      url: app.globalData.URL +  '/user/findbackpassword',
      //请求方法
      method:"POST",
      //请求参数
      data:{
        phonenumber:e.detail.value.phonenumber,
        password:e.detail.value.password,
        code:e.detail.value.code
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
          
          wx.showToast({
            title: '设置新成功',
            icon:'success',
            duration:2000
          });
            var pages = getCurrentPages(); // 当前页面
            var beforePage = pages[pages.length - 2]; // 前一个页面
            setTimeout(function(){
              wx.navigateBack({
                success: function() {
                    beforePage.onLoad(); // 执行前一个页面的onLoad方法
                }
                })
            },1000)
        }else {
          wx.showToast({
            title: '设置新密码失败，手机号或验证码错误',
            icon:'error',
            duration:2000
          });
        }
      },
      //请求失败回调
      fail(res){
        console.log(res.errMsg);
      }
    })
  },
  /**
   * 生命周期函数--监听页面加载
   */
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