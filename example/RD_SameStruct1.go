package main

import (
	"github.com/fananchong/go-redis-orm"
	"github.com/garyburd/redigo/redis"
	"github.com/golang/protobuf/proto"
)

type RD_SameStruct1 struct {
	SameStruct1
}

func NewRD_SameStruct1(id string) *RD_SameStruct1 {
	this := &RD_SameStruct1{}
	this.Id = id
	return this
}

func (this *RD_SameStruct1) Key() string {
	return "SameStruct1:" + this.Id
}

func (this *RD_SameStruct1) Value() ([]byte, error) {
	return proto.Marshal(&this.SameStruct1)
}

func (this *RD_SameStruct1) Load(dbName string) error {
	db := go_redis_orm.GetDB(dbName)
	val, err := redis.String(db.Do("GET", this.Key()))
	if err == nil {
		err = proto.Unmarshal([]byte(val), &this.SameStruct1)
	}
	return err
}

func (this *RD_SameStruct1) Save(dbName string) error {
	db := go_redis_orm.GetDB(dbName)
	val, err := this.Value()
	if err == nil {
		_, err = db.Do("SET", this.Key(), val)
	}
	return err
}
