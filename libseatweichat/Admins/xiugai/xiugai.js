// Admins/daoru/daoru.js
var app = getApp();//写在页面顶部page()外
Page({

  /**
   * 页面的初始数据
   */
  data: {
    istrue:false,
    id :"",

    oldfenguan:"",
    oldlouceng:"",
    oldbianhao:"",
    oldnumber:"",
    oldnumberdigit:0,
    oldseatimage:"", //暂时不支持

    fenguan:"",
    louceng:"",
    bianhao:"",
    number:"",
    numberdigit:0,
    status:"",
    statusdigit:-1,
    seatimage:"", //暂时不支持 
  },
  oldfenguanInput:function(e){
    var that = this;
    that.setData({
      oldfenguan:e.detail.value
    })
   },
   oldloucengInput:function(e){
     var that = this;
     that.setData({
      oldlouceng:e.detail.value
     })
    },
    oldbianhaoInput:function(e){
     var that = this;
     that.setData({
      oldbianhao:e.detail.value
     })
    },
    oldnumberInput:function(e){
     var that = this;
     var v1 = e.detail.value;
       var v2 = RegExp('^\-', 'g').exec(e.detail.value)
       var g = 1;
       if (v2) {
         g = -1;
       }
       var v3 = parseFloat(v1.replace(/\D/g, '')) * g
       that.setData({ 
         oldnumber: v3,
         oldnumberdigit:v3,
        });
    },
    oldsubmit:function(e){
      var that = this;
      if (that.data.oldfenguan == "") {
       wx.showModal({
         title: '提示',
         content: '馆号信息不能为空',
      });
     }else if (that.data.oldlouceng == "") {
       wx.showModal({
         title: '提示',
         content: '楼层信息不能为空',
      });
     }else if (that.data.oldbianhao == "") {
       wx.showModal({
         title: '提示',
         content: '室号信息不能为空',
      });
     }else if (that.data.oldnumber == "") {
       wx.showModal({
         title: '提示',
         content: '编号信息不能为空',
      });
     }else{
       wx.request({
         // 请求地址
         url: app.globalData.URL + '/admin/seatinfoExist',
         //请求方法
         method:"POST",
         //请求参数
         data:{
           fenguan:that.data.oldfenguan,
           louceng:that.data.oldlouceng,
           bianhao:that.data.oldbianhao,
           number:that.data.oldnumberdigit,
           seatimage:that.data.oldseatimage
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
               istrue:true,
               id:res.data.body[0].Id,
             });
             console.log(res.data.body[0].Id)
           }else {
             wx.showToast({
               title: '该座位不存在',
               icon:'error',
               duration:2000,
             });
             that.setData({
              istrue:false,
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
//-------------------------------------------------------
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
      numberdigit:v3,
    });
 },
 statusInput:function(e){
  var that = this;
  var v1 = e.detail.value;
    var v2 = RegExp('^\-', 'g').exec(e.detail.value)
    var g = 1;
    if (v2) {
      g = -1;
    }
    var v3 = parseFloat(v1.replace(/\D/g, '')) * g
    that.setData({ 
      status: v3,
      statusdigit:v3,
    });
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
  }else if (that.data.status == "") {
    wx.showModal({
      title: '提示',
      content: '状态信息不能为空',
   });
  }else if (that.data.istrue == true){
    wx.request({
      // 请求地址
      url: app.globalData.URL +  '/admin/changeSeatsInfo',
      //请求方法
      method:"POST",
      //请求参数
      data:{
        id:that.data.id,
        fenguan:that.data.fenguan,
        louceng:that.data.louceng,
        bianhao:that.data.bianhao,
        number:that.data.numberdigit,
        status:that.data.statusdigit,
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
            title: '修改座位成功',
            icon:'success',
            duration:2000,
          });
        }else {
          wx.showToast({
            title: '修改座位失败',
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
   }else{
    wx.showModal({
      title: '提示',
      content: '该座位不存在，无法修改',
   });
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