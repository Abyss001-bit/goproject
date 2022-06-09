// pages/changeinfo/cahngeinfo.js
var app = getApp();//写在页面顶部page()外
Page({

  /**
   * 页面的初始数据
   */
  data: {
    name:"",
    password:"",
    weichat:"",
    userimage:"",

    nametemp:"",
    passwordtemp:"",
    weichattemp:"",
    userimagetemp:"../../static/添加图片.png",
    inputValue:null,
  },
  name: function (e) {
    var that = this;
    that.setData({
      nametemp: e.detail.value
    })
  },
  password: function (e) {
    var that = this;
    that.setData({
      passwordtemp: e.detail.value
    })
  },
  weichat: function (e) {
    var that = this;
    that.setData({
      weichattemp: e.detail.value
    })
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
        userimagetemp: tempFilePaths
      });
      if(that.data.userimagetemp!=undefined){
          wx.getFileSystemManager().readFile({
            filePath:res.tempFilePaths[0],
            encoding:'base64',
            success(res){
              that.setData({
                filebase64:res.data
              });
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
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    var that = this;
    setInterval(function () {
      // that.requinfo();
    }, 1000)    //代表1秒钟发送一次请求
  },

  requinfo:function(){
    var that = this;
    wx.request({
      url: app.globalData.URL +  '/user/showUserInfo',
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
            name : res.data.body[0].Name,
            password : res.data.body[0].Password,
            weichat:res.data.body[0].Weichat,
            // userimage:res.data.body[0].Userimage,
            userimage:"../../static/头像.jpg",
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
    wx.request({
      url:   app.globalData.URL + '/user/changeUserInfo',
      method:"POST",
      data:{
        name:that.data.nametemp,
        password:that.data.passwordtemp,
        userimage:that.data.userimagetemp,
        weichat:that.data.weichattemp,
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
            name : res.data.body[0].Name,
            password : res.data.body[0].Password,
            weichat:res.data.body[0].Weichat,
            userimage:res.data.body[0].Userimage,
            // 清空input框 
            inputValue:''
          });
          wx.showToast({
            title:'修改成功',
            icon:'success',
            duration:2000
          });
        }else {
          wx.showToast({
            title:'修改失败',
            icon:'error',
            duration:2000
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