package initdb

import (
	"fmt"

	"github.com/astaxie/beego/logs"
	"github.com/go-redis/redis"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

// pg数据库配置
const (
	// host = "192.168.255.5"
	// host     = "127.0.0.1"
	host     = "120.25.153.180"
	port     = 5432
	user     = "postgres"
	password = "147852"
	dbname   = "dblib"
)

func GetDBEngine() *xorm.EngineGroup {
	//格式
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// 连接池
	conns := []string{
		psqlInfo,
		psqlInfo,
		psqlInfo,
	}
	//引擎
	engine, err := xorm.NewEngineGroup("postgres", conns, xorm.WeightRoundRobinPolicy([]int{2, 3}))
	if err != nil {
		logs.Error("NewEngine err:", err)
		// 连接失败过着是ping不同即直接退出
		panic(err)
	}

	//ping 判断是否链接
	err = engine.Ping()
	if err != nil {
		logs.Error("connect postgresql failed:", err)
		panic(err)
	}
	logs.Info("connect postgresql success")

	return engine
}

func GetRedisClient() *redis.Client {
	//redis数据库
	redisclient := redis.NewClient(&redis.Options{
		Addr: "120.25.153.180:6379",
		// Addr:     "localhost:6379",
		// Password: "147852",
		DB: 0,
	})

	pong, err := redisclient.Ping().Result()
	if err != nil {
		panic(err)
	}
	logs.Info(pong, err)
	return redisclient
}
