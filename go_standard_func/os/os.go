package main

import "os"

func main() {
	_, err := os.Create("/Users/qiuxi/workspace/qudian/toy/cache/ssss.txt")
	if err != nil {
		panic(err)
	}
}
