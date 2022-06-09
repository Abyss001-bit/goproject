package routers

import (
	"libseat/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//code
	beego.Router("/user/sendCode", &controllers.SendCode{})
	// user
	beego.Router("/user/registerBypassword", &controllers.RegisterController{})
	// admin
	beego.Router("/admin/registerBypassword", &controllers.RegisterAdminController{})
	//普通用户 找回密码、登录、显示个人信息、显示管理员信息
	beego.Router("/user/findbackpassword", &controllers.FindBackPasswordController{})
	beego.Router("/user/loginBypassword", &controllers.LoginBypasswordController{})
	beego.Router("/user/showUserInfo", &controllers.ShowUserInfoController{})
	beego.Router("/user/showAdminUserInfo", &controllers.ShowAdminUserInfoController{})
	// beego.Router("/userloginByweichat", &controllers.UserLoginController{})
	// beego.Router("/registerByweichat", &controllers.NewUserregisterController{})

	//修改个人信息
	beego.Router("/user/changeUserInfo", &controllers.ChangeUserInfoController{})
	//登出
	beego.Router("/user/loginOut", &controllers.UserLoginOutController{})

	// Status    int    `json:status`
	// Begintime string `json:begintime`
	// Endtime   string `json:endtime`
	// Limit     int    `json:limite`  默认20条
	beego.Router("/user/showSeatInfo", &controllers.ShowSeatInfoController{})

	//输入none 获取 fenguan名字 + null/full
	beego.Router("/user/ShowFenguanInfo", &controllers.ShowFenguanInfoController{})
	//预约
	//输入 none 获取fenguan的name
	beego.Router("/user/ShowFenguanName", &controllers.ShowFenguanNameInfoController{})
	//输入fenguan获取楼层
	beego.Router("/user/ShowFenguanLouceng", &controllers.ShowFenguanLoucengInfoController{})
	//输入fenguan + louceng 获取编号
	beego.Router("/user/ShowFenguanBianhao", &controllers.ShowFenguanBianhaoInfoController{})
	//输入fenguan + louceng + bianhao 加 status = 0 获取number
	beego.Router("/user/ShowFenguanNumber", &controllers.ShowFenguanNumberInfoController{})
	// // //user选座-自主选座
	//输入fenguan + louceng + bianhao 加 status = 0 + number + begintime + endtime 修改seat_info数据库信息 使得status = 1,并且 在user_hisinfo中插入数据
	beego.Router("/user/InsertChoiceSeat", &controllers.InsertChoiceSeatInfoController{})
	// 历史
	// 显示历史数据，按照id desc 排序
	beego.Router("/user/ShowHistory", &controllers.ShowHistoryInfoController{})
	//取消
	beego.Router("/user/QuxiaoYuyue", &controllers.QuxiaoYuyueInfoController{})

	// beego.Router("/admin/deleteUser", &controllers.UserDeleteController{})
	// --------------------------------------- seat ---------------------------------------- //
	// seat
	beego.Router("/admin/insertSeat", &controllers.InsertSeatController{})
	beego.Router("/admin/seatinfoExist", &controllers.SeatinfoExistController{})
	beego.Router("/admin/showSeatInfo", &controllers.ShowSeatInfoController{})
	beego.Router("/admin/deleteOneSeat", &controllers.DeleteOneSeatController{})
	beego.Router("/admin/deleteAllSeat", &controllers.DeleteAllSeatController{})
	beego.Router("/admin/changeSeatsInfo", &controllers.ChangeSeatsInfoController{})

	// // //user选座-推荐算法
	beego.Router("/user/Tuijian", &controllers.TuijianController{})
	// // //user扫码
	beego.Router("/user/signInSeat", &controllers.SignInSeatController{})
	// // //user反馈
	// beego.Router("/user/setRebacks", &controllers.SetRebacksController{})
	// beego.Router("/user/getRebacks", &controllers.GetRebacksController{})
	// 用户反馈信息
	beego.Router("/user/imageupload", &controllers.UserimageuploadController{}) //????????
}
