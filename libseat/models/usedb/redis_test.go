package usedb

import (
	"fmt"
	"libseat/models/initdb"
	"testing"
)

var redb = initdb.GetRedisClient()

// func TestRedisSet(t *testing.T) {

// 	type Val struct {
// 		Key   string
// 		Value string
// 		Time  time.Duration
// 	}
// 	type datatest struct {
// 		name string
// 		val  Val
// 	}

// 	// 测试用例:
// 	tests := []struct {
// 		name string
// 		val  Val
// 	}{
// 		{
// 			name: "test1",
// 			val: Val{
// 				Key:   "key1",
// 				Value: "value1",
// 			},
// 		},
// 		{
// 			name: "test2",
// 			val: Val{
// 				Key:   "key2",
// 				Value: "value2",
// 			},
// 		},
// 		{
// 			name: "test3",
// 			val: Val{
// 				Key:   "key3",
// 				Value: "value3",
// 			},
// 		},
// 	}

// 	for _, v := range tests {
// 		t.Run(v.name, func(t *testing.T) {
// 			err := RedisSet(redb, v.val.Key, v.val.Value, v.val.Time)
// 			if err != nil {
// 				fmt.Println("redis set key val time err:", err)
// 			}
// 			fmt.Println("redis set key val time success")
// 		})
// 	}
// }

// $ go test
// [xorm] [info]  2022/05/20 10:14:20.571069 PING DATABASE postgres
// [xorm] [info]  2022/05/20 10:14:20.585853 PING DATABASE postgres
// [xorm] [info]  2022/05/20 10:14:20.594654 PING DATABASE postgres
// 2022/05/20 10:14:20.600 [I]  connect postgresql success
// 2022/05/20 10:14:20.602 [I]  PONG <nil>
// 2022/05/20 10:14:20.603 [I]  redis set success
// redis set key val time success
// 2022/05/20 10:14:20.603 [I]  redis set success
// redis set key val time success
// 2022/05/20 10:14:20.604 [I]  redis set success
// redis set key val time success
// PASS
// ok      libseat/models/usedb    0.237s

// func TestRedisGet(t *testing.T) {

// 	type Val struct {
// 		Key string
// 	}

// 	// 测试用例:
// 	tests := []struct {
// 		Key string
// 	}{
// 		{
// 			Key: "key1",
// 		},
// 		{

// 			Key: "key2",
// 		},
// 		{

// 			Key: "key3",
// 		},
// 	}

// 	for _, v := range tests {
// 		t.Run("", func(t *testing.T) {
// 			val, err := RedisGet(redb, v.Key)
// 			if err != nil {
// 				fmt.Println("redis set key val time err:", err)
// 			}
// 			fmt.Println("redis set key val time success:", val)
// 		})
// 	}
// }

// $ go test
// [xorm] [info]  2022/05/23 13:00:15.729723 PING DATABASE postgres
// [xorm] [info]  2022/05/23 13:00:15.742108 PING DATABASE postgres
// [xorm] [info]  2022/05/23 13:00:15.751255 PING DATABASE postgres
// 2022/05/23 13:00:15.757 [I]  connect postgresql success
// 2022/05/23 13:00:15.758 [I]  PONG <nil>
// 2022/05/23 13:00:15.759 [I]  redis get success
// redis set key val time success: value1
// 2022/05/23 13:00:15.759 [I]  redis get success
// redis set key val time success: value2
// 2022/05/23 13:00:15.760 [I]  redis get success
// redis set key val time success: value3
// PASS
// ok      libseat/models/usedb    0.297s

func TestRedisDel(t *testing.T) {

	type Val struct {
		Key string
	}

	// 测试用例:
	tests := []struct {
		Key string
	}{
		{
			Key: "key1",
		},
		{

			Key: "key2",
		},
		{

			Key: "key3",
		},
	}

	for _, v := range tests {
		t.Run("", func(t *testing.T) {
			err := RedisDel(redb, v.Key)
			if err != nil {
				fmt.Println("redis set key val time err:", err)
			}
			fmt.Println("redis set key val time success:")
		})
	}
}

// $ go test
// [xorm] [info]  2022/05/23 13:02:46.446050 PING DATABASE postgres
// [xorm] [info]  2022/05/23 13:02:46.458541 PING DATABASE postgres
// [xorm] [info]  2022/05/23 13:02:46.465040 PING DATABASE postgres
// 2022/05/23 13:02:46.474 [I]  connect postgresql success
// 2022/05/23 13:02:46.475 [I]  PONG <nil>
// 2022/05/23 13:02:46.476 [I]  redis del success
// redis set key val time success:
// 2022/05/23 13:02:46.477 [I]  redis del success
// redis set key val time success:
// 2022/05/23 13:02:46.477 [I]  redis del success
// redis set key val time success:
// PASS
// ok      libseat/models/usedb    0.297s
