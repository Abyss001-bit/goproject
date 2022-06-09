package controllers

import (
	"encoding/json"
	"fmt"
	"libseat/models/table"
	"libseat/service"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/logs"
)

//管理员界面才回调用，故不作登录判断
func (this *InsertSeatController) Post() {
	var seat table.InsertSeatInfo
	body := this.Ctx.Input.RequestBody //这是获取到request的body 的json二进制数据
	err := json.Unmarshal(body, &seat) //解析二进制json，把结果放进ob中
	if err != nil {
		logs.Error(err)
		this.Data["json"] = map[string]interface{}{
			"body": seat,
			"msg":  "解析请求失败",
		}
		this.ServeJSON()
	} else {
		// 1. 先查询该座位是否存在,存在则不插入,不存在则插入
		s, err := service.SerachSeatsInfo(service.Engine, &seat)
		if err != nil {
			logs.Error("SerachSeatsInfo 查询出错")
			this.Data["json"] = map[string]interface{}{
				"body": "",
				"msg":  "插入查询出错",
			}
			this.ServeJSON()
		} else if s != nil {
			logs.Error("该座位已经存在，插入失败")
			this.Data["json"] = map[string]interface{}{
				"body": s,
				"msg":  "该座位已经存在",
			}
			this.ServeJSON()
		} else {
			err = service.InsertSeats(service.Engine, &seat)
			if err != nil {
				logs.Error("插入失败")
				this.Data["json"] = map[string]interface{}{
					"body": "",
					"msg":  "插入失败",
				}
				this.ServeJSON()
			} else {
				logs.Info("插入seatinfo成功")
				this.Data["json"] = map[string]interface{}{
					"body": seat,
					"msg":  "ok",
				}
				this.ServeJSON()
			}
		}
	}

}

func (this *ShowSeatInfoController) Post() {
	// 显示用户信息只需要cookie
	// 1. 获取 cookie
	key := this.Ctx.Input.Context.GetCookie("Cookie")
	logs.Info("------------key:", key) //值是该用户的phonenumber
	// 2. 判断cookie是否存在 redis中,存在即在登录状态 ,不存在则 不在登录状态
	if val, err := service.GetRedisKey(service.Redisclient, key); err != nil {
		logs.Info("Cookie值不存在")
		this.Data["json"] = map[string]interface{}{
			"body": key,
			"msg":  "Cookie不存在,登录超时",
		}
		this.ServeJSON()
	} else if val != "ok" {
		// 3. 判断用户是否在黑名单中
		logs.Info("用户信誉积分不足")
		this.Data["json"] = map[string]interface{}{
			"body": key,
			"msg":  "用户信誉积分不足,无权查看",
		}
		this.ServeJSON()
	} else {
		var seat table.Feng
		body := this.Ctx.Input.RequestBody
		err := json.Unmarshal(body, &seat)
		if err != nil {
			logs.Error(err)
			this.Data["json"] = map[string]interface{}{
				"body": seat,
				"msg":  "解析请求失败",
			}
			this.ServeJSON()
		} else {
			s, err := service.ShowSeatInfo(service.Engine, &seat)
			if err != nil {
				logs.Error(err)
				this.Data["json"] = map[string]interface{}{
					"body": "",
					"msg":  "ShowSeatInfo err",
				}
				this.ServeJSON()
			} else if s == nil {
				logs.Info("ShowSeatInfo return nil")
				this.Data["json"] = map[string]interface{}{
					"body": s,
					"msg":  "ok",
				}
				this.ServeJSON()
			} else {
				logs.Info("ShowSeatInfo ok")
				this.Data["json"] = map[string]interface{}{
					"body": s,
					"msg":  "ok",
				}
				this.ServeJSON()
			}
		}
	}

}

