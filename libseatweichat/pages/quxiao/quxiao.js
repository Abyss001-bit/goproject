// pages/history/history.js
var app = getApp();//写在页面顶部page()外
Page({

  /**
   * 页面的初始数据
   */
  data: {
    historylist:[],
  },
  /**
   * 生命周期函数--监听页面加载
   */

  onLoad: function (options) {
    var that = this;
    setInterval(function () {
      that.requinfo();
    }, 1000)    //代表1秒钟发送一次请求
  },
  requinfo:function(){
    var that = this;
    wx.request({
      url: app.globalData.URL +  '/user/ShowHistory',
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
            historylist:res.data.body,
          });
        }
      },
      fail(res){
        console.log(res.errMsg);
      }
    })
  },
  deleteyuyue:function(e){
    var that = this;
    var index = e.currentTarget.dataset.index;
    console.log(index)
    console.log(that.data.historylist[index].Fenguan)
    wx.showModal({
      title: '提示',
      content: '确认取消预约？',
      success (res) {
        if (res.confirm) {
          console.log('用户点击确定')
          wx.request({
                url: app.globalData.URL +  '/user/QuxiaoYuyue',
                method:"POST",
                data:{
                  fenguan:that.data.historylist[index].Fenguan,
                  louceng:that.data.historylist[index].Louceng,
                  bianhao:that.data.historylist[index].Bianhao,
                  number:that.data.historylist[index].Number,
                  status:that.data.historylist[index].Status,
                  date:that.data.historylist[index].Date,
                  begintime:that.data.historylist[index].Begintime,
                  endtime:that.data.historylist[index].Endtime,
                },
                dataType:"json",
                header:{
                  "content-type": "application/x-www-form-urlencoded",
                  "content-type": "application/json",
                  "Cookie":wx.getStorageSync('cookieKey')
                },
                success(res){
                  console.log(res.data);
                  wx.showToast({
                    title: '取消成功',
                    icon: 'success',
                    duration: 1500,
                    })
                },
                fail(res){
                  console.log(res.errMsg);
                }
              });
        } else if (res.cancel) {
          console.log('用户点击取消')
        }
      }
    })
   
  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady: function () {
    var that = this;

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