package usedb

import (
	"database/sql"
	"fmt"
	"libseat/models/table"
	"log"
	"strconv"
	"time"

	"github.com/beego/beego/logs"
	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
)

// 增
// 即用户注册 :支持用户单人注册
// 用户操作
func InsertUser(engine *xorm.EngineGroup, user *table.UserInfo) (uInfo *table.UserInfo, err error) {
	uInfo = &table.UserInfo{
		Id:          strconv.Itoa(int(time.Now().UnixMilli())),
		Name:        user.Name,
		Password:    user.Password,
		Phonenumber: user.Phonenumber,
		Admin:       false, //默认不能自己注册成为管理员
		Weichat:     user.Weichat,
		Integral:    100,
		Userimage:   "",
	}
	_, err = engine.InsertOne(uInfo)
	if err != nil {
		return nil, errors.Wrap(err, "pg insert err")
	}
	return uInfo, nil
}

func InsertAdmin(engine *xorm.EngineGroup, user *table.UserInfo) (uInfo *table.UserInfo, err error) {
	uInfo = &table.UserInfo{
		Id:          strconv.Itoa(int(time.Now().UnixMilli())),
		Name:        user.Name,
		Password:    user.Password,
		Phonenumber: user.Phonenumber,
		Admin:       true, //默认不能自己注册成为管理员
		Weichat:     user.Weichat,
		Integral:    100,
		Userimage:   "",
	}
	_, err = engine.InsertOne(uInfo)
	if err != nil {
		return nil, errors.Wrap(err, "pg insert err")
	}
	return uInfo, nil
}

// 删,单个删除和批量删除
//管理员操作
func DropUser(engine *xorm.EngineGroup, user []*table.UserInfo) (ok bool, err error) {
	var result sql.Result
	for _, v := range user {
		result, err = engine.Exec("delete from user_tbl where phonenumber=?", v.Phonenumber)
		if err != nil {
			log.Println(err)
			return false, errors.Wrap(err, "pg delete err")
		}
		// dblib=# delete from user_info where phonenumber='1212';
		// DELETE 0
		// 删除不存在的不报错,返回0
		rows, err := result.RowsAffected()
		if err == nil && rows > 0 {
			return true, nil
		}
	}
	return
}

//改 用户个人操作
func ChangeUser(engine *xorm.EngineGroup, user *table.UserInfo, phonenumber string) (users []table.UserInfo, err error) {
	updateInfo := &table.UserInfo{} //更新用户信息后暂时存放在这里
	if user.Name != "" {
		updateInfo.Name = user.Name
	}
	if user.Password != "" {
		updateInfo.Password = user.Password
	}
	if user.Weichat != "" {
		updateInfo.Weichat = user.Weichat
	}

	if user.Userimage != "" {
		updateInfo.Userimage = user.Userimage
	}

	_, err = engine.Where("phonenumber=?", phonenumber).Update(updateInfo)
	if err != nil {
		return nil, errors.Wrap(err, "pg update err")
	}
	user.Phonenumber = phonenumber
	users, err = InquireUser(engine, user)
	if err != nil {
		return nil, err
	}
	if users == nil {
		return nil, nil
	}
	return users, err
}

//查 用户登录，判断用户是否在用户信息表中
func InquireUser(engine *xorm.EngineGroup, user *table.UserInfo) (users []table.UserInfo, err error) {
	var checknumber int
	if user.Phonenumber == "" {
		checknumber = 1
	} else if user.Phonenumber != "" {
		checknumber = 0
	}
	//注册时使用phonenumber查询是否存在此人,存在返回用户,不存在返回nil
	switch {
	case checknumber == 0:
		err = engine.SQL("select * from user_info where phonenumber=?", user.Phonenumber).Find(&users)
		fmt.Println("inq--------------", users)
		if err != nil {
			return nil, errors.Wrap(err, "find by phonenumber err")
		}

		// 登录时查询用户是否存在使用,存在则返回用户,不存在返回nil
	case checknumber == 1:
		err = engine.SQL("select * from user_info where name=? and password=?", user.Name, user.Password).Find(&users)
		fmt.Println(users)
		if err != nil {
			return nil, errors.Wrap(err, "find by name and password err")
		}
	}
	return
}

