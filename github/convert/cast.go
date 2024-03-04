package convert

import (
	"fmt"

	"github.com/spf13/cast"
)

func ToString() {
	foo := cast.ToString(1)
	fmt.Println(foo)
}

func ToUint32() {
	str := "123"
	foo := cast.ToUint32(str)
	fmt.Println(foo)
}
