package go_redis_orm

import (
	"github.com/fananchong/goredis"
	"github.com/garyburd/redigo/redis"
)

type DefaultRedisClient struct {
	*goredis.Client
}

func (this *DefaultRedisClient) Get(key string) ([]byte, error) {
	reply, err := redis.String(this.Client.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return []byte(reply), nil
}

func (this *DefaultRedisClient) Set(key string, data []byte) error {
	_, err := this.Client.Do("SET", key, data)
	return err
}

func (this *DefaultRedisClient) Del(key string) error {
	_, err := this.Client.Do("DEL", key)
	return err
}

func NewDefaultRedisClient(dbName string, addrs []string, password string, dbindex int) (IClient, error) {
	this := &DefaultRedisClient{}
	option := goredis.NewDefaultOption()
	option.Password = password
	option.DBIndex = dbindex
	client, err := goredis.NewClient(dbName, addrs, option)
	if err != nil {
		return nil, err
	}
	this.Client = client
	return this, nil
}
