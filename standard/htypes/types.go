package htypes

import "fmt"

func Str() {
	sample := "我爱GO"
	runeSamp := []rune(sample)
	runeSamp[0] = '你'
	fmt.Println(string(runeSamp)) // "你爱GO"
	fmt.Println(len(runeSamp))    // 4

	sample1 := "我爱GO"
	runeSamp1 := []int32(sample1)
	runeSamp1[0] = '你'
	fmt.Printf("%v", runeSamp1[0])
}
