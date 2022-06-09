package controllers

import (
	"encoding/json"
	"fmt"
	"libseat/models/table"
	"libseat/service"
	"strconv"
	"time"

	"github.com/beego/beego/logs"
)

func (this *SendCode) Post() {
	// 1. 数据解析
	var co table.GetCodeMsg
	body := this.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &co)
	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"body": co,
			"msg":  "解析请求失败",
		}
		this.ServeJSON()
	} else {
		// 2.redis 存储code  key = code val = phone
		code := strconv.Itoa(service.GetCode())
		err = service.SetRedisKey(service.Redisclient, code, co.Phonenumber, 80*time.Second)
		if err != nil {
			this.Data["json"] = map[string]interface{}{
				"body": "",
				"msg":  "Redis set code err",
			}
			this.ServeJSON()
		} else {
			// 在redis中保存成功才发送code
			service.VerCode(code, co.Phonenumber)
			this.Data["json"] = map[string]interface{}{
				"body": code,
				"msg":  "ok",
			}
			this.ServeJSON()
		}
	}
}

// 注册需要 {
// 	name
// 	password
// 	phonenumber
// 	Code
// }
// 验证手机号不存在并且code存在且手机号对应相同 即可注册

func (this *RegisterController) Post() {
	// 1. 解析数据
	var newuser table.RegistNewUser
	body := this.Ctx.Input.RequestBody    //这是获取到request的body 的json二进制数据
	err := json.Unmarshal(body, &newuser) //解析二进制json，把结果放进ob中
	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"body": newuser,
			"msg":  "解析请求失败",
		}
		this.ServeJSON()
	} else if newuser.Name != "" && newuser.Password != "" && newuser.Phonenumber != "" && newuser.Code != "" {
		// 注册数据全不能为空
		logs.Info(newuser.Name, "\t", newuser.Password, "\t", newuser.Phonenumber, "\t", newuser.Code)

		// 2. 判断 code 是否存在
		val, err := service.GetRedisKey(service.Redisclient, newuser.Code)
		if err != nil {
			logs.Error("注册失败:", err)
			this.Data["json"] = map[string]interface{}{
				"body": "",
				"msg":  "验证码错误",
			}
			this.ServeJSON()
		} else if val != newuser.Phonenumber {
			// 3. 判断当前输入的手机号是否为用户获取code的同一个手机号
			logs.Error("注册失败:", err)
			this.Data["json"] = map[string]interface{}{
				"body": "",
				"msg":  "验证码错误",
			}
			this.ServeJSON()
		} else {
			// 4. 注册账号
			Info, err := service.RegistNewUser(service.Engine, &newuser)
			if err != nil {
				logs.Error("注册失败:", err)
				this.Data["json"] = map[string]interface{}{
					"body": newuser,
					"msg":  "注册账号失败",
				}
				this.ServeJSON()
			}
			if Info == nil {
				this.Data["json"] = map[string]interface{}{
					"body": newuser,
					"msg":  "此号码已经注册过了",
				}
				this.ServeJSON()
			} else {
				this.Data["json"] = map[string]interface{}{
					"body": newuser,
					"msg":  "ok",
				}
				this.ServeJSON()
			}

		}

	} else {
		logs.Error("注册失败:用户名或密码或电话或验证码为空")
		this.Data["json"] = map[string]interface{}{
			"body": newuser,
			"msg":  "用户名、密码、电话、验证码都不能为空",
		}
	}
}
func (this *RegisterAdminController) Post() {
	// 1. 解析数据
	var newuser table.RegistNewUser
	body := this.Ctx.Input.RequestBody    //这是获取到request的body 的json二进制数据
	err := json.Unmarshal(body, &newuser) //解析二进制json，把结果放进ob中
	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"body": newuser,
			"msg":  "解析请求失败",
		}
		this.ServeJSON()
	} else if newuser.Name != "" && newuser.Password != "" && newuser.Phonenumber != "" && newuser.Code != "" {
		// 注册数据全不能为空
		logs.Info(newuser.Name, "\t", newuser.Password, "\t", newuser.Phonenumber, "\t", newuser.Code)

		// 2. 判断 code 是否存在
		val, err := service.GetRedisKey(service.Redisclient, newuser.Code)
		if err != nil {
			logs.Error("注册失败:", err)
			this.Data["json"] = map[string]interface{}{
				"body": "",
				"msg":  "验证码错误",
			}
			this.ServeJSON()
		} else if val != newuser.Phonenumber {
			// 3. 判断当前输入的手机号是否为用户获取code的同一个手机号
			logs.Error("注册失败:", err)
			this.Data["json"] = map[string]interface{}{
				"body": "",
				"msg":  "验证码错误",
			}
			this.ServeJSON()
		} else {
			// 4. 注册账号
			Info, err := service.RegistNewAdmin(service.Engine, &newuser)
			if err != nil {
				logs.Error("注册失败:", err)
				this.Data["json"] = map[string]interface{}{
					"body": newuser,
					"msg":  "注册账号失败",
				}
				this.ServeJSON()
			}
			if Info == nil {
				this.Data["json"] = map[string]interface{}{
					"body": newuser,
					"msg":  "此号码已经注册过了",
				}
				this.ServeJSON()
			} else {
				this.Data["json"] = map[string]interface{}{
					"body": newuser,
					"msg":  "ok",
				}
				this.ServeJSON()
			}

		}

	} else {
		logs.Error("注册失败:用户名或密码或电话或验证码为空")
		this.Data["json"] = map[string]interface{}{
			"body": newuser,
			"msg":  "用户名、密码、电话、验证码都不能为空",
		}
	}
}

