// Admins/daoru/daoru.js
var app = getApp();//写在页面顶部page()外
Page({

  /**
   * 页面的初始数据
   */
  data: {
    fenguan:"",
    louceng:"",
    bianhao:"",
    numberfirst:"",
    numberfirstdigit:0,
    numberlast:"",
    numberlastdigit:0,
    seatimage:"", //暂时不支持 
  },
fenguanInput:function(e){
 var that = this;
 that.setData({
  fenguan:e.detail.value
 })
},
loucengInput:function(e){
  var that = this;
  that.setData({
    louceng:e.detail.value
  })
 },
 bianhaoInput:function(e){
  var that = this;
  that.setData({
    bianhao:e.detail.value
  })
 },
 numberfirstInput:function(e){
  var that = this;
  var v1 = e.detail.value;
    var v2 = RegExp('^\-', 'g').exec(e.detail.value)
    var g = 1;
    if (v2) {
      g = -1;
    }
    var v3 = parseFloat(v1.replace(/\D/g, '')) * g
    that.setData({ 
      numberfirst: v3,
      numberfirstdigit:v3,
     })
 },
 numberlastInput:function(e){
  var that = this;
  var v1 = e.detail.value;
    var v2 = RegExp('^\-', 'g').exec(e.detail.value)
    var g = 1;
    if (v2) {
      g = -1;
    }
    var v3 = parseFloat(v1.replace(/\D/g, '')) * g
    that.setData({ 
      numberlast: v3,
      numberlastdigit:v3,
     })
 },
 submit:function(e){
   var that = this;
   if (that.data.fenguan == "") {
    wx.showModal({
      title: '提示',
      content: '馆号信息不能为空',
   });
  }else if (that.data.louceng == "") {
    wx.showModal({
      title: '提示',
      content: '楼层信息不能为空',
   });
  }else if (that.data.bianhao == "") {
    wx.showModal({
      title: '提示',
      content: '室号信息不能为空',
   });
  }else if (that.data.numberfirst == "") {
    wx.showModal({
      title: '提示',
      content: '编号开始信息不能为空',
   });
  }else if (that.data.numberlast == "") {
    wx.showModal({
      title: '提示',
      content: '编号终止信息不能为空',
   });
  }else{
    console.log(that.data.numberfirstdigit)
    console.log(that.data.numberlastdigit)
      for (var i = that.data.numberfirstdigit; i <= that.data.numberlastdigit; ) {
        wx.setStorageSync('I', i)
        console.log(i)
        wx.request({
          // 请求地址
          url: app.globalData.URL +  '/admin/deleteOneSeat',
          //请求方法
          method:"POST",
          //请求参数
          data:{
            fenguan:that.data.fenguan,
            louceng:that.data.louceng,
            bianhao:that.data.bianhao,
            number:wx.getStorageSync('I'),
            seatimage:that.data.seatimage
          },
          //请求头
          header:{
            "content-type": "application/x-www-form-urlencoded",
            "content-type": "application/json"
          },
          //请求成功回调
          success(res){
            console.log(res.data);
            if(res.data.msg == "ok"){
              wx.showToast({
                title: '删除座位成功',
                icon:'success',
                duration:3000,
              });
              that.setData({
                // fenguan:"",
                // louceng:"",
                // bianhao:"",
                // // numbercount:"",
                // seatimage:"", //暂时不支持 
              });
            }else {
              wx.showToast({
                title: '删除座位失败',
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
        i++;
      }
   }

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