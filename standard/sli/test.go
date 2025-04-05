package main

import (
	"fmt"
)

func main() {
	sli := make([]int, 0)
	sli = append(sli, 1)
	fmt.Println(sli, len(sli), cap(sli))
}