func InquireAdminUser(engine *xorm.EngineGroup) (users []table.UserInfo, err error) {
	err = engine.SQL("select * from user_info where admin=true").Find(&users)
	fmt.Println("inq--------------", users)
	if err != nil {
		return nil, errors.Wrap(err, "find admin err")
	}
	return
}

//----------------------------------seat-----------------------------------------------------------//

// 增  一个
func InsertSeats(engine *xorm.EngineGroup, seatInfo *table.InsertSeatInfo) (err error) {
	seats := &table.SeatInfo{
		Id:        strconv.Itoa(int(time.Now().UnixMilli())),
		Fenguan:   seatInfo.Fenguan,
		Louceng:   seatInfo.Louceng,
		Bianhao:   seatInfo.Bianhao,
		Number:    seatInfo.Number,
		Status:    0, //默认是空
		Begintime: "",
		Endtime:   "",
		Date:      "",
		SeatImage: seatInfo.SeatImage,
	}
	_, err = engine.Insert(seats)
	if err != nil {
		return err
	}
	return
}

func InsertHis(engine *xorm.EngineGroup, l *table.UserHisinfo) (err error) {
	// _, err = engine.SQL("insert into user_hisinfo ( phonenumber , fenguan ,louceng ,bianhao ,number , begintime , endtime ) values(?,?,?,?,?,?,?)", l.Phonenumber, l.Fenguan, l.Louceng, l.Bianhao, l.Number, l.Begintime, l.Endtime).Insert()
	s := &table.UserHisinfo{
		Id:          strconv.Itoa(int(time.Now().UnixMilli())),
		Phonenumber: l.Phonenumber,
		Fenguan:     l.Fenguan,
		Bianhao:     l.Bianhao,
		Louceng:     l.Louceng,
		Number:      l.Number,
		Begintime:   l.Begintime,
		Endtime:     l.Endtime,
		Date:        l.Date,
		Status:      l.Status,
	}
	_, err = engine.Insert(s)
	if err != nil {
		return err
	}
	return

}

func InsertReback(engine *xorm.EngineGroup, l *table.UserReBack) (err error) {
	s := &table.UserReBack{
		Id:           strconv.Itoa(int(time.Now().UnixMilli())),
		Userphone:    l.Userphone,
		Usermsgtext:  l.Usermsgtext,
		Umsgimage:    l.Umsgimage,
		Adminphone:   l.Adminphone,
		Adminmsgtext: l.Adminmsgtext,
		Amsgimage:    l.Amsgimage,
		Status:       l.Status,
	}
	_, err = engine.Insert(s)
	if err != nil {
		return err
	}
	return

}

// 删
func DropSeat(engine *xorm.EngineGroup, seat *table.SeatInfo) (bool, error) {
	result, err := engine.Exec("delete from seat_info where fenguan=? and louceng=? and bianhao=? and number=?", seat.Fenguan, seat.Louceng, seat.Bianhao, seat.Number)
	if err != nil {
		log.Println(err)
		return false, err
	}
	rows, err := result.RowsAffected()
	if err == nil && rows > 0 {
		return true, nil
	}
	return false, err
}

func DropAllSeat(engine *xorm.EngineGroup) (bool, error) {
	result, err := engine.Exec("delete from seat_info")
	if err != nil {
		log.Println(err)
		return false, err
	}
	rows, err := result.RowsAffected()
	if err == nil && rows > 0 {
		return true, nil
	}
	return false, err
}

