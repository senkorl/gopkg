package github

import (
	"fmt"

	"github.com/fatih/structs"
	"github.com/spf13/cast"
)

type UserInfo struct {
	Name    string `structs:"name"`
	Age     int    `json:"age,omitempty"`
	Address string `json:"address,omitempty"`
}

func ToString() {
	foo := cast.ToString(1)
	fmt.Println(foo)
}

func ToUint32() {
	str := "123"
	foo := cast.ToUint32(str)
	fmt.Println(foo)
}

func ToMap() {
	user := structs.Map(UserInfo{
		Name:    "zhangsan",
		Age:     18,
		Address: "asdadaadasdas",
	})

	fmt.Printf("%v", user)
}
