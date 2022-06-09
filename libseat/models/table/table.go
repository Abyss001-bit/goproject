package table

// 用户全部信息

// id,用户名，电话，密码，admin，信誉分,头像
type UserInfo struct {
	Id          string `json:id`
	Name        string `json:name`
	Password    string `json:password`
	Phonenumber string `json:phonenumber`
	Admin       bool   `json:admin`
	Integral    int64  `json:integral`
	Weichat     string `json:weichat`
	Userimage   string `json:userimage`
}

type GetCodeMsg struct {
	Phonenumber string `json:phonenumber`
	Code        string `json:code`
}

// 注册 普通用户
type RegistNewUser struct {
	Name        string `json:name`
	Password    string `json:password`
	Phonenumber string `json:phonenumber`
	Code        string `json:code`
}

//找回密码
type FindBackPassword struct {
	Phonenumber string `json:phonenumber`
	Password    string `json:password`
	Code        string `json:code`
}

//用户名密码登录
type UserLogin struct {
	Name     string `json:name`
	Password string `json:password`
}

//用户登出
type UserLogout struct {
}

/////////////////////////////////////////-------seat--------////////////////////////////////////////
const (
	SeatNo_appointment   = iota //0
	SeatHave_appointment        //1
	SeatFull
	SeatFullButNil //(入座后人走了被二次扫码)
)

type SeatInfo struct {
	Id        string `json:id`
	Fenguan   string `json:fenguan`
	Louceng   string `json:louceng`
	Bianhao   string `json:bianhao`
	Number    int    `json:number`
	Status    int    `json:status`
	Begintime string `json:begintime`
	Endtime   string `json:endtime`
	SeatImage string `json:seatimage`
	Date      string `json:date`
}

type InsertSeatInfo struct {
	Fenguan   string `json:fenguan`
	Louceng   string `json:louceng`
	Bianhao   string `json:bianhao`
	Number    int    `json:number`
	SeatImage string `json:seatimage`
}

type DelSeatInfo struct {
	Fenguan   string `json:fenguan`
	Louceng   string `json:louceng`
	Bianhao   string `json:bianhao`
	Number    int    `json:number`
	SeatImage string `json:seatimage`
}

type SeatReqInfo struct {
	Status    int    `json:status`
	Begintime string `json:begintime`
	Endtime   string `json:endtime`
	Limit     int    `json:limite`
}

type FenguanInfo struct {
	Fenname       string `json:fenname`
	Fenallseatnum int    `json:fenallseatnum`
	Kongseatnum   int    `json:kongseatnum`
}

type FenguanNameInfo struct {
	Fenname string `json:fenname`
}

type TempFenguanInfo struct {
	Fenname       string `json:fenname`
	Fenallseatnum int    `json:fenallseatnum`
}

type TemKong struct {
	Kongseatnum int `json:kongseatnum`
}

//学生选座,通过传入的 fenguan、louceng、bianhao、number确定座位后,修改座位的status信息,并将该学生通过phonenumber 绑定学生信息 和 座位信息 成为一张新表,作为学生行为表

type SginCode struct {
	Fenguan   string `json:fenguan`
	Louceng   string `json:louceng`
	Bianhao   string `json:bianhao`
	Number    int    `json:number`
	Begintime string `json:begintime`
	Endtime   string `json:endtime`
	Signbegin bool   `json:sginbegin`
	Signend   bool   `json:sginend`
	Status    int    `json:status`
}

type UserHisinfo struct {
	Id          string `json:id`
	Phonenumber string `json:phonenumber`
	Fenguan     string `json:fenguan`
	Louceng     string `json:louceng`
	Bianhao     string `json:bianhao`
	Number      int    `json:number`
	Begintime   string `json:begintime`
	Endtime     string `json:endtime`
	Date        string `json:date`
	Status      int    `json:status`
}

type UserReBack struct {
	Id           string `json:id`
	Userphone    string `json:userphone`
	Usermsgtext  string `json:usermsgtext`
	Umsgimage    string `json:umsgimage`
	Adminphone   string `json:adminphone`
	Adminmsgtext string `json:adminmsgtext`
	Amsgimage    string `json:amsgimage`
	Status       bool   `json:status`
}

type Count struct {
	Count int `json:count`
}

type Feng struct {
	Fenguan string `json:fenguan`
}

type StatusStrings struct {
	Seatstrinngs string `json:seatstrinngs`
	Status       int    `json:status`
}
