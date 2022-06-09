// pages/problemback/problemback.js
var app = getApp();//写在页面顶部page()外
Page({

  /**
   * 页面的初始数据
   */
  data: {
    tempFilePath:["../../static/添加图片.png"],
    filebase64:"",
    rebacktext:"",
  },

  textareaInput2:function(e){
    var that = this;
    var temp = e.detail.value;
    that.setData({
      rebacktext:temp
    });
  },

  choice: function () {
    var that = this;
    wx.chooseImage({
     count: 1, // 默认9
     sizeType: ['original', 'compressed'], // 可以指定是原图还是压缩图，默认二者都有
     sourceType: ['album', 'camera'], // 可以指定来源是相册还是相机，默认二者都有
     success: function (res) {
      // 返回选定照片的本地文件路径列表，tempFilePath可以作为img标签的src属性显示图片
      var tempFilePaths = res.tempFilePaths;
      console.log(tempFilePaths)
      that.setData({
        tempFilePath: tempFilePaths
      });
      if(that.data.tempFilePath!=undefined){
        // console.log("wx.uploadFile");
        // console.log(wx.getStorageSync('cookieKey'));
        // console.log(that.data.tempFilePath[0]);
          wx.getFileSystemManager().readFile({
            filePath:res.tempFilePaths[0],
            encoding:'base64',
            success(res){
              // console.log('data:image/png;base64,'+res.data)
              that.setData({
                // filebase64:'data:image/png;base64,'+res.data
                filebase64:res.data
              });
              // console.log(that.data.filebase64);
            }
          });
        }else{
          wx.showToast({
            title: '未选择照片！',
            duration:2000
          })
        }
     }
    });
   },
submit:function(e){
  var that = this;
// 发送的数据为base64 + msg文字
  wx.request({
    // 请求地址
    url: app.globalData.URL +  '/user/imageupload',
    //请求方法
    method:"POST",
    //请求参数
    data:{
      umsgimage:that.data.filebase64,
      usermsgtext:that.data.rebacktext,
    },
    //请求头
    header:{
      "content-type": "application/x-www-form-urlencoded",
      "content-type": "application/json" ,
      "Cookie":wx.getStorageSync('cookieKey')
    },
    //请求成功回调
    success(res){
      console.log(res.data);
      console.log(that.data.rebacktext);
      if(res.data.msg == "ok"){
        wx.showToast({
          title: '成功',
          icon:'success',
          duration:2000
        });
      }else {
        wx.showToast({
          title: '失败',
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