package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	cli       *redis.Client
	addr, pwd string
	dbName    int
}

func New(addr, pwd string, dbName int) *Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,    // 没有密码，默认值
		DB:       dbName, // 默认DB 0
	})
	return &Redis{cli: rdb, addr: addr, pwd: pwd, dbName: dbName}
}

func (r *Redis) Get(key string) interface{} {
	val, err := r.cli.Get(context.Background(), key).Result()
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return val
}

func (r *Redis) Set(key string, val interface{}, exp int) error {
	_, err := r.cli.Set(context.Background(), key, val, time.Duration(exp)*time.Second).Result()
	return err
}
