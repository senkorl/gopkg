package main

import (
	"fmt"
	"sort"
)

// 判断 s1 是否是 s2 的子序列
func isSubsequence(s1, s2 string) bool {
	i := 0
	for j := 0; j < len(s2) && i < len(s1); j++ {
		if s1[i] == s2[j] {
			i++
		}
	}
	return i == len(s1)
}

func findLUSlength(strs []string) int {
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) > len(strs[j])
	})

	for i, s := range strs {
		isUnique := true
		for j, t := range strs {
			if i != j && isSubsequence(s, t) {
				isUnique = false
				break
			}
		}
		if isUnique {
			return len(s)
		}
	}

	return -1
}

func main() {
	strs := []string{"aba", "cdc", "eae"}
	fmt.Println(findLUSlength(strs)) // 输出: 3

	strs = []string{"aaa", "aaa", "aa"}
	fmt.Println(findLUSlength(strs)) // 输出: -1
}
