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

	dbName := "mysentinel"

	go_redis_orm.CreateDB(dbName, []string{"192.168.1.4:46379", "192.168.1.4:46380", "192.168.1.4:46381"})

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

### Redis哨兵、主从搭建

详细参见：http://blog.csdn.net/u013272009/article/details/78513251


### TODO:

  - go_redis_orm.CreateDB(...) 函数自识别redis类型，实例化不同redis客户端实例
