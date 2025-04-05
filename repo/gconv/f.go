package gconv

import (
	"fmt"

	"github.com/gogf/gf/util/gconv"
)

func fff() {
	a := 39.89861360128482
	b := 116.30869049580403
	c := 39.949515032087575
	d := 116.38207084146282
	fmt.Println(a, b, c, d)
	fmt.Println(gconv.String(a), gconv.String(b), gconv.String(c), gconv.String(d))
	e := fmt.Sprintf("rectangle(%s,%s,%s,%s)", gconv.String(a), gconv.String(b), gconv.String(c), gconv.String(d))
	fmt.Println(e)

}
