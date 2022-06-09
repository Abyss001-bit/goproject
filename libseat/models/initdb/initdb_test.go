package initdb

import (
	"fmt"
	"testing"
)

func TestGetDBEngine(t *testing.T) {
	t.Run("", func(t *testing.T) {
		db := GetDBEngine()
		err := db.Ping()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("ping成功，pg数据库连通")
		}
	})
}

// $ go test
// [xorm] [info]  2022/05/23 12:52:10.991427 PING DATABASE postgres
// [xorm] [info]  2022/05/23 12:52:11.006791 PING DATABASE postgres
// [xorm] [info]  2022/05/23 12:52:11.014703 PING DATABASE postgres
// 2022/05/23 12:52:11.023 [I]  connect postgresql success
// [xorm] [info]  2022/05/23 12:52:11.023765 PING DATABASE postgres
// [xorm] [info]  2022/05/23 12:52:11.024209 PING DATABASE postgres
// [xorm] [info]  2022/05/23 12:52:11.024706 PING DATABASE postgres
// ping成功，pg数据库连通
// 2022/05/23 12:52:11.026 [I]  PONG <nil>
// ping成功，pg数据库连通: PONG
// PASS
// ok      libseat/models/initdb   0.269s

func TestGetRedisClient(t *testing.T) {
	t.Run("", func(t *testing.T) {
		db := GetRedisClient()
		pong, err := db.Ping().Result()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("ping成功，pg数据库连通:", pong)
		}
	})
}

// $ go test
// 2022/05/23 12:53:37.099 [I]  PONG <nil>
// ping成功，pg数据库连通: PONG
// PASS
// ok      libseat/models/initdb   0.277s