func (this *ShowFenguanInfoController) Post() {

	// 显示用户信息只需要cookie
	// 1. 获取 cookie
	key := this.Ctx.Input.Context.GetCookie("Cookie")
	logs.Info("------------key:", key) //值是该用户的phonenumber
	// 2. 判断cookie是否存在 redis中,存在即在登录状态 ,不存在则 不在登录状态
	if val, err := service.GetRedisKey(service.Redisclient, key); err != nil {
		logs.Info("Cookie值不存在")
		this.Data["json"] = map[string]interface{}{
			"body": key,
			"msg":  "Cookie不存在,登录超时",
		}
		this.ServeJSON()
	} else if val != "ok" {
		// 3. 判断用户是否在黑名单中
		logs.Info("用户信誉积分不足")
		this.Data["json"] = map[string]interface{}{
			"body": key,
			"msg":  "用户信誉积分不足,无权查看",
		}
		this.ServeJSON()
	} else {
		FenguanInfo, err := service.ShowFenguanCountInfo(service.Engine)
		if err != nil {
			logs.Error("ShowFenguanInfo err")
			this.Data["json"] = map[string]interface{}{
				"body": FenguanInfo,
				"msg":  "ShowFenguanInfo err",
			}
			this.ServeJSON()
		} else if FenguanInfo == nil {
			logs.Info("ShowFenguanInfo FenguanInfo nil")
			this.Data["json"] = map[string]interface{}{
				"body": FenguanInfo,
				"msg":  "ShowFenguanInfo FenguanInfo nil",
			}
			this.ServeJSON()
		} else {
			logs.Info("ShowFenguanInfo FenguanInfo ok")
			this.Data["json"] = map[string]interface{}{
				"body": FenguanInfo,
				"msg":  "ok",
			}
			this.ServeJSON()
		}
	}

}

func (this *ShowFenguanNameInfoController) Post() {
	FenguanNameInfo, err := service.ShowFenguanNameInfo(service.Engine)
	if err != nil {
		logs.Error("ShowFenguanNameInfo err")
		this.Data["json"] = map[string]interface{}{
			"body": FenguanNameInfo,
			"msg":  "ShowFenguanNameInfo err",
		}
		this.ServeJSON()
	} else if FenguanNameInfo == nil {
		logs.Info("ShowFenguanNameInfo FenguanInfo nil")
		this.Data["json"] = map[string]interface{}{
			"body": FenguanNameInfo,
			"msg":  "ShowFenguanNameInfo FenguanInfo nil",
		}
		this.ServeJSON()
	} else {
		logs.Info("ShowFenguanNameInfo FenguanInfo ok")
		this.Data["json"] = map[string]interface{}{
			"body": FenguanNameInfo,
			"msg":  "ok",
		}
		this.ServeJSON()
	}
}

func (this *ShowFenguanLoucengInfoController) Post() {

	var seat table.SeatInfo
	body := this.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &seat)
	if err != nil {
		logs.Error(err)
		this.Data["json"] = map[string]interface{}{
			"body": seat,
			"msg":  "解析请求失败",
		}
		this.ServeJSON()
	} else {
		s, err := service.ShowFenguanLoucengInfo(service.Engine, &seat)
		if err != nil {
			logs.Error(err)
			this.Data["json"] = map[string]interface{}{
				"body": "",
				"msg":  "ShowSeatInfo err",
			}
			this.ServeJSON()
		} else if s == nil {
			logs.Info("ShowSeatInfo return nil")
			this.Data["json"] = map[string]interface{}{
				"body": s,
				"msg":  "ok",
			}
			this.ServeJSON()
		} else {
			logs.Info("ShowSeatInfo ok")
			this.Data["json"] = map[string]interface{}{
				"body": s,
				"msg":  "ok",
			}
			this.ServeJSON()
		}
	}

}
func (this *ShowFenguanBianhaoInfoController) Post() {

	var seat table.SeatInfo
	body := this.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &seat)
	if err != nil {
		logs.Error(err)
		this.Data["json"] = map[string]interface{}{
			"body": seat,
			"msg":  "解析请求失败",
		}
		this.ServeJSON()
	} else {
		s, err := service.ShowFenguanBianhaoInfo(service.Engine, &seat)
		if err != nil {
			logs.Error(err)
			this.Data["json"] = map[string]interface{}{
				"body": "",
				"msg":  "ShowSeatInfo err",
			}
			this.ServeJSON()
		} else if s == nil {
			logs.Info("ShowSeatInfo return nil")
			this.Data["json"] = map[string]interface{}{
				"body": s,
				"msg":  "ok",
			}
			this.ServeJSON()
		} else {
			logs.Info("ShowSeatInfo ok")
			this.Data["json"] = map[string]interface{}{
				"body": s,
				"msg":  "ok",
			}
			this.ServeJSON()
		}
	}

}
func (this *ShowFenguanNumberInfoController) Post() {
	var seat table.SeatInfo
	body := this.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &seat)
	if err != nil {
		logs.Error(err)
		this.Data["json"] = map[string]interface{}{
			"body": seat,
			"msg":  "解析请求失败",
		}
		this.ServeJSON()
	} else {
		s, err := service.ShowFenguanNumberInfo(service.Engine, &seat)
		if err != nil {
			logs.Error(err)
			this.Data["json"] = map[string]interface{}{
				"body": "",
				"msg":  "ShowSeatInfo err",
			}
			this.ServeJSON()
		} else if s == nil {
			logs.Info("ShowSeatInfo return nil")
			this.Data["json"] = map[string]interface{}{
				"body": s,
				"msg":  "ok",
			}
			this.ServeJSON()
		} else {
			logs.Info("ShowSeatInfo ok")
			this.Data["json"] = map[string]interface{}{
				"body": s,
				"msg":  "ok",
			}
			this.ServeJSON()
		}
	}

}

