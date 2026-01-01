package convert_type

import (
	"fmt"

	"github.com/fatih/structs"
)

type UserInfo1 struct {
	Name    string `structs:"name"`
	Age     int    `json:"age,omitempty"`
	Address string `json:"address,omitempty"`
}

func ToMap() {
	user := structs.Map(UserInfo1{
		Name:    "zhangsan",
		Age:     18,
		Address: "asdadaadasdas",
	})

	fmt.Printf("%v", user)
}
