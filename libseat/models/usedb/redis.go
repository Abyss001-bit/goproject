package usedb

import (
	"time"

	"github.com/beego/beego/logs"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

// 这里是数据库操作部分，不涉及业务逻辑部分

//添加键值对
func RedisSet(redisclient *redis.Client, key string, value string, time time.Duration) error {
	// 一周
	err := redisclient.Set(key, value, time).Err()
	if err != nil {
		logs.Error("redis set err:", err)
		return errors.Wrap(err, "redis set err")
	}
	logs.Info("redis set success")
	return nil
}

//查询键值对
func RedisGet(redisclient *redis.Client, key string) (string, error) {
	val, err := redisclient.Get(key).Result()
	if err != nil {
		logs.Error("redis get err:", err)
		return "", errors.Wrap(err, "redis get err")
	}
	logs.Info("redis get success")
	return val, nil
}

//删除键值对
func RedisDel(redisclient *redis.Client, key string) error {
	_, err := redisclient.Del(key).Result()
	if err != nil {
		logs.Error("rediss del err:", err)
		return errors.Wrap(err, "redis del err")
	}
	logs.Info("redis del success")
	return nil
}
