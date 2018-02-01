# go-redis-orm
redis orm base on protobuf

### 使用方法

1. 定义proto。**必须有id字段**
1. 通过redis2go.py，生产 redis orm 类

更详细内容，请参见：http://blog.csdn.net/u013272009/article/details/78574079


### 例子

```go
package main

import (
	"fmt"

	"github.com/fananchong/go-redis-orm"
)

func main() {

	dbName := "db1"

	go_redis_orm.SetNewRedisHandler(go_redis_orm.NewDefaultRedisClient)
	go_redis_orm.CreateDB(dbName, []string{"192.168.1.4:16379"}, "", 0)

	// key值为1的 TestStruct2 数据
	data2 := NewRD_TestStruct2(1)
	data2.F1 = "bbb"
	err := data2.Save(dbName)
	if err != nil {
		panic("#1")
	}

	err = data2.Load(dbName)
	if err == nil {
		if data2.F1 != "bbb" {
			panic("#2")
		}
	} else {
		panic("#3")
	}

	err = data2.Delete(dbName)
	if err != nil {
		panic("#5")
	}

	err = data2.Load(dbName)
	if data2.F1 != "" {
		panic("#6")
	}

	fmt.Println("OK")
}
```

### SetNewRedisHandler接口

本库支持第3方redis客户端整合进本项目，通过调用go_redis_orm.SetNewRedisHandler函数

需要实现go_redis_orm.IClient接口
```go
type IClient interface {
	Get(key string) ([]byte, error)
	Set(key string, data []byte) error
	Del(key string) error
}
```

例子参考：default_redis_client.go


### Redis单机、主从、哨兵、集群搭建

详细参见：http://blog.csdn.net/u013272009/article/details/78513251
