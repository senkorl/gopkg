package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // 自动注册pprof处理器
	"runtime"
	"sync"
	"time"
)

// 模拟CPU密集型任务
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 模拟内存分配
func allocateMemory() {
	var data [][]byte

	for i := 0; i < 100; i++ {
		data = append(data, make([]byte, 1024*1024)) // 每次分配1MB
		time.Sleep(100 * time.Millisecond)
	}
	sync.Once{}
}

func main() {
	// 启动HTTP服务用于pprof分析
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	go func() {
		a := 1
		tk := time.NewTicker(time.Second)
		for {
			select {
			case <-tk.C:
				a++
			}
		}
	}()

	// 模拟业务路由
	http.HandleFunc("/cpu", func(w http.ResponseWriter, r *http.Request) {
		result := fibonacci(40) // 故意使用低效的递归算法
		fmt.Fprintf(w, "Fibonacci(40) = %d", result)
	})

	http.HandleFunc("/mem", func(w http.ResponseWriter, r *http.Request) {
		allocateMemory()
		fmt.Fprintf(w, "Memory allocated")
	})

	http.HandleFunc("/leak", func(w http.ResponseWriter, r *http.Request) {
		// 模拟goroutine泄露
		go func() {
			for {
				time.Sleep(time.Second)
				runtime.Gosched()
			}
		}()
		fmt.Fprintf(w, "Goroutine started")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
