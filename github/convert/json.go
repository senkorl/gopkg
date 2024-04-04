package convert

import (
	"fmt"
	"log"

	gojson "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/gjson"
)

type Schema struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Contact   struct {
		Phone1 string `json:"phone_1"`
		Phone2 string `json:"phone_2"`
	} `json:"contact"`
}

var yourSchema = Schema{ID: 10, FirstName: "lin", Contact: struct {
	Phone1 string `json:"phone_1"`
	Phone2 string `json:"phone_2"`
}(struct {
	Phone1 string
	Phone2 string
}{Phone1: "12312", Phone2: "12312"})}

func Iterator() {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	output, err := json.Marshal(&yourSchema)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))
}

func GoJson() {
	output, err := gojson.Marshal(&yourSchema)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))
}

func GJson() {
	s := "{\"id\":10,\"first_name\":\"lin\",\"last_name\":\"\",\"age\":0, \"contact\": {\"phone_1\": \"3123412\"}}"
	res := gjson.Get(s, "first_name")
	fmt.Println(res)
	res1 := gjson.Get(s, "contact.phone_1")
	fmt.Println(res1)
}
