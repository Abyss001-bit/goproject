package service

import (
	"libseat/models/table"
	usedb "libseat/models/usedb"

	"github.com/astaxie/beego/logs"
	"github.com/go-xorm/xorm"
)

func SerachSeatsInfo(engine *xorm.EngineGroup, l *table.InsertSeatInfo) (seatInfos []table.SeatInfo, err error) {
	seatInfos, err = usedb.InquireSeatS(engine, l)
	if err != nil {
		logs.Error("inquire seat err:", err)
		return nil, err
	} else if seatInfos != nil {
		logs.Error("ShowSeatInfo查询该座位已经存在请跳过")
		return seatInfos, nil
	}
	return nil, nil
}

func ShowSeatInfo(engine *xorm.EngineGroup, l *table.Feng) (seatInfos []table.SeatInfo, err error) {
	seatInfos, err = usedb.InquireSeat(engine, l)
	if err != nil {
		logs.Error("inquire seat err:", err)
		return nil, err
	} else if seatInfos == nil {
		logs.Error("ShowSeatInfo查询无条件座位")
		return nil, nil
	}
	return seatInfos, nil
}

func ShowFenguanCountInfo(engine *xorm.EngineGroup) (FenguanInfo []table.FenguanInfo, err error) {
	FenguanInfo, err = usedb.InquireFenguanCountInfo(engine)
	if err != nil {
		logs.Error("inquire fenguan info err:", err)
		return nil, err
	}
	return
}

func ShowFenguanNameInfo(engine *xorm.EngineGroup) (seatInfo []table.SeatInfo, err error) {
	seatInfo, err = usedb.InquireFenguanNameInfo(engine)
	if err != nil {
		logs.Error("inquire fenguan info err:", err)
		return nil, err
	}
	return
}

func ShowFenguanLoucengInfo(engine *xorm.EngineGroup, s *table.SeatInfo) (seatInfo []table.SeatInfo, err error) {
	seatInfo, err = usedb.InquireFenguanLoucengInfo(engine, s)
	if err != nil {
		logs.Error("inquire fenguan info err:", err)
		return nil, err
	}
	return
}

func ShowFenguanBianhaoInfo(engine *xorm.EngineGroup, s *table.SeatInfo) (seatInfo []table.SeatInfo, err error) {
	seatInfo, err = usedb.InquireFenguanBianhaoInfo(engine, s)
	if err != nil {
		logs.Error("inquire fenguan info err:", err)
		return nil, err
	}
	return
}
func ShowFenguanNumberInfo(engine *xorm.EngineGroup, s *table.SeatInfo) (seatInfo []table.SeatInfo, err error) {
	seatInfo, err = usedb.InquireFenguanNumberInfo(engine, s)
	if err != nil {
		logs.Error("inquire fenguan info err:", err)
		return nil, err
	}
	return
}
func ShowHistoryInfo(engine *xorm.EngineGroup, phone string) (his []table.UserHisinfo, err error) {
	his, err = usedb.InquierHistory(engine, phone)
	if err != nil {
		logs.Error("inquire fenguan info err:", err)
		return nil, err
	}
	return
}
func ShowHistoryInfoCount(engine *xorm.EngineGroup, phone string) (count int, err error) {
	count, err = usedb.InquierHistoryCount(engine, phone)
	if err != nil {
		logs.Error("inquire fenguan info err:", err)
		return
	}
	return
}

func ShowSeatInfoTuijian2(engine *xorm.EngineGroup) (seatInfo []table.SeatInfo, err error) {
	seatInfo, err = usedb.InquireTuijian2SeatInfo(engine)
	if err != nil {
		logs.Error("inquire tuijian2 info err:", err)
		return nil, err
	}
	return
}

func DeleteHistoryInfo(engine *xorm.EngineGroup, seat *table.UserHisinfo) (ok bool, err error) {
	ok, err = usedb.DeleteHistoryInfo(engine, seat)
	if err != nil {
		logs.Error("delete history info err:", err)
		return false, err
	}
	return true, nil
}

func ChangeChoiceSeatInfo(engine *xorm.EngineGroup, s *table.SeatInfo) (err error) {
	err = usedb.ChangeChoiceSeatInfo(engine, s)
	if err != nil {
		logs.Error("inquire fenguan info err:", err)
		return err
	}
	return
}

func ChangeHisSeatInfo(engine *xorm.EngineGroup, s *table.UserHisinfo) (err error) {
	err = usedb.ChangeHisSeatInfo(engine, s)
	if err != nil {
		logs.Error("inquire fenguan info err:", err)
		return err
	}
	return
}

func ShowFenguanAllInfo(engine *xorm.EngineGroup, name string) (F []table.SeatInfo, err error) {
	F, err = usedb.InquireFenguanAllInfo(engine, name)
	if err != nil {
		logs.Error("inquire fenguan info err:", err)
		return nil, err
	}
	return
}

func ShowseatExistInfo(engine *xorm.EngineGroup, seat *table.SeatInfo) (s []table.SeatInfo, err error) {
	s, err = usedb.InquireseatExistInfo(engine, seat)
	if err != nil {
		logs.Error("inquire fenguan info err:", err)
		return nil, err
	}
	return
}

func DeleteSeats(engine *xorm.EngineGroup, seat *table.DelSeatInfo) (ok bool, err error) {
	var l = &table.SeatInfo{
		Fenguan:   seat.Fenguan,
		Louceng:   seat.Louceng,
		Bianhao:   seat.Bianhao,
		Number:    seat.Number,
		SeatImage: seat.SeatImage,
	}
	ok, err = usedb.DropSeat(engine, l)
	if err != nil {
		logs.Error("Delete Seats err")
		return false, err
	}
	logs.Info("Delete Seats success")
	return
}

func DeleteAllSeats(engine *xorm.EngineGroup) (ok bool, err error) {
	ok, err = usedb.DropAllSeat(engine)
	if err != nil {
		logs.Error("Delete Seats err")
		return false, err
	}
	logs.Info("Delete Seats success")
	return
}

func ChangeSeats(engine *xorm.EngineGroup, l *table.SeatInfo) (ok bool, err error) {
	ok, err = usedb.ChangeSeatOne(engine, l)
	if err != nil {
		logs.Error("Change Seats err")
		return false, err
	}
	logs.Info("Change Seats success")
	return
}

func InsertSeats(engine *xorm.EngineGroup, l *table.InsertSeatInfo) (err error) {
	err = usedb.InsertSeats(engine, l)
	if err != nil {
		logs.Error("Insert Seats err")
		return err
	}
	logs.Info("Insert Seats err")
	return
}

func InsertHis(engine *xorm.EngineGroup, l *table.UserHisinfo) (err error) {
	err = usedb.InsertHis(engine, l)
	if err != nil {
		logs.Error("Insert Seats err")
		return err
	}
	logs.Info("Insert Seats err")
	return
}

func InsertReback(engine *xorm.EngineGroup, l *table.UserReBack) (err error) {
	err = usedb.InsertReback(engine, l)
	if err != nil {
		logs.Error("Insert Reback err")
		return err
	}
	logs.Info("Insert Reback err")
	return
}

func ChoiceSeat(engine *xorm.EngineGroup, l *table.SeatInfo) (err error) {
	_, err = ChangeSeats(engine, l)
	if err != nil {
		logs.Error("Change Seat err")
		return err
	}
	logs.Info("Change Seat err")
	return nil
}
