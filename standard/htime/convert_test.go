package htime

import "testing"

func TestExcludeSeconds(t *testing.T) {
	ExcludeSeconds("1970-01-02 11:11:11")
}
