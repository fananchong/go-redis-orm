# go-redis-orm
redis orm base on protobuf

### 使用方法

1. 定义proto。**定义的结构，必须有id字段**
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

	data1 := NewRD_SameStruct1("a")
	data1.F1 = "aaaa"
	err := data1.Save(dbName)
	if err != nil {
		panic("#1")
	}

	err = data1.Load(dbName)
	if err == nil {
		if data1.F1 != "aaaa" {
			panic("#2")
		}
	} else {
		panic("#3")
	}

	data2 := NewRD_SameStruct2(1)
	data2.F1 = "bbb"
	err = data2.Save(dbName)
	if err != nil {
		panic("#4")
	}

	err = data2.Load(dbName)
	if err == nil {
		if data2.F1 != "bbb" {
			panic("#5")
		}
	} else {
		panic("#6")
	}
	fmt.Println("OK")
}

```
