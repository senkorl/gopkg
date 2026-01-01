package main

import (
	"fmt"
)

func main() {
	sli := make([]int, 0)
	sli = append(sli, 1)
	fmt.Println(sli, len(sli), cap(sli))
	a := make(map[int]int)
	fmt.Println(a)

	b := make(chan int, 1)
	fmt.Println(b)
}
