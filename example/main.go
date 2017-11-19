package main

import (
	"fmt"

	"github.com/fananchong/go-redis-orm"
)

func main() {

	dbName := "mysentinel"

	go_redis_orm.CreateDB(dbName, []string{"192.168.1.4:46379", "192.168.1.4:46380", "192.168.1.4:46381"})

	// key值为"a"的 TestStruct1 数据
	data1 := NewRD_TestStruct1("a")
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

	// key值为1的 TestStruct2 数据
	data2 := NewRD_TestStruct2(1)
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
