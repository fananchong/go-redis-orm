package go_redis_orm

import (
	"errors"
	"time"

	sentinel "github.com/FZambia/go-sentinel"
	"github.com/garyburd/redigo/redis"
)

type RedisClient struct {
	masters    *redis.Pool
	masterName string
}

func NewRedisClient(masterName string, addrs []string) *RedisClient {
	cli := &RedisClient{}
	cli.Init(masterName, addrs)
	return cli
}

func (this *RedisClient) Init(masterName string, addrs []string) {
	sntnl := &sentinel.Sentinel{
		Addrs:      addrs,
		MasterName: masterName,
		Dial: func(addr string) (redis.Conn, error) {
			timeout := 500 * time.Millisecond
			c, err := redis.DialTimeout("tcp", addr, timeout, timeout, timeout)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
	this.masterName = masterName
	this.masters = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   0,
		Wait:        true,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			masterAddr, err := sntnl.MasterAddr()
			if err != nil {
				return nil, err
			}
			c, err := redis.Dial("tcp", masterAddr)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if !sentinel.TestRole(c, "master") {
				return errors.New("[redis] Role check failed")
			} else {
				return nil
			}
		},
	}
}

func (this *RedisClient) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	conn := this.masters.Get()
	defer conn.Close()
	if conn != nil {
		return conn.Do(commandName, args...)
	} else {
		return nil, errors.New("[redis] Can't get master!")
	}
}
