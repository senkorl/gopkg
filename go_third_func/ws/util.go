package main

import (
	"fmt"
	"runtime"
	"time"

	"go.uber.org/zap"
)

func Retry(f func() error, times int, interval int) error {
	var err error
	for i := 0; i < times; i++ {
		err = f()
		if err == nil {
			return nil
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
	return err
}

func Recover(logIns *zap.Logger) {
	if r := recover(); r != nil {
		const size = 64 << 10
		buf := make([]byte, size)
		buf = buf[:runtime.Stack(buf, false)]
		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("%v", r)
		}
		logIns.Error("panic", zap.Error(err), zap.Any("stack", "...\n"+string(buf)))
	}
}
