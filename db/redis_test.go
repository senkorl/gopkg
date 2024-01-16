package db

import (
	"fmt"
	"testing"
)

func TestConnRds(t *testing.T) {
	val := ConnRds()
	if id, ok := val.(string); ok {
		fmt.Println(id)
	} else {
		fmt.Println("val")
	}
}
