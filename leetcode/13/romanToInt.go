package main

import "fmt"

func main() {
	fmt.Println(romanToInt("III"))
}

func romanToInt(s string) int {
	var m = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	t := 0
	l := len(s)
	for i := 0; i < l; {
		if i < l-1 && m[s[i]] < m[s[i+1]] {
			t += m[s[i+1]] - m[s[i]]
			i = i + 2
		} else {
			t += m[s[i]]
			i++
		}
	}
	return t
}
