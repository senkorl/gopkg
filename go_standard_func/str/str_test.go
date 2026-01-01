package str

import (
	"fmt"
	"log"
	"slices"
	"testing"
	"unicode/utf8"
)

func TestStr(t *testing.T) {
	Str()
}

func TestZh(t *testing.T) {
	log.Println(utf8.RuneCountInString("你好zh"))
}

func TestCalc(t *testing.T) {
	s := make([]float32, 0)
	fmt.Printf("%v", slices.Max(s))
}
