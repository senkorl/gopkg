package standard

import (
	"log"
	"testing"
	"unicode/utf8"
)

func TestStr(t *testing.T) {
	Str()
}

func TestZh(t *testing.T) {
	log.Println(utf8.RuneCountInString("zh"))
}