func (this *FindBackPasswordController) Post() {
	// 1. 数据解析
	var fb table.FindBackPassword
	body := this.Ctx.Input.RequestBody //这是获取到request的body 的json二进制数据
	err := json.Unmarshal(body, &fb)   //解析二进制json，把结果放进ob中
	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"body": fb,
			"msg":  "解析请求失败",
		}
		this.ServeJSON()
	} else if fb.Phonenumber != "" && fb.Password != "" && fb.Code != "" {
		// 2. 数据不能为空
		// 3. 验证手机号和验证码是否匹配
		val, err := service.GetRedisKey(service.Redisclient, fb.Code)
		if err != nil {
			logs.Error("找回密码失败:", err)
			this.Data["json"] = map[string]interface{}{
				"body": "",
				"msg":  "验证码获取错误",
			}
			this.ServeJSON()
		} else if val != fb.Phonenumber {
			logs.Error("找回密码失败:", err)
			this.Data["json"] = map[string]interface{}{
				"body": "",
				"msg":  "验证码匹配错误",
			}
			this.ServeJSON()
		} else {
			// 4. 检查Phonenumber是否存在
			_, err := service.ShowUserInfo(service.Engine, fb.Phonenumber)
			if err != nil {
				logs.Error("查询用户失败:", err)
				this.Data["json"] = map[string]interface{}{
					"body": "",
					"msg":  "查询用户失败",
				}
				this.ServeJSON()
			} else {
				// 5. 设置新密码
				// 用户存在
				// 设置新密码
				// u[0].Password = fb.Password
				var user = &table.UserInfo{
					Password: fb.Password,
				}
				// 更新用户密码
				users, err := service.ChangeUserInfo(service.Engine, *user, fb.Phonenumber)
				fmt.Println(">>>>>>>>>", users)
				if err != nil {
					logs.Error("更新用户密码失败:", err)
					this.Data["json"] = map[string]interface{}{
						"body": "",
						"msg":  "更新用户密码失败",
					}
					this.ServeJSON()
				} else {
					logs.Info("更新用户密码成功:", err)
					this.Data["json"] = map[string]interface{}{
						"body": users,
						"msg":  "ok",
					}
					this.ServeJSON()
				}
			}
		}
	}
}