func (this *InsertChoiceSeatInfoController) Post() {
	var seat table.SeatInfo
	body := this.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &seat)
	if err != nil {
		logs.Error(err)
		this.Data["json"] = map[string]interface{}{
			"body": "",
			"msg":  "解析请求失败",
		}
		this.ServeJSON()
	} else {
		fmt.Println(seat)
		// 1.获取cookie
		key := this.Ctx.Input.Context.GetCookie("Cookie")
		// 2. 插入user_Hisinfo
		var his = &table.UserHisinfo{
			Id:          strconv.Itoa(int(time.Now().UnixMilli())),
			Phonenumber: key,
			Fenguan:     seat.Fenguan,
			Louceng:     seat.Louceng,
			Bianhao:     seat.Bianhao,
			Number:      seat.Number,
			Begintime:   seat.Begintime,
			Endtime:     seat.Endtime,
			Date:        seat.Date,
			Status:      seat.Status,
		}
		err = service.InsertHis(service.Engine, his)
		if err != nil {
			logs.Error("InsertHis return err")
			this.Data["json"] = map[string]interface{}{
				"body": "",
				"msg":  "err",
			}
			this.ServeJSON()
		} else {
			logs.Info("InsertHis ok ")
			// 1. 修改seat的座位信息
			s := &table.SeatInfo{
				Fenguan:   seat.Fenguan,
				Louceng:   seat.Louceng,
				Bianhao:   seat.Bianhao,
				Number:    seat.Number,
				Begintime: seat.Begintime,
				Endtime:   seat.Endtime,
				Date:      seat.Date,
				Status:    seat.Status,
			}
			err = service.ChangeChoiceSeatInfo(service.Engine, s)
			if err != nil {
				logs.Error(err)
				this.Data["json"] = map[string]interface{}{
					"body": "",
					"msg":  "ChangeChoiceSeatInfo err",
				}
				this.ServeJSON()
			} else {

				this.Data["json"] = map[string]interface{}{
					"body": s,
					"msg":  "ok",
				}
				this.ServeJSON()
			}
		}
	}

}

func (this *ShowHistoryInfoController) Post() {

	key := this.Ctx.Input.Context.GetCookie("Cookie")
	phone := key
	s, err := service.ShowHistoryInfo(service.Engine, phone)
	if err != nil {
		logs.Error(err)
		this.Data["json"] = map[string]interface{}{
			"body": "",
			"msg":  "ShowSeatInfo err",
		}
		this.ServeJSON()
	} else if s == nil {
		logs.Info("ShowSeatInfo return nil")
		this.Data["json"] = map[string]interface{}{
			"body": s,
			"msg":  "ok",
		}
		this.ServeJSON()
	} else {
		logs.Info("ShowSeatInfo ok")
		this.Data["json"] = map[string]interface{}{
			"body": s,
			"msg":  "ok",
		}
		this.ServeJSON()
	}
}

