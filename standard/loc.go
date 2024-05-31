package standard

import (
	"fmt"
	"log"
	"testing"
	"time"
)

var city = []string{
	"Australia/Sydney",
	"Pacific/Auckland",
	"Europe/London",
	"Europe/Madrid",
	"America/New_York",
	"America/Toronto",
}

func tz() {
	for _, v := range city {
		loc, err := time.LoadLocation(v)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(time.Now().In(loc))
		fmt.Println(loc.String())
	}
	fmt.Println("==========")

	loc := time.FixedZone("UTC", 0)
	t1, _ := time.ParseInLocation(time.DateTime, "2024-03-15 15:32:33", loc)
	t2, _ := time.ParseInLocation(time.DateTime, "2024-03-15 15:32:33", time.Local)
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t2.In(loc))
	fmt.Println("==========")

	fmt.Println(t1.Sub(t2))
	if t1.Sub(t2) > 0 {
		fmt.Println(t1.Sub(t2).Seconds())
	}
	fmt.Println(t1.After(t2))
	fmt.Println(t1.Compare(t2))
}

func Compare(t *testing.T) {

}