func (this *LoginBypasswordController) Post() {
	// 1. 数据解析
	var userlogin table.UserLogin
	body := this.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &userlogin)
	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"body": userlogin,
			"msg":  "解析请求失败",
		}
		this.ServeJSON()
	} else if userlogin.Name != "" && userlogin.Password != "" {
		// 2. 登录数据不能为空
		fmt.Println(userlogin)
		// 3. 登录数据匹配
		u, err := service.UserLogin(service.Engine, &userlogin)
		if err != nil {
			logs.Error("用户名或密码错误:", err)
			this.Data["json"] = map[string]interface{}{
				"body": "",
				"msg":  "用户名或密码错误",
			}
			this.ServeJSON()
		} else {
			// 4. 登录设置 cookie
			//登录成功设置cookie
			// 使用表中的key键组合cookie的key
			phonenumber := u[0].Phonenumber

			// 1. 判断是否存在于黑名单中
			// 判断用户是否在黑名单中
			if val, _ := service.GetRedisKey(service.Redisclient, phonenumber); val == "nook" {
				// 3. 判断用户是在黑名单中
				logs.Info("用户信誉积分不足")
				this.Data["json"] = map[string]interface{}{
					"body": phonenumber,
					"msg":  "用户信誉积分不足,无权查看",
				}
				this.ServeJSON()
			}
			// 2. 根据获取的用户信息，添加cookie内

			//保存cookie到redis,并设置时间限制   val = ok 说明用户不在黑名单中，否侧val = no,说明用户在黑名单中，能登录但是不能查看数据
			err := service.SetRedisKey(service.Redisclient, phonenumber, "ok", 604800*time.Second) //key : val,一周
			if err != nil {
				logs.Error("保存cookie失败:", err)
				this.Data["json"] = map[string]interface{}{
					"body": userlogin,
					"msg":  "保存cookie失败,登陆失败",
				}
				this.ServeJSON()
			} else {
				// 5. header中设置cookie
				this.Ctx.Output.Cookie("Cookie", phonenumber) // Cookie : key
				logs.Info("cookie设置成功")

				this.Data["json"] = map[string]interface{}{
					"body": u,
					"msg":  "ok",
				}
				this.ServeJSON()
			}
		}

	} else {
		this.Data["json"] = map[string]interface{}{
			"body": userlogin,
			"msg":  "用户名或密码不能为空",
		}
		this.ServeJSON()
	}

}

func (this *ShowUserInfoController) Post() {
	// 显示用户信息只需要cookie
	// 1. 获取 cookie
	var userinfo []table.UserInfo
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
		// 4. 获取用户信息
		phonenumber := key
		fmt.Println(phonenumber)
		userinfo, err = service.ShowUserInfo(service.Engine, phonenumber)
		if err != nil {
			this.Data["json"] = map[string]interface{}{
				"body": "",
				"msg":  "获取用户信息失败",
			}
			this.ServeJSON()
		} else {
			this.Data["json"] = map[string]interface{}{
				"body": userinfo,
				"msg":  "ok",
			}
			this.ServeJSON()
		}
	}
}

func (this *ShowAdminUserInfoController) Post() {
	info, err := service.ShowAdminUserInfo(service.Engine)
	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"body": "",
			"msg":  "获取用户信息失败",
		}
		this.ServeJSON()
	} else {
		this.Data["json"] = map[string]interface{}{
			"body": info,
			"msg":  "ok",
		}
		this.ServeJSON()
	}
}

func (this *ChangeUserInfoController) Post() {
	var userinfo []table.UserInfo
	var userinfochange table.UserInfo
	// 1.获取cookie
	key := this.Ctx.Input.Context.GetCookie("Cookie")
	if _, err := service.GetRedisKey(service.Redisclient, key); err != nil {
		logs.Info("Cookie不存在")
		this.Data["json"] = map[string]interface{}{
			"body": "",
			"msg":  "Cookie不存在,登录超时",
		}
		this.ServeJSON()
	} else {
		// 2. 数据解析
		// 可修改数据 ： 用户名、密码、头像 、weichat
		phonenumber := key
		body := this.Ctx.Input.RequestBody
		err := json.Unmarshal(body, &userinfochange)
		// 3. 数据查找更新
		// 即不修改电话号码,不用重新设置cookie
		userinfo, err = service.ChangeUserInfo(service.Engine, userinfochange, phonenumber)
		if err != nil {
			this.Data["json"] = map[string]interface{}{
				"body": "",
				"msg":  "获取用户信息失败",
			}
			this.ServeJSON()
		}
		this.Data["json"] = map[string]interface{}{
			"body": userinfo,
			"msg":  "ok",
		}
		this.ServeJSON()
	}
}