func (this *SeatinfoExistController) Post() {
	var seat table.SeatInfo
	body := this.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &seat)
	if err != nil {
		logs.Error(err)
		this.Data["json"] = map[string]interface{}{
			"body": seat,
			"msg":  "解析请求失败",
		}
		this.ServeJSON()
	} else {
		// 1. 删除user_hisinfo数据
		s, err := service.ShowseatExistInfo(service.Engine, &seat)
		if err != nil {
			logs.Error(err)
			this.Data["json"] = map[string]interface{}{
				"body": "",
				"msg":  "ShowseatExistInfo err",
			}
			this.ServeJSON()
		} else {
			logs.Info("DeleteHistoryInfo ok")
			this.Data["json"] = map[string]interface{}{
				"body": s,
				"msg":  "ok",
			}
			this.ServeJSON()

		}

	}

}

func (this *QuxiaoYuyueInfoController) Post() {
	var seat table.UserHisinfo
	body := this.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &seat)
	if err != nil {
		logs.Error(err)
		this.Data["json"] = map[string]interface{}{
			"body": seat,
			"msg":  "解析请求失败",
		}
		this.ServeJSON()
	} else {
		key := this.Ctx.Input.Context.GetCookie("Cookie")
		phone := key
		seat.Phonenumber = phone
		// 1. 删除user_hisinfo数据
		_, err := service.DeleteHistoryInfo(service.Engine, &seat)
		if err != nil {
			logs.Error(err)
			this.Data["json"] = map[string]interface{}{
				"body": "",
				"msg":  "DeleteHistoryInfo err",
			}
			this.ServeJSON()
		} else {

			// 2. 修改seat_info状态
			sinfo := &table.SeatInfo{
				Fenguan:   seat.Fenguan,
				Louceng:   seat.Louceng,
				Bianhao:   seat.Bianhao,
				Number:    seat.Number,
				Status:    0,
				Begintime: "",
				Endtime:   "",
				Date:      "",
			}
			err = service.ChangeChoiceSeatInfo(service.Engine, sinfo)
			if err != nil {
				logs.Error(err)
				this.Data["json"] = map[string]interface{}{
					"body": "",
					"msg":  "CangeChoiceSeatInfo err",
				}
				this.ServeJSON()
			} else {
				logs.Info("DeleteHistoryInfo ok")
				this.Data["json"] = map[string]interface{}{
					"body": "",
					"msg":  "ok",
				}
				this.ServeJSON()

			}

		}
	}
}

func (this *DeleteOneSeatController) Post() {
	var seat table.DelSeatInfo
	body := this.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &seat)
	if err != nil {
		logs.Error(err)
		this.Data["json"] = map[string]interface{}{
			"body": seat,
			"msg":  "解析请求失败",
		}
		this.ServeJSON()
	} else {
		ok, err := service.DeleteSeats(service.Engine, &seat)
		if err != nil || !ok {
			logs.Error(err)
			this.Data["json"] = map[string]interface{}{
				"body": "",
				"msg":  "删除座位出错",
			}
			this.ServeJSON()
		} else {
			logs.Info(err)
			this.Data["json"] = map[string]interface{}{
				"body": ok,
				"msg":  "ok",
			}
			this.ServeJSON()
		}
	}
}

func (this *DeleteAllSeatController) Post() {

	ok, err := service.DeleteAllSeats(service.Engine)
	if err != nil || !ok {
		logs.Error(err)
		this.Data["json"] = map[string]interface{}{
			"body": "",
			"msg":  "删除座位出错",
		}
		this.ServeJSON()
	} else {
		logs.Info(err)
		this.Data["json"] = map[string]interface{}{
			"body": ok,
			"msg":  "ok",
		}
		this.ServeJSON()
	}

}

func (this *ChangeSeatsInfoController) Post() {
	var seat table.SeatInfo
	body := this.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &seat)
	if err != nil {
		logs.Error(err)
		this.Data["json"] = map[string]interface{}{
			"body": seat,
			"msg":  "解析请求失败",
		}
		this.ServeJSON()
	} else {
		ok, err := service.ChangeSeats(service.Engine, &seat)
		if err != nil || !ok {
			logs.Error(err)
			this.Data["json"] = map[string]interface{}{
				"body": "",
				"msg":  "修改座位出错",
			}
			this.ServeJSON()
		} else {
			logs.Info(err)
			this.Data["json"] = map[string]interface{}{
				"body": ok,
				"msg":  "ok",
			}
			this.ServeJSON()
		}
	}
}