//改  分管,楼层，编号，number  确定一个位置
func ChangeSeatOne(engine *xorm.EngineGroup, seat *table.SeatInfo) (ok bool, err error) {
	_, err = engine.Where("id=?", seat.Id).Update(seat)
	if err != nil {
		return false, err
	}
	return true, nil
}

// //改  分管,楼层，编号，number  确定一个位置
// func ChangeSeats(engine *xorm.EngineGroup, sOld *table.SeatInfo, sNew *table.SeatInfo) (ok bool, err error) {
// 	_, err = engine.Where("fenguan=? and louceng=? and bianhao=? and number=?", sOld.Fenguan, sOld.Louceng, sOld.Bianhao, sOld.Number).Update(sNew)
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }

// 条件查询 查
func InquireSeat(engine *xorm.EngineGroup, l *table.Feng) (seatInfos []table.SeatInfo, err error) {

	err = engine.SQL("select * from seat_info where status=0 and fenguan=?", l.Fenguan).Find(&seatInfos)
	if err != nil {
		return nil, err
	}

	fmt.Println(seatInfos)
	return
}

// 插入查询存在性
func InquireSeatS(engine *xorm.EngineGroup, l *table.InsertSeatInfo) (seatInfos []table.SeatInfo, err error) {

	err = engine.SQL("select * from seat_info where fenguan=? and louceng=? and bianhao=? and number=?", l.Fenguan, l.Louceng, l.Bianhao, l.Number).Find(&seatInfos)
	if err != nil {
		return nil, err
	}
	fmt.Println(seatInfos)
	return
}

// FenguanInfo 查询
func InquireFenguanCountInfo(engine *xorm.EngineGroup) (FenguanInfo []table.FenguanInfo, err error) {
	var t []table.TempFenguanInfo
	err = engine.SQL("select count(*) AS Fenallseatnum , fenguan as Fenname  from seat_info GROUP BY fenguan order by fenguan").Find(&t)
	if err != nil {
		logs.Error("TempFenguanInfo err")
		return nil, err
	}
	fmt.Println(t)
	for _, v := range t {
		var tk []table.TemKong
		err = engine.SQL("select count(*) AS Kongseatnum FROM seat_info where fenguan=? and status=0", v.Fenname).Find(&tk)
		if err != nil {
			logs.Error("Temp_status_FenguanInfo err")
			return nil, err
		}
		var F = table.FenguanInfo{
			Fenname:       v.Fenname,
			Fenallseatnum: v.Fenallseatnum,
			Kongseatnum:   tk[0].Kongseatnum,
		}
		FenguanInfo = append(FenguanInfo, F)
	}
	fmt.Println(FenguanInfo)
	return
}

func InquireFenguanNameInfo(engine *xorm.EngineGroup) (seatInfo []table.SeatInfo, err error) {

	err = engine.SQL("select fenguan as Fenguan from seat_info group by Fenguan").Find(&seatInfo)
	if err != nil {
		logs.Error("err")
		return nil, err
	}
	fmt.Println(seatInfo)
	return
}

func InquireFenguanLoucengInfo(engine *xorm.EngineGroup, s *table.SeatInfo) (seatinfo []table.SeatInfo, err error) {

	err = engine.SQL("select louceng as Louceng  from seat_info where fenguan=? group by Louceng", s.Fenguan).Find(&seatinfo)
	if err != nil {
		logs.Error("err")
		return nil, err
	}
	fmt.Println(seatinfo)
	return
}

func InquireFenguanBianhaoInfo(engine *xorm.EngineGroup, s *table.SeatInfo) (seatinfo []table.SeatInfo, err error) {

	err = engine.SQL("select bianhao as Bianhao  from seat_info where fenguan=? and louceng=? group by Bianhao", s.Fenguan, s.Louceng).Find(&seatinfo)
	if err != nil {
		logs.Error("err")
		return nil, err
	}
	fmt.Println(seatinfo)
	return
}
func InquireFenguanNumberInfo(engine *xorm.EngineGroup, s *table.SeatInfo) (seatinfo []table.SeatInfo, err error) {

	err = engine.SQL("select number as Number  from seat_info where fenguan=? and louceng=? and bianhao=? and status=0", s.Fenguan, s.Louceng, s.Bianhao).Find(&seatinfo)
	if err != nil {
		logs.Error("err")
		return nil, err
	}
	fmt.Println(seatinfo)
	return
}