func (this *UserLoginOutController) Post() {
	key := this.Ctx.Input.Context.GetCookie("Cookie")
	if err := service.DelRedisKey(service.Redisclient, key); err != nil {
		logs.Info("Del Cookie 失败")
		this.Data["json"] = map[string]interface{}{
			"body": key,
			"msg":  "Del Cookie 失败",
		}
		this.ServeJSON()
	} else {
		this.Data["json"] = map[string]interface{}{
			"body": "",
			"msg":  "ok",
		}
		this.ServeJSON()
	}
}

func (this *UserimageuploadController) Post() {
	// 1. 判断cookie
	key := this.Ctx.Input.Context.GetCookie("Cookie")
	if _, err := service.GetRedisKey(service.Redisclient, key); err != nil {
		logs.Info("Cookie不存在")
		this.Data["json"] = map[string]interface{}{
			"body": "",
			"msg":  "Cookie不存在,登录超时",
		}
		this.ServeJSON()
	} else {
		var rebckmsg table.UserReBack
		// 2. 获取文件
		body := this.Ctx.Input.RequestBody
		err = json.Unmarshal(body, &rebckmsg)
		// if err != nil {
		// 	this.Data["json"] = map[string]interface{}{
		// 		"body": "",
		// 		"msg":  "解析错误",
		// 	}
		// 	this.ServeJSON()
		// }
		// 3. 文件保存本地
		// a. 把值凑好
		phonenumber := key
		rebckmsg.Userphone = phonenumber
		// b. 重组数据图片存云,数据库存连接
		// fmt.Println(rebckmsg.Umsgimage)
		url := service.Base64Tourl(rebckmsg.Umsgimage)
		// fmt.Println(url)
		rebckmsg.Umsgimage = url
		fmt.Println(rebckmsg.Umsgimage)
		rebckmsg.Status = false
		// c. 插入数据库
		err = service.InsertReback(service.Engine, &rebckmsg)
		if err != nil {
			logs.Error("插入数据库失败")
			this.Data["json"] = map[string]interface{}{
				"body": "rebckmsg",
				"msg":  "insert err",
			}
			this.ServeJSON()
		} else {
			logs.Info("接收imagebase64")
			this.Data["json"] = map[string]interface{}{
				"body": rebckmsg,
				"msg":  "ok",
			}
			this.ServeJSON()
		}
	}
}

// func (this *UserDeleteController) Post() {

// }

// func (this *SignInSeatController) Post() {
// 	var sign table.SginCode
// 	body := this.Ctx.Input.RequestBody
// 	err := json.Unmarshal(body, &sign)
// 	if err != nil {
// 		this.Data["json"] = map[string]interface{}{
// 			"body": sign,
// 			"msg":  "解析请求失败",
// 		}
// 		this.ServeJSON()
// 	} else {
// 		fmt.Println(sign)
// 		var l = &table.SeatInfo{
// 			Fenguan:   sign.Fenguan,
// 			Louceng:   sign.Louceng,
// 			Bianhao:   sign.Bianhao,
// 			Number:    sign.Number,
// 			Status:    sign.Status,
// 			Begintime: sign.Begintime,
// 			Endtime:   sign.Endtime,
// 		}
// 		//扫码一次
// 		if sign.Signbegin && !sign.Signend {
// 			ok, err := service.ChangeSeats(service.Engine, l)
// 			if err != nil || !ok {
// 				logs.Error(err)
// 				this.Data["json"] = map[string]interface{}{
// 					"body": "",
// 					"msg":  "err",
// 				}
// 				this.ServeJSON()