func (this *TuijianController) Post() {
	var seatInfo []table.SeatInfo
	key := this.Ctx.Input.Context.GetCookie("Cookie")
	phonenumber := key
	if _, err := service.GetRedisKey(service.Redisclient, phonenumber); err != nil {
		logs.Info("Cookie不存在")
		this.Data["json"] = map[string]interface{}{
			"body": "",
			"msg":  "Cookie不存在,登录超时",
		}
		this.ServeJSON()
	} else {

		count, err := service.ShowHistoryInfoCount(service.Engine, phonenumber)
		if err != nil {
			logs.Info("获取座位信息失败")
			this.Data["json"] = map[string]interface{}{
				"body": "",
				"msg":  "获取座位信息失败",
			}
			this.ServeJSON()
		} else if count > 10 {
			// 推荐相近的作为
			// ok, err = service.ShowSeatInfo(service.Engine, &)

			// logs.Info("获取座位信息成功")
			// this.Data["json"] = map[string]interface{}{
			// 	"body": seatInfo,
			// 	"msg":  "ok",
			// }
			// this.ServeJSON()
		} else {
			// 随机推荐
			seatInfo, err = service.ShowSeatInfoTuijian2(service.Engine)
			if err != nil {
				logs.Info("获取座位信息失败")
				this.Data["json"] = map[string]interface{}{
					"body": seatInfo,
					"msg":  "获取座位信息失败",
				}
				this.ServeJSON()
			} else {
				logs.Info("获取座位信息成功")
				this.Data["json"] = map[string]interface{}{
					"body": seatInfo,
					"msg":  "ok",
				}
				this.ServeJSON()
			}
		}

	}

}

// func (this *RecomSeatController) Post() {

// }
func (this *SignInSeatController) Post() {
	var seat table.SeatInfo
	var hisSeat table.UserHisinfo
	var s table.StatusStrings
	body := this.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &s)
	if err != nil {
		logs.Error(err)
		this.Data["json"] = map[string]interface{}{
			"body": s,
			"msg":  "解析请求失败",
		}
		this.ServeJSON()
	} else {
		key := this.Ctx.Input.Context.GetCookie("Cookie")
		phone := key
		hisSeat.Phonenumber = phone
		// 1. 获取seat信息
		list := strings.Split(s.Seatstrinngs, "_")
		seat.Fenguan = list[0]
		seat.Louceng = list[1]
		seat.Bianhao = list[2]
		seat.Number, _ = strconv.Atoi(list[3])
		seat.Status = s.Status

		hisSeat.Fenguan = list[0]
		hisSeat.Louceng = list[1]
		hisSeat.Bianhao = list[2]
		hisSeat.Number, _ = strconv.Atoi(list[3])
		hisSeat.Status = s.Status
		// 2. 修改seat 的status
		err = service.ChangeChoiceSeatInfo(service.Engine, &seat)
		if err != nil {
			logs.Error(err)
			this.Data["json"] = map[string]interface{}{
				"body": "",
				"msg":  "CangeChoiceSeatInfo err",
			}
			this.ServeJSON()
		} else {
			// 3. 修改his_seat 座位状态
			err = service.ChangeHisSeatInfo(service.Engine, &hisSeat)
			if err != nil {
				logs.Error(err)
				this.Data["json"] = map[string]interface{}{
					"body": "",
					"msg":  "CangeChoiceSeatInfo err",
				}
				this.ServeJSON()
			} else {
				logs.Info("DeleteHistoryInfo ok")
				this.Data["json"] = map[string]interface{}{
					"body": "",
					"msg":  "ok",
				}
				this.ServeJSON()
			}

		}
	}
}

// func (this *SignOutSeatController) Post() {

// }

// func (this *SetRebacksController) Post() {

// }

// func (this *GetRebacksController) Post() {

// }

// func (this *GetHistoryController) Post() {

// }

// func (this *UserDeleteController) Post() {

// }