func InquierHistory(engine *xorm.EngineGroup, phone string) (his []table.UserHisinfo, err error) {
	err = engine.SQL("select * from user_hisinfo where phonenumber=?", phone).Find(&his)
	if err != nil {
		logs.Error("err")
		return nil, err
	}
	fmt.Println(his)
	return
}

func InquierHistoryCount(engine *xorm.EngineGroup, phone string) (count int, err error) {
	var c = []table.Count{}
	err = engine.SQL("select count(*) AS count from user_hisinfo where phonenumber=?", phone).Find(&c)
	if err != nil {
		logs.Error("err")
		return
	}
	count = c[0].Count
	fmt.Println(c[0].Count)
	return
}

func InquireTuijian2SeatInfo(engine *xorm.EngineGroup) (seatinfo []table.SeatInfo, err error) {
	err = engine.SQL("select * from seat_info where status=0 limit 5").Find(&seatinfo)
	if err != nil {
		logs.Error("err")
		return nil, err
	}
	fmt.Println(seatinfo)
	return
}

func DeleteHistoryInfo(engine *xorm.EngineGroup, seat *table.UserHisinfo) (ok bool, err error) {
	result, err := engine.Exec("delete from user_hisinfo where phonenumber=? and fenguan=? and louceng=? and bianhao=? and number=? and begintime=? and endtime=? and date=? and status=?", seat.Phonenumber, seat.Fenguan, seat.Louceng, seat.Bianhao, seat.Number, seat.Begintime, seat.Endtime, seat.Date, seat.Status)
	if err != nil {
		log.Println(err)
		return false, err
	}
	rows, err := result.RowsAffected()
	if err == nil && rows > 0 {
		return true, nil
	}
	return false, err
}

func ChangeChoiceSeatInfo(engine *xorm.EngineGroup, s *table.SeatInfo) (err error) {
	ss := make([]table.SeatInfo, 1)
	err = engine.SQL("update seat_info set status=?,begintime=?, endtime=? , date=? where fenguan=? and louceng=? and bianhao=? and number=? ", s.Status, s.Begintime, s.Endtime, s.Date, s.Fenguan, s.Louceng, s.Bianhao, s.Number).Find(&ss)
	if err != nil {
		logs.Error("err")
		return err
	}
	return
}

func ChangeHisSeatInfo(engine *xorm.EngineGroup, s *table.UserHisinfo) (err error) {
	ss := make([]table.UserHisinfo, 1)
	err = engine.SQL("update user_hisinfo set status=? where fenguan=? and louceng=? and bianhao=? and number=? ", s.Status, s.Fenguan, s.Louceng, s.Bianhao, s.Number).Find(&ss)
	if err != nil {
		logs.Error("err")
		return err
	}
	return
}

func InquireFenguanAllInfo(engine *xorm.EngineGroup, name string) (F []table.SeatInfo, err error) {
	err = engine.SQL("select * from seat_info where fenguan=? order by id", name).Find(&F)
	if err != nil {
		logs.Error("TempFenguanInfo err")
		return nil, err
	}
	fmt.Println(F)
	return
}

func InquireseatExistInfo(engine *xorm.EngineGroup, seat *table.SeatInfo) (s []table.SeatInfo, err error) {
	err = engine.SQL("select * from seat_info where fenguan=? and louceng=? and bianhao=? and number=?", seat.Fenguan, seat.Louceng, seat.Bianhao, seat.Number).Find(&s)

	if err != nil {
		logs.Error("InquireseatExistInfo err")
		return nil, err
	}
	fmt.Println(s)
	return
}
