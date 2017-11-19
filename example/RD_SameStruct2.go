package main

import (
	"github.com/fananchong/go-redis-orm"
	"github.com/garyburd/redigo/redis"
	"github.com/golang/protobuf/proto"
	"strconv"
)

type RD_SameStruct2 struct {
	SameStruct2
}

func NewRD_SameStruct2(id int32) *RD_SameStruct2 {
	this := &RD_SameStruct2{}
	this.Id = id
	return this
}

func (this *RD_SameStruct2) Key() string {
	return "SameStruct2:" + strconv.FormatInt(int64(this.Id), 10)
}

func (this *RD_SameStruct2) Value() ([]byte, error) {
	return proto.Marshal(&this.SameStruct2)
}

func (this *RD_SameStruct2) Load(dbName string) error {
	db := go_redis_orm.GetDB(dbName)
	val, err := redis.String(db.Do("GET", this.Key()))
	if err == nil {
		err = proto.Unmarshal([]byte(val), &this.SameStruct2)
	}
	return err
}

func (this *RD_SameStruct2) Save(dbName string) error {
	db := go_redis_orm.GetDB(dbName)
	val, err := this.Value()
	if err == nil {
		_, err = db.Do("SET", this.Key(), val)
	}
	return err
}
