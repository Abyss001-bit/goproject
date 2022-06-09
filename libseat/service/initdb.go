package service

import (
	"libseat/models/initdb"
	"time"

	"github.com/arthurkiller/rollingwriter"
	"github.com/go-redis/redis"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

var Engine *xorm.EngineGroup
var Redisclient *redis.Client

// 初始化数据库
func init() {
	//pg引擎组
	Engine = initdb.GetDBEngine()
	//pg改变时区
	Engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	//redis引擎
	Redisclient = initdb.GetRedisClient()
}

// 初始化日志
func init() {
	//xorm操作数据库的日志
	config := rollingwriter.Config{
		LogPath:                "./log",                     //日志路径
		TimeTagFormat:          "060102150405",              //时间格式串
		FileName:               "log",                       //日志文件名
		MaxRemain:              3,                           //配置日志最大存留数
		RollingPolicy:          rollingwriter.VolumeRolling, //配置滚动策略 norolling timerolling volumerolling
		RollingTimePattern:     "* * * * * *",               //配置时间滚动策略
		RollingVolumeSize:      "1M",                        //配置截断文件下限大小
		WriterMode:             "none",
		BufferWriterThershould: 256,
		Compress:               true,
	}

	writer, err := rollingwriter.NewWriterFromConfig(&config)
	if err != nil {
		panic(err)
	}

	//xorm日志存储等级 info
	var logger *xorm.SimpleLogger = xorm.NewSimpleLogger(writer)
	Engine.SetLogger(logger)
	Engine.ShowSQL(true)
}
