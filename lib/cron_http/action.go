package main

import (
	"time"
)

type Action struct{}

func NewAction() *Action {
	return &Action{}
}

func (a *Action) Test(p map[string]interface{}) {
	println("cron test, time:", time.Now().Format(time.DateTime))
	println(p)
}
