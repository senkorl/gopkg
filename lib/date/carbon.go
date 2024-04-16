package date

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-module/carbon/v2"
)

func DateString() {
	fmt.Println(carbon.Now().ToString()) // 2020-08-05 13:14:15 +0800 CST
}

func DateTimeString() {
	s := "1704940744274"
	ackMsgTimeStamp, _ := strconv.ParseInt(s, 10, 64)
	ackMsgTime := time.Unix(ackMsgTimeStamp/1000, 0)
	fmt.Println(ackMsgTime)

	t := "2024-01-12 08:28:54"
	tt, _ := time.Parse(time.DateTime, t)
	fmt.Println(tt.Unix())
	//loc := time.FixedZone("UTC", 3600*4)

	cur := carbon.ParseByLayout("2015-04-12 09:00:00", time.DateTime, carbon.Sydney).ToStdTime()
	fmt.Println(cur)
	fmt.Println(carbon.Now().ToDateTimeString()) // 2020-08-05 13:14:15 +0800 CST
}
