package service

import (
	"libseat/models/usedb"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/go-redis/redis"
)

//redis主要用来保存code 和 cookie 的
// 如果可以，黑名单机制也存在里面，即用户信用度过低，将会被加入黑名单禁止访问  黑名单存储的值 key = phonenumber ,value = "随便"

func SetRedisKey(Redisclient *redis.Client, key string, val string, time time.Duration) error {
	err := usedb.RedisSet(Redisclient, key, val, time)
	if err != nil {
		logs.Error("保存RedisKey失败:", err)
		return err
	}
	logs.Info("保存RedisKey成功", key)
	return err
}

func GetRedisKey(Redisclient *redis.Client, key string) (string, error) {
	val, err := usedb.RedisGet(Redisclient, key)
	if err != nil {
		logs.Error("获取RedisKey值失败:", err)
		return "", err
	}
	logs.Info("获取RedisKey值成功")
	return val, nil
}

func DelRedisKey(Redisclient *redis.Client, key string) error {
	err := usedb.RedisDel(Redisclient, key)
	if err != nil {
		logs.Error("Del RedisKey失败:", err)
		return err
	}
	logs.Info("Del RedisKey成功")
	return nil
}
