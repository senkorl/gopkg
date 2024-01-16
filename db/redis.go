package db

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func ConnRds() interface{} {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
	val, err := rdb.Get(context.Background(), "test").Result()
	if err != nil {
		fmt.Println(err.Error())
	}
	return val
}
