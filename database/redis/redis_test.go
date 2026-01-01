package redis

import (
	"fmt"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	rdb := New("127.0.0.1:6379", "", 0)
	str := rdb.Get("lock:order_id:1")
	fmt.Println(str)
}

func TestSet(t *testing.T) {
	rdb := New("127.0.0.1:6379", "", 0)
	err := rdb.Set("lock:order_id:1", time.Now().UTC().String(), 3600)
	if err != nil {
		fmt.Println("err:", err)
	}
}
