package github

import (
	"fmt"
	"time"

	"github.com/golang-module/carbon/v2"
)

func DateString() {
	fmt.Println(carbon.Now().ToString()) // 2020-08-05 13:14:15 +0800 CST
}

func DateTimeString() {
	//loc := time.FixedZone("UTC", 3600*4)

	cur := carbon.ParseByLayout("2015-04-12 09:00:00", time.DateTime, carbon.Sydney).ToStdTime()
	fmt.Println(cur)
	fmt.Println(carbon.Now().ToDateTimeString()) // 2020-08-05 13:14:15 +0800 CST
}