// 			} else {
// 				this.Data["json"] = map[string]interface{}{
// 					"body": ok,
// 					"msg":  "ok",
// 				}
// 				this.ServeJSON()
// 			}
// 		}
// 		//二次扫码
// 		if sign.Signbegin && sign.Signend {
// 			ok, err := service.ChangeSeats(service.Engine, l)
// 			if err != nil || !ok {
// 				logs.Error(err)
// 				this.Data["json"] = map[string]interface{}{
// 					"body": "",
// 					"msg":  "err",
// 				}
// 				this.ServeJSON()

// 			} else {
// 				this.Data["json"] = map[string]interface{}{
// 					"body": ok,
// 					"msg":  "ok",
// 				}
// 				this.ServeJSON()
// 			}
// 		}
// 	}
// }

func (this *ChoiceSeatController) Post() {
	key := this.Ctx.Input.Context.GetCookie("Cookie")
	if _, err := service.GetRedisKey(service.Redisclient, key); err != nil {
		logs.Info("Cookie不存在")
		this.Data["json"] = map[string]interface{}{
			"body": "",
			"msg":  "Cookie不存在,登录超时",
		}
		this.ServeJSON()
	} else {
		phonenumber := key
		var t table.SeatInfo
		body := this.Ctx.Input.RequestBody
		err := json.Unmarshal(body, &t)
		if err != nil {
			this.Data["json"] = map[string]interface{}{
				"body": t,
				"msg":  "解析请求失败",
			}
			this.ServeJSON()
		} else {
			if t.Fenguan != "" && t.Louceng == "" {
				f, err := service.ShowFenguanAllInfo(service.Engine, t.Fenguan)
				if err != nil || f == nil {
					this.Data["json"] = map[string]interface{}{
						"body": f,
						"msg":  "1 err",
					}
					this.ServeJSON()
				} else {
					this.Data["json"] = map[string]interface{}{
						"body": f,
						"msg":  "ok",
					}
					this.ServeJSON()
				}
			}
			if t.Fenguan != "" && t.Louceng != "" && t.Bianhao == "" {
				f, err := service.ShowFenguanAllInfo(service.Engine, t.Fenguan)
				if err != nil || f == nil {
					this.Data["json"] = map[string]interface{}{
						"body": f,
						"msg":  "2 err",
					}
					this.ServeJSON()
				} else {
					this.Data["json"] = map[string]interface{}{
						"body": f,
						"msg":  "ok",
					}
					this.ServeJSON()
				}
			}
			if t.Fenguan != "" && t.Louceng != "" && t.Bianhao != "" && t.Number == 0 {
				f, err := service.ShowFenguanAllInfo(service.Engine, t.Fenguan)
				if err != nil || f == nil {
					this.Data["json"] = map[string]interface{}{
						"body": f,
						"msg":  "3 err",
					}
					this.ServeJSON()
				} else {
					this.Data["json"] = map[string]interface{}{
						"body": f,
						"msg":  "ok",
					}
					this.ServeJSON()
				}
			}
			if t.Fenguan != "" && t.Louceng != "" && t.Bianhao != "" && t.Number != 0 {
				f, err := service.ShowFenguanAllInfo(service.Engine, t.Fenguan)
				if err != nil || f == nil {
					this.Data["json"] = map[string]interface{}{
						"body": f,
						"msg":  "4 err",
					}
					this.ServeJSON()
				} else {
					histable := &table.UserHisinfo{
						Fenguan:     f[0].Fenguan,
						Louceng:     f[0].Louceng,
						Bianhao:     f[0].Bianhao,
						Number:      f[0].Number,
						Begintime:   f[0].Begintime,
						Endtime:     f[0].Endtime,
						Phonenumber: phonenumber,
					}
					// 将hitable添加进表user_hisinfo中
					err = service.InsertHis(service.Engine, histable)
					if err != nil {
						this.Data["json"] = map[string]interface{}{
							"body": histable,
							"msg":  "insert err",
						}
						this.ServeJSON()
					} else {
						this.Data["json"] = map[string]interface{}{
							"body": histable,
							"msg":  "ok",
						}
						this.ServeJSON()
					}
				}
			}
		}

	}

}

func (this *GetHistoryController) Post() {
	// cookie 得到电话返回
}
