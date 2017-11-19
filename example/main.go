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
