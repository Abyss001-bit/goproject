package service

import (
	"fmt"
	"libseat/models/table"
	usedb "libseat/models/usedb"

	"github.com/astaxie/beego/logs"
	"github.com/go-xorm/xorm"
)

// 用户注册,一个电话号码只能注册一个，phonenumber都key
func RegistNewUser(engine *xorm.EngineGroup, u *table.RegistNewUser) (uInfo *table.UserInfo, err error) {
	var userinfo table.UserInfo
	userinfo.Phonenumber = u.Phonenumber
	// 1. 查询 phonenumber 是否存在
	users, err := usedb.InquireUser(engine, &userinfo)
	if err != nil {
		logs.Error("inquire user err:", err)
	}
	//1.1号码存在，注册失败，返回nil
	if users != nil {
		logs.Info("此号码已注册")
		return nil, nil
	}
	// 1.2 号码不存在，注册成功，返回用户全部信息
	userinfo.Name = u.Name
	userinfo.Password = u.Password
	uInfo, err = usedb.InsertUser(engine, &userinfo)
	if err != nil {
		logs.Error("insert user err:", err)
		return nil, err
	}
	return uInfo, nil
}

// 用户注册,一个电话号码只能注册一个，phonenumber都key
func RegistNewAdmin(engine *xorm.EngineGroup, u *table.RegistNewUser) (uInfo *table.UserInfo, err error) {
	var userinfo table.UserInfo
	userinfo.Phonenumber = u.Phonenumber
	// 1. 查询 phonenumber 是否存在
	users, err := usedb.InquireUser(engine, &userinfo)
	if err != nil {
		logs.Error("inquire user err:", err)
	}
	//1.1号码存在，注册失败，返回nil
	if users != nil {
		logs.Info("此号码已注册")
		return nil, nil
	}
	// 1.2 号码不存在，注册成功，返回用户全部信息
	userinfo.Name = u.Name
	userinfo.Password = u.Password
	uInfo, err = usedb.InsertAdmin(engine, &userinfo)
	if err != nil {
		logs.Error("insert user err:", err)
		return nil, err
	}
	return uInfo, nil
}

//登录,判断用户名和密码是否同时满足存在，满足则登录成功，否则查无此人登录失败
func UserLogin(engine *xorm.EngineGroup, u *table.UserLogin) (users []table.UserInfo, err error) {
	var userinfo = &table.UserInfo{
		Name:     u.Name,
		Password: u.Password,
	}
	fmt.Println(userinfo)
	users, err = usedb.InquireUser(engine, userinfo)
	fmt.Println("-------------------------------------------")
	if err != nil {
		logs.Error("inquire user err:", err)
		return nil, err
	} else if users == nil {
		logs.Error("查询无此人")
		return nil, nil
	}
	return users, nil
}

//在登录状态下显示个人全部信息
func ShowUserInfo(engine *xorm.EngineGroup, phonenumber string) (users []table.UserInfo, err error) {
	var userinfo table.UserInfo
	userinfo.Phonenumber = phonenumber
	users, err = usedb.InquireUser(engine, &userinfo)
	if err != nil {
		logs.Error("inquire user err:", err)
		return nil, err
	} else if users == nil {
		logs.Error("ShowUserInfo查询无此人")
		return nil, nil
	}
	return users, nil
}

func ShowAdminUserInfo(engine *xorm.EngineGroup) (users []table.UserInfo, err error) {
	users, err = usedb.InquireAdminUser(engine)
	if err != nil {
		logs.Error("inquire user err:", err)
		return nil, err
	} else if users == nil {
		logs.Error("ShowUserInfo查询无此人")
		return nil, nil
	}
	return users, nil
}

//在登录状态下,进行修改用户信息
func ChangeUserInfo(engine *xorm.EngineGroup, userinfochange table.UserInfo, phonenumber string) (users []table.UserInfo, err error) {

	users, err = usedb.ChangeUser(engine, &userinfochange, phonenumber)
	if err != nil {
		logs.Error("ChangeUser user err:", err)
		return nil, err
	} else if users == nil {
		logs.Error("ChangeUserInfo查询无此人")
		return nil, nil
	}
	return
}

// //用户是否存在,即电话号码是否存在
// func ExistUser(engine *xorm.EngineGroup, phonenumber string) (users []*table.UserInfo, err error) {
// 	var userinfo table.UserInfo
// 	userinfo.Phonenumber = phonenumber
// 	users, err = usedb.InquireUser(engine, &userinfo)
// 	if err != nil {
// 		logs.Error("inquire user err:", err)
// 		return nil, err
// 	} else if users == nil {
// 		logs.Error("ExistUser查询无此人")
// 		return nil, nil
// 	}
// 	return users, nil
// }
