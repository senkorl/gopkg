package time_frequency

import (
	"fmt"
	"testing"
	"time"
)

func TestNewTimeFrequency(t *testing.T) {
	timeFrequency := NewTimeFrequency(1 * time.Second) // 每秒执行一次

	for i := 0; i < 10; i++ {
		timeFrequency.TryExecute(func() {
			fmt.Println("Action executed.")
		}, func() {
			fmt.Println("Fallback logic executed.")
		})
		time.Sleep(200 * time.Millisecond) // 模拟操作间隔
	}
}
