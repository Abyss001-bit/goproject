// Admins/daoru/daoru.js
var app = getApp();//写在页面顶部page()外
Page({

  /**
   * 页面的初始数据
   */
  data: {

  },
  tooneseat:function(){
    wx.navigateTo({
      url: '../onedeleteseat/onedeleteseat',
    })
  },
  tomoreseat:function(e){
    wx.navigateTo({
      url: '../moredeleteseat/moredeleteseat',
    })
  },
  toall:function(e){
    wx.showModal({
      title: '提示',
      content: '确定删除所有座位？',
      success (res) {
        if (res.confirm) {
          console.log("User clicks OK.");
          wx.request({
            // 请求地址
            url: app.globalData.URL +  '/admin/deleteAllSeat',
            //请求方法
            method:"POST",
            //请求参数
            data:{
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
                  title: '座位删除成功',
                  icon:'success',
                  duration:3000,
                });
              }else {
                wx.showToast({
                  title: '座位删除失败',
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
        } else if (res.cancel) {
          console.log("User clicks to cancel.");
        }
      }
   });

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