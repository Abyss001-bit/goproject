package controllers

import "github.com/astaxie/beego"

// ------------------------------------------- user --------------------------------------------//

type SendCode struct {
	beego.Controller
}

type RegisterController struct {
	beego.Controller
}
type RegisterAdminController struct {
	beego.Controller
}

type FindBackPasswordController struct {
	beego.Controller
}

type LoginBypasswordController struct {
	beego.Controller
}

type ShowUserInfoController struct {
	beego.Controller
}
type ShowAdminUserInfoController struct {
	beego.Controller
}

type ChangeUserInfoController struct {
	beego.Controller
}

type UserLoginOutController struct {
	beego.Controller
}
type UserimageuploadController struct {
	beego.Controller
}

// type UserDeleteController struct {
// 	beego.Controller
// }

type SignInSeatController struct {
	beego.Controller
}

type ChoiceSeatController struct {
	beego.Controller
}

// --------------------------------------------------------show seat ----------------------------------------//
type InsertSeatController struct {
	beego.Controller
}

type ShowSeatInfoController struct {
	beego.Controller
}
type ShowFenguanInfoController struct {
	beego.Controller
}
type ShowFenguanLoucengInfoController struct {
	beego.Controller
}
type ShowFenguanBianhaoInfoController struct {
	beego.Controller
}
type ShowFenguanNumberInfoController struct {
	beego.Controller
}

type InsertChoiceSeatInfoController struct {
	beego.Controller
}

type ShowHistoryInfoController struct {
	beego.Controller
}
type QuxiaoYuyueInfoController struct {
	beego.Controller
}

type DeleteOneSeatController struct {
	beego.Controller
}

type ChangeSeatsInfoController struct {
	beego.Controller
}

type GetHistoryController struct {
	beego.Controller
}

type ShowFenguanNameInfoController struct {
	beego.Controller
}

type TuijianController struct {
	beego.Controller
}

type SeatinfoExistController struct {
	beego.Controller
}
type DeleteAllSeatController struct {
	beego.Controller
}

// type UserDeleteController struct {
// 	beego.Controller
// }

// type RecomSeatController struct {
// 	beego.Controller
// }

// type SetRebacksController struct {
// 	beego.Controller
// }
// type GetRebacksController struct {
// 	beego.Controller
// }

// type DeletesSeatsController struct {
// 	beego.Controller
// }
