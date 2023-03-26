package tests

import (
	"fmt"
	"testing"
)

func TestGetToken(t *testing.T) {
	GetToken()
}

func TestQuote(t *testing.T) {
	Quote()
}

func TestBar(t *testing.T) {
	Bar()
}

func TestA(t *testing.T) {
	b := 5.0000
	a, _ := FormatFloat(b, 1)
	fmt.Println(a)
	fmt.Println(b)
}

func TestFoo(t *testing.T) {
	Foo()
}

func TestBar11(t *testing.T) {
	Bar11()
}
