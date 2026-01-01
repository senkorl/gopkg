package main

import (
	"reflect"
	"sync"

	"github.com/robfig/cron/v3"
)

type CrontabManager struct {
	cli    *cron.Cron
	jobs   map[string]cron.EntryID
	lock   sync.Mutex
	action *Action
}

func NewCronManager() *CrontabManager {
	return &CrontabManager{
		cli:    cron.New(cron.WithSeconds()),
		jobs:   make(map[string]cron.EntryID),
		action: &Action{},
	}
}

func (c *CrontabManager) Serve() {
	c.cli.Start()
}

type ActionFunc func(param map[string]interface{})

// AddCronJob
// @Description: 添加定时任务的处理函数
func (c *CrontabManager) AddCronJob(spec, name, action string, params map[string]interface{}) error {
	if params == nil {
		params = make(map[string]interface{})
	}
	id, err := c.addJob(spec, func() {
		f := reflect.ValueOf(c.action).MethodByName(action)
		f.Call([]reflect.Value{
			reflect.ValueOf(params),
		})
	})
	if err != nil {
		return err
	}
	c.lock.Lock()
	c.jobs[name] = id
	c.lock.Unlock()
	return nil
}

// DelCronJob
// @Description: 删除定时任务的处理函数
func (c *CrontabManager) DelCronJob(name string) error {
	if c.jobs[name] == 0 {
		return nil
	}
	c.delJob(c.jobs[name])
	c.lock.Lock()
	delete(c.jobs, name)
	c.lock.Unlock()
	return nil
}

// addJob
// @Description: 添加定时任务的辅助函数
func (c *CrontabManager) addJob(spec string, cmd func()) (cron.EntryID, error) {
	id, err := c.cli.AddFunc(spec, cmd)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// delJob
// @Description: 删除定时任务的辅助函数
func (c *CrontabManager) delJob(id cron.EntryID) {
	c.cli.Remove(id)
}

func (c *CrontabManager) Stop() {
	c.cli.Stop()
}
