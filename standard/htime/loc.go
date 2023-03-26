package htime

import (
	"fmt"
	"time"
)

const (
	TimeFormatStr     = "2006-01-02 15:04:05"
	DateFormatStr     = "2006-01-02"
	TimeFormatStrHHMM = "2006-01-02 15:04:00"
)

func loc() error {
	//timeStr := "2023-12-08 00:00:00"
	//timezone := time.FixedZone("UTC", 3600*+11)
	//stopTime, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	//if err != nil {
	//	return err
	//}
	//fmt.Printf("%d\n", stopTime.Unix())

	//stopTime1, err1 := time.ParseInLocation("2006-01-02 15:04:05", timeStr, timezone)
	//if err1 != nil {
	//	return err1
	//}
	//fmt.Printf("%v\n", stopTime1)

	//fmt.Printf("%s\n", stopTime)

	//st := time.Now().In(timezone).Format(DefaultFormat)
	//fmt.Printf("%s", st)
	var (
		scheduleTimeStr = time.Now().Format(TimeFormatStrHHMM)
		scheduleTime, _ = time.ParseInLocation(TimeFormatStrHHMM, scheduleTimeStr, time.Local)
	)
	loc := time.FixedZone("UTC", 3600*(10))

	//duration := int64(5 * 60)

	//todayStr := scheduleTime.Format(DateFormatStr)
	fmt.Println(scheduleTimeStr)
	today, _ := time.ParseInLocation(TimeFormatStrHHMM, scheduleTimeStr, loc)
	fmt.Println(today.In(time.Local).Format(TimeFormatStrHHMM))
	//fmt.Println(scheduleTime)
	//fmt.Println(scheduleTime.In(loc))

	startTime, _ := time.ParseInLocation(TimeFormatStr, today.Format(DateFormatStr)+" "+"19:34:00", loc)
	stopTime, _ := time.ParseInLocation(TimeFormatStr, today.Format(DateFormatStr)+" "+"23:59:59", loc)
	//fmt.Println(startTime.Unix())
	fmt.Println(startTime.In(time.Local).Format(TimeFormatStr))
	//fmt.Println(startTime.Format(TimeFormatStr))
	if scheduleTime.Before(startTime) || scheduleTime.After(stopTime) {
		fmt.Println("aaa")
	}
	timeDiff := scheduleTime.Unix() - startTime.Unix()
	fmt.Println(timeDiff)

	return nil
}
