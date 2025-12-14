package main

import "fmt"

func main() {
	fmt.Println(LongestCommonPrefix([]string{"Large", "Laccc", "Lss"}))
}

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	count := len(strs)
	prefix := strs[0]
	for i := 1; i < count; i++ {
		prefix = lcp(prefix, strs[i])
		if prefix == "" {
			break
		}
	}
	return prefix
}

func lcp(s1, s2 string) string {
	length := min(len(s1), len(s2))
	index := 0
	for index < length && s1[index] == s2[index] {
		index++
	}
	return s1[:index]
}
