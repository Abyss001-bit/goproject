// Admins/code/code.js
// 将 dist 目录下，weapp.qrcode.esm.js 复制到项目utils目录下
// 直接引入 js 文件
import drawQrcode from '../../utils/weapp.qrcode.esm.js'
let QR = require('../../utils/qrcode')
var app = getApp();//写在页面顶部page()外
Page({

  /**
   * 页面的初始数据
   */
  data: {
    qrcodeWidth: 0,
    istrue:true,

    // fenguan:"",
    // louceng:"",
    // bianhao:"",
    // number:"",
    // numberdigit:0,
    // seatimage:"", //暂时不支持 
    fenguan:"图书馆一",
    louceng:"一楼",
    bianhao:"a1-1室",
    number:"1",
    numberdigit:1,
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
    numberInput:function(e){
     var that = this;
     var v1 = e.detail.value;
       var v2 = RegExp('^\-', 'g').exec(e.detail.value)
       var g = 1;
       if (v2) {
         g = -1;
       }
       var v3 = parseFloat(v1.replace(/\D/g, '')) * g
       that.setData({ 
         number: v3,
         numberdigit:v3
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
     }else if (that.data.number == "") {
       wx.showModal({
         title: '提示',
         content: '编号信息不能为空',
      });
     }else{
       wx.request({
         // 请求地址
         url: app.globalData.URL +  '/admin/seatinfoExist',
         //请求方法
         method:"POST",
         //请求参数
         data:{
           fenguan:that.data.fenguan,
           louceng:that.data.louceng,
           bianhao:that.data.bianhao,
           number:that.data.numberdigit,
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
           if(res.data.msg == "ok" && res.data.body != null){
            wx.showToast({
              title: '该座位存在',
              icon:'success',
              duration:2000,
            });
            that.setData({
              istrue:false
            });
           }else {
            wx.showToast({
              title: '该座位不存在',
              icon:'error',
              duration:2000,
            });
           }
         },
         //请求失败回调
         fail(res){
           console.log(res.errMsg);
         }
       })
      }
   
    },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    var that = this;
    drawQrcode({
      width: 200, // 必须，二维码宽度，与canvas的width保持一致
      height: 200, // 必须，二维码高度，与canvas的height保持一致2
      canvasId: 'myQrcode',
      background:'#ffffff', //	非必须，二维码背景颜色，默认值白色
      foreground: '#2bb15e', // 非必须，二维码前景色，默认值黑色 	'#000000'
      // ctx: wx.createCanvasContext('myQrcode'), // 非必须，绘图上下文，可通过 wx.createCanvasContext('canvasId') 获取，v1.0.0+版本支持
      text: "图书馆一_一楼_a1-1室_1",  // 必须，二维码内容
      // v1.0.0+版本支持在二维码上绘制图片
      image: {
        // imageResource: '../../images/icon.png', // 指定二维码小图标
        dx: 70,
        dy: 70,
        dWidth: 60,
        dHeight: 60
      }
    });
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