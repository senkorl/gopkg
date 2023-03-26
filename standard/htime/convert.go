package htime

import (
	"fmt"
	"strings"
	"time"
)

// ExcludeSeconds 剔除秒数
func ExcludeSeconds(t string) {
	timeSli := strings.SplitAfter(t, "")
	timeStr := strings.Join(timeSli, "")
	fmt.Println(timeStr)

	targetTime, _ := time.Parse(TimeFormatStr, t)
	s5 := targetTime.Format(TimeFormatStrHHMM)
	fmt.Println(s5)
}
