package hreg

import "testing"

func TestGetCrawlUsername(t *testing.T) {
	dt := GetCrawlUsername("https://www.pinterest.com/dicasdemulher/")
	println(dt)
}

func TestCheck(t *testing.T) {
	Check()
}
