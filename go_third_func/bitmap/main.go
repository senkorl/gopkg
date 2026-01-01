package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 设置整数范围
const MaxInt = 1 << 28 // 2^28 = 268,435,456，约2.5亿

func main() {
	// 模拟数据：随机生成一些重复的数字，部分只出现一次
	input := generateTestData()

	// 创建2位状态的位图，每4个数字需要1字节，因此总大小是 MaxInt / 4 字节
	bitmap := make([]byte, MaxInt/4)

	// 第一次遍历：更新状态位图
	for _, num := range input {
		updateBitmap(bitmap, num)
	}

	// 第二次遍历：找出只出现一次的数字
	result := []int{}
	for i := 0; i < MaxInt; i++ {
		if getState(bitmap, i) == 1 {
			result = append(result, i)
		}
	}

	fmt.Printf("找到 %d 个只出现一次的整数\n", len(result))
	for _, num := range result {
		fmt.Println(num)
	}
}

// 更新数字的状态：00->01, 01->10, 10->10
func updateBitmap(bitmap []byte, num int) {
	idx := num / 4
	offset := (num % 4) * 2

	mask := byte(0b11 << offset)
	state := (bitmap[idx] & mask) >> offset

	if state == 0 {
		// 00 -> 01
		bitmap[idx] |= (0b01 << offset)
	} else if state == 1 {
		// 01 -> 10
		bitmap[idx] &= ^mask            // 清除原位
		bitmap[idx] |= (0b10 << offset) // 设置为10
	}
	// 如果已经是10，保持不变
}

// 获取某个数字的状态（00/01/10）
func getState(bitmap []byte, num int) byte {
	idx := num / 4
	offset := (num % 4) * 2

	return (bitmap[idx] >> offset) & 0b11
}

// 模拟生成数据
func generateTestData() []int {
	rand.Seed(time.Now().UnixNano())
	data := []int{}

	// 插入重复数字
	for i := 0; i < 1000000; i++ {
		val := rand.Intn(MaxInt)
		data = append(data, val, val) // 插入两次
	}

	// 插入只出现一次的数字
	for i := 0; i < 100; i++ {
		data = append(data, MaxInt-1-i)
	}

	return data
}
